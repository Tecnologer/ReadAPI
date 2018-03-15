package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Post struct of data from JSon
type Post struct {
	UserID int    `json:"userid"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}

	url := "http://jsonplaceholder.typicode.com/posts"

	resp, err := client.Get(url)

	panicErr(err)

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)

	panicErr(err)

	fmt.Println(posts)
}

func panicErr(e error) {
	if e != nil {
		panic(e.Error())
	}
}
