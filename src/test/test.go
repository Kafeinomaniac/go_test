package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	/*c := make(chan string)
	go tick(c)
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", <-c)
	}*/

	i1, err := newInterval2(time.Now(), time.Duration(time.Hour)*24)
	i2, err := newInterval2(i1.end.Add(time.Duration(time.Hour)), time.Duration(time.Hour)*24)
	i3, err := newInterval2(time.Now().Add(time.Duration(time.Hour)), time.Duration(time.Hour)*24*7)
	if err != nil {
		panic(err)
	}
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)

	if i1.overlaps(i2) {
		fmt.Println("i1 et i2 se chevauchent")
	}

	if i1.overlaps(i3) {
		fmt.Println("i1 et i3 se chevauchent")
	}
}

func tick(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s", time.Now())
		time.Sleep(time.Second)
	}
}

type dateTimeInterval struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

func (d dateTimeInterval) String() string {
	return fmt.Sprintf("%v => %v [%v]", d.start.Format("02/01/2006 15:04:05"), d.end.Format("02/01/2006 15:04:05"), d.duration)
}

func newInterval(start, end time.Time) (result *dateTimeInterval) {
	result = new(dateTimeInterval)
	if end.Before(start) {
		result.start = start
		result.end = end
		result.duration = start.Sub(end)
		return
	}
	result.start = start
	result.duration = end.Sub(start)
	return
}

func newInterval2(start time.Time, duration time.Duration) (result *dateTimeInterval, err error) {
	if duration.Nanoseconds() < 0 {
		return nil, errors.New("duration must not be negative")
	}
	result = new(dateTimeInterval)
	result.start = start
	result.end = start.Add(duration)
	result.duration = duration
	return
}

func (d *dateTimeInterval) contains(interval *dateTimeInterval) bool {
	if !d.containsTime(interval.start) {
		return false
	}
	return d.containsTime(interval.end)
}

func (d *dateTimeInterval) containsTime(time time.Time) bool {
	if d.start.Nanosecond() > time.Nanosecond() {
		return false
	}
	return d.end.Nanosecond() >= time.Nanosecond()
}

func (d *dateTimeInterval) overlaps(interval *dateTimeInterval) bool {
	if interval.start.Nanosecond() > d.end.Nanosecond() {
		return false
	}
	return interval.end.Nanosecond() >= d.start.Nanosecond()
}
