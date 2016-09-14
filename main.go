package main

import (
	"fmt"
	"gopkg.in/redis.v4"
	"io"
	"net/http"
)

var (
	client *redis.Client
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		io.WriteString(w, "favicon")
		return
	}
	count, err := client.Incr("count").Result()
	if err != nil {
		io.WriteString(w, "Redis is unhappy")
	} else {
		io.WriteString(w, fmt.Sprintln("View Count: ", count))
	}
}

func main() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	http.HandleFunc("/", hello)
	http.ListenAndServe(":80", nil)
}
