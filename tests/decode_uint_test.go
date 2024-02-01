package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeUint8(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint8_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xf206d150df34037c8e2a1c292d55c236dd29db8e8049f0e7982058f9adcaf49c"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000ff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint8",
			Name:  "newValue",
			Value: "255",
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

func TestDecodeUint16(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint16_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xa858038ccd49495a0fb8c5a45fbdb691b7de721527ab994a441a774b612de191"),
		},
		Data: "000000000000000000000000000000000000000000000000000000000000ffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint16",
			Name:  "newValue",
			Value: "65535",
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

func TestDecodeUint24(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint24_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xf2f23b4fb011c3a4c505f58a63cb4767c5ac7548527ed750d9e97c171dd07b7a"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000ffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint24",
			Name:  "newValue",
			Value: "16777215",
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

func TestDecodeUint32(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint32_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xcae5cb91c4f3e770525817826068609504a105a4af04861b4333fac29d3f9225"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000ffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint32",
			Name:  "newValue",
			Value: "4294967295",
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

func TestDecodeUint40(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint40_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xc411deed415cb87859e97eb1ac863f6b0ba8f94674a0d5b2f04ffc511dd64bd0"),
		},
		Data: "000000000000000000000000000000000000000000000000000000ffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint40",
			Name:  "newValue",
			Value: "1099511627775",
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

func TestDecodeUint48(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint48_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x94eba54952f83263a77539923a6edd4287664bd576a468cf36bb29a5f9b2cee4"),
		},
		Data: "0000000000000000000000000000000000000000000000000000ffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint48",
			Name:  "newValue",
			Value: "281474976710655",
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

func TestDecodeUint56(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint56_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xe6c597a2a555eb58aefbc1bf8471ec9f7f302750355daa96fcd4dacd90930fef"),
		},
		Data: "00000000000000000000000000000000000000000000000000ffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint56",
			Name:  "newValue",
			Value: "72057594037927935",
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

func TestDecodeUint64(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint64_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x6486a095b9503e3d4c10450bcfab4c46eb84850fe537a7a498e3f17b793492bd"),
		},
		Data: "000000000000000000000000000000000000000000000000ffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint64",
			Name:  "newValue",
			Value: "18446744073709551615",
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

func TestDecodeUint72(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint72_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x8525b2be05a8ec622ada4858d4369132700e3603021e162940328621b7ebc2d8"),
		},
		Data: "0000000000000000000000000000000000000000000000ffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint72",
			Name:  "newValue",
			Value: "4722366482869645213695",
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

func TestDecodeUint80(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint80_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x1ac4ecbe8ad28342a89289dec5cb216fadfb78510675bd9d107d78c4a9e7069d"),
		},
		Data: "00000000000000000000000000000000000000000000ffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint80",
			Name:  "newValue",
			Value: "1208925819614629174706175",
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

func TestDecodeUint88(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint88_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x5835c664ca059a712792e54182e6dc71ca54becb3420dcff77109623a20bdc76"),
		},
		Data: "000000000000000000000000000000000000000000ffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint88",
			Name:  "newValue",
			Value: "309485009821345068724781055",
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

func TestDecodeUint96(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint96_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x769f9f8d5e50bb510f2521da6dc0efdad9b0c56ba73e1de7a4475eb41027a8a5"),
		},
		Data: "0000000000000000000000000000000000000000ffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint96",
			Name:  "newValue",
			Value: "79228162514264337593543950335",
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

func TestDecodeUint104(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint104_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x9e78719f2de202ce6122649dd1ebd598c5f3b7509ad633f147a05e35b1ce2a04"),
		},
		Data: "00000000000000000000000000000000000000ffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint104",
			Name:  "newValue",
			Value: "20282409603651670423947251286015",
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

func TestDecodeUint112(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint112_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x0ffbb3934a6d04cc2c8bf8b89946749d1a7ef5f9e70bfa3cac1dccfea3a94d3f"),
		},
		Data: "000000000000000000000000000000000000ffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint112",
			Name:  "newValue",
			Value: "5192296858534827628530496329220095",
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

func TestDecodeUint120(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint120_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xc8131d2e81beedc8edde2351dba7dcf97e6cc7c44d46aa3335b12cb0372bdbb2"),
		},
		Data: "0000000000000000000000000000000000ffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint120",
			Name:  "newValue",
			Value: "1329227995784915872903807060280344575",
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

func TestDecodeUint128(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint128_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x1d1d63ae641aa4127862a72519434ef3bdad156890b39bb785e270448193ff48"),
		},
		Data: "00000000000000000000000000000000ffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint128",
			Name:  "newValue",
			Value: "340282366920938463463374607431768211455",
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

func TestDecodeUint136(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint136_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xbf2e7d67cd770b9a67003ae18d0237cd70c80c3fcd084a74c856541a73c4a2be"),
		},
		Data: "000000000000000000000000000000ffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint136",
			Name:  "newValue",
			Value: "87112285931760246646623899502532662132735",
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

func TestDecodeUint144(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint144_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xe065245bca6701de89fa215224170ffd177dab7105a0f0205ef73e3e3691cf04"),
		},
		Data: "0000000000000000000000000000ffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint144",
			Name:  "newValue",
			Value: "22300745198530623141535718272648361505980415",
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

func TestDecodeUint152(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint152_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xa92cc7f08e9133eafa34be22d7003df4aa17e86bc5ee12bbe18d5f174fd0ad73"),
		},
		Data: "00000000000000000000000000ffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint152",
			Name:  "newValue",
			Value: "5708990770823839524233143877797980545530986495",
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

func TestDecodeUint160(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint160_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x4f91204294cf2023ea631f0838a599013e281043c9241fa772cef5709512e273"),
		},
		Data: "000000000000000000000000ffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint160",
			Name:  "newValue",
			Value: "1461501637330902918203684832716283019655932542975",
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

func TestDecodeUint168(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint168_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xe08c3567f6f4363e651f5777572d03a22d90d3c2a55bda9e61fb15f07d53ff22"),
		},
		Data: "0000000000000000000000ffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint168",
			Name:  "newValue",
			Value: "374144419156711147060143317175368453031918731001855",
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

func TestDecodeUint176(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint176_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x3f1145d6571ceae4b853f9eede39ed3748718c9d11525d601ab4e3a49db6e87b"),
		},
		Data: "00000000000000000000ffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint176",
			Name:  "newValue",
			Value: "95780971304118053647396689196894323976171195136475135",
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

func TestDecodeUint184(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint184_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xb4796332d0513fd9449647bf7fd68ff15dafa58fb8dc9d7f9a729091bf2402a8"),
		},
		Data: "000000000000000000ffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint184",
			Name:  "newValue",
			Value: "24519928653854221733733552434404946937899825954937634815",
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

func TestDecodeUint192(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint192_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xe6c36ce161e25229b8d8a50fd93638c02068511f8cc202ce07cf8786dbb9cf2b"),
		},
		Data: "0000000000000000ffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint192",
			Name:  "newValue",
			Value: "6277101735386680763835789423207666416102355444464034512895",
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

func TestDecodeUint200(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint200_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x76f3976d46a760507964800324c348da72fe9698a45e0b031dd68811cb9495a1"),
		},
		Data: "00000000000000ffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint200",
			Name:  "newValue",
			Value: "1606938044258990275541962092341162602522202993782792835301375",
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

func TestDecodeUint208(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint208_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xc8c6fb90b1d6cfdb30f6677349cdcb768167e4abcefa14f119a5b046af68d75f"),
		},
		Data: "000000000000ffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint208",
			Name:  "newValue",
			Value: "411376139330301510538742295639337626245683966408394965837152255",
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

func TestDecodeUint216(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint216_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xd45ae2ac5a5a2efbc788302b9dcf540ca48fe5526d911cb572d27ad181bf4f8d"),
		},
		Data: "0000000000ffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint216",
			Name:  "newValue",
			Value: "105312291668557186697918027683670432318895095400549111254310977535",
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

func TestDecodeUint224(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint224_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x6f405c224ef71d8250bbfb6c1638614c37683bdc48d988f3d4102ac67f79fe7f"),
		},
		Data: "00000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint224",
			Name:  "newValue",
			Value: "26959946667150639794667015087019630673637144422540572481103610249215",
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

func TestDecodeUint232(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint232_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x3c1624fb854dbddba868ad8b860192ad11b64049be52f04141707f9e736a3351"),
		},
		Data: "000000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint232",
			Name:  "newValue",
			Value: "6901746346790563787434755862277025452451108972170386555162524223799295",
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

func TestDecodeUint240(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint240_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x302825b63313fa43935078e805fe5aed078e8f9be59cb96ba4327c53c5c6c4e6"),
		},
		Data: "0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint240",
			Name:  "newValue",
			Value: "1766847064778384329583297500742918515827483896875618958121606201292619775",
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

func TestDecodeUint248(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint248_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x1a476f3b4597e110048ae8640c68372a33059377ef7f13b51b29eee415acc67e"),
		},
		Data: "00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "uint248",
			Name:  "newValue",
			Value: "452312848583266388373324160190187140051835877600158453279131187530910662655",
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
