package desumasu

import (
	"fmt"
	"regexp"
	"sort"
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
	{"ります", "る", "1"},
	{"ます", "る"},
	{"です", "だ"},
	{"思います", "思う"},
	{"感じます", "感じる"},
	{"知りました", "知った"},
	{"なりました", "なった"},
	{"しまいます", "しまう"},
	{"ておきましょう", "ておく"},
	{"のでしょうか", "のだろうか"},
	{"のでしょうか", "のか"},
	{"でしょうか", "だろうか"},
}

// Convert converts Keitai to Jotai or Jotai to Keitai
// toJotai: true = Keitai to Jotai, false = Jotai to Keitai
// checkNe: true = Check "ね" at the end of the sentence
// removeNe: true = Remove "ね" at the end of the sentence
func Convert(output string, toJotai bool, checkNe bool, removeNe bool) string {
	separator := "、。（）()！？"

	list := [][]string{}
	if toJotai {
		list = convertList
	} else {
		for _, v := range convertList {
			if len(v) > 2 && v[2] == "1" {
				continue
			}

			list = append(list, []string{v[1], v[0]})
		}
	}

	sort.SliceStable(list, func(i, j int) bool {
		return len(list[i][0]) > len(list[j][0])
	})

	for _, i := range list {
		src := i[0]
		dest := i[1]

		output = replaceAllString(output, src, separator, dest)
		if checkNe {
			if !removeNe {
				dest = dest + "ね"
			}
			output = replaceAllString(output, src+"ね", separator, dest)
		}
	}
	return output
}

func replaceAllString(output, src, separator, dest string) string {
	left := fmt.Sprintf("(%s)([%s])", src, separator)
	right := fmt.Sprintf("%s$2", dest)
	return regexp.MustCompile(left).ReplaceAllString(output, right)
}
