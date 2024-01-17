package types

import (
	"encoding/json"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type byteArray struct {
	Values []string `json:"values"`
}

type byteMatrix struct {
	Values [][]string `json:"values"`
}

func ParseBytesMatrix(solidityArgument abi.Argument, input interface{}) string {
	if solidityArgument.Type.Elem.Elem.T == abi.BytesTy || solidityArgument.Type.Elem.Elem.T == abi.FixedBytesTy {
		return toByteMatrixSlice(solidityArgument, input)
	}
	return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
}

func toByteMatrixSlice(solidityArgument abi.Argument, input interface{}) string {
	v := reflect.ValueOf(input)

	bMatrix := byteMatrix{Values: make([][]string, v.Len())}
	for i := 0; i < v.Len(); i++ {
		row := v.Index(i)
		bMatrix.Values[i] = make([]string, row.Len())
		for j := 0; j < row.Len(); j++ {
			bMatrix.Values[i][j] = hexutil.Encode(row.Index(j).Bytes())
		}
	}

	str, err := json.Marshal(bMatrix)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), reflect.TypeOf(input).String())
	}

	return string(str)
}

func toByteSlices(data interface{}) [][]byte {
	var res [][]byte
	v := reflect.ValueOf(data)
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i)
		buf := make([]byte, item.Len())
		reflect.Copy(reflect.ValueOf(buf), item)
		res = append(res, buf)
	}
	return res
}

func encodeBytesToHashArray(solidityArgument abi.Argument, data [][]byte) string {
	bArr := byteArray{Values: make([]string, len(data))}
	for i, v := range data {
		bArr.Values[i] = hexutil.Encode(v[:])
	}

	str, err := json.Marshal(bArr)
	if err != nil {
		return handleNotSupported(solidityArgument.Type.String(), "[][]byte")
	}

	return string(str)
}
