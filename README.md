Streamed Tesla vehicle data comes in over websocket as a protobuf. The only reference is their [`vehicle_data.proto`](https://github.com/teslamotors/fleet-telemetry/blob/main/protos/vehicle_data.proto).

A rules file consists of elements that look like this

```yaml
- teslaField: DCChargingPower
  teslaType: string
  teslaUnit: kW
  vssSignal: Vehicle.Powertrain.TractionBattery.CurrentPower
  automations:
  - PARSE_FLOAT
  - CONVERT_UNIT
```
