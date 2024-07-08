package errors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/akutz/lem"
	"github.com/stretchr/testify/assert"
)

func Wrapped() error {
	return Wrap(fmt.Errorf("wrapped error"), WithMsg("upper error"))
}

func NewNested() error {
	fmt.Println("nested first line")
	return New("nested error")
}

func UpperNew() error {
	fmt.Println("upper first line")
	return NewNested()
}

func TestTraceNew(t *testing.T) {
	err := UpperNew()
	fmt.Printf("\n%+v", err)
}

func TestTraceWrapped(t *testing.T) {
	err := Wrapped()
	fmt.Printf("\n%+v", err)
}

func TestNewWithHttpStatus(t *testing.T) {
	err := New("new error", WithHttpStatus(http.StatusNotFound))
	var e Error
	if !errors.As(err, &e) {
		t.Fatal("wrong error type")
	}
	assert.Equal(t, http.StatusNotFound, e.httpStatus)
}

func TestEscapeAnalysis(t *testing.T) {
	lem.Run(t)
}
