package timetable

import (
	"regexp"
	"strings"
)

func decodeDescription(desc string) map[string]string {
	arr := strings.Split(desc, "\\n")
	reg := regexp.MustCompile("[^0-9]+")
	result := make(map[string]string)
	for _, value := range arr {
		k, v := strings.Split(value, ":")[0], strings.Split(value, ":")[1:]
		kTrim := strings.ToLower(strings.TrimSpace(k))
		vTrim := strings.TrimSpace(strings.Join(v, ""))
		result[kTrim] = vTrim
	}

	result["dozent"] = strings.ReplaceAll(result["dozent"], "\\", "")
	pause := reg.ReplaceAllString(result["pause"], "")

	if pause == "" {
		pause = "0"
	}
	result["pause"] = pause

	result["anmerkung"] = strings.ReplaceAll(result["anmerkung"], "\\", "")

	return result
}
