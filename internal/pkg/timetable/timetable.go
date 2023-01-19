package timetable

import (
	"os"
	"strings"
	"time"

	"github.com/apognu/gocal"
)

func Parse() ([]Lesson, error) {
	lessons := make([]Lesson, 0)
	file, err := os.Open("./internal/pkg/timetable/inf.ics")

	if err != nil {
		return nil, err
	}

	defer file.Close()
	start, end := time.Now(), time.Now().Add(12*30*24*time.Hour)

	c := gocal.NewParser(file)

	c.Start, c.End = &start, &end
	err = c.Parse()

	if err != nil {
		return nil, err
	}

	for _, value := range c.Events {
		organizer := ""
		arr := strings.Split(value.Summary, ";")
		eventType := arr[0]
		lessonName := arr[1]
		if len(arr) >= 3 {
			organizer = arr[2]
		}
		l := Lesson{
			Start:       value.Start,
			End:         value.End,
			Description: value.Description,
			Summary:     value.Summary,
			Location:    value.Location,
			Organizer:   organizer,
			Type:        eventType,
			Name:        lessonName,
		}

		lessons = append(lessons, l)
	}

	return lessons, nil
}
