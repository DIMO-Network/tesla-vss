package convert

import (
	"fmt"
	"testing"
	"time"

	"github.com/teslamotors/fleet-telemetry/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvert(t *testing.T) {
	ts, err := time.Parse(time.RFC3339, "2025-01-01T09:00:00Z")
	if err != nil {
		t.Fatal("Failed to create test timestamp.")
	}

	vin := "5YJYGDEF2LFR00942"

	pl := protos.Payload{
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

	fmt.Println(pl)
}
