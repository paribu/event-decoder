package event

import "github.com/ethereum/go-ethereum/common"

type Event struct {
	Topics []common.Hash `json:"topics"`
	Data   string        `json:"data"`
}

func (e *Event) HasTopics() bool {
	return len(e.Topics) > 0
}
