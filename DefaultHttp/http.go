package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	svr := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(time.Hour)
			},
		),
	)
	defer svr.Close()

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	fmt.Println("Making the request")
	_, err := client.Get(svr.URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Finished Request")
}
