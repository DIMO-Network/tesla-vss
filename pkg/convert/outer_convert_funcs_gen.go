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
				val, err := ConvertLocationToCurrentLocationLatitudeOuter(v.LocationValue)
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
				val, err := ConvertLocationToCurrentLocationLongitudeOuter(v.LocationValue)
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
				val, err := ConvertACChargingPowerToPowertrainTractionBatteryCurrentPowerOuter(v.StringValue)
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
				val, err := ConvertDCChargingPowerToPowertrainTractionBatteryCurrentPowerOuter(v.StringValue)
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
		if d.GetKey() == protos.Field_ACChargingEnergyIn {
			if v, ok := d.GetValue().Value.(*protos.Value_StringValue); ok {
				val, err := ConvertACChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyOuter(v.StringValue)
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
				val, err := ConvertDCChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyOuter(v.StringValue)
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
				val, err := ConvertSocToPowertrainTractionBatteryStateOfChargeCurrentOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureFlToChassisAxleRow1WheelLeftTirePressureOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureFrToChassisAxleRow1WheelRightTirePressureOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureRlToChassisAxleRow2WheelLeftTirePressureOuter(v.StringValue)
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
				val, err := ConvertTpmsPressureRrToChassisAxleRow2WheelRightTirePressureOuter(v.StringValue)
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
				val, err := ConvertOutsideTempToExteriorAirTemperatureOuter(v.StringValue)
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
				val, err := ConvertEstBatteryRangeToPowertrainRangeOuter(v.StringValue)
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
				val, err := ConvertChargeLimitSocToPowertrainTractionBatteryChargingChargeLimitOuter(v.StringValue)
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
				val, err := ConvertVehicleSpeedToSpeedOuter(v.StringValue)
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

func ConvertLocationToCurrentLocationLatitudeOuter(wrap *protos.LocationValue) (float64, error) {
	return ConvertLocationToCurrentLocationLatitudeInner(wrap)
}

func ConvertLocationToCurrentLocationLongitudeOuter(wrap *protos.LocationValue) (float64, error) {
	return ConvertLocationToCurrentLocationLongitudeInner(wrap)
}

func ConvertACChargingPowerToPowertrainTractionBatteryCurrentPowerOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.KilowattsToWatts(fp)

	return ConvertACChargingPowerToPowertrainTractionBatteryCurrentPowerInner(fp)
}

func ConvertDCChargingPowerToPowertrainTractionBatteryCurrentPowerOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.KilowattsToWatts(fp)

	return ConvertDCChargingPowerToPowertrainTractionBatteryCurrentPowerInner(fp)
}

func ConvertACChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertACChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyInner(fp)
}

func ConvertDCChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertDCChargingEnergyInToPowertrainTractionBatteryChargingAddedEnergyInner(fp)
}

func ConvertSocToPowertrainTractionBatteryStateOfChargeCurrentOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertSocToPowertrainTractionBatteryStateOfChargeCurrentInner(fp)
}

func ConvertTpmsPressureFlToChassisAxleRow1WheelLeftTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureFlToChassisAxleRow1WheelLeftTirePressureInner(fp)
}

func ConvertTpmsPressureFrToChassisAxleRow1WheelRightTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureFrToChassisAxleRow1WheelRightTirePressureInner(fp)
}

func ConvertTpmsPressureRlToChassisAxleRow2WheelLeftTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureRlToChassisAxleRow2WheelLeftTirePressureInner(fp)
}

func ConvertTpmsPressureRrToChassisAxleRow2WheelRightTirePressureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.BarsToKilopascals(fp)

	return ConvertTpmsPressureRrToChassisAxleRow2WheelRightTirePressureInner(fp)
}

func ConvertOutsideTempToExteriorAirTemperatureOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertOutsideTempToExteriorAirTemperatureInner(fp)
}

func ConvertEstBatteryRangeToPowertrainRangeOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertEstBatteryRangeToPowertrainRangeInner(fp)
}

func ConvertChargeLimitSocToPowertrainTractionBatteryChargingChargeLimitOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	return ConvertChargeLimitSocToPowertrainTractionBatteryChargingChargeLimitInner(fp)
}

func ConvertVehicleSpeedToSpeedOuter(wrap string) (float64, error) {
	fp, err := strconv.ParseFloat(wrap, 64)
	if err != nil {
		var tmpOut float64
		return tmpOut, fmt.Errorf("failed to parse float: %w", err)
	}

	fp = unit.MilesPerHourToKilometersPerHour(fp)

	return ConvertVehicleSpeedToSpeedInner(fp)
}
