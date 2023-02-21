package timetable

import "time"

type Lesson struct {
	Start      *time.Time `json:"start"`
	End        *time.Time `json:"end"`
	Room       string     `json:"room"`
	Teacher    string     `json:"teacher"`
	Kind       string     `json:"kind"`
	Name       string     `json:"name"`
	Pause      int        `json:"pause"`
	Notice     string     `json:"notice"`
	Profession string     `json:"profession"`
	Semester   string     `json:"semester"`
	Group      rune       `json:"group"`
}
