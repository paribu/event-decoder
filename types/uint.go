package types

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type uintArray struct {
	Values []string `json:"values"`
}

type uintMatrix struct {
	Values [][]string `json:"values"`
}

func ParseUint(solidityArgument abi.Argument, input interface{}) string {
	return fmt.Sprintf("%d", input)
}

func ParseUintArray(solidityArgument abi.Argument, input interface{}) string {
	if solidityArgument.Type.T == abi.BytesTy || solidityArgument.Type.T == abi.FixedBytesTy {
		return parseByteArray(input)
	}
	return toUintSlice(solidityArgument, input)
}

func ParseUintMatrix(solidityArgument abi.Argument, input interface{}) string {
	if solidityArgument.Type.Elem.T == abi.BytesTy || solidityArgument.Type.Elem.T == abi.FixedBytesTy {
		return parseByteMatrix(solidityArgument, input)
	} else {
		return toUintMatrixSlice(solidityArgument, input)
	}
}

func parseByteArray(input interface{}) string {
	return hexutil.Encode(toByteSlice(input))
}

func parseByteMatrix(solidityArgument abi.Argument, input interface{}) string {
	return encodeBytesToHashArray(solidityArgument, toByteSlices(input))
}

func toByteSlice(data interface{}) []byte {
	var res []byte

	v := reflect.ValueOf(data)
	buf := make([]byte, v.Len())

	reflect.Copy(reflect.ValueOf(buf), v)
	res = append(res, buf...)

	return res
}

func toUintSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	uArr := uintArray{Values: make([]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		uArr.Values[i] = fmt.Sprintf("%d", v.Index(i).Uint())
	}

	str, err := json.Marshal(uArr)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}

func toUintMatrixSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	uMatrix := uintMatrix{Values: make([][]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		uMatrix.Values[i] = make([]string, row.Len())
		for j := 0; j < row.Len(); j++ {
			uMatrix.Values[i][j] = fmt.Sprintf("%d", row.Index(j).Uint())
		}
	}

	str, err := json.Marshal(uMatrix)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}
