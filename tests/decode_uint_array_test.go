package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeUint8Array(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint8_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x76640e0454e37adebc7db07bebbf0a58e839afcd9d16633bde2e97ef58abb28a"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000ff00000000000000000000000000000000000000000000000000000000000000fe",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "uint8[]",
			Name:  "newValue",
			Value: []string{"255", "254"},
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

func TestDecodeUint24Array(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint24_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x61d6f1ccffbf1f5cf9e52c7f0f401939894f101e19a5bf4bf87eb4d6b774c736"),
		},
		Data: "000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000ffffff0000000000000000000000000000000000000000000000000000000000fffffe",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "uint24[]",
			Name:  "newValue",
			Value: []string{"16777215", "16777214"},
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

func TestDecodeUint32Array(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint32_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xf9593a566df02e076cadd0dc9fdc7a2f68124d2f1f3cb53d2e4fcc2f70abfe1e"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000ffffffff00000000000000000000000000000000000000000000000000000000fffffffe",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "uint32[]",
			Name:  "newValue",
			Value: []string{"4294967295", "4294967294"},
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

func TestDecodeUint40Array(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint40_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x6f6fb5d65cb0be2fe170574db1b48aca3e49faafb44bc5e6558f3a9361ab66c3"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000ffffffffff000000000000000000000000000000000000000000000000000000fffffffffe",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "uint40[]",
			Name:  "newValue",
			Value: []string{"1099511627775", "1099511627774"},
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

func TestDecodeUint64Array(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint64_array_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xd77aedeb4984236c8ec1d875c89b8e3cf7e01203f0ef7ea5f7ee35c279d8c92e"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000ffffffffffffffff000000000000000000000000000000000000000000000000fffffffffffffffe",
	}

	expectedParameters := []ArrayParameter{
		{
			Type:  "uint64[]",
			Name:  "newValue",
			Value: []string{"18446744073709551615", "18446744073709551614"},
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

func TestDecodeUint256Array(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint256_array_abi.json"

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
