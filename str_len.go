package main

import (
	"fmt"
)

func main() {
	s := "$2a$10$wYolC.QoIxjSFAZK2/R5y.1izmhhLYcJgoY2CxMa9mXUDvij41UOG"
	s1 := "$2a$10$txBJA9UfZQrX8.2OXLyzlO7Ft08cEy/YcmgUok74ENvbLZd2qqcci"
	s2 := "$2a$10$c23NfOLVAEJ/R/fRx1CBu.NTDkR5vmkJ/gvpEVQbydFFJptRHZw7."
	s3 := "$2a$10$oOJ/X5kbyC15llWImgbfY.aY9Aviy.aDTZVPlUnDpR8cBz5eXB/Ui"
	s4 := "$2a$10$CDFUYkToiA/5No8KIB7AlOUVyC62k660MCVdUhyln0qHfPrimX6le"

	fmt.Printf("%d %d %d %d %d \n", len(s), len(s1), len(s2), len(s3), len(s4))
}
