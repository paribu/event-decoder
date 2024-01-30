package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

type Matrix struct {
	Values [][]string `json:"values"`
}

type MatrixParameter struct {
	Type  string     `json:"type"`
	Name  string     `json:"name"`
	Value [][]string `json:"value"`
}

func TestDecodeBytes32Matrix(t *testing.T) {
	abiFile := "decode_bytes32_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xd0675676ee0952ac2c3f0c55d267ec392734bb465a5273d849f4c186c88d576e"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000230783163366637323635366432303639373037333735366432303636363631313078326336663732363536643230363937303733373536643230363636363232000000000000000000000000000000000000000000000000000000000000000230783363366637323635366432303639373037333735366432303636363633333078346336663732363536643230363937303733373536643230363636363434",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "bytes32[][]",
			Name: "newValue",
			Value: [][]string{
				{"0x1c6f72656d20697073756d20666611", "0x2c6f72656d20697073756d20666622"},
				{"0x3c6f72656d20697073756d20666633", "0x4c6f72656d20697073756d20666644"},
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

		for i, row := range decodedMatrix.Values {
			for j := range row {
				decodedMatrix.Values[i][j], err = hexToASCII(decodedMatrix.Values[i][j])
				if err != nil {
					t.Error(err)
				}
			}
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}
