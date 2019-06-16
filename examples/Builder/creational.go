package creational

/*
BuildProcess implements all available builder methods
*/
type BuildProcess interface {
    SetWheels() BuildProcess
    SetSeats() BuildProcess
    SetStructure() BuildProcess
    GetVehicle() VehicleProduct
}

/*
ManufacturingDirector orchestra the creation process
*/
type ManufacturingDirector struct {
    builder BuildProcess
}

/*
SetBuilder configures the construction mechanism
*/
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
    f.builder = b
}

/*
Construct implements the construction process
*/
func (f *ManufacturingDirector) Construct() {
    f.builder.SetSeats().SetStructure().SetWheels()
}

/*
VehicleProduct general type
*/
type VehicleProduct struct {
    Wheels    int
    Seats     int
    Structure string
}