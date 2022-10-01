package main

import "github.com/youichiro/go-slack-my-unipos/pkg/router"

func main() {
	r := router.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
