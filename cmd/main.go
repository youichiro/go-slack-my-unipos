package main

import (
	"github.com/youichiro/go-slack-my-unipos/internal/repositories"
	"github.com/youichiro/go-slack-my-unipos/internal/router"
)

func main() {
	dbRepo := &repositories.PostgresRepository{}
	err := dbRepo.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer dbRepo.Close()

	r := router.SetupRouter()
	err = r.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
}
