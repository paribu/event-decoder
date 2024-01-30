package tests

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	event_decoder "github.com/paribu/event-decoder"
	"github.com/paribu/event-decoder/event"
)

var (
	_, b, _, _ = runtime.Caller(0)
	currentDir = filepath.Dir(b)
)

func decodeEventWithABI(abiFile string, encodedEvent *event.Event) ([]*event.Parameter, error) {
	// ABI for our event. We read it from a json file
	abiBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", currentDir, abiFile))
	if err != nil {
		return nil, err
	}

	// Parse read json.
	contractABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return nil, err
	}

	// Actually decode event parameters
	parameters, err := event_decoder.Decode(encodedEvent, &contractABI)
	if err != nil {
		return nil, err
	}

	return parameters, nil
}

func hexToASCII(str string) (string, error) {
	str = strings.TrimPrefix(str, "0x")
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
