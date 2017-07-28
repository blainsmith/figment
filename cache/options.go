package cache

import (
	"time"
)

// Option sets an optional parameter for cache or items.
type Option func(*Cache, *Item)

// NotifyHandler is used to be notified of access to cache items
type NotifyHandler func(string)

// TriggerHandler is used to perform concurrent logic
type TriggerHandler func(string, chan []byte, chan error)

// ErrorHandler is called when and error occurs
type ErrorHandler func(string, error)

// Expire is a option that is called before the item is feched from the Backend cache and checks if it has expired.
// If the item has been in the cache past its expiration time it will delete it
func Expire(d time.Duration, f NotifyHandler) Option {
	return func(c *Cache, i *Item) {
		i.before = append(i.before, func(i *Item) {
			if time.Since(i.CreatedAt) > d {
				c.Delete(i.Key, nil)
				if f != nil {
					f(i.Key)
				}
			}
		})
	}
}

// Stale is an option to implement stale-while-revalidate and stale-if-error logic for updating the cache from
// an upstream source. After the revalidation duration `rd` time as passed the TriggerHandler will be called
// and whatever is sent back over the TriggerHandler channel will get set in the cache while also resetting
// the revalidation duration. If an error is passed over the channel the ErrorHandler will be called for as long
// as the error duration is valid. After that they item will be deleted.
func Stale(rd time.Duration, f TriggerHandler, ed time.Duration, e ErrorHandler) Option {
	return func(c *Cache, i *Item) {
		i.trigger = append(i.trigger, func(i *Item) {
			if time.Since(i.CreatedAt) > rd {
				valuechan := make(chan []byte, 1)
				errchan := make(chan error, 1)

				for {
					f(i.Key, valuechan, errchan)
					value := <-valuechan
					err := <-errchan

					if err != nil {
						if time.Since(i.CreatedAt) > ed {
							c.Delete(i.Key, nil)
							e(i.Key, err)
							break
						}
					} else {
						c.Set(i.Key, value, Stale(rd, f, ed, e))
					}
				}
			}
		})
	}
}
