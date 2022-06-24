package repository

import "github.com/kitabisa/golang-poc/internal/app/commons"

// Option anything any repo object needed
type Option struct {
	commons.Options
}

// Repository all repo object injected here
type Repository struct {
	// User IUserRepository
	Cache ICacheRepo
}
