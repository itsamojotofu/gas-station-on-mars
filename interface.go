type GasStationService interface {
	RegisterUFO(licensePlate string, fuelType FuelType, tankCapacity float64) error
	GetFuelPrices() (map[FuelType]float64, error)
	GetRegisteredUFOs() ([]RegisteredUFO, error)
	RefuelUFO(licensePlate string, fuelType FuelType, refuelAmount float64) (float64, error)
	GetAllRefuelingHistories() ([]RefuelingHistory, error)
	GetRefuelingHistoryByLicensePlate(licensePlate string) ([]RefuelingHistory, error)
	GetSalesInfo() (map[FuelType]float64, float64, error)
}