package delivery

import (
	"OverflowBackend/internal/config"
	"OverflowBackend/internal/usecase"
)

type Delivery struct {
	uc usecase.UseCaseInterface
	config *config.Config
}

func (d *Delivery) Init(uc usecase.UseCaseInterface, config *config.Config) {
	d.uc = uc
	d.config = config
}