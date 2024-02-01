package tests

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeInt256Matrix(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int256_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x5c1bacf54dafce3497c26f76c8391ef2971c09102c841e47d96848150be4fdd8"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000010fa4a62c4dffff90000000000000000000000000000000000000000000000001edb01166c43fffa00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000002cbbb7ca13a7fffb0000000000000000000000000000000000000000000000003a9c6e7dbb0bfffc",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "int256[][]",
			Name: "newValue",
			Value: [][]string{
				{"1223372036854775801", "2223372036854775802"},
				{"3223372036854775803", "4223372036854775804"},
			},
		},
	}

	parameters, err := decodeEventWithABI(abiFile, encodedEvent)
	if err != nil {
		t.Error(err)
	}

	for i, decoded := range parameters {
		expected := expectedParameters[i]

		if expected.Name != decoded.Name {
			t.Errorf("expected name %s, got %s", expected.Name, decoded.Name)
		}
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}

		var decodedMatrix Matrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}

func TestDecodeInt256InnerMatrix(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int256_inner_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x2e4732b55604708b53b44afec72352e633e046bbbe2b4d95f22a7011d4c1054a"),
		},
		Data: "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000003a9c6e7dbb0bfff9000000000000000000000000000000000000000000000000487d2531626ffffa000000000000000000000000000000000000000000000000565ddbe509d3fffb000000000000000000000000000000000000000000000000643e9298b137fffc",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "int256[2][]",
			Name: "newValue",
			Value: [][]string{
				{"4223372036854775801", "5223372036854775802"},
				{"6223372036854775803", "7223372036854775804"},
			},
		},
	}

	parameters, err := decodeEventWithABI(abiFile, encodedEvent)
	if err != nil {
		t.Error(err)
	}

	for i, decoded := range parameters {
		expected := expectedParameters[i]

		if expected.Name != decoded.Name {
			t.Errorf("expected name %s, got %s", expected.Name, decoded.Name)
		}
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}

		var decodedMatrix Matrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}

func TestDecodeInt256OuterMatrix(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int256_outer_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x7c73e41ce741367c584346648a2a65d372370439ebbb28f9e00a7626d7b3a84a"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000007fffffffffffffff000000000000000000000000000000000000000000000000721f494c589c000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000007fffffffffffffff000000000000000000000000000000000000000000000000721f494c589c0000",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "int256[][2]",
			Name: "newValue",
			Value: [][]string{
				{"9223372036854775807", "8223372036854775808"},
				{"9223372036854775807", "8223372036854775808"},
			},
		},
	}

	parameters, err := decodeEventWithABI(abiFile, encodedEvent)
	if err != nil {
		t.Error(err)
	}

	for i, decoded := range parameters {
		expected := expectedParameters[i]

		if expected.Name != decoded.Name {
			t.Errorf("expected name %s, got %s", expected.Name, decoded.Name)
		}
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}

		fmt.Println(decoded)

		var decodedMatrix Matrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}
