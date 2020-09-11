package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	SPORT = iota
	CRUISE
)

type MotorBikeFactory struct{}

func (m *MotorBikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SPORT:
		return new(SportMotorbike), nil
	case CRUISE:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
