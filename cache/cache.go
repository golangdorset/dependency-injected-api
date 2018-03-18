package cache

// Fake is a no-op cache
type Fake struct{}

// Set is a no-op
func (f *Fake) Set(interface{}) {}

// Get is a no-op
func (f *Fake) Get(key string) interface{} { return nil }

// Del is a no-op
func (f *Fake) Del(key string) bool { return true }
