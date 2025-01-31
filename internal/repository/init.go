package repository

import "github.com/AwesomeXjs/iq-progress/pkg/dbClient"

type IRepository interface {
}

type Repository struct {
	db dbClient.Client
}

func New(db dbClient.Client) IRepository {
	return &Repository{}
}
