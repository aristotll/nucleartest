package main

import (
	"structz/pkga"
	"structz/pkgb"
)

func main() {
	topa := &pkga.Top{}
	topb := &pkgb.Top{}

	pkga.Do((*pkga.Top)(topb))
	pkgb.Do((*pkgb.Top)(topa))	
}