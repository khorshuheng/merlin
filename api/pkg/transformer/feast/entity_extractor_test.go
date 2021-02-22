package feast

import (
	"errors"
	"github.com/antonmedv/expr/vm"
	"github.com/mmcloughlin/geohash"
	"testing"

	feast "github.com/feast-dev/feast/sdk/go"
	feastType "github.com/feast-dev/feast/sdk/go/protos/feast/types"
	"github.com/stretchr/testify/assert"

	"github.com/gojek/merlin/pkg/transformer"
)

func TestGetValuesFromJSONPayload(t *testing.T) {
	testData := []byte(`{
		"integer" : 1234,
		"float" : 1234.111,
		"string" : "1234",
		"boolean" : true,
		"booleanString" : "false",
		"latitude": 1.0,
		"longitude": 2.0,
		"struct" : {
                "integer" : 1234,
				"float" : 1234.111,
				"string" : "value",
				"boolean" : true
		},
		"array" : [
		{
                "integer" : 1111,
				"float" : 1111.1111,
				"string" : "value1",
				"boolean" : true
		},
		{
                "integer" : 2222,
				"float" : 2222.2222,
				"string" : "value2",
				"boolean" : false
		}]
	}`)

	tests := []struct {
		name         string
		entityConfig *transformer.Entity
		compiledUdfs map[string]*vm.Program
		expValues    []*feastType.Value
		expError     error
	}{
		{
			"integer to int32",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT32",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.integer",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int32Val(1234),
			},
			nil,
		},
		{
			"integer to int64",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT64",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.integer",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int64Val(1234),
			},
			nil,
		},
		{
			"integer to float",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "FLOAT",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.integer",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.FloatVal(1234),
			},
			nil,
		},
		{
			"integer to double",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "DOUBLE",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.integer",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.DoubleVal(1234),
			},
			nil,
		},
		{
			"integer to string",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "STRING",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.integer",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.StrVal("1234"),
			},
			nil,
		},
		{
			"float to int32",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT32",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.float",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int32Val(1234),
			},
			nil,
		},
		{
			"float to int64",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT64",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.float",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int64Val(1234),
			},
			nil,
		},
		{
			"float to float",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "FLOAT",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.float",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.FloatVal(1234.111),
			},
			nil,
		},
		{
			"float to double",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "DOUBLE",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.float",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.DoubleVal(1234.111),
			},
			nil,
		},
		{
			"float to string",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "STRING",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.float",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.StrVal("1234.111"),
			},
			nil,
		},
		{
			"string to int32",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT32",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.string",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int32Val(1234),
			},
			nil,
		},
		{
			"string to int64",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT64",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.string",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int64Val(1234),
			},
			nil,
		},
		{
			"string to float",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "FLOAT",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.string",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.FloatVal(1234),
			},
			nil,
		},
		{
			"string to double",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "DOUBLE",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.string",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.DoubleVal(1234),
			},
			nil,
		},
		{
			"string to string",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "STRING",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.string",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.StrVal("1234"),
			},
			nil,
		},
		{
			"boolean to boolean",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "BOOL",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.boolean",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.BoolVal(true),
			},
			nil,
		},
		{
			"string to boolean",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "BOOL",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.booleanString",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.BoolVal(false),
			},
			nil,
		},
		{
			"array of integer",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT32",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.array[*].integer",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int32Val(1111),
				feast.Int32Val(2222),
			},
			nil,
		},
		{
			"struct integer",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "INT32",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.struct.integer",
				},
			},
			map[string]*vm.Program{},
			[]*feastType.Value{
				feast.Int32Val(1234),
			},
			nil,
		},
		{
			"Geohash udf",
			&transformer.Entity{
				Name:      "my_geohash",
				ValueType: "STRING",
				Extractor: &transformer.Entity_Udf{
					Udf: "Geohash(\"$.latitude\", \"$.longitude\")",
				},
			},
			map[string]*vm.Program{
				"my_geohash": mustCompileUdf("Geohash(\"$.latitude\", \"$.longitude\")"),
			},
			[]*feastType.Value{
				feast.StrVal(geohash.Encode(1.0, 2.0)),
			},
			nil,
		},
		{
			"unsupported feast type",
			&transformer.Entity{
				Name:      "my_entity",
				ValueType: "BYTES",
				Extractor: &transformer.Entity_JsonPath{
					JsonPath: "$.booleanString",
				},
			},
			map[string]*vm.Program{},
			nil,
			errors.New("unsupported type BYTES"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actual, err := getValuesFromJSONPayload(testData, test.entityConfig, test.compiledUdfs)
			if err != nil {
				if test.expError != nil {
					assert.EqualError(t, err, test.expError.Error())
					return
				} else {
					assert.Fail(t, err.Error())
				}
			}
			assert.Equal(t, test.expValues, actual)
		})
	}
}

func BenchmarkGetValuesFromJSONPayload100Entity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Result, _ = getValuesFromJSONPayload(benchData, &transformer.Entity{
			Name:      "",
			ValueType: "INT32",
			Extractor: &transformer.Entity_JsonPath{
				JsonPath: "$.array[*].id",
			},
		}, map[string]*vm.Program{})
	}
}

func BenchmarkGetValuesFromJSONPayload1StringEntity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Result, _ = getValuesFromJSONPayload(benchData, &transformer.Entity{
			Name:      "",
			ValueType: "STRING",
			Extractor: &transformer.Entity_JsonPath{
				JsonPath: "$.string",
			},
		}, map[string]*vm.Program{})
	}
}

func BenchmarkGetValuesFromJSONPayload1IntegerEntity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Result, _ = getValuesFromJSONPayload(benchData, &transformer.Entity{
			Name:      "",
			ValueType: "INT32",
			Extractor: &transformer.Entity_JsonPath{
				JsonPath: "$.integer",
			},
		}, map[string]*vm.Program{})
	}
}

func BenchmarkGetValuesFromJSONPayload1FloatEntity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Result, _ = getValuesFromJSONPayload(benchData, &transformer.Entity{
			Name:      "",
			ValueType: "DOUBLE",
			Extractor: &transformer.Entity_JsonPath{
				JsonPath: "$.float",
			},
		}, map[string]*vm.Program{})
	}
}

func BenchmarkGetValuesFromJSONPayloadGeohashUdf(b *testing.B) {
	b.ReportAllocs()
	udfString := "Geohash(\"$.latitude\", \"$.longitude\")"
	compiledUdf := mustCompileUdf(udfString)
	for i := 0; i < b.N; i++ {
		Result, _ = getValuesFromJSONPayload(benchData, &transformer.Entity{
			Name:      "my_geohash",
			ValueType: "STRING",
			Extractor: &transformer.Entity_Udf{
				Udf: udfString,
			},
		},
			map[string]*vm.Program{
				"my_geohash": compiledUdf,
			})
	}
}

var Result []*feastType.Value
var benchData = []byte(`{
  "latitude": 1.0,
  "longitude": 1.0,
  "string": "string_value",
  "integer" : 1234,
  "float" : 1234.111,
  "array": [
    {
      "id": 0
    },
    {
      "id": 1
    },
    {
      "id": 2
    },
    {
      "id": 3
    },
    {
      "id": 4
    },
    {
      "id": 5
    },
    {
      "id": 6
    },
    {
      "id": 7
    },
    {
      "id": 8
    },
    {
      "id": 9
    },
    {
      "id": 10
    },
    {
      "id": 11
    },
    {
      "id": 12
    },
    {
      "id": 13
    },
    {
      "id": 14
    },
    {
      "id": 15
    },
    {
      "id": 16
    },
    {
      "id": 17
    },
    {
      "id": 18
    },
    {
      "id": 19
    },
    {
      "id": 20
    },
    {
      "id": 21
    },
    {
      "id": 22
    },
    {
      "id": 23
    },
    {
      "id": 24
    },
    {
      "id": 25
    },
    {
      "id": 26
    },
    {
      "id": 27
    },
    {
      "id": 28
    },
    {
      "id": 29
    },
    {
      "id": 30
    },
    {
      "id": 31
    },
    {
      "id": 32
    },
    {
      "id": 33
    },
    {
      "id": 34
    },
    {
      "id": 35
    },
    {
      "id": 36
    },
    {
      "id": 37
    },
    {
      "id": 38
    },
    {
      "id": 39
    },
    {
      "id": 40
    },
    {
      "id": 41
    },
    {
      "id": 42
    },
    {
      "id": 43
    },
    {
      "id": 44
    },
    {
      "id": 45
    },
    {
      "id": 46
    },
    {
      "id": 47
    },
    {
      "id": 48
    },
    {
      "id": 49
    },
    {
      "id": 50
    },
    {
      "id": 51
    },
    {
      "id": 52
    },
    {
      "id": 53
    },
    {
      "id": 54
    },
    {
      "id": 55
    },
    {
      "id": 56
    },
    {
      "id": 57
    },
    {
      "id": 58
    },
    {
      "id": 59
    },
    {
      "id": 60
    },
    {
      "id": 61
    },
    {
      "id": 62
    },
    {
      "id": 63
    },
    {
      "id": 64
    },
    {
      "id": 65
    },
    {
      "id": 66
    },
    {
      "id": 67
    },
    {
      "id": 68
    },
    {
      "id": 69
    },
    {
      "id": 70
    },
    {
      "id": 71
    },
    {
      "id": 72
    },
    {
      "id": 73
    },
    {
      "id": 74
    },
    {
      "id": 75
    },
    {
      "id": 76
    },
    {
      "id": 77
    },
    {
      "id": 78
    },
    {
      "id": 79
    },
    {
      "id": 80
    },
    {
      "id": 81
    },
    {
      "id": 82
    },
    {
      "id": 83
    },
    {
      "id": 84
    },
    {
      "id": 85
    },
    {
      "id": 86
    },
    {
      "id": 87
    },
    {
      "id": 88
    },
    {
      "id": 89
    },
    {
      "id": 90
    },
    {
      "id": 91
    },
    {
      "id": 92
    },
    {
      "id": 93
    },
    {
      "id": 94
    },
    {
      "id": 95
    },
    {
      "id": 96
    },
    {
      "id": 97
    },
    {
      "id": 98
    },
    {
      "id": 99
    }
  ]
}`)
