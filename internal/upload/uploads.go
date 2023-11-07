package upload

import "github.com/blocky/timekeeper/internal/hash"

type Uploads map[hash.Hexdigest]bool

func MakeUploads() Uploads {
	return make(map[hash.Hexdigest]bool)
}
