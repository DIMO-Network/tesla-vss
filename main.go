package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"io"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/teslamotors/fleet-telemetry/protos"
	"gopkg.in/yaml.v3"
)

type Rule struct {
	TeslaField string `yaml:"teslaField"`
	// TeslaType is the protobuf type of the value field on the Datum. If not specified, it is
	// assumed to be string. This is the dominant case.
	TeslaType string `yaml:"teslaType"`
	// VSSSignal is the full path to the VSS signal that will be produced from this Datum. For
	// example, one might write "Vehicle.Cabin.Door.Row1.Left.IsLocked".
	VSSSignal string `yaml:"vssSignal"`
	TeslaUnit string `yaml:"teslaUnit"`

	Automations []string `yaml:"automations"`
}

const (
	ParseFloatFlag  = "PARSE_FLOAT"
	ConvertUnitFlag = "CONVERT_UNIT"
)

//go:embed process.tmpl
var tmpl string

func main() {
	vss := schema.VssRel42DIMO()

	signalInfos, err := schema.LoadSignalsCSV(strings.NewReader(vss))
	if err != nil {
		panic(err)
	}

	signalInfoBySignal := make(map[string]*schema.SignalInfo, len(signalInfos))
	for _, s := range signalInfos {
		signalInfoBySignal[s.Name] = s
	}

	f, err := os.Open("rules.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fb, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var rules []Rule

	err = yaml.Unmarshal(fb, &rules)
	if err != nil {
		panic(err)
	}

	var tmplInput TemplateInput

	for _, r := range rules {
		signalInfo, ok := signalInfoBySignal[r.VSSSignal]
		if !ok {
			panic("don't recognize the VSS signal " + r.VSSSignal)
		}

		_, ok = protos.Field_value[r.TeslaField]
		if !ok {
			panic("don't recognize the Tesla field " + r.TeslaField)
		}

		parseFloat := false

		var convertUnit string

		var wrapperName, wrapperField, inputType string
		if r.TeslaType == "string" {
			wrapperName = "Value_StringValue"
			wrapperField = "StringValue"
			inputType = "string"
			// Just plain parse okay too
			if slices.Contains(r.Automations, ParseFloatFlag) {
				parseFloat = true
			}
		} else if r.TeslaType == "LocationValue" {
			wrapperName = "Value_LocationValue"
			wrapperField = "LocationValue"
			inputType = "*proto.LocationValue"
		}

		if slices.Contains(r.Automations, ConvertUnitFlag) && r.TeslaUnit != signalInfo.Unit {
			m, ok := conversions[r.TeslaUnit]
			if !ok {
				panic("no conversion from unit " + r.TeslaUnit)
			}

			n, ok := m[signalInfo.Unit]
			if !ok {
				panic("no conversion from unit " + r.TeslaUnit + " to " + signalInfo.Unit)
			}

			convertUnit = n
		}

		tmplInput.Conversions = append(tmplInput.Conversions, Conversion{
			TeslaField:       r.TeslaField,
			WrapperName:      wrapperName,
			WrapperFieldName: wrapperField,
			GoVSSSignalName:  signalInfo.GOName,
			OuterInputType:   inputType,
			JSONName:         signalInfo.JSONName,
			OutputType:       signalInfo.GOType(),
			ParseFloat:       parseFloat,
			UnitFunc:         convertUnit,
		})
	}

	t, err := template.New("xdd").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, tmplInput)
	if err != nil {
		panic(err)
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(out)
}

type Field struct {
	Name  string
	Types []Type
}

type Type struct {
	GoType      string
	Conversions []Conversion
}

type Conversion struct {
	TeslaField       string
	WrapperName      string
	WrapperFieldName string
	GoVSSSignalName  string
	OuterInputType   string
	JSONName         string
	OutputType       string

	ParseFloat bool
	UnitFunc   string
}

type TemplateInput struct {
	Conversions []Conversion
}

var conversions = map[string]map[string]string{
	"kW": {
		"W": "KilowattsToWatts",
	},
	"bar": {
		"kPa": "BarsToKilopascals",
	},
	"mi": {
		"km": "MilesToKilometers",
	},
	"mph": {
		"km/h": "MilesPerHourToKilometersPerHour",
	},
}
