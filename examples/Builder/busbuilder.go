package creational

type BusBuilder struct {
    v VehicleProduct
}

func (c *BusBuilder) SetWheels() BuildProcess {
    c.v.Wheels = 8
    return c
}

func (c *BusBuilder) SetSeats() BuildProcess {
    c.v.Seats = 32
    return c
}

func (c *BusBuilder) SetStructure() BuildProcess {
    c.v.Structure = "Bus"
    return c
}

func (c *BusBuilder) GetVehicle() VehicleProduct {
	return c.v
}