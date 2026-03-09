package logging

import "sync"

// resetGlobalState resets the global logger state for testing purposes.
func resetGlobalState() {
	defaultLogger = nil
	once = sync.Once{}
	initErr = nil
}
