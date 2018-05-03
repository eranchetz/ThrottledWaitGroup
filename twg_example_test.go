package throttledwaitgroup_test

import (
	"math/rand"
	"net/http"
	"time"

	TWG "github.com/eranchetz/ThrottledWaitGroup-go"
)

const exampleURL = "https://exmaple.com"

func ExampleTWG() {

	rand.Seed(time.Now().UnixNano())

	twg := TWG.New(8)
	for i := 0; i < 50; i++ {
		twg.Add()
		go func(i int) {
			defer twg.Done()
			query(exampleURL)
		}(i)
	}

	twg.Wait()
}

func query(url string) {
	http.Get(url)
}
