package main

func DoWork(c chan<- string) {
	c <- "Finished!"
}
