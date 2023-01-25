package timetable

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	ics "github.com/arran4/golang-ical"
)

type Timetable struct {
	URL     string
	Lessons []*Lesson
}

func New(url string) *Timetable {
	return &Timetable{
		URL:     url,
		Lessons: make([]*Lesson, 0),
	}
}

func (tt *Timetable) Parse() error {
	resp, err := http.Get(tt.URL)

	if err != nil {
		return nil
	}

	cal, err := ics.ParseCalendar(resp.Body)

	if err != nil {
		return err
	}

	for _, event := range cal.Events() {
		const Description = 6
		v := event.Properties[Description].Value
		desc := tt.decodeDescription(v)
		start, err := event.GetStartAt()

		if err != nil {
			return err
		}

		end, err := event.GetEndAt()

		if err != nil {
			return err
		}

		pause, err := strconv.Atoi(desc["pause"])

		if err != nil {
			return err
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

		tt.Lessons = append(tt.Lessons, &l)

	}

	return nil
}

func (Timetable) decodeDescription(desc string) map[string]string {
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
