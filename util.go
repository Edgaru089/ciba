package ciba

import "strings"

//var digits [...]string ={"零","一","二","三","四","五","六","七","八","九","十"}

func ChineseNumber(n int) (s string) {

	if n == 0 {
		return "零"
	}

	var sb strings.Builder

	if n < 0 {
		sb.WriteString("负")
		n = -n
	}

	switch {

	case n < 10:

	}

	return sb.String()
}
