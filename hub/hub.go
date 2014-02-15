package hub

import (
	"math/rand"
	"os"
	"path"
	// "strings"
)

func init() {
	HOME = os.Getenv("HOME")
	if HOME == "" {
		panic("can't find HOME env variable")
	}
	HUBS_PATH = path.Join(HOME, HUBS_FOLDER)

	_, err := os.Stat(HUBS_PATH)

}

const (
	NAME_LEN    = 6
	FILE_EXT    = "hub"
	HUBS_FOLDER = ".gonethubs"
)

var (
	HOME      = ""
	HUBS_PATH = ""
)

type Hub struct {
	name    string
	running bool
}

func New() *Hub {
	ret := new(Hub)
	ret.name = randString(NAME_LEN)
	return nil
}

func (h *Hub) Start() error {
	return nil
}

func (h *Hub) Stop() error {
	return nil
}
