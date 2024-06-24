package main

func main() {
	for c1 := 'a'; c1 <= 'z'; c1++ {
		print(string(c1))
	}
	for c2 := 'a'-1 ; c2 <= 'z'; c2++ {
		println(string(c2))
		println()
	}
}

// package main

// func main() {
// 	for c := 'z'; c >= 'a'; c-- {
// 		print(string(c))
// 	}
// }