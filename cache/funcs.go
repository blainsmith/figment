package cache

// ItemFunc may take in an Item and manipulate
type ItemFunc func(*Item)

// Trigger functions are executed on the Item while concurrently accessing it from the Backend cache.
func Trigger(trigger ...ItemFunc) Option {
	return func(c *Cache, i *Item) { i.trigger = append(i.trigger, trigger...) }
}

// Before functions are executed on the Item before accessing it from the Backend cache.
func Before(before ...ItemFunc) Option {
	return func(c *Cache, i *Item) { i.before = append(i.before, before...) }
}

// After functions are executed on the Item after accessing it from the Backend cache.
func After(after ...ItemFunc) Option {
	return func(c *Cache, i *Item) { i.after = append(i.after, after...) }
}
