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
				val, err := ConvertLocationLocationValueToCurrentLocationLatitudeOuter(v.LocationValue)
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
				val, err := ConvertLocationLocationValueToCurrentLocationLongitudeOuter(v.LocationValue)
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
		if d.GetKey() == protos.Field_ACChargingPower {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPowerOuter(v.StringValue)
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
		if d.GetKey() == protos.Field_ACChargingPower {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertACChargingPowerStringToPowertrainTractionBatteryChargingIsChargingOuter(v.StringValue)
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
		if d.GetKey() == protos.Field_DCChargingPower {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPowerOuter(v.StringValue)
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
				val, err := ConvertDCChargingPowerStringToPowertrainTractionBatteryChargingIsChargingOuter(v.StringValue)
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
		if d.GetKey() == protos.Field_ACChargingEnergyIn {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertACChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyOuter(v.StringValue)
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
		if d.GetKey() == protos.Field_DCChargingEnergyIn {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyOuter(v.StringValue)
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
		if d.GetKey() == protos.Field_Soc {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrentOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressureOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressureOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressureOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressureOuter(v.StringValue)
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
				val, err := ConvertOutsideTempStringToExteriorAirTemperatureOuter(v.StringValue)
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
				val, err := ConvertEstBatteryRangeStringToPowertrainRangeOuter(v.StringValue)
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
				val, err := ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimitOuter(v.StringValue)
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
				val, err := ConvertVehicleSpeedStringToSpeedOuter(v.StringValue)
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
	}

	return out, outErr
}

func ConvertLocationLocationValueToCurrentLocationLatitudeOuter(wrap *protos.LocationValue) (float64, error) {
	return ConvertLocationLocationValueToCurrentLocationLatitudeInner(wrap)
}

func ConvertLocationLocationValueToCurrentLocationLongitudeOuter(wrap *protos.LocationValue) (float64, error) {
	return ConvertLocationLocationValueToCurrentLocationLongitudeInner(wrap)
}

func ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPowerOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.KilowattsToWatts(fp)

	return ConvertACChargingPowerStringToPowertrainTractionBatteryCurrentPowerInner(fp)
}

func ConvertACChargingPowerStringToPowertrainTractionBatteryChargingIsChargingOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertACChargingPowerStringToPowertrainTractionBatteryChargingIsChargingInner(fp)
}

func ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPowerOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.KilowattsToWatts(fp)

	return ConvertDCChargingPowerStringToPowertrainTractionBatteryCurrentPowerInner(fp)
}

func ConvertDCChargingPowerStringToPowertrainTractionBatteryChargingIsChargingOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertDCChargingPowerStringToPowertrainTractionBatteryChargingIsChargingInner(fp)
}

func ConvertACChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertACChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyInner(fp)
}

func ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertDCChargingEnergyInStringToPowertrainTractionBatteryChargingAddedEnergyInner(fp)
}

func ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrentOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertSocStringToPowertrainTractionBatteryStateOfChargeCurrentInner(fp)
}

func ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureFlStringToChassisAxleRow1WheelLeftTirePressureInner(fp)
}

func ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureFrStringToChassisAxleRow1WheelRightTirePressureInner(fp)
}

func ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureRlStringToChassisAxleRow2WheelLeftTirePressureInner(fp)
}

func ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureRrStringToChassisAxleRow2WheelRightTirePressureInner(fp)
}

func ConvertOutsideTempStringToExteriorAirTemperatureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertOutsideTempStringToExteriorAirTemperatureInner(fp)
}

func ConvertEstBatteryRangeStringToPowertrainRangeOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertEstBatteryRangeStringToPowertrainRangeInner(fp)
}

func ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimitOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertChargeLimitSocStringToPowertrainTractionBatteryChargingChargeLimitInner(fp)
}

func ConvertVehicleSpeedStringToSpeedOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.MilesPerHourToKilometersPerHour(fp)

	return ConvertVehicleSpeedStringToSpeedInner(fp)
}
