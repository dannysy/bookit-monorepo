package error

import (
	"fmt"

	"github.com/ztrue/tracerr"
)

func New(msg string) error {
	return tracerr.New(msg)

}

func Errorf(format string, a ...any) error {
	return tracerr.Errorf(format, a...)
}

func Wrap(err error, msg string) error {
	return tracerr.Wrap(fmt.Errorf("%s: %w", msg, err))
}
