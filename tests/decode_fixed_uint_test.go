package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeFixedUInt(t *testing.T) {
	abiFile := "decode_fixed_uint_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xb5c42705c958db0824a1d28dc4acbfee71066ee56ee60e9720e1088b5dc9ce41"),
		},
		Data: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []ArrayParameter{
		{
			Type: "uint256[2]",
			Name: "newValue",
			Value: []string{"115792089237316195423570985008687907853269984665640564039457584007913129639935",
				"115792089237316195423570985008687907853269984665640564039457584007913129639935"},
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
