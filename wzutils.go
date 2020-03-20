package wzlib

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// MakeJid creates a new job ID, based on ULID
func MakeJid() string {
	ts := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(ts.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(ts), entropy).String()
}
