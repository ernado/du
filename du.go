// Package du provides cross-platform disc usage stats.
package du

// Info represents a disk usage statistics. All values are in bytes.
type Info struct {
	Available int64
	Free      int64
	Total     int64
}

// Get returns disk usage info and error if any.
func Get(path string) (*Info, error) {
	return getUsage(path)
}
