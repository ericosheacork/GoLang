package main

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, index := range numbers {
		if index%2 == 0 {
			println(index, "is Even")
		} else {
			println(index, "is odd")
		}

	}
}
