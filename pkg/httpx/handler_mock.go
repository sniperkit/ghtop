// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package httpx

import (
	"net/http"
	"sync"
)

var (
	lockHandlerMockServeHTTP sync.RWMutex
)

// HandlerMock is a mock implementation of Handler.
//
//     func TestSomethingThatUsesHandler(t *testing.T) {
//
//         // make and configure a mocked Handler
//         mockedHandler := &HandlerMock{
//             ServeHTTPFunc: func(in1 http.ResponseWriter, in2 *http.Request)  {
// 	               panic("TODO: mock out the ServeHTTP method")
//             },
//         }
//
//         // TODO: use mockedHandler in code that requires Handler
//         //       and then make assertions.
//
//     }
type HandlerMock struct {
	// ServeHTTPFunc mocks the ServeHTTP method.
	ServeHTTPFunc func(in1 http.ResponseWriter, in2 *http.Request)

	// calls tracks calls to the methods.
	calls struct {
		// ServeHTTP holds details about calls to the ServeHTTP method.
		ServeHTTP []struct {
			// In1 is the in1 argument value.
			In1 http.ResponseWriter
			// In2 is the in2 argument value.
			In2 *http.Request
		}
	}
}

// ServeHTTP calls ServeHTTPFunc.
func (mock *HandlerMock) ServeHTTP(in1 http.ResponseWriter, in2 *http.Request) {
	if mock.ServeHTTPFunc == nil {
		panic("moq: HandlerMock.ServeHTTPFunc is nil but Handler.ServeHTTP was just called")
	}
	callInfo := struct {
		In1 http.ResponseWriter
		In2 *http.Request
	}{
		In1: in1,
		In2: in2,
	}
	lockHandlerMockServeHTTP.Lock()
	mock.calls.ServeHTTP = append(mock.calls.ServeHTTP, callInfo)
	lockHandlerMockServeHTTP.Unlock()
	mock.ServeHTTPFunc(in1, in2)
}

// ServeHTTPCalls gets all the calls that were made to ServeHTTP.
// Check the length with:
//     len(mockedHandler.ServeHTTPCalls())
func (mock *HandlerMock) ServeHTTPCalls() []struct {
	In1 http.ResponseWriter
	In2 *http.Request
} {
	var calls []struct {
		In1 http.ResponseWriter
		In2 *http.Request
	}
	lockHandlerMockServeHTTP.RLock()
	calls = mock.calls.ServeHTTP
	lockHandlerMockServeHTTP.RUnlock()
	return calls
}
