package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

type replacement struct {
	Old string
	New string
}

var fieldReps = []replacement{
	{"Dhcpdv6", "DHCPDV6"},

	{"Dhcpd", "DHCPD"},
	{"Idx", "IDX"},
	{"Ipsec", "IPSec"},
	{"Ipv6", "IPV6"},
	{"Openvpn", "OpenVPN"},
	{"Tftp", "TFTP"},
	{"Wlangroup", "WLANGroup"},

	{"Bc", "Broadcast"},
	{"Dhcp", "DHCP"},
	{"Dns", "DNS"},
	{"Dpi", "DPI"},
	{"Dtim", "DTIM"},
	{"Firewallgroup", "FirewallGroup"},
	{"Fixedip", "FixedIP"},
	{"Icmp", "ICMP"},
	{"Id", "ID"},
	{"Igmp", "IGMP"},
	{"Ip", "IP"},
	{"Leasetime", "LeaseTime"},
	{"Mac", "MAC"},
	{"Mcastenhance", "MulticastEnhance"},
	{"Minrssi", "MinRSSI"},
	{"Monthdays", "MonthDays"},
	{"Nat", "NAT"},
	{"Networkconf", "Network"},
	{"Networkgroup", "NetworkGroup"},
	{"Pd", "PD"},
	{"Pmf", "PMF"},
	{"Qos", "QOS"},
	{"Radiusprofile", "RADIUSProfile"},
	{"Radius", "RADIUS"},
	{"Ssid", "SSID"},
	{"Startdate", "StartDate"},
	{"Starttime", "StartTime"},
	{"Stopdate", "StopDate"},
	{"Stoptime", "StopTime"},
	{"Tcp", "TCP"},
	{"Udp", "UDP"},
	{"Usergroup", "UserGroup"},
	{"Utc", "UTC"},
	{"Vlan", "VLAN"},
	{"Vpn", "VPN"},
	{"Wan", "WAN"},
	{"Wep", "WEP"},
	{"Wlan", "WLAN"},
	{"Wpa", "WPA"},
}

var fileReps = []replacement{
	{"WlanConf", "WLAN"},
	{"Dhcp", "DHCP"},
	{"Wlan", "WLAN"},
	{"NetworkConf", "Network"},
	{"RadiusProfile", "RADIUSProfile"},
}

var embedTypes bool

type Resource struct {
	StructName     string
	ResourcePath   string
	Types          map[string]*FieldInfo
	FieldProcessor func(name string, f *FieldInfo) error
}

type FieldInfo struct {
	FieldName       string
	JSONName        string
	FieldType       string
	FieldValidation string
	OmitEmpty       bool
	IsArray         bool
	Fields          map[string]*FieldInfo
}

func NewResource(structName string, resourcePath string) *Resource {
	baseType := NewFieldInfo(structName, resourcePath, "struct", "", false, false)
	resource := &Resource{
		StructName:   structName,
		ResourcePath: resourcePath,
		Types: map[string]*FieldInfo{
			structName: baseType,
		},
		FieldProcessor: func(name string, f *FieldInfo) error { return nil },
	}

	// Since template files iterate through map keys in sorted order, these initial fields
	// are named such that they stay at the top for consistency. The spacer items create a
	// blank line in the resulting generated file.
	//
	// This hack is here for stability of the generatd code, but can be removed if desired.
	baseType.Fields = map[string]*FieldInfo{
		"   ID":      NewFieldInfo("ID", "_id", "string", "", true, false),
		"   SiteID":  NewFieldInfo("SiteID", "site_id", "string", "", true, false),
		"   _Spacer": nil,

		"  Hidden":   NewFieldInfo("Hidden", "attr_hidden", "bool", "", true, false),
		"  HiddenID": NewFieldInfo("HiddenID", "attr_hidden_id", "string", "", true, false),
		"  NoDelete": NewFieldInfo("NoDelete", "attr_no_delete", "bool", "", true, false),
		"  NoEdit":   NewFieldInfo("NoEdit", "attr_no_edit", "bool", "", true, false),
		"  _Spacer":  nil,

		" _Spacer": nil,
	}

	if resource.IsSetting() {
		resource.ResourcePath = strcase.ToSnake(strings.TrimPrefix(structName, "Setting"))
		baseType.Fields[" Key"] = NewFieldInfo("Key", "key", "string", "", false, false)
	}

	if resource.StructName == "Device" {
		baseType.Fields[" MAC"] = NewFieldInfo("MAC", "mac", "string", "", false, false)
	}

	if resource.StructName == "User" {
		baseType.Fields[" IP"] = NewFieldInfo("IP", "ip", "string", "non-generated field", true, false)
	}

	return resource
}

func NewFieldInfo(fieldName string, jsonName string, fieldType string, fieldValidation string, omitempty bool, isArray bool) *FieldInfo {
	return &FieldInfo{
		FieldName:       fieldName,
		JSONName:        jsonName,
		FieldType:       fieldType,
		FieldValidation: fieldValidation,
		OmitEmpty:       omitempty,
		IsArray:         isArray,
	}
}

func cleanName(name string, reps []replacement) string {
	for _, rep := range reps {
		name = strings.ReplaceAll(name, rep.Old, rep.New)
	}

	return name
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] versionDir outputDir\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}

func main() {

	flag.Usage = usage

	noEmbeddedTypes := flag.Bool("no-embedded-types", false, "Whether to generate top-level type definitions for embedded type definitions")
	flag.Parse()

	versionDir := flag.Arg(0)
	outputDir := flag.Arg(1)
	embedTypes = !*noEmbeddedTypes

	if versionDir == "" {
		fmt.Print("error: no version directory specified\n\n")
		usage()
		os.Exit(1)
	}

	if outputDir == "" {
		fmt.Print("error: no output directory specified\n\n")
		usage()
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fieldsDir := filepath.Join(wd, versionDir)
	outDir := filepath.Join(wd, outputDir)

	fieldsFiles, err := ioutil.ReadDir(fieldsDir)
	if err != nil {
		panic(err)
	}

	for _, fieldsFile := range fieldsFiles {
		name := fieldsFile.Name()
		ext := filepath.Ext(name)

		switch name {
		case "AuthenticationRequest.json", "Setting.json", "Wall.json":
			continue
		}

		if filepath.Ext(name) != ".json" {
			continue
		}

		name = name[:len(name)-len(ext)]

		urlPath := strings.ToLower(name)
		structName := cleanName(name, fileReps)

		goFile := strcase.ToSnake(structName) + ".generated.go"
		fieldsFilePath := filepath.Join(fieldsDir, fieldsFile.Name())
		b, err := ioutil.ReadFile(fieldsFilePath)
		if err != nil {
			fmt.Printf("skipping file %s: %s", fieldsFile.Name(), err)
			continue
		}

		resource := NewResource(structName, urlPath)

		switch resource.StructName {
		case "Account":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				if name == "IP" {
					f.OmitEmpty = true
				}
				return nil
			}
		case "Device":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "X", "Y":
					f.FieldType = "float64"
				case "Channel", "BackupChannel", "TxPower":
					f.FieldType = "int"
				case "StpPriority", "Ht":
					f.FieldType = "string"
				}

				f.OmitEmpty = true
				return nil
			}
		case "SettingUsg":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				if strings.HasSuffix(name, "Timeout") && name != "ArpCacheTimeout" {
					f.FieldType = "int"
				}
				return nil
			}
		case "User":
			resource.FieldProcessor = func(name string, f *FieldInfo) error {
				switch name {
				case "Blocked":
					f.FieldType = "bool"
				case "LastSeen":
					f.FieldType = "int"
				}
				return nil
			}
		}

		err = resource.processJSON(b)
		if err != nil {
			fmt.Printf("skipping file %s: %s", fieldsFile.Name(), err)
			continue
		}

		code, _ := resource.generateCode()

		_ = os.Remove(filepath.Join(outDir, goFile))
		ioutil.WriteFile(filepath.Join(outDir, goFile), ([]byte)(code), 0644)
	}

	fmt.Printf("%s\n", outDir)
}

func (r *Resource) IsSetting() bool {
	return strings.HasPrefix(r.StructName, "Setting")
}

func (r *Resource) processFields(fields map[string]interface{}) {
	t := r.Types[r.StructName]
	for name, validation := range fields {
		fieldInfo, err := r.fieldInfoFromValidation(name, validation)
		if err != nil {
			continue
		}

		t.Fields[fieldInfo.FieldName] = fieldInfo
	}
}

func (r *Resource) fieldInfoFromValidation(name string, validation interface{}) (fieldInfo *FieldInfo, err error) {
	fieldName := strcase.ToCamel(name)
	fieldName = cleanName(fieldName, fieldReps)

	empty := &FieldInfo{}

	switch validation := validation.(type) {
	case []interface{}:
		if len(validation) == 0 {
			fieldInfo, err = NewFieldInfo(fieldName, name, "string", "", false, true), nil
			err = r.FieldProcessor(fieldName, fieldInfo)
			return fieldInfo, err
		}
		if len(validation) > 1 {
			return empty, fmt.Errorf("unknown validation %#v", validation)
		}

		fieldInfo, err := r.fieldInfoFromValidation(name, validation[0])
		if err != nil {
			return empty, err
		}

		fieldInfo.OmitEmpty = true
		fieldInfo.IsArray = true

		err = r.FieldProcessor(fieldName, fieldInfo)
		return fieldInfo, err

	case map[string]interface{}:
		typeName := r.StructName + fieldName

		result := NewFieldInfo(fieldName, name, typeName, "", true, false)
		result.Fields = make(map[string]*FieldInfo)

		for name, fv := range validation {
			child, err := r.fieldInfoFromValidation(name, fv)
			if err != nil {
				return empty, err
			}

			result.Fields[child.FieldName] = child
		}

		err = r.FieldProcessor(fieldName, result)
		r.Types[typeName] = result
		return result, err

	case string:
		fieldValidation := validation
		normalized := normalizeValidation(validation)

		omitEmpty := false

		switch {
		case normalized == "falsetrue" || normalized == "truefalse":
			fieldInfo, err = NewFieldInfo(fieldName, name, "bool", "", omitEmpty, false), nil
			return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
		default:
			if _, err := strconv.ParseFloat(normalized, 64); err == nil {

				if normalized == "09" || normalized == "09.09" {
					fieldValidation = ""
				}

				if strings.Contains(normalized, ".") {
					if strings.Contains(validation, "\\.){3}") {
						break
					}

					omitEmpty = true
					fieldInfo, err = NewFieldInfo(fieldName, name, "float64", fieldValidation, omitEmpty, false), nil
					return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
				}

				omitEmpty = true
				fieldInfo, err = NewFieldInfo(fieldName, name, "int", fieldValidation, omitEmpty, false), nil
				return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
			}
		}
		if validation != "" && normalized != "" {
			fmt.Printf("normalize %q to %q\n", validation, normalized)
		}

		omitEmpty = omitEmpty || (!strings.Contains(validation, "^$") && !strings.HasSuffix(fieldName, "ID"))
		fieldInfo, err = NewFieldInfo(fieldName, name, "string", fieldValidation, omitEmpty, false), nil
		return fieldInfo, r.FieldProcessor(fieldName, fieldInfo)
	}

	return empty, fmt.Errorf("unable to determine type from validation %q", validation)
}

func (r *Resource) processJSON(b []byte) error {
	var fields map[string]interface{}
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	r.processFields(fields)

	return nil
}

func (r *Resource) generateCode() (string, error) {
	var err error
	var buf bytes.Buffer
	writer := io.Writer(&buf)

	tpl := template.Must(template.New("api.go.tmpl").Funcs(template.FuncMap{
		"embedTypes": func() bool { return embedTypes },
	}).ParseFiles("api.go.tmpl"))

	tpl.Execute(writer, r)

	return buf.String(), err
}

func normalizeValidation(re string) string {
	re = strings.ReplaceAll(re, "\\d", "[0-9]")
	re = strings.ReplaceAll(re, "[-+]?", "")
	re = strings.ReplaceAll(re, "[+-]?", "")
	re = strings.ReplaceAll(re, "[-]?", "")
	re = strings.ReplaceAll(re, "\\.", ".")
	re = strings.ReplaceAll(re, "[.]?", ".")

	quants := regexp.MustCompile(`\{\d*,?\d*\}|\*|\+|\?`)
	re = quants.ReplaceAllString(re, "")

	control := regexp.MustCompile(`[\(\[\]\)\|\-\$\^]`)
	re = control.ReplaceAllString(re, "")

	re = strings.TrimPrefix(re, "^")
	re = strings.TrimSuffix(re, "$")

	return re
}
