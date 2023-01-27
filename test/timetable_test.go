package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestCleanData(t *testing.T) {
	test := []string{"PÜ;IT2141-Labor SWE II: Gruppe 1;Stricker;CL: 6B.052 (IT-L);", "SU;IT1151-Datenanalyse;Höhne;ONLINE", "R;Wegezeit Online/Präsenz;", "I;Tag des Dualen Studiums;"}

	for _, value := range test {
		arr := strings.Split(value, ";")
		organizer := ""
		kind := arr[0]
		name := arr[1]
		if len(value) >= 3 {
			organizer = arr[2]
		}

		fmt.Println(kind, name, organizer)
	}

}
