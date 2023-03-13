//go:build (darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris) && (arm || 386 || riscv)

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
		Total:     int64(stat.Bsize) * int64(stat.Blocks),
		Available: int64(stat.Bsize) * int64(stat.Bavail),
		Free:      int64(stat.Bsize) * int64(stat.Bfree),
	}, nil
}
