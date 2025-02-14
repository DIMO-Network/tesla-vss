package convert

import (
	"github.com/DIMO-Network/tesla-vss/pkg/unit"
	"github.com/teslamotors/fleet-telemetry/protos"
)

// ConvertLocationLocationValueToCurrentLocationLatitude converts a telemetry datum with key Location to the VSS signal CurrentLocationLatitude.
func ConvertLocationLocationValueToCurrentLocationLatitude(val *protos.LocationValue) (float64, error) {
	return val.Latitude, nil
}

// ConvertLocationLocationValueToCurrentLocationLongitude converts a telemetry datum with key Location to the VSS signal CurrentLocationLongitude.
func ConvertLocationLocationValueToCurrentLocationLongitude(val *protos.LocationValue) (float64, error) {
	return val.Longitude, nil
}

// ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsCharging converts a telemetry datum with key DetailedChargeState to the VSS signal PowertrainTractionBatteryChargingIsCharging.
func ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsCharging(val protos.DetailedChargeStateValue) (float64, error) {
	switch val {
	case protos.DetailedChargeStateValue_DetailedChargeStateStarting, protos.DetailedChargeStateValue_DetailedChargeStateCharging:
		return 1, nil
	default:
		return 0, nil
	}
}

// ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrent converts a telemetry datum with key Soc to the VSS signal PowertrainTractionBatteryStateOfChargeCurrent.
func ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrent(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressure converts a telemetry datum with key TpmsPressureFl to the VSS signal ChassisAxleRow1WheelLeftTirePressure.
func ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressure converts a telemetry datum with key TpmsPressureFr to the VSS signal ChassisAxleRow1WheelRightTirePressure.
func ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressure converts a telemetry datum with key TpmsPressureRl to the VSS signal ChassisAxleRow2WheelLeftTirePressure.
func ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressure converts a telemetry datum with key TpmsPressureRr to the VSS signal ChassisAxleRow2WheelRightTirePressure.
func ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressure(val float64) (float64, error) {
	return val, nil
}

// ConvertOutsideTempStringToExteriorAirTemperature converts a telemetry datum with key OutsideTemp to the VSS signal ExteriorAirTemperature.
func ConvertOutsideTempStringToExteriorAirTemperature(val float64) (float64, error) {
	return val, nil
}

// ConvertEstBatteryRangeStringToPowertrainRange converts a telemetry datum with key EstBatteryRange to the VSS signal PowertrainRange.
func ConvertEstBatteryRangeStringToPowertrainRange(val float64) (float64, error) {
	// We have to do this because we've violated the VSS spec and used km here instead of m.
	return unit.MilesToKilometers(val), nil
}

// ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimit converts a telemetry datum with key ChargeLimitSoc to the VSS signal PowertrainTractionBatteryChargingChargeLimit.
func ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimit(val float64) (float64, error) {
	return val, nil
}

// ConvertVehicleSpeedStringToSpeed converts a telemetry datum with key VehicleSpeed to the VSS signal Speed.
func ConvertVehicleSpeedStringToSpeed(val float64) (float64, error) {
	return val, nil
}
