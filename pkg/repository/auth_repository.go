package repository

import "gitlab.chakravuyha.com/zerohero/cero_hero_go_backend/pkg/mongo"

type authRepository struct {
	database   mongo.Database
	collection string
}
