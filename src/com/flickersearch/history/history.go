package history

import (
	"encoding/json"
)

var history map[string][]string

func init() {
	history = make(map[string][]string)
}

// AddHistory allows the client to set the user history
func AddHistory(name, search string) {
	//store last 10
	if h, ok := history[name]; ok {
		//ensure that it not duplicate
		for _, k := range h {
			if k == search {
				return
			}
		}
		if len(h) == 10 {
			history[name] = append(h[1:], search)
		} else {
			history[name] = append(h, search)
		}

	} else {
		history[name] = []string{
			search,
		}
	}
}

type historyData struct {
	Searches []string `json:"searches"`
}

// GetHistory retrieves history for a user
func GetHistory(name string) ([]byte, error) {
	if h, ok := history[name]; ok {
		return json.Marshal(h)

	}

	return []byte{}, nil
}
