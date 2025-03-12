package main

import (
	"fmt"
)

func main() {
	var (
		userInput int
		err       error
	)

	fmt.Print("E2E scripts \n\n")
	fmt.Println(`Choose script to run:
	1 - create superuser
	2 - fill db with test data
	3- clear all tables`)

	_, err = fmt.Scanf("%d", &userInput)
	if err != nil {
		panic(err)
	}

	fmt.Println(userInput, err)
}
