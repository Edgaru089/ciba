package ciba

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type SampleSentence struct {
	ID string `json:"idKey"` // The IDKey used by iciba

	AudioURL  string `json:"tts_mp3"`  // The URL pointing to the audio file of the sentence. Empty if not exist.
	AudioSize string `json:"tts_size"` // Estimated size of the audio file. Empty if not exist.

	SourceType  int    `json:"source_type"`  // Unknown. source_type in JSON.
	SourceID    int    `json:"source_id"`    // Unknown. source_id in JSON.
	SourceTitle string `json:"source_title"` // Source title of the sentence.

	NetworkID string `json:"Network_id"` // Unknown. Network_id in JSON.
	TextEN    string `json:"Network_en"` // Sentence text in foreign text.
	TextCN    string `json:"Network_cn"` // Sentence text in Chinese text.
}

// TODO: Synonym represents a class of synonyms of a certain word.
// ... and of course antonyms.
type Synonym struct {
	Part string // Part of the synonym.(v. for verb, n. for noun, etc.). Might be empty.
}

type Symbol struct {
	PronounceEN    string `json:"ph_en"`
	PronounceAM    string `json:"ph_am"`
	PronounceOther string `json:"ph_other"`

	AudioURLEN  string `json:"ph_en_mp3"`
	AudioURLAM  string `json:"ph_am_mp3"`
	AudioURLTTS string `json:"ph_tts_mp3"`

	// There's also those with a "_bk" suffix
	// ph_en_mp3_bk
	// ph_am_mp3_bk
	// ph_tts_mp3_bk

	Parts []struct {
		Part     string   // ".n" for noun, "v." for verb, etc.
		Meanings []string `json:"means"`
	}
}

type BaseInfo struct {
	Name string `json:"word_name"` // Word name

	Exchange *struct {
		Plural []string `json:"word_pl"`
		Third  []string `json:"word_third"`
		Past   []string `json:"word_past"`
		Done   []string `json:"word_done"`
		Ing    []string `json:"word_ing"`
	}

	Symbols []Symbol
}

type Collins struct {
}

type WordDetail struct {
	SampleSentence []SampleSentence `json:"sentence"` // Sample sentences.

	BaseInfo BaseInfo `json:"baesInfo"` // Maybe it's their typo?

}

func GetDetails(word string) (dt WordDetail, err error) {

	resp, err := http.Post(
		"http://dict-mobile.iciba.com/interface/index.php?client=8&type=1&uuid=BD23E471-60B9-5C75-823B-4468BB70CE9B&c=word&m=index&v=1.1.4&sv=10.15.7&sign=a0d31d85c7120974&list=300,100,7,8,10,12,13,9,2,5,14,6,17",
		"application/x-www-form-urlencoded",
		strings.NewReader("word="+word),
	)
	if err != nil {
		return dt, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&struct{ Message *WordDetail }{Message: &dt})
	if err != nil {
		return dt, errors.New("ciba.GetDetails Level 1(msg):" + err.Error())
	}

	//fmt.Fprintln(os.Stderr, dt.SampleSentence[0])
	//fmt.Fprintln(os.Stderr, dt.BaseInfo.Symbols[0])

	return
}
