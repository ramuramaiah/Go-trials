package go_routine_trial

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func Run_Func_Concurrently() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("End. Type any key to continue..")
	var input string
	fmt.Scanln(&input)
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func Channel_Test() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	time.Sleep(2 * time.Second)
	fmt.Println("End. Type any key to continue..")
	var input string
	fmt.Scanln(&input)
}

func Channel_Select_Test() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				//default:
				//	fmt.Println("nothing ready")
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("End. Type any key to continue..")
	var input string
	fmt.Scanln(&input)
}

func Biggest_Home_Page_Size() {
	type HomePageSize struct {
		URL  string
		Size int
	}

	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	results := make(chan HomePageSize, 4)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("The biggest home page:", biggest.URL)
}
