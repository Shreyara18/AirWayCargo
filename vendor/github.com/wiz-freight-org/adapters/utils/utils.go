package utils

import "time"

func AddWorkingDaysUnix(t int64, days int) int64 {
	ti := time.Unix(t, 0)
	return AddWorkingDays(ti, days).Unix()
}

func AddWorkingDays(t time.Time, days int) time.Time {
	m := time.Duration(1)
	if days < 0 {
		m = -1
		days = -days
	}

	if days <= 0 {
		return t
	}

	for {
		t = t.Add(time.Hour * 24 * m)

		if t.Weekday() != time.Saturday && t.Weekday() != time.Sunday {
			days--

			if days <= 0 {
				break
			}
		}
	}

	return t
}
