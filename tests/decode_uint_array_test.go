package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeUintArray(t *testing.T) {
	abiFile := "decode_uint_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x4ccac2261fec75801fe2fcf44ae5496037107033c82b28a503a8d40d9c1275de"),
		},
		Data: "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000bac12c02f0f3bd2ce2734477f4cc6c33364ca40e0ea7cfe900000000000000008de4a2d4944a2d95f875b19afaad48a2cb4ca40e0ea7cfde",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "uint256[]",
			Name:  "newValue",
			Value: []string{"4579208923731619542357098500868790785326998466564056403945", "3479208923731619542357098500868790785326998466564056403934"},
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

		var decodedArray Array
		if err := json.Unmarshal([]byte(decoded.Value), &decodedArray); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedArray.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedArray.Values)
		}
	}
}
