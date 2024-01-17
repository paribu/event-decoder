package types

import (
	"reflect"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var SolidityTypeMap = map[interface{}]func(solidityArgument abi.Argument, input interface{}) string{}

func InitSolidityTypeMap() {
	SolidityTypeMap["bool"] = ParseBool
	SolidityTypeMap["[]bool"] = ParseBoolArray
	SolidityTypeMap["[][]bool"] = ParseBoolMatrix
	SolidityTypeMap["string"] = ParseString
	SolidityTypeMap["[]string"] = ParseStringArray
	SolidityTypeMap["[][]string"] = ParseStringMatrix
	SolidityTypeMap["common.Address"] = ParseAddress
	SolidityTypeMap["[]common.Address"] = ParseAddressArray
	SolidityTypeMap["[][]common.Address"] = ParseAddressMatrix
	SolidityTypeMap["uint8"] = ParseUint
	SolidityTypeMap["[]uint8"] = ParseUintArray
	SolidityTypeMap["[][]uint8"] = ParseUintMatrix
	SolidityTypeMap["[][][]uint8"] = ParseBytesMatrix
	SolidityTypeMap["uint16"] = ParseUint
	SolidityTypeMap["[]uint16"] = ParseUintArray
	SolidityTypeMap["[][]uint16"] = ParseUintMatrix
	SolidityTypeMap["uint32"] = ParseUint
	SolidityTypeMap["[]uint32"] = ParseUintArray
	SolidityTypeMap["[][]uint32"] = ParseUintMatrix
	SolidityTypeMap["uint64"] = ParseUint
	SolidityTypeMap["[]uint64"] = ParseUintArray
	SolidityTypeMap["[][]uint64"] = ParseUintMatrix
	SolidityTypeMap["int8"] = ParseInt
	SolidityTypeMap["[]int8"] = ParseIntArray
	SolidityTypeMap["[][]int8"] = ParseIntMatrix
	SolidityTypeMap["int16"] = ParseInt
	SolidityTypeMap["[]int16"] = ParseIntArray
	SolidityTypeMap["[][]int16"] = ParseIntMatrix
	SolidityTypeMap["int32"] = ParseInt
	SolidityTypeMap["[]int32"] = ParseIntArray
	SolidityTypeMap["[][]int32"] = ParseIntMatrix
	SolidityTypeMap["int64"] = ParseInt
	SolidityTypeMap["[]int64"] = ParseIntArray
	SolidityTypeMap["[][]int64"] = ParseIntMatrix
	SolidityTypeMap["*big.Int"] = ParseBigInt
	SolidityTypeMap["[]*big.Int"] = ParseBigIntArray
	SolidityTypeMap["[][]*big.Int"] = ParseBigIntMatrix
	SolidityTypeMap["struct"] = ParseStruct
}

func GetTypeName(v interface{}) string {
	return getSolidityTypeName(reflect.TypeOf(v).String())
}

func getSolidityTypeName(v string) string {
	re := regexp.MustCompile(`\[\d+\]`)
	return re.ReplaceAllString(v, "[]")
}
