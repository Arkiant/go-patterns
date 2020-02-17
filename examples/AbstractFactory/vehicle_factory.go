package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

const (
	CAR = iota
	MOTORBIKE
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CAR:
		return new(CarFactory), nil
	case MOTORBIKE:
		return new(MotorBikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}
