package main

func main() {
	exit := make(chan bool)

	go selectGrib()

	go hold(exit)

	for {
		if <-exit {
			return
		}
	}
}
