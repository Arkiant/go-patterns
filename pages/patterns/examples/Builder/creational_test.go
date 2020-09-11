package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()

	cases := []struct {
		name string
		wheels int
		structure string
		seats int
		vehicle *VehicleProduct
	}{
		{"Testing car", 4, "Car", 5, &car},
		{"Testing bus", 8, "Bus", 32, &bus},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T){
			
			if c.vehicle.Wheels != c.wheels {
				t.Errorf("Wheels on a %s must be %d and they were %d\n", c.structure, c.wheels, c.vehicle.Wheels)
			}
	
			if c.vehicle.Structure != c.structure {
				t.Errorf("Structure on a %s must be %s and they were %s\n", c.structure, c.structure, c.vehicle.Structure)
			}
	
			if c.vehicle.Seats != c.seats {
				t.Errorf("Seats on a %s must be %d and they were %d\n", c.structure, c.seats, c.vehicle.Seats)
			}
			
		})
		
	}
}