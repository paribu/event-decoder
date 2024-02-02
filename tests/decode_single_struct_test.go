package tests

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeSingleStruct(t *testing.T) {
	abiFile := "decode_single_struct_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x8fb467591981a04a877379ddb550cccac1341acba5fb4ba2341901f658932715"),
		},
		Data: "000000000000000000000000a508dd875f10c33c52a8abb20e16fc68e981f186",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "address",
			Name:  "newValue",
			Value: `{"A":"0xa508dD875f10C33C52a8abb20E16fc68E981F186"}`,
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

		decoded.Type = strings.ReplaceAll(decoded.Type, "(", "")
		decoded.Type = strings.ReplaceAll(decoded.Type, ")", "")
		if expected.Type != decoded.Type {
			t.Errorf("expected type %s, got %s", expected.Type, decoded.Type)
		}

		if !reflect.DeepEqual(expected.Value, decoded.Value) {
			t.Errorf("expected value %v, got %v", expected.Value, decoded.Value)
		}
	}
}
