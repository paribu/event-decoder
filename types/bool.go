package types

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type boolArray struct {
	Values []bool `json:"values"`
}

type boolMatrix struct {
	Values [][]bool `json:"values"`
}

func ParseBool(solidityArgument abi.Argument, input interface{}) string {
	return fmt.Sprintf("%t", input)
}

func ParseBoolArray(solidityArgument abi.Argument, input interface{}) string {
	return toBoolSlice(solidityArgument, input)
}

func ParseBoolMatrix(solidityArgument abi.Argument, input interface{}) string {
	return toBoolMatrixSlice(solidityArgument, input)
}

func toBoolSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	bArr := boolArray{Values: make([]bool, v.Len())}
	for i := 0; i < v.Len(); i++ {
		bArr.Values[i] = v.Index(i).Bool()
	}

	str, err := json.Marshal(bArr)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}

func toBoolMatrixSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	bMatrix := boolMatrix{Values: make([][]bool, v.Len())}
	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		bMatrix.Values[i] = make([]bool, row.Len())
		for j := 0; j < row.Len(); j++ {
			bMatrix.Values[i][j] = row.Index(j).Bool()
		}
	}

	str, err := json.Marshal(bMatrix)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}
