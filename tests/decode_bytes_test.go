package tests

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

type BytesData struct {
	Data []byte `json:"data"`
}

type BytesParameter struct {
	Type  string    `json:"type"`
	Name  string    `json:"name"`
	Value BytesData `json:"value"`
}

func TestDecodeBytes(t *testing.T) {
	abiFile := "decode_bytes_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xc4891e866dc73e3161ba05c7e6f6392ec67fd5ffc89c1ffd29ebe4995b1660cf"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000014a508dd875f10c33c52a8abb20e16fc68e981f186000000000000000000000000",
	}

	expectedParameters := []BytesParameter{
		{
			Type: "bytes",
			Name: "newValue",
			Value: BytesData{
				Data: []byte("0xa508dd875f10c33c52a8abb20e16fc68e981f186"),
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

		decodedValueBytes := []byte(decoded.Value)
		if !reflect.DeepEqual(expected.Value.Data, decodedValueBytes) {
			t.Errorf("expected value %v, got %v", expected.Value.Data, decodedValueBytes)
		}
	}
}
