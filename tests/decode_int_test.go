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

func TestDecodeInt104(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int104_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xe89728a762abe6dd071df520131427e24d633cb9dd26abd33ae884cd8a7e46dd"),
		},
		Data: "000000000000000000000000000000000000007fffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int104",
			Name:  "newValue",
			Value: "10141204801825835211973625643007",
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

func TestDecodeInt112(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int112_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xcfe0955ac00cbbab049b757f6b89c2cb65e995702eb680c327bd206439a5d956"),
		},
		Data: "0000000000000000000000000000000000007fffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int112",
			Name:  "newValue",
			Value: "2596148429267413814265248164610047",
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

func TestDecodeInt120(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int120_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x94b9af8dbf0438d3d1e3a2d223b5677aaad3146d45cc36eb927b513bb2135d24"),
		},
		Data: "00000000000000000000000000000000007fffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int120",
			Name:  "newValue",
			Value: "664613997892457936451903530140172287",
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

func TestDecodeInt128(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int128_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xed3ad701959a6432b9cfaf4e20e5aabab59577a7ea247bec17a8ada1629a1cab"),
		},
		Data: "000000000000000000000000000000007fffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int128",
			Name:  "newValue",
			Value: "170141183460469231731687303715884105727",
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

func TestDecodeInt136(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int136_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xb2b71135be932f7fb21390472f155e319881038e5974abe7c3e866e71885a9bb"),
		},
		Data: "0000000000000000000000000000007fffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int136",
			Name:  "newValue",
			Value: "43556142965880123323311949751266331066367",
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

func TestDecodeInt144(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int144_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x91460285300ecee24366b9f896520164756ced47d8700136412544b47c0b1ce9"),
		},
		Data: "00000000000000000000000000007fffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int144",
			Name:  "newValue",
			Value: "11150372599265311570767859136324180752990207",
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

func TestDecodeInt152(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int152_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x6183bcaca48f9e41597ac6d4faa53d4f610341893655fd3c3a71cd966018b9fa"),
		},
		Data: "000000000000000000000000007fffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int152",
			Name:  "newValue",
			Value: "2854495385411919762116571938898990272765493247",
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

func TestDecodeInt160(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int160_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x82d3a57b5d9c22d377100b2244cdb3060e125304ab23912468dd10ac7f631cc3"),
		},
		Data: "0000000000000000000000007fffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int160",
			Name:  "newValue",
			Value: "730750818665451459101842416358141509827966271487",
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

func TestDecodeInt168(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int168_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x205f17acc2a1042bf8878104530c30cd7282802f6f8e9598f5ccfbf9f85b4606"),
		},
		Data: "00000000000000000000007fffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int168",
			Name:  "newValue",
			Value: "187072209578355573530071658587684226515959365500927",
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

func TestDecodeInt176(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int176_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x2ecbcf02d12df05eae2a9a97992cf47022395df365e4bffa4d255c3c7f628943"),
		},
		Data: "000000000000000000007fffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int176",
			Name:  "newValue",
			Value: "47890485652059026823698344598447161988085597568237567",
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

func TestDecodeInt184(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int184_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x9b5195e6a1aec7e4775b44f3e2f26c4b07c3e32b257b67b70bb32775ebcf8bf5"),
		},
		Data: "0000000000000000007fffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int184",
			Name:  "newValue",
			Value: "12259964326927110866866776217202473468949912977468817407",
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

func TestDecodeInt192(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int192_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xae0dd40e71374463bbb7d971199ff881110260bf4cd26162613067e09af67620"),
		},
		Data: "00000000000000007fffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int192",
			Name:  "newValue",
			Value: "3138550867693340381917894711603833208051177722232017256447",
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

func TestDecodeInt200(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int200_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x298d8f63cf7ba30c1c48669f8cb0c2ef3db4933e3649121157381d888fdc3a06"),
		},
		Data: "000000000000007fffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int200",
			Name:  "newValue",
			Value: "803469022129495137770981046170581301261101496891396417650687",
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

func TestDecodeInt208(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int208_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xd153bdb1cee8c6edf92a83b789a6e34908552fcb6046500d66e541277ab351c9"),
		},
		Data: "0000000000007fffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int208",
			Name:  "newValue",
			Value: "205688069665150755269371147819668813122841983204197482918576127",
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

func TestDecodeInt216(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int216_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x306df8fd1508ece5066ebd655d828a815d2cccec330477ce891c8f0eb69ac574"),
		},
		Data: "00000000007fffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int216",
			Name:  "newValue",
			Value: "52656145834278593348959013841835216159447547700274555627155488767",
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

func TestDecodeInt224(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int224_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x3291dbf400ee2a8a53333300cfae489560c2a1ae4b1a795b96b91c2ccdebb79a"),
		},
		Data: "000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int224",
			Name:  "newValue",
			Value: "13479973333575319897333507543509815336818572211270286240551805124607",
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

func TestDecodeInt232(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int232_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x70dc5f322a2c19829315b600c57ecc9960102646850346bfb831536136b6365d"),
		},
		Data: "0000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int232",
			Name:  "newValue",
			Value: "3450873173395281893717377931138512726225554486085193277581262111899647",
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

func TestDecodeInt240(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int240_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x2a50a91d2dd77b271916fc7e30b4caa662bf02a7a77a5e353237dc7862728f00"),
		},
		Data: "00007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int240",
			Name:  "newValue",
			Value: "883423532389192164791648750371459257913741948437809479060803100646309887",
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

func TestDecodeInt248(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int248_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xb968374c9e64f76ced0fe7fc301a125a5674933adb4a48a8eee97790e2e1aa20"),
		},
		Data: "007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int248",
			Name:  "newValue",
			Value: "226156424291633194186662080095093570025917938800079226639565593765455331327",
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

func TestDecodeInt256(t *testing.T) {
	abiFile := "./decode-int-abi/decode_int256_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x99c91f120643dba4ad0dc3770e16881ab6b1c902059670ccd3f2ae1528993959"),
		},
		Data: "7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	}

	expectedParameters := []event.Parameter{
		{
			Type:  "int256",
			Name:  "newValue",
			Value: "57896044618658097711785492504343953926634992332820282019728792003956564819967",
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
