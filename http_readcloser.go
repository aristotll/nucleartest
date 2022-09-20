package main

import (
	"fmt"
	"net/http"
	"io"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(os.Stdout, "read request body from io.Copy: ")
		io.Copy(os.Stdout, r.Body)
		io.WriteString(os.Stdout, "\n")

        io.NopCloser(r.Body)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("read request body error: ", err)
			return
		}
		fmt.Println("read request body from io.ReadAll: ", body)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
