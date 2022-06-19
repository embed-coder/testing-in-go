package mytime

import (
	"fmt"
	"time"
)

type MyTime interface {
	SetTime()
	GetTime() time.Time
	AddTime()
	SubtractTime()
}

type mytime struct {
	currentTime time.Time
}

func NewMyTime() MyTime {
	return &mytime{}
}

func (mt *mytime) SetTime() {
	mt.currentTime = time.Now()
}

func (mt *mytime) GetTime() time.Time {
	return mt.currentTime
}

func (mt *mytime) AddTime(number int) {
	// Add one hour to current time
	mt.currentTime.Add(time.Hour * 1)
	// Add 30 minutes to currentTime
	mt.currentTime.Add(time.Minute * 30)
	// Add 10 seconds to currentTime
	mt.currentTime.Add(time.Second * 10)
}

func (mt *mytime) SubtractTime() {
	// Subtract one hour from current time
	mt.currentTime.Add(-time.Hour * 1)
	// Subtract 30 minutes from currentTime
	mt.currentTime.Add(-time.Minute * 30)
	// Subtract 10 seconds from currentTime
	mt.currentTime.Add(-time.Second * 10)
}

func (mt *mytime) PrintTime() {
	fmt.Println(mt.currentTime)
	// Add one day to current time
	fmt.Println(mt.currentTime.AddDate(0, 0, 1))
	// Add one month to current time
	fmt.Println(mt.currentTime.AddDate(0, 1, 0))
	// Add one year to current time
	fmt.Println(mt.currentTime.AddDate(1, 0, 0))
}
