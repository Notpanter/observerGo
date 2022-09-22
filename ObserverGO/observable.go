package main

type Observable interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	sendAll()
}
