package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeInt8(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int8_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x76b362fd9358724017c5399c1d512abb9f6aa21d12f9579415990d07e39d1fbf"),
		},
		Data: "000000000000000000000000000000000000000000000000000000000000007f",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int8",
			Name:  "newValue",
			Value: "127",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt16(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int16_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xdb49ce849ac17a85ccd6c5f14d0d624bb1bd1b6cfc459fdae7f405f64406b598"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000007fff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int16",
			Name:  "newValue",
			Value: "32767",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt24(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int24_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x8583a5f61c99eac9ae4774e4341631b295a7dccf24b741e140bb996d57353b88"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000007fffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int24",
			Name:  "newValue",
			Value: "8388607",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt32(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int32_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x87559020b32bf87a1f50ac7af331fb4f896d4fa6cf1020f8328f789090b58099"),
		},
		Data: "000000000000000000000000000000000000000000000000000000007fffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int32",
			Name:  "newValue",
			Value: "2147483647",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt40(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int40_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x8beebf1133fd2d180d5a1a0d8d77fb43ba55677aa73c7bb88ca354a236e01df6"),
		},
		Data: "0000000000000000000000000000000000000000000000000000007fffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int40",
			Name:  "newValue",
			Value: "549755813887",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt48(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int48_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xfd14982a2ce92ba124e994396474dd85d1dc39af99c824914cc6b805e86e0eb1"),
		},
		Data: "00000000000000000000000000000000000000000000000000007fffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int48",
			Name:  "newValue",
			Value: "140737488355327",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt56(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int56_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xa5ddf03b159377ac9c65d0fb6f632108d9cf1010bc35ac528bb24849a56fddd7"),
		},
		Data: "000000000000000000000000000000000000000000000000007fffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int56",
			Name:  "newValue",
			Value: "36028797018963967",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt64(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int64_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x2aef5fb56e7f7c9390386634efc124ed184b724f65dc40c48095d69c26a56f3d"),
		},
		Data: "0000000000000000000000000000000000000000000000007fffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int64",
			Name:  "newValue",
			Value: "9223372036854775807",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt72(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int72_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xeae8c0f6b886d2f41f773ca182591c03de76df66da4d8b5bd796b6bf4652d748"),
		},
		Data: "00000000000000000000000000000000000000000000007fffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int72",
			Name:  "newValue",
			Value: "2361183241434822606847",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt80(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int80_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x8da03080ca1e1f977075bc0c51cfff84eb9dd35489efd9ee35abaf6ee74a76e9"),
		},
		Data: "000000000000000000000000000000000000000000007fffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int80",
			Name:  "newValue",
			Value: "604462909807314587353087",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt88(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int88_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xdef56bd9358ba86d876358d1f0bb9f429685e47494961acc4238c5392e507760"),
		},
		Data: "0000000000000000000000000000000000000000007fffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int88",
			Name:  "newValue",
			Value: "154742504910672534362390527",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}

func TestDecodeInt96(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int96_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xd0627436d530d76e2f51cccddbac41646d6cf791804c302afddd6c15a93bff2c"),
		},
		Data: "00000000000000000000000000000000000000007fffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int96",
			Name:  "newValue",
			Value: "39614081257132168796771975167",
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

		if expected.Value != decoded.Value {
			t.Errorf("expected value %s, got %s", expected.Value, decoded.Value)
		}
	}
}
