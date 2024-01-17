package types

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type intArray struct {
	Values []string `json:"values"`
}

type intMatrix struct {
	Values [][]string `json:"values"`
}

func ParseInt(solidityArgument abi.Argument, input interface{}) string {
	return fmt.Sprintf("%d", input)
}

func ParseIntArray(solidityArgument abi.Argument, input interface{}) string {
	return toIntSlice(solidityArgument, input)
}

func ParseIntMatrix(solidityArgument abi.Argument, input interface{}) string {
	return toIntMatrixSlice(solidityArgument, input)
}

func toIntSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	iArr := intArray{Values: make([]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		iArr.Values[i] = fmt.Sprintf("%d", v.Index(i).Int())
	}

	str, err := json.Marshal(iArr)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}

func toIntMatrixSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	iMatrix := intMatrix{Values: make([][]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		iMatrix.Values[i] = make([]string, row.Len())
		for j := 0; j < row.Len(); j++ {
			iMatrix.Values[i][j] = fmt.Sprintf("%d", row.Index(j).Int())
		}
	}

	str, err := json.Marshal(iMatrix)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}
