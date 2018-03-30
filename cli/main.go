package main

import (
	"fmt"

	"github.com/rwirdemann/3skills.time/cli/presenter"
	"github.com/rwirdemann/3skills.time/rest"

	"github.com/rwirdemann/3skills.time/database"
	"github.com/rwirdemann/3skills.time/usecase"
)

func main() {
	presenter := presenter.NewCLIPresenter()
	consumer := rest.NewQueryConsumer()
	repository := database.NewMySQLRepository()
	usecase := usecase.NewGetProjects(consumer, presenter, repository)
	result := usecase.Run(nil)
	fmt.Println(result)
}
