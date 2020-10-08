package main

import "fmt"

const usixeenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal float64
	brake_pedal float64
	steering_wheel float64
	top_speed_kmh float64
}

func (c car) kmh() float64 {
	return float64 (c.gas_pedal) * (c.top_speed_kmh/ usixeenbitmax)
}
func (c car) mph() float64 {
	return float64 (c.gas_pedal) * (c.top_speed_kmh / usixeenbitmax / kmh_multiple)
}

func main () {
	a_car := car{
		gas_pedal:      65000,
		brake_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  255.0,
	}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}