// Package error provides error handling utilities.
// escape analysis command: go build -gcflags '-m=3 -l'  ./pkg/error/error.go
package errors

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

// DefaultSkip is a default number of frames to skip.
var DefaultSkip = 3

// DefaultCap is a default cap for frames to store.
var DefaultCap = 10

type WithFn func(*Error)

type Error struct {
	msg        string
	httpStatus int
	frames     []Frame
	stackTrace string
	innerErr   error
}

// Frame is a single step in stack trace.
type Frame struct {
	// Func contains a function name.
	Func string
	// Line contains a line number.
	Line int
	// Path contains a file path.
	Path string
}

// New creates a new Error with message and options.
func New(msg string, optFns ...WithFn) error {
	err := inst() // lem.New.m=moved to heap: err
	err.msg = msg
	for _, fn := range optFns {
		fn(&err)
	}
	return err
}

func (e Error) HttpStatus() int {
	return e.httpStatus

}

func (e Error) Msg() string {
	return e.msg
}

func (e Error) StackTrace() string {
	return e.stackTrace
}

func (e Error) Unwrap() error {
	return e.innerErr
}

func (e Error) Error() string {
	var s string // lem.Error.m!=(leak|escape|move)
	if e.innerErr != nil {
		s = fmt.Sprintf("%s: ", e.innerErr)
	}
	return fmt.Sprintf("%s%s\n%s", s, e.msg, e.stackTrace)
}

func WithHttpStatus(httpStatus int) WithFn {
	return func(e *Error) {
		e.httpStatus = httpStatus
	}
}

func WithMsg(msg string) WithFn {
	return func(e *Error) {
		e.msg = msg
	}
}

func Wrap(err error, optFns ...WithFn) error {
	e := inst()
	e.innerErr = err
	for _, fn := range optFns {
		fn(&e)
	}
	return e
}

func inst() Error {
	frames := trace(DefaultSkip)
	return Error{
		httpStatus: http.StatusInternalServerError,
		frames:     frames,
		stackTrace: stacktrace(frames),
	}
}

func stacktrace(frames []Frame) string {
	sb := strings.Builder{}
	for _, f := range frames {
		sb.WriteString(fmt.Sprintf("%s:%d %s()\n", f.Path, f.Line, f.Func))
	}
	return sb.String()
}

func trace(skip int) []Frame {
	frames := make([]Frame, 0, DefaultCap)
	for len(frames) < cap(frames) {
		pc, path, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		frame := Frame{
			Func: fn.Name(),
			Line: line,
			Path: path,
		}
		frames = append(frames, frame)
		skip++
	}
	return frames
}
