// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"pkg.berachain.dev/jinx/lib/types"
	"sync"
)

// Ensure, that SnapshottableMock does implement types.Snapshottable.
// If this is not the case, regenerate this file with moq.
var _ types.Snapshottable = &SnapshottableMock{}

// SnapshottableMock is a mock implementation of types.Snapshottable.
//
//	func TestSomethingThatUsesSnapshottable(t *testing.T) {
//
//		// make and configure a mocked types.Snapshottable
//		mockedSnapshottable := &SnapshottableMock{
//			RevertToSnapshotFunc: func(n int)  {
//				panic("mock out the RevertToSnapshot method")
//			},
//			SnapshotFunc: func() int {
//				panic("mock out the Snapshot method")
//			},
//		}
//
//		// use mockedSnapshottable in code that requires types.Snapshottable
//		// and then make assertions.
//
//	}
type SnapshottableMock struct {
	// RevertToSnapshotFunc mocks the RevertToSnapshot method.
	RevertToSnapshotFunc func(n int)

	// SnapshotFunc mocks the Snapshot method.
	SnapshotFunc func() int

	// calls tracks calls to the methods.
	calls struct {
		// RevertToSnapshot holds details about calls to the RevertToSnapshot method.
		RevertToSnapshot []struct {
			// N is the n argument value.
			N int
		}
		// Snapshot holds details about calls to the Snapshot method.
		Snapshot []struct {
		}
	}
	lockRevertToSnapshot sync.RWMutex
	lockSnapshot         sync.RWMutex
}

// RevertToSnapshot calls RevertToSnapshotFunc.
func (mock *SnapshottableMock) RevertToSnapshot(n int) {
	if mock.RevertToSnapshotFunc == nil {
		panic("SnapshottableMock.RevertToSnapshotFunc: method is nil but Snapshottable.RevertToSnapshot was just called")
	}
	callInfo := struct {
		N int
	}{
		N: n,
	}
	mock.lockRevertToSnapshot.Lock()
	mock.calls.RevertToSnapshot = append(mock.calls.RevertToSnapshot, callInfo)
	mock.lockRevertToSnapshot.Unlock()
	mock.RevertToSnapshotFunc(n)
}

// RevertToSnapshotCalls gets all the calls that were made to RevertToSnapshot.
// Check the length with:
//
//	len(mockedSnapshottable.RevertToSnapshotCalls())
func (mock *SnapshottableMock) RevertToSnapshotCalls() []struct {
	N int
} {
	var calls []struct {
		N int
	}
	mock.lockRevertToSnapshot.RLock()
	calls = mock.calls.RevertToSnapshot
	mock.lockRevertToSnapshot.RUnlock()
	return calls
}

// Snapshot calls SnapshotFunc.
func (mock *SnapshottableMock) Snapshot() int {
	if mock.SnapshotFunc == nil {
		panic("SnapshottableMock.SnapshotFunc: method is nil but Snapshottable.Snapshot was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSnapshot.Lock()
	mock.calls.Snapshot = append(mock.calls.Snapshot, callInfo)
	mock.lockSnapshot.Unlock()
	return mock.SnapshotFunc()
}

// SnapshotCalls gets all the calls that were made to Snapshot.
// Check the length with:
//
//	len(mockedSnapshottable.SnapshotCalls())
func (mock *SnapshottableMock) SnapshotCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSnapshot.RLock()
	calls = mock.calls.Snapshot
	mock.lockSnapshot.RUnlock()
	return calls
}
