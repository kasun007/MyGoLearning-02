package data

import "fmt"

// ///////////////////////////////// Types //////////////
type integer = int
type float = float64
type json = map[string]string

type distance float64   // distance in meters
type distanceKm float64 // distance in kilometers

type Location string

func (location Location) DistanceTo(destination Location) distance {
	fmt.Print("DistanceTo")

	return 10
}
func locationTest() {

	nyc := Location("3.34")

	nyc.DistanceTo(Location("3.34 ,67.6"))
	print(nyc)

}

func (miles distance) toKm() distanceKm { // miles to km
	return distanceKm(miles * 1.60934)
}

func (miles distanceKm) toMiles() distance { // km to miles
	return distance(miles / 1.60934)
}

func locationtest() {
	var d distance = 1.2
	var d2 distanceKm = 1.2
	d.toKm()
	d2.toMiles()
}

////////////////////////Types Struts ////////////////////////
