package main

import (
	//...
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/pkg/profile"
)

var cache = make(map[int]int)
var m sync.Mutex

// ...
func main() {
	defer profile.Start(profile.MemProfile).Stop()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		http.ListenAndServe(":8080", nil)
		wg.Done()
	}()
	go func() {
		t := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-t.C:
				r := rand.Int()
				m.Lock()
				cache[r] = r
				m.Unlock()
				println(cache[r])
			}
		}
	}()
	go func() {
		t := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-t.C:
				r := rand.Int()
				m.Lock()
				cache[r] = r
				m.Unlock()
				println(cache[r])
			}
		}
	}()

	wg.Wait()

}
