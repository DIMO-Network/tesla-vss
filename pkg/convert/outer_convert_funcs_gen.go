package convert

import (
	"fmt"
	"strconv"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/DIMO-Network/tesla-vss/pkg/unit"
	"github.com/teslamotors/fleet-telemetry/protos"
)

func ProcessPayload(payload *protos.Payload, tokenID uint32, source string) ([]vss.Signal, []error) {
	var out []vss.Signal
	var outErr []error

	ts := payload.GetCreatedAt().AsTime()

	for _, d := range payload.GetData() {
		if d.GetKey() == protos.Field_Location {
			if v, ok := d.GetValue().Value.(*protos.Value_LocationValue); ok {
				val, err := ConvertLocationLocationValueToCurrentLocationLatitudeWrapper(v.LocationValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "currentLocationLatitude",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_Location {
			if v, ok := d.GetValue().Value.(*protos.Value_LocationValue); ok {
				val, err := ConvertLocationLocationValueToCurrentLocationLongitudeWrapper(v.LocationValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "currentLocationLongitude",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_DetailedChargeState {
			if v, ok := d.GetValue().Value.(*protos.Value_DetailedChargeStateValue); ok {
				val, err := ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsChargingWrapper(v.DetailedChargeStateValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainTractionBatteryChargingIsCharging",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_ACChargingPower {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPowerWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainTractionBatteryCurrentPower",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_DCChargingPower {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPowerWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainTractionBatteryCurrentPower",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_DCChargingEnergyIn {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainTractionBatteryChargingAddedEnergy",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_EnergyRemaining {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertEnergyRemainingStringToPowertrainTractionBatteryStateOfChargeCurrentEnergyWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainTractionBatteryStateOfChargeCurrentEnergy",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_Soc {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrentWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainTractionBatteryStateOfChargeCurrent",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_TpmsPressureFl {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressureWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "chassisAxleRow1WheelLeftTirePressure",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_TpmsPressureFr {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressureWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "chassisAxleRow1WheelRightTirePressure",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_TpmsPressureRl {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressureWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "chassisAxleRow2WheelLeftTirePressure",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_TpmsPressureRr {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressureWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "chassisAxleRow2WheelRightTirePressure",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_OutsideTemp {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertOutsideTempStringToExteriorAirTemperatureWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "exteriorAirTemperature",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_EstBatteryRange {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertEstBatteryRangeStringToPowertrainRangeWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainRange",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_ChargeLimitSoc {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimitWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "powertrainTractionBatteryChargingChargeLimit",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_VehicleSpeed {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertVehicleSpeedStringToSpeedWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "speed",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_DoorState {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDoorStateStringToCabinDoorRow1DriverSideIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow1DriverSideIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_DoorState {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDoorStateStringToCabinDoorRow1PassengerSideIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow1PassengerSideIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_DoorState {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDoorStateStringToCabinDoorRow2DriverSideIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow2DriverSideIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_DoorState {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDoorStateStringToCabinDoorRow2PassengerSideIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow2PassengerSideIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_FdWindow {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertFdWindowStringToCabinDoorRow1DriverSideWindowIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow1DriverSideWindowIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_FpWindow {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertFpWindowStringToCabinDoorRow1PassengerSideWindowIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow1PassengerSideWindowIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_RdWindow {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertRdWindowStringToCabinDoorRow2DriverSideWindowIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow2DriverSideWindowIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
		if d.GetKey() == protos.Field_RpWindow {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertRpWindowStringToCabinDoorRow2PassengerSideWindowIsOpenWrapper(v.StringValue)
				if err != nil {
					outErr = append(outErr, err)
				} else {
					sig := vss.Signal{
						TokenID:   tokenID,
						Name:      "cabinDoorRow2PassengerSideWindowIsOpen",
						Timestamp: ts,
						Source:    source,
					}
					sig.SetValue(val)
					out = append(out, sig)
				}
			}
		}
	}

	return out, outErr
}

func ConvertLocationLocationValueToCurrentLocationLatitudeWrapper(wrap *protos.LocationValue) (float64, error) {
	return ConvertLocationLocationValueToCurrentLocationLatitude(wrap)
}

func ConvertLocationLocationValueToCurrentLocationLongitudeWrapper(wrap *protos.LocationValue) (float64, error) {
	return ConvertLocationLocationValueToCurrentLocationLongitude(wrap)
}

func ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsChargingWrapper(wrap protos.DetailedChargeStateValue) (float64, error) {
	return ConvertDetailedChargeStateDetailedChargeStateValueToPowertrainTractionBatteryChargingIsCharging(wrap)
}

func ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPowerWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.KilowattsToWatts(fp)

	return ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPower(fp)
}

func ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPowerWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.KilowattsToWatts(fp)

	return ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPower(fp)
}

func ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergy(fp)
}

func ConvertEnergyRemainingStringToPowertrainTractionBatteryStateOfChargeCurrentEnergyWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertEnergyRemainingStringToPowertrainTractionBatteryStateOfChargeCurrentEnergy(fp)
}

func ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrentWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrent(fp)
}

func ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressureWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.AtmospheresToKilopascals(fp)

	return ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressure(fp)
}

func ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressureWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.AtmospheresToKilopascals(fp)

	return ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressure(fp)
}

func ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressureWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.AtmospheresToKilopascals(fp)

	return ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressure(fp)
}

func ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressureWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.AtmospheresToKilopascals(fp)

	return ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressure(fp)
}

func ConvertOutsideTempStringToExteriorAirTemperatureWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertOutsideTempStringToExteriorAirTemperature(fp)
}

func ConvertEstBatteryRangeStringToPowertrainRangeWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertEstBatteryRangeStringToPowertrainRange(fp)
}

func ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimitWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimit(fp)
}

func ConvertVehicleSpeedStringToSpeedWrapper(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.MilesPerHourToKilometersPerHour(fp)

	return ConvertVehicleSpeedStringToSpeed(fp)
}

func ConvertDoorStateStringToCabinDoorRow1DriverSideIsOpenWrapper(wrap string) (float64, error) {
	return ConvertDoorStateStringToCabinDoorRow1DriverSideIsOpen(wrap)
}

func ConvertDoorStateStringToCabinDoorRow1PassengerSideIsOpenWrapper(wrap string) (float64, error) {
	return ConvertDoorStateStringToCabinDoorRow1PassengerSideIsOpen(wrap)
}

func ConvertDoorStateStringToCabinDoorRow2DriverSideIsOpenWrapper(wrap string) (float64, error) {
	return ConvertDoorStateStringToCabinDoorRow2DriverSideIsOpen(wrap)
}

func ConvertDoorStateStringToCabinDoorRow2PassengerSideIsOpenWrapper(wrap string) (float64, error) {
	return ConvertDoorStateStringToCabinDoorRow2PassengerSideIsOpen(wrap)
}

func ConvertFdWindowStringToCabinDoorRow1DriverSideWindowIsOpenWrapper(wrap string) (float64, error) {
	return ConvertFdWindowStringToCabinDoorRow1DriverSideWindowIsOpen(wrap)
}

func ConvertFpWindowStringToCabinDoorRow1PassengerSideWindowIsOpenWrapper(wrap string) (float64, error) {
	return ConvertFpWindowStringToCabinDoorRow1PassengerSideWindowIsOpen(wrap)
}

func ConvertRdWindowStringToCabinDoorRow2DriverSideWindowIsOpenWrapper(wrap string) (float64, error) {
	return ConvertRdWindowStringToCabinDoorRow2DriverSideWindowIsOpen(wrap)
}

func ConvertRpWindowStringToCabinDoorRow2PassengerSideWindowIsOpenWrapper(wrap string) (float64, error) {
	return ConvertRpWindowStringToCabinDoorRow2PassengerSideWindowIsOpen(wrap)
}
