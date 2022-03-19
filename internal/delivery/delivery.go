package delivery

import (
	"OverflowBackend/internal/usecase"
)

type Delivery struct {
	uc usecase.UseCaseInterface
}

func (d *Delivery) Init(uc usecase.UseCaseInterface) {
	d.uc = uc
}