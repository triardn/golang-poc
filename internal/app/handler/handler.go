package handler

import (
	"github.com/kitabisa/golang-poc/internal/app/commons"
	"github.com/kitabisa/golang-poc/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}
