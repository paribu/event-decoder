package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/paribu/event-decoder/event"
)

func TestDecodeUint256Matrix(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint256_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x3fceb0c030fe50463f38317524d6884481eee353f4dee6790c605d742654849f"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000200000000000000004067b5e197254cd71e79e27a61ebdca9584ca40e0ea7cfe50000000000000000693032975fbf72490a77adcee836b72ca24ca40e0ea7cfe60000000000000000000000000000000000000000000000000000000000000002000000000000000091f8af4d285997baf67579236e8191afec4ca40e0ea7cfe70000000000000000bac12c02f0f3bd2ce2734477f4cc6c33364ca40e0ea7cfe8",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "uint256[][]",
			Name: "newValue",
			Value: [][]string{
				{"1579208923731619542357098500868790785326998466564056403941", "2579208923731619542357098500868790785326998466564056403942"},
				{"3579208923731619542357098500868790785326998466564056403943", "4579208923731619542357098500868790785326998466564056403944"},
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

		var decodedMatrix Matrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}

func TestDecodeUint256InnerMatrix(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint256_inner_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0xd0eccc6726e9951f45b233daef361d8bb3cc049f593e3b85e00707735c1bba84"),
		},
		Data: "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000001fc5eb2e17b09dbf6a65683df5b3fff900000000000000000000000000000000274bd93eed8ae4990401789626fffffa000000000000000000000000000000002ed1c74fc3652b729d9d88ee584bfffb000000000000000000000000000000003657b560993f724c373999468997fffc",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "uint256[2][]",
			Name: "newValue",
			Value: [][]string{
				{"42233720368547758014223372036854775801", "52233720368547758025223372036854775802"},
				{"62233720368547758036223372036854775803", "72233720368547758047223372036854775804"},
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

		var decodedMatrix Matrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}

func TestDecodeUint256OuterMatrix(t *testing.T) {
	abiFile := "./decode-uint-abi/decode_uint256_outer_matrix_abi.json"

	encodedEvent := &event.Event{
		Topics: []common.Hash{
			common.HexToHash("0x7ab485edd498e72912309da8f7f32060004b915bb72cbc4845745b571b41ded3"),
		},
		Data: "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000020000000000000000bac12c02f0f3bd2ce2734477f4cc6c33364ca40e0ea7cfe900000000000000008de4a2d4944a2d95f875b19afaad48a2cb4ca40e0ea7cfde00000000000000000000000000000000000000000000000000000000000000020000000000000000bac12c02f0f3bd2ce2734477f4cc6c33364ca40e0ea7cfe900000000000000008de4a2d4944a2d95f875b19afaad48a2cb4ca40e0ea7cfde",
	}

	expectedParameters := []MatrixParameter{
		{
			Type: "uint256[][2]",
			Name: "newValue",
			Value: [][]string{
				{"4579208923731619542357098500868790785326998466564056403945", "3479208923731619542357098500868790785326998466564056403934"},
				{"4579208923731619542357098500868790785326998466564056403945", "3479208923731619542357098500868790785326998466564056403934"},
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

		var decodedMatrix Matrix
		if err := json.Unmarshal([]byte(decoded.Value), &decodedMatrix); err != nil {
			t.Errorf("error decoding JSON: %v", err)
		}

		if !reflect.DeepEqual(expected.Value, decodedMatrix.Values) {
			t.Errorf("expected value %v, got %v", expected.Value, decodedMatrix.Values)
		}
	}
}
