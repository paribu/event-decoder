package tests

import (
	"encoding/hex"
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

func TestDecodeBytesMatrix(t *testing.T) {
	abiFile := "decode_bytes_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x88190b8ca4680df66be6b9f39c7a998b51ae5811a725652c35d2c16d68230498"),
		},
		Data: "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000023161000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000232610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000002336100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000023461000000000000000000000000000000000000000000000000000000000000",
	}

	expectedParameters := []MatrixParameter{{
		Type: "bytes[][]",
		Name: "newValue",
		Value: [][]string{
			{"1a", "2a"},
			{"3a", "4a"},
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

		decodedMatrixAscii := convertHexToAscii(t, decodedMatrix)
		if !reflect.DeepEqual(expected.Value, decodedMatrixAscii) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrixAscii)
		}
	}

}

func convertHexToAscii(t *testing.T, decodedMatrix Matrix) [][]string {
	rows := len(decodedMatrix.Values)
	cols := len(decodedMatrix.Values[0])

	result := make([][]string, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]string, cols)

		for j := 0; j < cols; j++ {
			hexString := decodedMatrix.Values[i][j]
			hexString = hexString[2:]

			hexBytes, err := hex.DecodeString(hexString)
			if err != nil {
				t.Error(err)
				return nil
			}

			result[i][j] = string(hexBytes)
		}
	}

	return result
}
