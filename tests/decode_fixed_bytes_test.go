package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeFixedBytes(t *testing.T) {
	abiFile := "decode_fixed_bytes_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x7087fa02d60ede2ba66fb9dcf1302fb90f16d47e0fff16f30c35ba136140dd06"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000014a508dd875f10c33c52a8abb20e16fc68e981f1860000000000000000000000000000000000000000000000000000000000000000000000000000000000000014a508dd875f10c33c52a8abb20e16fc68e981f186000000000000000000000000",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "bytes[2]",
			Name:  "newValue",
			Value: []string{"0xa508dD875f10C33C52a8abb20E16fc68E981F186", "0xa508dD875f10C33C52a8abb20E16fc68E981F186"},
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

		for j := 0; j < len(decodedArray.Values); j++ {
			if !reflect.DeepEqual(common.FromHex(expected.Value[j]), common.FromHex(decodedArray.Values[j])) {
				t.Errorf("expected value %v, got %v", expected.Value[j], decodedArray.Values[j])
			}
		}
	}
}
