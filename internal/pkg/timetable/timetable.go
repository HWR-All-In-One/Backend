package timetable

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	ics "github.com/arran4/golang-ical"
)

const Description = 6

func Parse(url string) (*ics.Calendar, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return ics.ParseCalendar(resp.Body)
}

func DecodeLessons(tt *ics.Calendar) ([]*Lesson, error) {
	lessons := make([]*Lesson, 0)

	for _, event := range tt.Events() {
		v := event.Properties[Description].Value
		desc := decodeDescription(v)
		start, err := event.GetStartAt()

		if err != nil {
			return nil, err
		}

		end, err := event.GetEndAt()

		if err != nil {
			return nil, err
		}

		pause, err := strconv.Atoi(desc["pause"])

		if err != nil {
			return nil, err
		}

		l := Lesson{
			Start:   &start,
			End:     &end,
			Room:    desc["raum"],
			Teacher: desc["dozent"],
			Kind:    desc["art"],
			Notice:  desc["anmerkung"],
			Name:    desc["veranstaltung"],
			Pause:   pause,
		}

		lessons = append(lessons, &l)

	}

	return lessons, nil
}

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
