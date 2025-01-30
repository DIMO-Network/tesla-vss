package codegen

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"html/template"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/teslamotors/fleet-telemetry/protos"
	"gopkg.in/yaml.v3"
)

const (
	ParseFloatFlag  = "PARSE_FLOAT"
	ConvertUnitFlag = "CONVERT_UNIT"
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

//go:embed inner.tmpl
var innerTmpl string

//go:embed outer.tmpl
var outerTmpl string

func Generate(packageName, rulesPath, outerOutputPath, innerOutputPath string) error {
	signalInfos, err := schema.LoadSignalsCSV(strings.NewReader(schema.VssRel42DIMO()))
	if err != nil {
		log.Fatalf("Failed to load VSS schema: %v", err)
	}

	signalInfoBySignal := make(map[string]*schema.SignalInfo, len(signalInfos))
	for _, s := range signalInfos {
		signalInfoBySignal[s.Name] = s
	}

	rules, err := loadRules(rulesPath)
	if err != nil {
		return fmt.Errorf("failed to load rules: %w", err)
	}

	tmplInput := &TemplateInput{
		Package: packageName,
	}

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

		innerInputType := teslaType.ValueType
		if slices.Contains(r.Automations, ParseFloatFlag) {
			innerInputType = "float64"
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
			InnerInputType:   innerInputType,
		})
	}

	err = writeOuter(tmplInput, outerOutputPath)
	if err != nil {
		return err
	}

	err = writeInner(tmplInput, innerOutputPath)
	if err != nil {
		return err
	}

	return nil
}

func writeOuter(tmplInput *TemplateInput, outerPath string) error {
	t, err := template.New("outer").Parse(outerTmpl)
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

	f, err := os.Create(outerPath)
	if err != nil {
		return fmt.Errorf("error opening outer output file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(out)
	if err != nil {
		return fmt.Errorf("error writing outer output file: %w", err)
	}

	return nil
}

func writeInner(tmplInput *TemplateInput, innerPath string) error {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, innerPath, nil, parser.ParseComments)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		for i, decl := range astFile.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				var buf bytes.Buffer
				err := format.Node(&buf, fset, &printer.CommentedNode{
					Node:     fn.Body,
					Comments: astFile.Comments,
				})
				if err != nil {
					panic(err)
				}
				fmt.Println(i, string(buf.Bytes()))
			}
		}
	}

	t, err := template.New("inner").Parse(innerTmpl)
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

	f, err := os.Create(innerPath)
	if err != nil {
		return fmt.Errorf("error opening outer output file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(out)
	if err != nil {
		return fmt.Errorf("error writing outer output file: %w", err)
	}

	return nil
}

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

type Conversion struct {
	TeslaField       string
	WrapperName      string
	WrapperFieldName string
	GoVSSSignalName  string
	OuterInputType   string
	JSONName         string
	InnerInputType   string
	OutputType       string
	Body             string

	ParseFloat bool
	UnitFunc   string
}

type TemplateInput struct {
	Package     string
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
	NiceName              string
}

var teslaTypeToAttributes = map[string]TeslaTypeDescription{
	"string": {
		TeslaWrapperType:      "Value_StringValue",
		TeslaWrapperFieldName: "StringValue",
		ValueType:             "string",
		NiceName:              "String",
	},
	"LocationValue": {
		TeslaWrapperType:      "Value_LocationValue",
		TeslaWrapperFieldName: "LocationValue",
		ValueType:             "*protos.LocationValue",
		NiceName:              "LocationValue",
	},
}
