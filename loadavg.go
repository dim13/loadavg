// +build darwin dragonfly freebsd linux netbsd openbsd solaris
// +build cgo

// Package loadavg implements get system load averages on POSIX compatible platforms
package loadavg

// #include <stdlib.h>
import "C"
import "unsafe"

// Get the number of processes in the system run queue averaged over various
// periods of time.  The system imposes a maximum of 3 samples, representing
// averages over the last 1, 5, and 15 minutes, respectively.
func Get() [3]float64 {
	var avg [3]C.double
	C.getloadavg(&avg[0], 3)
	return *(*[3]float64)(unsafe.Pointer(&avg[0]))
}
