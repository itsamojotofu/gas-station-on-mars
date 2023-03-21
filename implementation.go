type gasStationService struct {
	fuelPrices map[FuelType]float64
	registeredUFOs map[string]RegisteredUFO
	refuelingHistories []RefuelingHistory
}

func NewGasStationService() GasStationService {
	fuelPrices := map[FuelType]float64{
			Diesel: 1.2,
			Gasoline87: 1.5,
			Gasoline90: 1.7,
			Gasoline94: 2.0,
	}
	return &gasStationService{
			fuelPrices: fuelPrices,
			registeredUFOs: make(map[string]RegisteredUFO),
			refuelingHistories: make([]RefuelingHistory, 0),
	}
}

func (s *gasStationService) RegisterUFO(licensePlate string, fuelType FuelType, tankCapacity float64) error {
	if _, ok := s.fuelPrices[fuelType]; !ok {
			return fmt.Errorf("invalid fuel type: %v", fuelType)
	}
	s.registeredUFOs[licensePlate] = RegisteredUFO{
			FuelType: fuelType,
			TankCapacity: tankCapacity,
	}
	return nil
}

func (s *gasStationService) GetFuelPrices() (map[FuelType]float64, error) {
	return s.fuelPrices, nil
}

func (s *gasStationService) GetRegisteredUFOs() ([]RegisteredUFO, error) {
	ufos := make([]RegisteredUFO, 0, len(s.registeredUFOs))
	for _, ufo := range s.registeredUFOs {
			ufos = append(ufos, ufo)
	}
	return ufos, nil
}

func (s *gasStationService) RefuelUFO(licensePlate string, fuelType FuelType, refuelAmount float64) (float64, error) {
	ufo, ok := s.registeredUFOs[licensePlate]
	if !ok {
			return 0, fmt.Errorf("unknown license plate: %v", licensePlate)
	}
	if fuelType != ufo.FuelType {
			return 0, fmt.Errorf("incorrect fuel type for UFO: %v", fuelType)
	}
	if refuelAmount <= 0 {
			return 0, fmt.Errorf("refuel amount must be greater than 0")
	}
	if refuelAmount > ufo.TankCapacity {
