package common

import (
	"regexp"
	"strings"
	"errors"
)

func ExtractCid(s string) (string, error) {
	if s == "" {
		return s, errors.New("empty CID")
	}
	// refact
	// re := regexp.MustCompile("^[a-z]\\d{3} | [a-z]\\d{3} |^[a-z]\\d{3}$| [a-z]\\d{3}$|^[a-z]\\d{2} \\d| [a-z]\\d{2} \\d|^[a-z] \\d{2} \\d$| [a-z] \\d{2} \\d$|^[a-z] \\d{2} \\d| [a-z] \\d{2} \\d|^[a-z]\\d{2} \\d$| [a-z]\\d{2} \\d$|^[a-z]\\d{2} | [a-z]\\d{2} |^[a-z]\\d{2}$| [a-z]\\d{2}$|^[a-z] \\d{3} | [a-z] \\d{3} |^[a-z] \\d{3}$| [a-z] \\d{3}$|^[a-z] \\d{2} | [a-z] \\d{2} |^[a-z] \\d{2}$| [a-z] \\d{2}$|^[a-z] \\d \\d \\d | [a-z] \\d \\d \\d | [a-z] \\d \\d \\d$|^[a-z] \\d \\d \\d$")
	re := regexp.MustCompile("^[a-z]\\d{3} | [a-z]\\d{3} |^[a-z]\\d{3}$| [a-z]\\d{3}$|^[a-z]\\d{2} \\d| [a-z]\\d{2} \\d|^[a-z] \\d{2} \\d$| [a-z] \\d{2} \\d$|^[a-z] \\d{2} \\d| [a-z] \\d{2} \\d|^[a-z]\\d{2} \\d$| [a-z]\\d{2} \\d$|^[a-z]\\d{2} | [a-z]\\d{2} |^[a-z]\\d{2}$| [a-z]\\d{2}$|^[a-z] \\d{3} | [a-z] \\d{3} |^[a-z] \\d{3}$| [a-z] \\d{3}$|^[a-z] \\d{2} | [a-z] \\d{2} |^[a-z] \\d{2}$| [a-z] \\d{2}$|^[a-z] \\d \\d \\d | [a-z] \\d \\d \\d | [a-z] \\d \\d \\d$|^[a-z] \\d \\d \\d$")
	res := re.FindStringSubmatch(s)
	if len(res) > 0 {
		return strings.Replace(res[0], " ", "", -1), nil
	 } else {
		return "", nil
	 }
}
