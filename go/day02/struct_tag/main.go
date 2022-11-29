package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string  `json:"username"`
	Sex      string  `json:"sex"`
	Score    float32 `json:"score"`
}

func main() {
	user := &User{
		UserName: "user01",
		Sex:      "男",
		Score:    100,
	}

	data, _ := json.Marshal(user)
	fmt.Printf("json str: %s \n", data)
}
