package ed_error

import "errors"

var (
	ErrAnonymous      = errors.New("anonymous events are not supported")
	ErrMustHaveTopics = errors.New("event should have topics")
)
