package main

import . "github.com/ahmetb/go-linq/v3"

type Book struct {
	id      int
	title   string
	authors []string
}

func f() {
	author := From(books).SelectMany( // make a flat array of authors
		func(book interface{}) Query {
			return From(book.(Book).authors)
		}).GroupBy( // group by author
		func(author interface{}) interface{} {
			return author // author as key
		}, func(author interface{}) interface{} {
			return author // author as value
		}).OrderByDescending( // sort groups by its length
		func(group interface{}) interface{} {
			return len(group.(Group).Group)
		}).Select( // get authors out of groups
		func(group interface{}) interface{} {
			return group.(Group).Key
		}).First() // take the first author
}

func ff() {
	author := From(books).
		SelectMany((book) => { return From(book.(Book).authors) }).
		GroupBy((author) => { return author }, (author) => { return author }).
		OrderByDescending((group) => { return len(group.(Group).Group) }).
		Select((group) => { return group.(Group).Key }).
		First()
}

type Car struct {
	year         int
	owner, model string
}

var owners []string

func a() {
	From(cars).Where(func(c interface{}) bool {
		return c.(Car).year >= 2015
	}).Select(func(c interface{}) interface{} {
		return c.(Car).owner
	}).ToSlice(&owners)
}

func aa() {
	From(cars).
		Where((c) => { return c.(Car).year >= 2015 }).
		Select((c) => { return c.(Car).owner }).
		ToSlice(&owners)
}

