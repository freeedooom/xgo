package xutil

import "time"

func GetRealDaysFromWorkday(from time.Time, n int) int {
	var res int
	for i := 0; i < n; {
		res++
		t := from.AddDate(0, 0, res)

		switch t.Weekday() {
		case time.Saturday, time.Sunday:
			continue
		default:
			// workday
			i++
		}
	}

	return res
}

func WorkdayAfter(from time.Time, n int) time.Time {
	return from.AddDate(0, 0, GetRealDaysFromWorkday(from, n))
}
