package abstract_factory

import (
	"errors"
	"fmt"
)

const (
	LUXURY = iota
	FAMILY
)

type CarFactory struct{}

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LUXURY:
		return new(LuxuryCar), nil
	case FAMILY:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))

	}
}
