package main

type ParkingSystem struct {
	slots [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		slots: [4]int{0, big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType] <= 0 {
		return false
	}
	this.slots[carType]--
	return true
}
