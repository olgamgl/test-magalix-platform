package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func fib(number int) int {
	if number == 0 || number == 1 {
		return number
	}

	return fib(number-2) + fib(number-1)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//log the call to the standard I/O. You can read these logs at Application web console
	log.Println(fmt.Sprintf("request recieved - path:  %s", r.URL.Path))

	number, err := strconv.Atoi(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintf(w, "Bad Input: %s. Error message: %s", r.URL.Path[1:], err.Error())
		log.Println(fmt.Sprintf("bad input  %s. Error: %s", r.URL.Path, err.Error()))
	} else {
		result := fib(number)
		fmt.Fprintf(w, fmt.Sprintf("Fibonacci result of %s is %s", strconv.Itoa(number), strconv.Itoa(result)))
	}

}

func main() {
	//http request handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
	//listen to port 8080
	http.ListenAndServe(":8080", nil)
}

//Message is the http parameter
type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {

	//send json text back when user calls the /about
	m := Message{"Welcome to the Simple Service API, v1.0"}
	log.Println("request recieve - path:  about")
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}
