package types

import (
	"encoding/json"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type bigIntArray struct {
	Values []string `json:"values"`
}

type bigIntMatrix struct {
	Values [][]string `json:"values"`
}

func ParseBigInt(solidityArgument abi.Argument, input interface{}) string {
	if val, ok := input.(*big.Int); ok {
		return val.String()
	}
	return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
}

func ParseBigIntArray(solidityArgument abi.Argument, input interface{}) string {
	return toBigIntSlice(solidityArgument, input)
}

func ParseBigIntMatrix(solidityArgument abi.Argument, input interface{}) string {
	return toBigIntMatrixSlice(solidityArgument, input)
}

func toBigIntSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	bArr := bigIntArray{Values: make([]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		bArr.Values[i] = v.Index(i).Interface().(*big.Int).String()
	}

	str, err := json.Marshal(bArr)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}

func toBigIntMatrixSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	bMatrix := bigIntMatrix{Values: make([][]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		bMatrix.Values[i] = make([]string, row.Len())
		for j := 0; j < row.Len(); j++ {
			bMatrix.Values[i][j] = row.Index(j).Interface().(*big.Int).String()
		}
	}

	str, err := json.Marshal(bMatrix)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}
