package common

import (
	"regexp"
	"strings"
	"errors"
)

func ExtractCid(s string) ([]string, error) {
	var res []string
	if s == "" {
		res = append(res, s)
		return res, errors.New("empty CID")
	}
	var regexs = []string{
		"^[a-z]\\d{3} ",
		" [a-z]\\d{3} ",
		"^[a-z]\\d{3}$",
		" [a-z]\\d{3}$",
		"^[a-z]\\d{2} \\d",
		" [a-z]\\d{2} \\d",
		"^[a-z] \\d{2} \\d$",
		" [a-z] \\d{2} \\d$",
		"^[a-z] \\d{2} \\d",
		" [a-z] \\d{2} \\d",
		"^[a-z]\\d{2} \\d$",
		" [a-z]\\d{2} \\d$",
		"^[a-z]\\d{2} ",
		" [a-z]\\d{2} ",
		"^[a-z]\\d{2}$",
		" [a-z]\\d{2}$",
		"^[a-z] \\d{3} ",
		" [a-z] \\d{3} ",
		"^[a-z] \\d{3}$",
		" [a-z] \\d{3}$",
		"^[a-z] \\d{2} ",
		" [a-z] \\d{2} ",
		"^[a-z] \\d{2}$",
		" [a-z] \\d{2}$",
		"^[a-z] \\d \\d \\d ",
		" [a-z] \\d \\d \\d ",
		" [a-z] \\d \\d \\d$",
		"^[a-z] \\d \\d \\d$",
	}
	for _, regex := range regexs {
		re := regexp.MustCompile(regex)
		found := re.FindAllString(s, -1)
		if len(found) > 0 {
			for _, cid := range found {
				res = append(res, strings.Replace(cid, " ", "", -1))
			}
		}
	}

	return res, nil

}