package upload

import "github.com/blocky/timekeeper/internal/hash"

type Upload struct {
	ID       hash.Hexdigest `json:"id"`
	Uploaded bool           `json:"uploaded"`
}
