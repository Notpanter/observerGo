package main

import "fmt"

type Person string

func (name Person) HandleEvent(vacancies []string) {
	fmt.Printf("Hi, dear %s", name)
	fmt.Println("Vacancies updated:")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}
