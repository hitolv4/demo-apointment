package main

import (
	"fmt"
	"time"
)

//appointment time.Time  sql time  format RFC3339
func EvaluateSameDay(app string) bool {
	current := time.Now()
	const shortForm = "2006-Jan-02"
	t, _ := time.Parse(time.RFC3339, app)
	if current.Format(shortForm) == t.Format(shortForm) {
		return true
	}

	return false

}
func EvaluateFiveDaysMin (app string) bool {
	current := time.Now()
	const shortForm = "2006-Jan-02"
	t, _ := time.Parse(time.RFC3339, app)
	if current
}

func main() {
	fmt.Println(EvaluateSameDay("2021-03-19T10:12:13.914527064-04:00"))
}

