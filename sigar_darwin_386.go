//go:build !ios
// +build !ios
package gosigar

import (
	"syscall"
	"time"
)

func (self *Uptime) Get() error {
	tv := syscall.Timeval{}

	if err := sysctlbyname("kern.boottime", &tv); err != nil {
		return err
	}

	self.Length = time.Since(time.Unix(int64(tv.Sec), int64(tv.Usec)*1000)).Seconds()

	return nil
}
