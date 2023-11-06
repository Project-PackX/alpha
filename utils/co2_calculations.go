package utils

import "math"

const (
	earthRadius = 6371 // Earth's radius in kilometers
)

func CalculateEmissionDifference(distance float64) float64 {
	electricCarEnergyConsumption := 0.2 // kWh per mile
	electricCarEmissions := 0.3         // kg CO2 per kWh

	petrolCarFuelConsumption := 30 // miles per gallon (mpg)
	petrolCarEmissions := 8.8      // kg CO2 per gallon of gasoline

	electricCarEnergy := electricCarEnergyConsumption * distance
	electricCarTotalEmissions := electricCarEnergy * electricCarEmissions

	petrolCarFuel := distance / float64(petrolCarFuelConsumption)
	petrolCarTotalEmissions := petrolCarFuel * petrolCarEmissions

	emissionSavings := petrolCarTotalEmissions - electricCarTotalEmissions

	return emissionSavings
}

func CalculateDistance(latitude1, longitude1, latitude2, longitude2 float64) float64 {
	// Convert latitude and longitude from degrees to radians
	latitude1 = latitude1 * math.Pi / 180
	longitude1 = longitude1 * math.Pi / 180
	latitude2 = latitude2 * math.Pi / 180
	longitude2 = longitude2 * math.Pi / 180

	// Haversine formula
	dlat := latitude2 - latitude1
	dlon := longitude2 - longitude1
	squaredChordLength := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(latitude1)*math.Cos(latitude2)*math.Sin(dlon/2)*math.Sin(dlon/2)

	centralAngle := 2 * math.Atan2(math.Sqrt(squaredChordLength), math.Sqrt(1-squaredChordLength))

	// Calculate the distance
	distance := earthRadius * centralAngle
	return distance
}
