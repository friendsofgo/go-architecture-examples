package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func New() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
