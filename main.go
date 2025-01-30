package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
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

//go:embed outer.tmpl
var tmpl string

func loadRules(path string) ([]Rule, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	fb, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var rules []Rule

	err = yaml.Unmarshal(fb, &rules)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rules YAML: %w", err)
	}

	return rules, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing required argument: path to rules file")
	}

	rulesFilePath := os.Args[1]

	flag.String("output", "", "Path ")

	signalInfos, err := schema.LoadSignalsCSV(strings.NewReader(schema.VssRel42DIMO()))
	if err != nil {
		log.Fatalf("Failed to load VSS schema: %v", err)
	}

	signalInfoBySignal := make(map[string]*schema.SignalInfo, len(signalInfos))
	for _, s := range signalInfos {
		signalInfoBySignal[s.Name] = s
	}

	rules, err := loadRules(rulesFilePath)
	if err != nil {
		log.Fatalf("Failed to load rules: %v", err)
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

		parseFloat := r.TeslaType == "string" && slices.Contains(r.Automations, ParseFloatFlag)

		var convertUnit string

		teslaType, ok := teslaTypeToAttributes[r.TeslaType]
		if !ok {
			panic("unrecognized Tesla type " + r.TeslaType)
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
			WrapperName:      teslaType.TeslaWrapperType,
			WrapperFieldName: teslaType.TeslaWrapperFieldName,
			GoVSSSignalName:  signalInfo.GOName,
			OuterInputType:   teslaType.ValueType,
			JSONName:         signalInfo.JSONName,
			OutputType:       signalInfo.GOType(),
			ParseFloat:       parseFloat,
			UnitFunc:         convertUnit,
		})
	}

	t, err := template.New("process").Parse(tmpl)
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

type TeslaTypeDescription struct {
	TeslaWrapperType      string
	TeslaWrapperFieldName string
	ValueType             string
}

var teslaTypeToAttributes = map[string]TeslaTypeDescription{
	"string": {
		TeslaWrapperType:      "Value_StringValue",
		TeslaWrapperFieldName: "StringValue",
		ValueType:             "string",
	},
	"LocationValue": {
		TeslaWrapperType:      "LocationValue",
		TeslaWrapperFieldName: "LocationValue",
		ValueType:             "*proto.LocationValue",
	},
}
