package main

import (
	"github.com/aobco/log"
	cron2 "github.com/robfig/cron/v3"
	"time"
	"xsky.com/ocpf/pkg/cron"
)

func main() {
	//scheduleInterface, err := cron.CustomCronParser.Parse("0 0 */4 * * ?")
	//if err != nil {
	//	log.Fatalf("parse cron expression failed: %v", err)
	//}
	//schedule := scheduleInterface.(*cron2.SpecSchedule)
	//// subtract one period to get the last run time
	//previousTime := Previous(schedule, time.Now())
	//nextTime := scheduleInterface.Next(time.Now())
	//log.Infof("previous run time: %v, next run time: %v", previousTime, nextTime)
	times, err := getNextRunTimes("0 0 0-23/4 * * ?", 20)
	if err != nil {
		log.Fatalf("get next run times failed: %v", err)
	}
	for i, t := range times {
		log.Infof("next run time %d: %v", i, t)
	}
}

func Previous(s *cron2.SpecSchedule, t time.Time) time.Time {
	// General approach
	//
	// For Month, Day, Hour, Minute, Second:
	// Check if the time value matches.  If yes, continue to the next field.
	// If the field doesn't match the schedule, then increment the field until it matches.
	// While incrementing the field, a wrap-around brings it back to the beginning
	// of the field list (since it is necessary to re-verify previous field
	// values)

	// Convert the given time into the schedule's timezone, if one is specified.
	// Save the original timezone so we can convert back after we find a time.
	// Note that schedules without a time zone specified (time.Local) are treated
	// as local to the time provided.
	origLocation := t.Location()
	loc := s.Location
	if loc == time.Local {
		loc = t.Location()
	}
	if s.Location != time.Local {
		t = t.In(s.Location)
	}

	// Start at the earliest possible time (the upcoming second).
	t = t.Add(-time.Duration(t.Nanosecond()) * time.Nanosecond)

	// If no time is found within five years, return zero.
	yearLimit := t.Year() - 5

WRAP:
	if t.Year() < yearLimit {
		return time.Time{}
	}

	// Find the first applicable month.
	// If it's this month, then do nothing.
	for 1<<uint(t.Month())&s.Month == 0 {
		t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc).Add(-time.Second)

		// Wrapped around.
		if t.Month() == time.December {
			goto WRAP
		}
	}

	// Now get a day in that month.
	//
	// NOTE: This causes issues for daylight savings regimes where midnight does
	// not exist.  For example: Sao Paulo has DST that transforms midnight on
	// 11/3 into 1am. Handle that by noticing when the Hour ends up != 0.
	for !dayMatches(s, t) {
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc).Add(-time.Second)
		date := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, loc).AddDate(0, 0, -1)
		if t.Day() == date.Day() {
			goto WRAP
		}
	}

	for 1<<uint(t.Hour())&s.Hour == 0 {
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, loc).Add(-time.Second)

		if t.Hour() == 23 {
			goto WRAP
		}
	}

	for 1<<uint(t.Minute())&s.Minute == 0 {
		t = t.Truncate(time.Minute).Add(-time.Second)

		if t.Minute() == 59 {
			goto WRAP
		}
	}

	for 1<<uint(t.Second())&s.Second == 0 {
		t = t.Add(-time.Second)

		if t.Second() == 59 {
			goto WRAP
		}
	}

	return t.In(origLocation)
}

const (
	// Set the top bit if a star was included in the expression.
	starBit = 1 << 63
)

func dayMatches(s *cron2.SpecSchedule, t time.Time) bool {
	var (
		domMatch bool = 1<<uint(t.Day())&s.Dom > 0
		dowMatch bool = 1<<uint(t.Weekday())&s.Dow > 0
	)
	if s.Dom&starBit > 0 || s.Dow&starBit > 0 {
		return domMatch && dowMatch
	}
	return domMatch || dowMatch
}

func getNextRunTimes(cronExpr string, count int) ([]time.Time, error) {
	schedule, err := cron.CustomCronParser.Parse(cronExpr)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	runTimes := make([]time.Time, 0, count)
	for i := 0; i < count; i++ {
		nextRun := schedule.Next(now)
		runTimes = append(runTimes, nextRun)
		now = nextRun // 更新当前时间为下一个运行时间
	}

	return runTimes, nil
}
