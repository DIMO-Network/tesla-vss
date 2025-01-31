package convert

import (
	"github.com/DIMO-Network/tesla-vss/pkg/unit"
	"github.com/teslamotors/fleet-telemetry/protos"
)

func ConvertLocationLocationValueToCurrentLocationLatitudeInner(val *protos.LocationValue) (float64, error) {
	return val.Latitude, nil
}

func ConvertLocationLocationValueToCurrentLocationLongitudeInner(val *protos.LocationValue) (float64, error) {
	return val.Longitude, nil
}

func ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPowerInner(val float64) (float64, error) {
	return val, nil
}

func ConvertACChargingPowerStringToPowertrainTractionBatteryChargingIsChargingInner(val float64) (float64, error) {
	if val > 0 {
		return 1, nil
	}
	return 0, nil
}

func ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPowerInner(val float64) (float64, error) {
	return val, nil
}

func ConvertDCChargingPowerStringToPowertrainTractionBatteryChargingIsChargingInner(val float64) (float64, error) {
	if val > 0 {
		return 1, nil
	}
	return 0, nil
}

func ConvertACChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyInner(val float64) (float64, error) {
	return val, nil
}

func ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyInner(val float64) (float64, error) {
	return val, nil
}

func ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrentInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertOutsideTempStringToExteriorAirTemperatureInner(val float64) (float64, error) {
	return val, nil
}

func ConvertEstBatteryRangeStringToPowertrainRangeInner(val float64) (float64, error) {
	return unit.MilesToKilometers(val), nil
}

func ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimitInner(val float64) (float64, error) {
	return val, nil
}

func ConvertVehicleSpeedStringToSpeedInner(val float64) (float64, error) {
	return val, nil
}
