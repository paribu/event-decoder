package event_decoder

import (
	"fmt"
	"log"
	"strings"

	ed_error "github.com/paribu/event-decoder/ed-error"
	"github.com/paribu/event-decoder/event"
	"github.com/paribu/event-decoder/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const structField = "struct"

// Decode returns decoded parameters for given event.
func Decode(e *event.Event, contractABI *abi.ABI) ([]*event.Parameter, error) {
	if !e.HasTopics() {
		return nil, ed_error.ErrMustHaveTopics
	}

	abiEvent, err := getEventByID(e, contractABI)
	if err != nil {
		return nil, err
	}

	decodedData, err := decodeData(e.Data)
	if err != nil {
		return nil, err
	}

	eventData, err := unpackData(contractABI, abiEvent.Name, decodedData)
	if err != nil {
		return nil, err
	}

	decodedParameters := make([]*event.Parameter, 0)
	topicIndex := 1

	for _, argument := range abiEvent.Inputs {
		var value string

		if argument.Indexed {
			switch argument.Type.T {
			case abi.BoolTy:
				value = fmt.Sprintf("%t", e.Topics[topicIndex].Big().Uint64() == 1)
			case abi.AddressTy:
				value = common.HexToAddress(e.Topics[topicIndex].Hex()).Hex()
			case abi.IntTy, abi.UintTy:
				value = e.Topics[topicIndex].Big().String()
			case abi.FixedBytesTy:
				value = hexutil.Encode(e.Topics[topicIndex].Bytes()[:argument.Type.Size])
			case abi.BytesTy, abi.StringTy, abi.SliceTy, abi.ArrayTy:
				// If a String or Bytes parameter is indexed, it is stored as a hash in the topic.
				// We don't actually decode them (because there is no way to do it),
				// we just return the hashed values
				value = e.Topics[topicIndex].Hex()
			default:
				log.Printf("Indexed type %b not supported", argument.Type.T)
			}

			topicIndex++
		} else {
			input := eventData[argument.Name]
			typeName := types.GetTypeName(input)

			if strings.Contains(typeName, structField) {
				typeName = structField
			}

			value = types.SolidityTypeMap[typeName](argument, input)
		}

		decodedParameters = append(decodedParameters, &event.Parameter{
			Name:  argument.Name,
			Type:  argument.Type.String(),
			Value: value,
		})
	}

	return decodedParameters, nil
}

func getEventByID(e *event.Event, contractABI *abi.ABI) (*abi.Event, error) {
	abiEvent, err := contractABI.EventByID(e.Topics[0])
	if err != nil {
		return nil, err
	}

	if abiEvent.Anonymous {
		return nil, ed_error.ErrAnonymous
	}

	return abiEvent, nil
}

func decodeData(data string) ([]byte, error) {
	// Normalize data by prepending "0x"
	return hexutil.Decode(fmt.Sprintf("0x%s", strings.TrimPrefix(data, "0x")))
}

func unpackData(contractABI *abi.ABI, eventName string, decodedData []byte) (map[string]interface{}, error) {
	eventData := make(map[string]interface{})

	err := contractABI.UnpackIntoMap(eventData, eventName, decodedData)
	if err != nil {
		return eventData, err
	}

	return eventData, nil
}
