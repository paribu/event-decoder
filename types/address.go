package types

import (
	"encoding/json"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type addressArray struct {
	Values []string `json:"values"`
}

type addressMatrix struct {
	Values [][]string `json:"values"`
}

func ParseAddress(solidityArgument abi.Argument, input interface{}) string {
	if val, ok := input.(common.Address); ok {
		return val.Hex()
	}
	return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
}

func ParseAddressArray(solidityArgument abi.Argument, input interface{}) string {
	return toAddressSlice(solidityArgument, input)
}

func ParseAddressMatrix(solidityArgument abi.Argument, input interface{}) string {
	return toAddressMatrixSlice(solidityArgument, input)
}

func toAddressSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	aArr := addressArray{Values: make([]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		aArr.Values[i] = v.Index(i).Interface().(common.Address).Hex()
	}

	str, err := json.Marshal(aArr)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}

func toAddressMatrixSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	aMatrix := addressMatrix{Values: make([][]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		aMatrix.Values[i] = make([]string, row.Len())
		for j := 0; j < row.Len(); j++ {
			aMatrix.Values[i][j] = row.Index(j).Interface().(common.Address).Hex()
		}
	}

	str, err := json.Marshal(aMatrix)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}
