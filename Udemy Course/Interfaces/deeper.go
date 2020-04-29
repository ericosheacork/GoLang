package main

type user struct {
	name string
}

type bot interface {
	getGreeting(string, int) (string, error)
	getBotversion() float64
	respondToUser(user) string
}

func main() {

}
