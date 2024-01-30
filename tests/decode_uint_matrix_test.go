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

func TestDecodeIntMatrix(t *testing.T) {
	abiFile := "decode_uint_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x3fceb0c030fe50463f38317524d6884481eee353f4dee6790c605d742654849f"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000200000000000000004067b5e197254cd71e79e27a61ebdca9584ca40e0ea7cfe50000000000000000693032975fbf72490a77adcee836b72ca24ca40e0ea7cfe60000000000000000000000000000000000000000000000000000000000000002000000000000000091f8af4d285997baf67579236e8191afec4ca40e0ea7cfe70000000000000000bac12c02f0f3bd2ce2734477f4cc6c33364ca40e0ea7cfe8",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "uint256[][]",
			Name: "newValue",
			Value: [][]string{
				{"1579208923731619542357098500868790785326998466564056403941", "2579208923731619542357098500868790785326998466564056403942"},
				{"3579208923731619542357098500868790785326998466564056403943", "4579208923731619542357098500868790785326998466564056403944"},
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
