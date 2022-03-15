package parkingSystem

const (
	bigCar    = 1
	mediumCar = 2
	smallCar  = 3
)

type ParkingSystem struct {
	smallSpaces  int
	mediumSpaces int
	bigSpaces    int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{smallSpaces: small, mediumSpaces: medium, bigSpaces: big}
}

func (this *ParkingSystem) AddCar(carType int) bool {

	switch carType {
	case bigCar:
		if this.bigSpaces > 0 {
			this.bigSpaces--
			return true
		}
	case mediumCar:
		if this.mediumSpaces > 0 {
			this.mediumSpaces--
			return true
		}
	case smallCar:
		if this.smallSpaces > 0 {
			this.smallSpaces--
			return true
		}
	default:
		return false
	}
	return false
}
