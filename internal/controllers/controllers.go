package controllers

import (
	"finworker/internal/repository"
)

type Controller struct {
	repo repository.Repositories
}

func NewController(repo repository.Repositories) *Controller {
	return &Controller{
		repo: repo,
	}
}
