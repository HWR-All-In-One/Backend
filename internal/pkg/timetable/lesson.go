package timetable

import "time"

type Lesson struct {
	Start       *time.Time
	End         *time.Time
	Description string
	Summary     string
	Location    string
	Organizer   string
	Type        string
	Name        string
}
