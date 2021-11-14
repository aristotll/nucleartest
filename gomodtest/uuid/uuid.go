package main

import "github.com/google/uuid"

func fileName2Uuid(filenames []string) {
	for _, name := range filenames {
		u := uuid.New()
		_ = u
	}
}

func main() {

}
