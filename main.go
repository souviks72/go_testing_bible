package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Getting started with testify")
	devs := []Developer{
		{
			"Souvik", 25,
		},
		{
			"Souvik", 26,
		},
		{
			"Soumi", 25,
		},
		{
			"Souradeep", 25,
		},
		{
			"Shubhadeep", 25,
		},
		{
			"Shubhadeep", 25,
		},
	}

	uniques := FilterUnique(devs)
	fmt.Println(uniques)
}

func FilterUnique(developers []Developer) []string {
	var uniques []string
	check := make(map[string]int)

	for _, dev := range developers {
		check[dev.Name] = 1
	}

	for name := range check {
		uniques = append(uniques, name)
	}

	return uniques
}
