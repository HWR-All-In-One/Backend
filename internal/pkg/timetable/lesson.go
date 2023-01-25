package timetable

import "time"

type Lesson struct {
	Start   *time.Time `mapstructure:"start"`
	End     *time.Time `mapstructure:"end"`
	Room    string     `mapstructure:"raum"`
	Teacher string     `mapstructure:"dozent"`
	Kind    string     `mapstructure:"art"`
	Name    string     `mapstructure:"veranstaltung"`
	Pause   int        `mapstructure:"pause"`
}
