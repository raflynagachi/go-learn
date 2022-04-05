package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"sort"
	"strconv"
	"time"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (user Users) Len() int {
	return len(user)
}
func (user Users) Less(i, j int) bool {
	return user[i].Name < user[j].Name
}
func (user Users) Swap(i, j int) {
	user[i], user[j] = user[j], user[i]
}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	dateTime, _ := time.Parse("1/2/2006 15:04:05", date)
	return dateTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	dateTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return dateTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	dateTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := dateTime.Hour()
	return (hour >= 12) && (hour < 18)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	schedule := Schedule(date).Format("Monday, January 2, 2006, at 15:04.")
	return "You have an appointment on " + schedule
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	/////////////////////LIST//////////////////////////
	var lst = list.New()
	lst.PushBack(40)
	lst.PushBack(2.3)
	lst.PushFront("Nag")
	fmt.Println(*lst.Back())

	for i := lst.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	for e := lst.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	/////////////////////RING//////////////////////////
	fmt.Println("======RING======")
	var rng *ring.Ring = ring.New(5)
	for i := 0; i < rng.Len(); i++ {
		rng.Value = "Data " + strconv.FormatInt(int64(i), 10)
		rng = rng.Next()
	}
	rng.Do(func(i interface{}) {
		fmt.Println(i)
	})

	//////////////////////SORT////////////////////////
	fmt.Println("=====SORT=====")
	users := Users{
		{Name: "Ruslan", Age: 21},
		{Name: "Nagachi", Age: 20},
	}
	fmt.Println(users)
	sort.Sort(users)
	fmt.Println(users)

	//////////////////////TIME////////////////////////
	fmt.Println("=====TIME=====")
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Clock())

	timeGo := time.Date(2006, time.January, 31, 0, 0, 0, 0, time.Local)
	fmt.Println(timeGo)

	layout := "2006-01-02"
	timeParsed, _ := time.Parse(layout, "2016-01-22")
	fmt.Println(timeParsed)
}
