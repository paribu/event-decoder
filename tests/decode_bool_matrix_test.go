package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

type BoolMatrix struct {
	Values [][]bool `json:"values"`
}

type BoolMatrixParameter struct {
	Type  string   `json:"type"`
	Name  string   `json:"name"`
	Value [][]bool `json:"value"`
}

func TestDecodeBoolMatrix(t *testing.T) {
	abiFile := "decode_bool_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x32bda4a9513bd493d5eda3f1a830f94a0425c8b664a783294a658f3bb105fce6"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000",
	}

	expectedParameters := []BoolMatrixParameter{
		{
			Type: "bool[][]",
			Name: "newValue",
			Value: [][]bool{
				{false, true},
				{true, false},
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

		var decodedMatrix BoolMatrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}
