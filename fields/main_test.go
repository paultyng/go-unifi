package main

import (
	"fmt"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldInfoFromValidation(t *testing.T) {
	for i, c := range []struct {
		expectedType      string
		expectedComment   string
		expectedOmitEmpty bool
		validation        interface{}
	}{
		{"string", "", true, ""},
		{"string", "default|custom", true, "default|custom"},
		{"string", ".{0,32}", true, ".{0,32}"},
		{"string", "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$", false, "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$"},

		{"int", "^([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$", true, "^([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$"},
		{"int", "", true, "^[0-9]*$"},

		{"float64", "", true, "[-+]?[0-9]*\\.?[0-9]+"},
		// this one is really an error as the . is not escaped
		{"float64", "", true, "^([-]?[\\d]+[.]?[\\d]*)$"},
		{"float64", "", true, "^([\\d]+[.]?[\\d]*)$"},

		{"bool", "", false, "false|true"},
		{"bool", "", false, "true|false"},
	} {
		t.Run(fmt.Sprintf("%d %s %s", i, c.expectedType, c.validation), func(t *testing.T) {
			resource := &Resource{StructName: "TestType", Types: make(map[string]*FieldInfo)}
			fieldInfo, err := resource.fieldInfoFromValidation("fieldName", c.validation)
			//actualType, actualComment, actualOmitEmpty, err := fieldInfoFromValidation(c.validation)
			if err != nil {
				t.Fatal(err)
			}
			if fieldInfo.FieldType != c.expectedType {
				t.Fatalf("expected type %q got %q", c.expectedType, fieldInfo.FieldType)
			}
			if fieldInfo.FieldValidation != c.expectedComment {
				t.Fatalf("expected comment %q got %q", c.expectedComment, fieldInfo.FieldValidation)
			}
			if fieldInfo.OmitEmpty != c.expectedOmitEmpty {
				t.Fatalf("expected omitempty %t got %t", c.expectedOmitEmpty, fieldInfo.OmitEmpty)
			}
		})
	}
}

func TestResourceTypes(t *testing.T) {
	testData := `
{
  "note": ".{0,1024}",
  "date": "^$|^(20[0-9]{2}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])Z?$",
  "mac": "^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$",
  "number": "\\d+",
  "boolean": "true|false",
	"nested_type": {
    "nested_field": "^$"
  },
  "nested_type_array": [{
    "nested_field": "^$"
  }]
}
	`
	expectedFields := map[string]*FieldInfo{
		"Note":    NewFieldInfo("Note", "note", "string", ".{0,1024}", true, false),
		"Date":    NewFieldInfo("Date", "date", "string", "^$|^(20[0-9]{2}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])Z?$", false, false),
		"MAC":     NewFieldInfo("MAC", "mac", "string", "^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$", true, false),
		"Number":  NewFieldInfo("Number", "number", "int", "", true, false),
		"Boolean": NewFieldInfo("Boolean", "boolean", "bool", "", false, false),
		"NestedType": &FieldInfo{
			FieldName:       "NestedType",
			JSONName:        "nested_type",
			FieldType:       "Struct_NestedType",
			FieldValidation: "",
			OmitEmpty:       true,
			IsArray:         false,
			Fields: map[string]*FieldInfo{
				"NestedField": NewFieldInfo("NestedField", "nested_field", "string", "^$", false, false),
			},
		},
		"NestedTypeArray": &FieldInfo{
			FieldName:       "NestedTypeArray",
			JSONName:        "nested_type_array",
			FieldType:       "Struct_NestedTypeArray",
			FieldValidation: "",
			OmitEmpty:       true,
			IsArray:         true,
			Fields: map[string]*FieldInfo{
				"NestedField": NewFieldInfo("NestedField", "nested_field", "string", "^$", false, false),
			},
		},
	}

	expectedStruct := map[string]*FieldInfo{
		"Struct": &FieldInfo{
			FieldName:       "Struct",
			JSONName:        "path",
			FieldType:       "struct",
			FieldValidation: "",
			OmitEmpty:       false,
			IsArray:         false,
			Fields: map[string]*FieldInfo{
				"   ID":      NewFieldInfo("ID", "_id", "string", "", true, false),
				"   SiteID":  NewFieldInfo("SiteID", "site_id", "string", "", true, false),
				"   _Spacer": nil,
				"  Hidden":   NewFieldInfo("Hidden", "attr_hidden", "bool", "", true, false),
				"  HiddenID": NewFieldInfo("HiddenID", "attr_hidden_id", "string", "", true, false),
				"  NoDelete": NewFieldInfo("NoDelete", "attr_no_delete", "bool", "", true, false),
				"  NoEdit":   NewFieldInfo("NoEdit", "attr_no_edit", "bool", "", true, false),
				"  _Spacer":  nil,
				" _Spacer":   nil,
			},
		},
	}

	for k, v := range expectedFields {
		expectedStruct["Struct"].Fields[k] = v
	}

	expectation := &Resource{
		StructName:   "Struct",
		ResourcePath: "path",

		Types: map[string]*FieldInfo{
			"Struct":                 expectedStruct["Struct"],
			"Struct_NestedType":      expectedStruct["Struct"].Fields["NestedType"],
			"Struct_NestedTypeArray": expectedStruct["Struct"].Fields["NestedTypeArray"],
		},
	}

	t.Run("structural test", func(t *testing.T) {
		resource := NewResource("Struct", "path")
		err := resource.processJSON(([]byte)(testData))

		assert.Empty(t, err, "No error processing JSON")
		assert.Equal(t, expectation, resource)
	})
}
