package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	rte "github.com/go-guru-academy/rte-magic"
	"github.com/gorilla/mux"
)

// Custom input to be passed down the request chain
type MyInput struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	fmt.Println("example-01: started")
	r := rte.New(rte.GORILLA)
	r.Get(
		"/example/{id}",
		example,
		&MyInput{
			Hobbies: []string{
				"coding",
				"debugging",
			},
		},
		[]rte.Middleware{
			middleware01,
			middleware02,
		})
	if err := http.ListenAndServe(":8888", r); err != nil {
		fmt.Println(err)
		return
	}
}

func middleware01(next http.Handler) http.Handler {
	fmt.Println("middleware01: appended to middleware chain")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware01: pre-request")
		input := rte.GetInput(w)
		i := input.(*MyInput)
		i.Name = "eric"
		next.ServeHTTP(w, r)
		fmt.Println("middleware01: post-request")
	})
}

func middleware02(next http.Handler) http.Handler {
	fmt.Println("middleware02: appended to middleware chain")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware02: pre-request")
		input := rte.GetInput(w)
		i := input.(*MyInput)
		i.Age = 36
		next.ServeHTTP(w, r)
		fmt.Println("middleware02: post-request")
	})
}

func example(w http.ResponseWriter, r *http.Request) {
	fmt.Println("status: request received")
	// Get path vars
	vars := mux.Vars(r)
	fmt.Println(vars["id"])

	// Get custom input
	input := rte.GetInput(w)
	i := input.(*MyInput)
	fmt.Printf("input: %+v\n", i)

	// Write response
	response, _ := json.Marshal(i)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
