package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}

	distribution = make(map[string]int, len(users))
)

func calcCoin(username string) int {
	coin := 0

	for _, value := range username {
		switch value {
		case 'a', 'A':
			coin = coin + 1
		case 'e', 'E':
			coin = coin + 1
		case 'i', 'I':
			coin = coin + 2
		case 'o', 'O':
			coin = coin + 3
		case 'u', 'U':
			coin = coin + 5
		}
	}
	return coin
}

func dispathCoin() int {
	left := coins
	for _, username := range users {
		allCoin := calcCoin(username)
		left = left - allCoin
		value, ok := distribution[username]
		if !ok {
			distribution[username] = allCoin
		} else {
			distribution[username] = value + allCoin
		}
	}
	return left
}

func main() {
	left := dispathCoin()
	for username, coin := range distribution {
		fmt.Printf("user:%s have %d coins \n", username, coin)

	}
	fmt.Printf("left coin:%d\n", left)
}
