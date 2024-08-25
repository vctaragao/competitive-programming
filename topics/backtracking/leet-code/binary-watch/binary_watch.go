package binarywatch

import "fmt"

const (
	MaxHourLeds    = 3
	MaxMinutesLeds = 5
	MaxLeds        = MaxHourLeds + MaxMinutesLeds
)

var (
	hours   = []int{1, 2, 4, 8}
	minutes = []int{1, 2, 4, 8, 16, 32}
)

type Times struct {
	hourLeds   int
	minuteLeds int
	times      []string
}

func Solve(turnedOn int) []string {
	if turnedOn > MaxLeds {
		return []string{}
	}

	return solve(turnedOn, []string{})
}

func solve(leds int, times []string) []string {
	for hourLeds := 0; hourLeds <= MaxHourLeds; hourLeds++ {
		t := Times{
			hourLeds:   hourLeds,
			minuteLeds: leds - hourLeds,
			times:      times,
		}

		t.generateHours(0, 0)
		times = t.times
	}

	return times
}

func (t *Times) generateHours(cI int, hour int) {
	if t.hourLeds <= 0 {
		t.generateMinutes(0, hour, 0)
		return
	}

	t.hourLeds--
	for i := cI; i < len(hours); i++ {
		if hour+hours[i] > 11 {
			continue
		}

		hour += hours[i]
		t.generateHours(i+1, hour)
		hour -= hours[i]
	}

	t.hourLeds++
}

func (t *Times) generateMinutes(cI int, hour int, minute int) {
	if t.minuteLeds == 0 {
		fmtString := "%d:%d"
		if minute < 10 {
			fmtString = "%d:0%d"
		}

		t.times = append(t.times, fmt.Sprintf(fmtString, hour, minute))
		return
	}

	t.minuteLeds--
	for i := cI; i < len(minutes); i++ {
		if minute+minutes[i] > 59 {
			continue
		}

		minute += minutes[i]
		t.generateMinutes(i+1, hour, minute)
		minute -= minutes[i]
	}

	t.minuteLeds++
}
