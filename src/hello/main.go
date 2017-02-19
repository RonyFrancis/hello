package main

import (
	"controller/customer"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()
	f, _ := os.Create("logger2.log")
	var MyFile *log.Logger
	var ErrorLogger *log.Logger
	f2, _ := os.Create("logger.log")
	MyFile = log.New(f,
		"PREFIX: ",
		log.Ldate|log.Ltime|log.Lshortfile) // satndard logger
	ErrorLogger = log.New(f2,
		"PREFIX: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	MyFile.Println("myfile is a logger")
	ErrorLogger.Println("error logger is a log file") // Error logger
	// Logger := log.SetOutput(f)
	// Logger.Println("log has begun")
	var c chan string = make(chan string)
	for i := 1; i <= 6; i++ {
		go postcall(c)
	}
	for i := 1; i <= 6; i++ {
		msg := <-c
		fmt.Printf("$")
		fmt.Printf(msg)
	}
	t2 := time.Now()
	fmt.Println(t1.Sub(t2))
	http.HandleFunc("/uk/users/register", customer.RegisterHandler)
	http.HandleFunc("/uk/users/1/verify", customer.OtpVerifyHandler)
	http.HandleFunc("/uk/users/:id/accounts", customer.AccountHandler)
	http.HandleFunc("/uk/users/1/accounts/1/transactions", customer.TransHandler)
	http.HandleFunc("/uk/messages", customer.InboxHandler)
	http.ListenAndServe(":8085", nil)

}

func postcall(c chan string) {

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
	c <- string(robots)

}
