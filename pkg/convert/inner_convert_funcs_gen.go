package convert

import (
	"github.com/teslamotors/fleet-telemetry/protos"
)

func ConvertLocationToCurrentLocationLatitudeInner(val *protos.LocationValue) (float64, error) {
	panic("not implemented")
}

func ConvertLocationToCurrentLocationLongitudeInner(val *protos.LocationValue) (float64, error) {
	panic("not implemented")
}

func ConvertACChargingPowerToPowertrainTractionBatteryCurrentPowerInner(val float64) (float64, error) {
	return val, nil
}

func ConvertDCChargingPowerToPowertrainTractionBatteryCurrentPowerInner(val float64) (float64, error) {
	return val, nil
}

func ConvertACChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyInner(val float64) (float64, error) {
	return val, nil
}

func ConvertDCChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyInner(val float64) (float64, error) {
	return val, nil
}

func ConvertSocToPowertrainTractionBatteryStateOfChargeCurrentInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureFlToChassisAxleRow1WheelLeftTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureFrToChassisAxleRow1WheelRightTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureRlToChassisAxleRow2WheelLeftTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureRrToChassisAxleRow2WheelRightTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertOutsideTempToExteriorAirTemperatureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertEstBatteryRangeToPowertrainRangeInner(val float64) (float64, error) {
	return val, nil
}

func ConvertChargeLimitSocToPowertrainTractionBatteryChargingChargeLimitInner(val float64) (float64, error) {
	return val, nil
}

func ConvertVehicleSpeedToSpeedInner(val float64) (float64, error) {
	return val, nil
}
