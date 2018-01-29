package persistence

import "thmlogwork.com/go-time-zone-app/domain"

type Dao struct {
	TimezoneRepository domain.TimezoneRepository
}

func CreateRepositories() *Dao {
	repo := NewTimezoneDao()
	dao := &Dao{TimezoneRepository: repo}
	return dao
}
