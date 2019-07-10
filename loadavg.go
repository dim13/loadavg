// +build darwin dragonfly freebsd linux netbsd nacl openbsd solaris
// +build cgo

// Package loadavg implements get system load averages on POSIX compatible platforms
package loadavg

// #include <stdlib.h>
import "C"
import "unsafe"

// Get system load averages
func Get() [3]float64 {
	var avg [3]C.double
	C.getloadavg(&avg[0], 3)
	return *(*[3]float64)(unsafe.Pointer(&avg[0]))
}
