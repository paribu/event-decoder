package types

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type stringArray struct {
	Values []string `json:"values"`
}

type stringMatrix struct {
	Values [][]string `json:"values"`
}

func ParseString(solidityArgument abi.Argument, input interface{}) string {
	return fmt.Sprintf("%s", input)
}

func ParseStringArray(solidityArgument abi.Argument, input interface{}) string {
	return toStringSlice(solidityArgument, input)
}

func ParseStringMatrix(solidityArgument abi.Argument, input interface{}) string {
	return toStringMatrixSlice(solidityArgument, input)
}

func toStringSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	sArr := stringArray{Values: make([]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		sArr.Values[i] = v.Index(i).String()
	}

	str, err := json.Marshal(sArr)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}

func toStringMatrixSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	sMatrix := stringMatrix{Values: make([][]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		sMatrix.Values[i] = make([]string, row.Len())
		for j := 0; j < row.Len(); j++ {
			sMatrix.Values[i][j] = row.Index(j).String()
		}
	}

	str, err := json.Marshal(sMatrix)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}
