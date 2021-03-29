package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Edgaru089/ciba"
)

func wrapDetailsRequest(word string) string {
	dt, err := ciba.GetDetails(word)
	if err != nil {
		return "    [#ff4f4f]内部错误：[-]" + err.Error()
	}
	if len(dt.BaseInfo.Symbols) == 0 {
		return fmt.Sprintf("    [#ff4f4f]错误：[-]单词 %s 不存在", word)
	}

	return buildDetailTUI(dt.BaseInfo.Name, dt, dt.BaseInfo.Symbols)
}

func buildExchangeLine(b ciba.BaseInfo) string {
	if b.Exchange == nil {
		return ""
	}

	var str strings.Builder
	str.WriteString("\n")

	if len(b.Exchange.Plural) > 0 {
		fmt.Fprintf(&str, "  [#8f8fff]复数 [-]")
		for i := 0; i < len(b.Exchange.Plural); i++ {
			if i != 0 {
				str.WriteString(", ")
			}
			str.WriteString(b.Exchange.Plural[i])
		}
		str.WriteString(";")
	}

	if len(b.Exchange.Ing) > 0 {
		fmt.Fprintf(&str, "  [#8f8fff]进行时 [-]")
		for i := 0; i < len(b.Exchange.Ing); i++ {
			if i != 0 {
				str.WriteString(", ")
			}
			str.WriteString(b.Exchange.Ing[i])
		}
		str.WriteString(";")
	}

	if len(b.Exchange.Past) > 0 {
		fmt.Fprintf(&str, "  [#8f8fff]过去式 [-]")
		for i := 0; i < len(b.Exchange.Past); i++ {
			if i != 0 {
				str.WriteString(", ")
			}
			str.WriteString(b.Exchange.Past[i])
		}
		str.WriteString(";")
	}

	if len(b.Exchange.Done) > 0 {
		fmt.Fprintf(&str, "  [#8f8fff]过去分词 [-]")
		for i := 0; i < len(b.Exchange.Done); i++ {
			if i != 0 {
				str.WriteString(", ")
			}
			str.WriteString(b.Exchange.Done[i])
		}
		str.WriteString(";")
	}

	if len(b.Exchange.Third) > 0 {
		fmt.Fprintf(&str, "  [#8f8fff]第三人称单数 [-]")
		for i := 0; i < len(b.Exchange.Third); i++ {
			if i != 0 {
				str.WriteString(", ")
			}
			str.WriteString(b.Exchange.Third[i])
		}
		str.WriteString(";")
	}

	str.WriteString("\n")
	ans := str.String()
	if ans == "\n\n" {
		return ""
	}
	return str.String()
}

func buildDetailTUI(word string, d ciba.WordDetail, sc []ciba.Symbol) string {

	var str strings.Builder

	fmt.Fprintf(&str, " [black:#b7af00] %s \n", word)

	for is, s := range sc {

		if len(sc) > 1 {
			fmt.Fprintf(&str, "\n[#8b8b8b:-]释义 %d:\n", is+1)
		}

		s.PronounceAM = strings.ReplaceAll(s.PronounceAM, "ː", ":")
		s.PronounceEN = strings.ReplaceAll(s.PronounceEN, "ː", ":")
		s.PronounceAM = strings.ReplaceAll(s.PronounceAM, " ", ":")
		s.PronounceEN = strings.ReplaceAll(s.PronounceEN, " ", ":")

		fmt.Fprintf(&str, "\n[#8b8b8b:-]  英: /[white]%s[#8b8b8b]/; 美: /[white]%s[#8b8b8b]/\n\n", s.PronounceAM, s.PronounceEN)

		maxPartLen := 0
		for _, p := range s.Parts {
			if maxPartLen < len(p.Part) {
				maxPartLen = len(p.Part)
			}
		}

		for i := 0; i < len(s.Parts); i++ {
			fmt.Fprintf(&str, "[#8b8b8b] %s [white]", s.Parts[i].Part)
			str.Write(bytes.Repeat([]byte{' '}, maxPartLen-len(s.Parts[i].Part)))

			tab := bytes.Repeat([]byte{' '}, len(s.Parts[i].Part)+2+maxPartLen-len(s.Parts[i].Part))
			for j := 0; j < len(s.Parts[i].Meanings); j++ {
				if j != 0 {
					str.Write(tab)
				}
				fmt.Fprintf(&str, " %s\n", s.Parts[i].Meanings[j])
			}
		}
	}

	if ex := buildExchangeLine(d.BaseInfo); len(ex) > 0 {
		str.WriteString(buildExchangeLine(d.BaseInfo))
	}

	str.Write([]byte("\n\n[black:#b7af00]例句[-:-]"))

	for _, se := range d.SampleSentence {

		fmt.Fprintf(&str, "\n  - %s\n    %s\n", strings.Replace(se.TextEN, word, "[#ff8f8f]"+word+"[-]", -1), se.TextCN)
	}

	return str.String()
}

func buildCollins(str *strings.Builder, d ciba.WordDetail) {

}
