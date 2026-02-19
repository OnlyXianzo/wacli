//go:build windows

package lock

import (
	"syscall"
)

// Windows locking constants
const (
	winLockfileExclusiveLock   = 0x00000002
	winLockfileFailImmediately = 0x00000001
)

func flock(fd int, how int) error {
	var overlapped syscall.Overlapped
	handle := syscall.Handle(fd)
	
	if how == lockUn {
		return syscall.UnlockFileEx(handle, 0, 1, 0, &overlapped)
	}
	
	var flags uint32
	if how&lockEx != 0 {
		flags |= winLockfileExclusiveLock
	}
	if how&lockNb != 0 {
		flags |= winLockfileFailImmediately
	}
	
	return syscall.LockFileEx(handle, flags, 0, 1, 0, &overlapped)
}

const (
	lockEx = 2
	lockNb = 1
	lockUn = 0
)
