package tests

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

type BytesArrayParameter struct {
	Type  string   `json:"type"`
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

func TestDecodeBytesArray(t *testing.T) {
	abiFile := "decode_bytes_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x8d736a171bdd0307586b03a8a0b63e605e3cdfec9e20c8237c4cd0459569b56d"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000014a508dd875f10c33c52a8abb20e16fc68e981f186000000000000000000000000000000000000000000000000000000000000000000000000000000000000001493c4c1e86434ea4e831d8a13e64ac288c49b7b76000000000000000000000000",
	}

	expectedParameters := []BytesArrayParameter{
		{
			Type: "bytes[]",
			Name: "newValue",
			Value: []string{
				("0xa508dd875f10c33c52a8abb20e16fc68e981f186"),
				("0x93c4c1e86434ea4e831d8a13e64ac288c49b7b76"),
			},
		},
	}

	parameters, err := decodeEventWithABI(abiFile, encodedEvent)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Actual Result:", parameters)

	for i, decoded := range parameters {
		expected := expectedParameters[i]
		if expected.Name != decoded.Name {
			t.Errorf("expected name %s, got %s", expected.Name, decoded.Name)
		}
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}

		var decodedByteArray struct{ Values []string }
		if err := json.Unmarshal([]byte(decoded.Value), &decodedByteArray); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedByteArray.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedByteArray.Values)
		}

	}
}
