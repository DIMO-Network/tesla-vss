package convert

import (
	"testing"
	"time"

	"github.com/DIMO-Network/model-garage/pkg/vss"
	"github.com/stretchr/testify/assert"
	"github.com/teslamotors/fleet-telemetry/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvert(t *testing.T) {
	teslaConnection := "0xc4035Fecb1cc906130423EF05f9C20977F643722" // This is the real value in dev and prod.

	ts, err := time.Parse(time.RFC3339, "2025-01-01T09:00:00Z")
	if err != nil {
		t.Fatal("Failed to create test timestamp.")
	}

	vin := "5YJYGDEF2LFR00942"

	pl := &protos.Payload{
		Data: []*protos.Datum{
			{Key: protos.Field_Location, Value: &protos.Value{Value: &protos.Value_LocationValue{LocationValue: &protos.LocationValue{Latitude: 30.267222, Longitude: -97.743056}}}},
			{Key: protos.Field_ACChargingPower, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "5.700000084936619"}}},
			{Key: protos.Field_ACChargingEnergyIn, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.380388924359452"}}},
			{Key: protos.Field_Soc, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "18.155283129013426"}}},
			{Key: protos.Field_TpmsPressureFl, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.9250000435858965"}}},
			{Key: protos.Field_TpmsPressureFr, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.425000036135316"}}},
			{Key: protos.Field_TpmsPressureRl, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.8000000417232513"}}},
			{Key: protos.Field_TpmsPressureRr, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.8000000417232513"}}},
			{Key: protos.Field_OutsideTemp, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "2.5"}}},
			{Key: protos.Field_EstBatteryRange, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "19.80471193262205"}}},
			{Key: protos.Field_ChargeLimitSoc, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "80"}}},
			{Key: protos.Field_VehicleSpeed, Value: &protos.Value{Value: &protos.Value_StringValue{StringValue: "21"}}},
		},
		CreatedAt: timestamppb.New(ts),
		Vin:       vin,
	}

	signals, errors := ProcessPayload(pl, 7, teslaConnection)
	if len(errors) != 0 {
		t.Fatalf("Unexpected errors from conversion: %v", err)
	}

	expectedSignals := []vss.Signal{
		{TokenID: 7, Timestamp: ts, Name: "currentLocationLatitude", ValueNumber: 30.267222, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "currentLocationLongitude", ValueNumber: -97.743056, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "powertrainTractionBatteryCurrentPower", ValueNumber: 5700.000084936619, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "powertrainTractionBatteryChargingIsCharging", ValueNumber: 1, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "powertrainTractionBatteryChargingAddedEnergy", ValueNumber: 2.380388924359452, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "powertrainTractionBatteryStateOfChargeCurrent", ValueNumber: 18.155283129013426, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "chassisAxleRow1WheelLeftTirePressure", ValueNumber: 292.50000435858965, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "chassisAxleRow1WheelRightTirePressure", ValueNumber: 242.5000036135316, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "chassisAxleRow2WheelLeftTirePressure", ValueNumber: 280.00000417232513, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "chassisAxleRow2WheelRightTirePressure", ValueNumber: 280.00000417232513, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "exteriorAirTemperature", ValueNumber: 2.5, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "powertrainRange", ValueNumber: 31.872594320493704, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "powertrainTractionBatteryChargingChargeLimit", ValueNumber: 80, Source: teslaConnection},
		{TokenID: 7, Timestamp: ts, Name: "speed", ValueNumber: 33.796224, Source: teslaConnection},
	}

	assert.ElementsMatch(t, expectedSignals, signals)
}
