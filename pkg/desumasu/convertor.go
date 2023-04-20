package desumasu

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var convertList = [][]string{
	{"しましょう", "しよう"},
	{"きましょう", "こう"},
	{"りましょう", "ろう"},
	{"出来ました", "出来た"},
	{"できました", "できた"},
	{"出来ます", "出来る"},
	{"できます", "できる"},
	{"あります", "ある"},
	{"なります", "なる"},
	{"られますが", "られるが"},
	{"きました", "きた"},
	{"ませんが", "ないが"},
	{"でしょう", "だろう"},
	{"りません", "らない"},
	{"みました", "みた"},
	{"ましょう", "よう"},
	{"でした", "だった"},
	{"ですが", "だが"},
	{"います", "いる"},
	{"かります", "かる"},
	{"えました", "えた"},
	{"いいです", "いい"},
	{"ないです", "ない"},
	{"無いです", "無い"},
	{"れます", "れる"},
	{"きます", "くる"},
	{"します", "する"},
	{"ません", "ない"},
	{"ていました", "ていた"},
	{"しまいました", "しまった"},
	{"にしました", "にした"},
	{"されました", "された"},
	{"れませんので", "れないので"},
	{"ました", "た", "1"},
	{"ります", "る"},
	{"りますが", "るが"},
	{"思います", "思う"},
	{"感じます", "感じる"},
	{"知りました", "知った"},
	{"なりました", "なった"},
	{"しまいます", "しまう"},
	{"ておきましょう", "ておく"},
	{"のでしょうか", "のか"},
	{"のでしょうか", "のだろうか"},
	{"でしょうか", "だろうか"},
	{"食べます", "食べる"},
	{"です", "だ"},
}

type conversionPair struct {
	Src  string
	Dest string
}

// Convert converts Keitai to Jotai or Jotai to Keitai
// input: Input string
// toJotai: true = Keitai to Jotai, false = Jotai to Keitai
// checkNe: true = Check "ね" at the end of the sentence
// removeNe: true = Remove "ね" at the end of the sentence
// return: Converted string
func Convert(input string, toJotai bool, checkNe bool, removeNe bool) string {
	separator := "、。（）()！？"

	table := generateConversionTable(toJotai)

	// Sort by length of Src and Dest.
	// Priority: Src's length > Dest's length, so Src's length is multiplied by 10.
	sort.Slice(table, func(i, j int) bool {
		return len(table[i].Src)*10+len(table[i].Dest) > len(table[j].Src)*10+len(table[j].Dest)
	})

	parts := splitString(input, separator)
	for i, part := range parts {
		parts[i] = convert(part, table, separator, checkNe, removeNe)
	}

	return strings.Join(parts, "")
}

func generateConversionTable(toJotai bool) []conversionPair {
	table := make([]conversionPair, len(convertList))

	for i, pair := range convertList {
		if len(pair) > 2 && pair[2] == "1" {
			continue
		}

		if toJotai {
			table[i] = conversionPair{pair[0], pair[1]}
		} else {
			table[i] = conversionPair{pair[1], pair[0]}
		}
	}

	return table
}

func splitString(input string, separator string) []string {
	var (
		r []string
		p int
	)

	re := regexp.MustCompile(fmt.Sprintf("[%s]", separator))

	is := re.FindAllStringIndex(input, -1)
	if is == nil {
		return append(r, input)
	}
	for _, i := range is {
		r = append(r, input[p:i[1]])
		p = i[1]
	}
	return append(r, input[p:])
}

func convert(input string, table []conversionPair, separator string, checkNe bool, removeNe bool) string {
	output := input

	for _, pair := range table {
		src := pair.Src
		dest := pair.Dest

		output = replaceAllString(output, src, dest, separator)

		if checkNe {
			if !removeNe {
				dest = dest + "ね"
			}
			output = replaceAllString(output, src+"ね", dest, separator)
		}

		if output != input {
			break
		}
	}

	return output
}

func replaceAllString(input, src, dest, separator string) string {
	left := "(" + regexp.QuoteMeta(src) + ")([" + regexp.QuoteMeta(separator) + "])"
	right := dest + "$2"
	re := regexp.MustCompile(left)
	return re.ReplaceAllString(input, right)
}
