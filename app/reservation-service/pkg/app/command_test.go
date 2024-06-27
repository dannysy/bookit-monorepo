package app

import (
	"fmt"
	"sync"
	"syscall"
	"testing"
	"time"
)

type Resource struct {
	DoneCh chan struct{}
}

func (c *Resource) Start() error {
	fmt.Println("starting...")
	<-c.DoneCh
	return nil
}

func (c *Resource) Stop() error {
	fmt.Println("stopping...")
	c.DoneCh <- struct{}{}
	return nil
}

func (c *Resource) Close() error {
	fmt.Println("closing...")
	return nil
}

type ErrResource struct {
}

func (c ErrResource) Start() error {
	fmt.Println("starting...")
	return fmt.Errorf("failed to start")
}

func (c ErrResource) Stop() error {
	fmt.Println("stopping...")
	return nil
}

func (c ErrResource) Close() error {
	fmt.Println("closing...")
	return fmt.Errorf("failed to close")
}

func Test_ShouldNotBlockWithoutResources(t *testing.T) {
	cmd := NewCommand(nil, nil)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wgroup *sync.WaitGroup) {
		fmt.Println("before Do")
		cmd.Do()
		fmt.Println("after Do")
		wgroup.Done()
		fmt.Println("after wg.Done")
	}(&wg)
	wg.Wait()
	fmt.Println("finished")
}

func Test_ShouldCloseResources(t *testing.T) {
	cmd := NewCommand([]Closable{&Resource{}, &Resource{}}, nil)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wgroup *sync.WaitGroup) {
		fmt.Println("before Do")
		cmd.Do()
		fmt.Println("after Do")
		wgroup.Done()
		fmt.Println("after wg.Done")
	}(&wg)
	wg.Wait()
	fmt.Println("finished")
}

func TestShouldNotCloseResources(t *testing.T) {
	cmd := NewCommand([]Closable{ErrResource{}, &Resource{}}, nil)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wgroup *sync.WaitGroup) {
		fmt.Println("before Do")
		cmd.Do()
		fmt.Println("after Do")
		wgroup.Done()
		fmt.Println("after wg.Done")
	}(&wg)
	wg.Wait()
	fmt.Println("finished")
}

func Test_ShouldStopResources(t *testing.T) {
	cmd := NewCommand(nil, []Runnable{&Resource{DoneCh: make(chan struct{})}})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wgroup *sync.WaitGroup) {
		fmt.Println("before Do")
		cmd.Do()
		fmt.Println("after Do")
		wgroup.Done()
		fmt.Println("after wg.Done")
	}(&wg)
	time.Sleep(50 * time.Millisecond)
	fmt.Println("sending sigterm")
	err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	fmt.Println("finished")
}

func Test_ShouldNotStartResources(t *testing.T) {
	cmd := NewCommand(nil, []Runnable{ErrResource{}})
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wgroup *sync.WaitGroup) {
		fmt.Println("before Do")
		cmd.Do()
		fmt.Println("after Do")
		wgroup.Done()
		fmt.Println("after wg.Done")
	}(&wg)
	wg.Wait()
	fmt.Println("finished")
}
