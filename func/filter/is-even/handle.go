package function

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/cloudevents/sdk-go/v2/event"
)

// Handle an event.
func Handle(ctx context.Context, e event.Event) (*event.Event, error) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */
	time.Sleep(time.Duration(rand.Intn(51)+50) * time.Millisecond)
	val, err := strconv.Atoi(string(e.Data()))
	if err != nil {
		return nil, err
	}
	if val%2 != 0 {
		return nil, nil
	} else {
		return &e, nil
	}

}

/*
Other supported function signatures:

	Handle()
	Handle() error
	Handle(context.Context)
	Handle(context.Context) error
	Handle(event.Event)
	Handle(event.Event) error
	Handle(context.Context, event.Event)
	Handle(context.Context, event.Event) error
	Handle(event.Event) *event.Event
	Handle(event.Event) (*event.Event, error)
	Handle(context.Context, event.Event) *event.Event
	Handle(context.Context, event.Event) (*event.Event, error)

*/
