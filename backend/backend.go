package backend

// Backend defines the interface for the Cache to use for storing data.
type Backend interface {
	Get(interface{}) (interface{}, error)
	Set(interface{}, interface{}) error
	Delete(interface{}) error
}
