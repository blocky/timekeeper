package hash

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Hexdigest string

func makeHexdigestFromBytes(bytes []byte) Hexdigest {
	digest := hex.EncodeToString(bytes)
	return Hexdigest(digest)
}

func MakeHexdigestFromHexString(s string) (Hexdigest, error) {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("not hex encoded: %w", err)
	}
	return makeHexdigestFromBytes(bytes), err
}

func MakeHexdigestFromHash(hash Hash) Hexdigest {
	return makeHexdigestFromBytes(hash)
}

func (d Hexdigest) String() string {
	return string(d)
}

func (h *Hexdigest) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)

	h2, err := MakeHexdigestFromHexString(s)
	if err != nil {
		return fmt.Errorf("could not unmarshal hexdigest: %s", err)
	}

	*h = h2
	return nil
}
