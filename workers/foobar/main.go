package foobar

import (
	"context"

	"github.com/inn4sc/vcg-go-common/routines"
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
