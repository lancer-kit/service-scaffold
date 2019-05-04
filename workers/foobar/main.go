package foobar

import (
	"context"

	"github.com/lancer-kit/armory/routines"
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
