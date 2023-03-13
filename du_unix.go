//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package du

import (
	"syscall"
)

func getUsage(path string) (*Info, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return nil, err
	}
	return &Info{
		Total:     stat.Bsize * int64(stat.Blocks),
		Available: stat.Bsize * int64(stat.Bavail),
		Free:      stat.Bsize * int64(stat.Bfree),
	}, nil
}
