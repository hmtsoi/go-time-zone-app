package main

import (
	. "thmlogwork.com/go-time-zone-app/domain"
	"thmlogwork.com/go-time-zone-app/persistence"
	"thmlogwork.com/go-time-zone-app/rest"
)

func main() {

	repos := persistence.CreateRepositories()

	service := &TimezoneServiceImpl{Repository: repos.TimezoneRepository}

	controller := &rest.TimezoneController{TimezoneService: service}
	controller.Create()

}
