package timetable

import "time"

type Lesson struct {
	Start   *time.Time
	End     *time.Time
	Room    string
	Teacher string
	Kind    string
	Name    string
	Pause   int
}
