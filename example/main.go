package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/blainsmith/figment/backend"
	"github.com/blainsmith/figment/cache"
)

type item struct {
	key   string
	value string
}

func (i item) Key() interface{} {
	return i.key
}
func (i item) Value() interface{} {
	return i.value
}

func main() {
	s := backend.Map()
	c := cache.New(s)

	// expiredFunc := func(k string) {
	// 	fmt.Println("key expired", k)
	// }
	staleHandler := func(k string, v chan []byte, err chan error) {
		time.Sleep(time.Duration(5) * time.Second)
		v <- nil
		err <- errors.New("failed")
	}
	errorHandler := func(k string, e error) {
		fmt.Println(k, e)
	}
	// loadFunc := func(k interface{}) interface{} {
	// 	return nil
	// }

	c.Set("test", []byte("testing"), cache.Stale(time.Duration(2)*time.Second, staleHandler, time.Duration(10)*time.Second, errorHandler))
	i, err := c.Get("test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	time.Sleep(time.Duration(4) * time.Second)

	i, err = c.Get("test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	time.Sleep(time.Duration(4) * time.Second)

	i, err = c.Get("test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	time.Sleep(time.Duration(4) * time.Second)

	i, err = c.Get("test", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
