package parkingSystem

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	got := Constructor(9, 8, 7)
	want := ParkingSystem{smallSpaces: 7, mediumSpaces: 8, bigSpaces: 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Constructor failed %+v vs %+v", got, want)
	}
}

func TestAddCar(t *testing.T) {
	parkingSystem := Constructor(8, 8, 8)
	for i := 0; i < 8; i++ {
		if got := parkingSystem.AddCar(1); !got {
			t.Errorf("Failed removing big car: %+v", parkingSystem)
		}
		if got := parkingSystem.AddCar(2); !got {
			t.Errorf("Failed removing medium car: %+v", parkingSystem)
		}
		if got := parkingSystem.AddCar(3); !got {
			t.Errorf("Failed removing small car: %+v", parkingSystem)
		}
	}
}
