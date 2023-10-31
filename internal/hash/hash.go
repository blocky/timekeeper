package hash

import (
	"fmt"

	"golang.org/x/crypto/sha3"

	"github.com/blocky/timekeeper/internal/chronos"
)

type Hash []byte

func MakeHashFromBytes(in []byte) Hash {
	h := sha3.New256()
	h.Write(in)
	out := h.Sum(nil)

	return Hash(out)
}

func MakeHashFromDate(date chronos.Date) Hash {
	start := date.StartDateAndTime()
	stop := date.StopDateAndTime()
	in := fmt.Sprintf("%s:%s", start, stop)

	return MakeHashFromBytes([]byte(in))
}
