package ciba

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrWordNotExist error = errors.New("word does not exist")
)

type WordPart struct {
	Part     string   `json:"part"`  // The part of the word (v. for verb, n. for noun, etc.)
	Meanings []string `json:"means"` // Meanings of the word
}

type WordMessage struct {
	Key        string     // The key of the word (word itself)
	Paraphrase string     // A briefing of the meaning of the word
	Value      int        // Unknown
	Meanings   []WordPart `json:"means"` // Meanings of the word
}

func GetBriefings(name string, count int) (msgs []WordMessage, err error) {

	resp, err := http.Get(fmt.Sprintf("http://dict-mobile.iciba.com/interface/index.php?c=word&is_need_mean=1&m=getsuggest&nums=%d&word=%s", count, name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pack := struct{ Message []WordMessage }{Message: msgs}
	dc := json.NewDecoder(resp.Body)
	err = dc.Decode(&pack)
	msgs = pack.Message

	return msgs, err
}
