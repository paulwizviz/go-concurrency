package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Data struct {
	ID    int
	Field string
	Date  time.Time
}

func monthToInt(month string) time.Month {
	if month == "Jan" {
		return time.January
	}
	if month == "Feb" {
		return time.February
	}
	if month == "Mar" {
		return time.March
	}
	if month == "Apr" {
		return time.April
	}
	if month == "May" {
		return time.May
	}
	if month == "Jun" {
		return time.June
	}
	if month == "Jul" {
		return time.July
	}
	if month == "Aug" {
		return time.August
	}
	if month == "Sep" {
		return time.September
	}
	if month == "Oct" {
		return time.October
	}
	if month == "Nov" {
		return time.December
	}
	if month == "Dec" {
		return time.December
	}
	return time.Month(0)
}

func dateToTime(date string) (time.Time, error) {

	d := strings.Split(date, "-")
	yr, err := strconv.Atoi(d[2])
	if err != nil {
		return time.Time{}, err
	}
	mth := monthToInt(d[1])
	day, err := strconv.Atoi(d[0])
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Date(yr, mth, day, 0, 0, 0, 0, time.UTC)

	return tm, nil
}

func parse(reader *csv.Reader) (c chan Data) {
	c = make(chan Data)
	go func() {
	loop:
		for {
			rec, err := reader.Read()
			if err == io.EOF {

				break
			}
			id, err := strconv.Atoi(rec[0])
			if err != nil {
				continue loop
			}
			dt, err := dateToTime(rec[2])
			if err != nil {
				continue loop
			}
			d := Data{
				ID:    id,
				Field: rec[1],
				Date:  dt,
			}
			c <- d
		}
		close(c)
	}()
	return
}

const (
	ClassOdd = iota + 1
	ClassEven
)

type Classification struct {
	Class int
	Data
}

func OddsEven(dc chan Data) chan Classification {
	c := make(chan Classification)
	go func() {
		for d := range dc {
			if d.ID%2 == 0 {
				c <- Classification{
					Class: ClassEven,
					Data:  d,
				}
			} else {
				c <- Classification{
					Class: ClassOdd,
					Data:  d,
				}
			}
		}
		close(c)
	}()
	return c
}

func Printer(cs ...chan Classification) chan string {
	result := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(cc chan Classification) {
			for c := range cc {
				if c.Class == ClassOdd {
					result <- fmt.Sprintf("Class: Odds Data: %v", c.Data)
				} else {
					result <- fmt.Sprintf("Class: Even Data: %v", c.Data)
				}
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}

func main() {
	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)

	// Extract items from CSV file and pass it through
	// channel d
	d := parse(reader)

	// Fan out
	c1 := OddsEven(d)
	c2 := OddsEven(d)

	// Fan in
	result := Printer(c1, c2)

	// Final process
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for r := range result {
			fmt.Println(r)
		}
		wg.Done()
	}()
	wg.Wait()
}
