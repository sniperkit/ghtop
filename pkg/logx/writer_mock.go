// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package logx

import (
	"sync"
)

var (
	lockWriterMockWrite sync.RWMutex
)

// WriterMock is a mock implementation of Writer.
//
//     func TestSomethingThatUsesWriter(t *testing.T) {
//
//         // make and configure a mocked Writer
//         mockedWriter := &WriterMock{
//             WriteFunc: func(p []byte) (int, error) {
// 	               panic("TODO: mock out the Write method")
//             },
//         }
//
//         // TODO: use mockedWriter in code that requires Writer
//         //       and then make assertions.
//
//     }
type WriterMock struct {
	// WriteFunc mocks the Write method.
	WriteFunc func(p []byte) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// Write holds details about calls to the Write method.
		Write []struct {
			// P is the p argument value.
			P []byte
		}
	}
}

// Write calls WriteFunc.
func (mock *WriterMock) Write(p []byte) (int, error) {
	if mock.WriteFunc == nil {
		panic("moq: WriterMock.WriteFunc is nil but Writer.Write was just called")
	}
	callInfo := struct {
		P []byte
	}{
		P: p,
	}
	lockWriterMockWrite.Lock()
	mock.calls.Write = append(mock.calls.Write, callInfo)
	lockWriterMockWrite.Unlock()
	return mock.WriteFunc(p)
}

// WriteCalls gets all the calls that were made to Write.
// Check the length with:
//     len(mockedWriter.WriteCalls())
func (mock *WriterMock) WriteCalls() []struct {
	P []byte
} {
	var calls []struct {
		P []byte
	}
	lockWriterMockWrite.RLock()
	calls = mock.calls.Write
	lockWriterMockWrite.RUnlock()
	return calls
}
