package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"bookit/pkg/logger"
)

type Command struct {
	closableResources []Closable
	runnableResources []Runnable
}

func NewCommand(closableResources []Closable, runnableResources []Runnable) *Command {
	return &Command{
		closableResources: closableResources,
		runnableResources: runnableResources,
	}
}

func (c *Command) Do() {
	doneCh := make(chan struct{})
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)
	go c.run(doneCh)
	select {
	case <-doneCh:
		c.release()
	case <-quitCh:
		c.release()
	}
}

func (c *Command) release() {
	for _, resource := range c.closableResources {
		err := resource.Close()
		if err != nil {
			logger.Gist(context.Background()).
				Err(err).
				Msg("failed to close resource")
		}
	}
	for _, resource := range c.runnableResources {
		err := resource.Stop()
		if err != nil {
			logger.Gist(context.Background()).
				Err(err).
				Msg("failed to stop resource")
		}

	}
}

func (c *Command) run(doneCh chan<- struct{}) {
	wg := sync.WaitGroup{}
	wg.Add(len(c.runnableResources))
	for _, resource := range c.runnableResources {
		go func(resource Runnable) {
			defer wg.Done()
			err := resource.Start()
			if err != nil {
				logger.Gist(context.Background()).
					Err(err).
					Msg("failed to run resource")
			}
		}(resource)
	}
	wg.Wait()
	doneCh <- struct{}{}
}
