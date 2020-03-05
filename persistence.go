package persistence

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/325gerbils/go-files"
)

var state map[string]string
var mutex sync.Mutex

// Set stores a string in a persistent key-value map
func Set(key string, value string) {
	if state == nil {
		state = make(map[string]string)
	}
	mutex.Lock()
	state[key] = value
	if !files.Exists("data/persistence") {
		os.Mkdir("data", 0777)
		jsondata, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}
		files.Save("data/persistence", string(jsondata))
	}
	mutex.Unlock()
}

// Get returns a string from a persistent key-value map
func Get(key string) string {
	mutex.Lock()
	if files.Exists("data/persistence") {
		ps := files.Open("data/persistence")
		json.Unmarshal([]byte(ps), &state)
	}
	v := state[key]
	mutex.Unlock()
	return v
}
