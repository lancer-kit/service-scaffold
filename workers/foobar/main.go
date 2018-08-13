package foobar

import (
	"context"

	"gitlab.inn4science.com/gophers/service-kit/routines"
)

type Service struct {
	ctx context.Context
}

func (s *Service) Init(ctx context.Context) routines.Worker {
	return &Service{
		ctx: ctx,
	}
}

func (s *Service) Run() {

}

func (s *Service) RestartOnFail() bool {
	return false
}
