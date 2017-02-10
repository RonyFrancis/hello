package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()
	var c chan string = make(chan string)
	for i := 1; i <= 10; i++ {

		fmt.Print("#")
		fmt.Println(postcall())
	}
	t2 := time.Now()
	fmt.Println(t1.Sub(t2))

}

func postcall() string {
	v := url.Values{}
	v.Set("no", strconv.Itoa(rand.Intn(100))) // encrypted msisdn 1212121212
	//custid := "vPTGEkzVsvLGOcF77PVa9g=="
	//phoneNo := "8J0sE0LtvIAxRfubIDBiGg=="

	resp, err := http.PostForm("http://localhost:3000/uk/welcome", v)
	// post call
	if err != nil {
		os.Exit(0)
	}

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(0)
	}
	//resp.Body.close()

	return string(robots)

}
