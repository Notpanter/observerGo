package main

type Observer interface {
	HandleEvent(vacancies []string)
}