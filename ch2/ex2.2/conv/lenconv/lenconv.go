package lenconv

import "fmt"

type Foot float64
type Meter float64

func (ft Foot) String() string {
	return fmt.Sprintf("%.3f Ft. ", ft)
}
func (m Meter) String() string {
	return fmt.Sprintf("%.3f Meter", m)
}

// FtToMet делим длину на 3.281
func FtToMet(ft Foot) Meter {
	return Meter(ft / 3.281)
}

// MetToFt умножаем длину на 3.281
func MetToFt(m Meter) Foot {
	return Foot(m * 3.281)
}
