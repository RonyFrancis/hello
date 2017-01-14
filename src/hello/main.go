package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	v := url.Values{}
	v.Set("msisdn", "8J0sE0LtvIAxRfubIDBiGg==") // encrypted msisdn 1212121212
	//custid := "vPTGEkzVsvLGOcF77PVa9g=="
	//phoneNo := "8J0sE0LtvIAxRfubIDBiGg=="

	resp, err := http.PostForm("https://send.bobsecure.com/v3/users/register", v)
	// post call
	if err != nil {
		os.Exit(0)
	}

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(0)
	}
	//resp.Body.close()

	fmt.Printf(string(robots))

}
