package main

import (
	"fmt"
	"os"
	configs "github.com/fwchen/jellyfish/config"
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/domain/visitor/repository/impl"
	visitorService "github.com/fwchen/jellyfish/domain/visitor/service"
)

func createUser(username, password string) {
	config, err := configs.LoadConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}

	datasource, err := database.GetDatabase(config.DataSource)
	if err != nil {
		panic(err)
	}
	service := visitorService.NewApplicationService(impl.NewVisitorRepository(datasource), &config.Application)
	err = service.SignUp(username, password)
	if err != nil {
		panic(err)
	}

	fmt.Print("Create user successful:\n")
	fmt.Print("username: ", username, "\npassword: ", password, "\n")
}

func main() {
	var command string = os.Args[1]
	
	if command == "create-user" {
		var username = os.Args[2]
		var password = os.Args[3]
		createUser(username, password)
	} else {
		fmt.Println("input invalid")
	}
}