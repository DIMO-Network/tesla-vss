package convert

import (
	"github.com/DIMO-Network/tesla-vss/pkg/unit"
	"github.com/teslamotors/fleet-telemetry/protos"
)

func ConvertLocationLocationValueToCurrentLocationLatitude(val *protos.LocationValue) (float64, error) {
	return val.Latitude, nil
}

func ConvertLocationLocationValueToCurrentLocationLongitude(val *protos.LocationValue) (float64, error) {
	return val.Longitude, nil
}

func ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsCharging(val protos.DetailedChargeStateValue) (float64, error) {
	switch val {
	case protos.DetailedChargeStateValue_DetailedChargeStateStarting, protos.DetailedChargeStateValue_DetailedChargeStateCharging:
		return 1, nil
	default:
		return 0, nil
	}
}

func ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrent(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

func ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

func ConvertOutsideTempStringToExteriorAirTemperature(val float64) (float64, error) {
	return val, nil
}

func ConvertEstBatteryRangeStringToPowertrainRange(val float64) (float64, error) {
	return unit.MilesToKilometers(val), nil
}

func ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimit(val float64) (float64, error) {
	return val, nil
}

func ConvertVehicleSpeedStringToSpeed(val float64) (float64, error) {
	return val, nil
}
