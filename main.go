package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	test("https://rydii.com/", 1, 100)
}

func test(url string, round int, concurrent int) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	for i := 1; i <= round; i++ {
		fmt.Println(fmt.Sprintf("\n==== Round %d ====", i))
		roundHit(req, concurrent)
	}
}

func roundHit(req *http.Request, concurrent int) {
	var wg sync.WaitGroup
	wg.Add(concurrent)
	for i := 0; i < concurrent; i++ {
		go concurrentHit(req, &wg)
	}
	wg.Wait()
}

func concurrentHit(req *http.Request, wg *sync.WaitGroup) {
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("x")
	} else {
		fmt.Printf(".")
	}
	wg.Done()
}
