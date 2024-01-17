package types

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func ParseStruct(solidityArgument abi.Argument, input interface{}) string {
	val := reflect.ValueOf(input)

	if val.Kind() == reflect.Struct {
		result := make(map[string]interface{})

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i).Interface()
			typeName := GetTypeName(field)

			if strings.Contains(typeName, "struct") {
				typeName = "struct"
			}

			result[val.Type().Field(i).Name] = SolidityTypeMap[typeName](solidityArgument, field)
		}

		str, err := json.Marshal(result)
		if err != nil {
			return handleNotSupported(solidityArgument.Type.String(), val.Type().String())
		}

		return string(str)
	}

	return handleNotSupported(solidityArgument.Type.String(), val.Type().String())
}
