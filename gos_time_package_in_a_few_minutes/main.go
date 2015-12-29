package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	RunAllExamples()
	fmt.Printf("\n\n")
}

func RunAllExamples() {
	// Create times.
	Example1()
	Example2()
	Example3()
	Example4()
	Example5()
	Example6()
}

func Example1() {
	var (
		t1, t2, t3     time.Time
		err            error
		estLoc, utcLoc *time.Location
	)

	estLoc, err = time.LoadLocation("EST")
	if err != nil {
		log.Fatalln(err)
	}

	utcLoc, err = time.LoadLocation("UTC")
	if err != nil {
		log.Fatalln(err)
	}

	//Time format, string to interpret format on, and a *time.Location
	t1, err = time.ParseInLocation(time.Kitchen, "6:03PM", estLoc)
	if err != nil {
		log.Fatalln(err)
	}

	// Return a time.Time object based on current environment settings.
	t2 = time.Now()

	/* year int, month Month, day, hour,
	   min, sec, nsec int, loc *Location) Time */
	t3 = time.Date(200, 4, 1, 0, 4, 0, 3, utcLoc)

	fmt.Printf("\n====EXAMPLE 1====")
	fmt.Printf("\nCreating times in three different ways.")

	fmt.Printf("\nt1 = %s \nt2 = %s \nt3 = %s",
		t1.String(), t2.String(), t3.String())
}

func Example2() {
	var (
		t1, t2, t3     time.Time
		err            error
		estLoc, utcLoc *time.Location
	)

	estLoc, err = time.LoadLocation("EST")
	if err != nil {
		log.Fatalln(err)
	}

	utcLoc, err = time.LoadLocation("UTC")
	if err != nil {
		log.Fatalln(err)
	}

	t1, err = time.ParseInLocation(time.Kitchen, "6:03PM", estLoc)
	if err != nil {
		log.Fatalln(err)
	}

	t2, err = time.ParseInLocation(time.Kitchen, "11:03PM", utcLoc)

	fmt.Printf("\n\n====EXAMPLE 2====")
	fmt.Printf("\nComparing two times.")

	fmt.Printf("\nt1 = %v \nt2 = %v",
		t1.String(), t3.String())

	if t1 == t2 {
		fmt.Printf("\nt1 and t2 are equal using `==`.")
	} else {
		fmt.Printf("\nt1 and t2 are not equal using `==`.")
	}

	if t1.Equal(t2) {
		fmt.Printf("\nt1 and t2 are equal using the Equal method.")
	} else {
		fmt.Printf("\nt1 and t2 are not equal using the Equal method.")
	}

	return
}

// Comparing whether two dates occur before or after each other.
func Example3() {
	var (
		t1, t2 time.Time
	)

	// Get current time and date information.
	t1 = time.Now()

	// assign t2 to the current time and date information + 1 day.
	t2 = t1.Add(time.Duration(24) * time.Hour)

	fmt.Printf("\n\n====EXAMPLE 3====")
	fmt.Printf("\nComparing one time occuring before or after another.")

	if t1.After(t2) {
		fmt.Printf("\nt1(%s) occurs after t2(%s)",
			t1.String(), t2.String())
	} else {
		fmt.Printf("\nt1(%s) occurs before t2(%s)",
			t1.String(), t2.String())
	}
}

// Representing a time in different formats.
func Example4() {
	var (
		t1     time.Time
		err    error
		estLoc *time.Location
	)

	estLoc, err = time.LoadLocation("EST")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\n====EXAMPLE 4====")
	fmt.Printf("\nRepresenting time in different formats.")

	t1, err = time.ParseInLocation(time.RFC1123,
		"Fri, 25 Dec 2015 12:00:00 EST",
		estLoc)

	fmt.Printf("\nt1 has been set to Xmas day 2015 Eastern Time.")
	fmt.Printf("\nt1 in RFC3339: %s", t1.Format(time.RFC3339))
	fmt.Printf("\nt1 in RFC1123: %s", t1.Format(time.RFC1123))

	fmt.Printf("\nt1 in custom format: %v %v %v, %v",
		t1.Weekday().String(), t1.Month().String(),
		t1.Day(), t1.Year())
}

// Representing a time from one place to another.
func Example5() {
	var (
		t1     time.Time
		err    error
		estLoc *time.Location
	)

	estLoc, _ = time.LoadLocation("EST")

	fmt.Printf("\n\n====EXAMPLE 5====")
	fmt.Printf("\nRepresenting equivalent times in different time zones.")

	t1, err = time.ParseInLocation(time.RFC1123,
		"Fri, 25 Dec 2015 12:00:00 EST",
		estLoc)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\nt1 (set to Xmas 2015 EST at noon): %s",
		t1.String())

	londonLoc, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nt1 equivalent time in London: %s",
		t1.In(londonLoc).String())
}

func Example6() {
	var (
		t1, t2 time.Time
	)

	fmt.Printf("\n\n====Example 6====")
	fmt.Printf("\nTiming Code")

	t1 = time.Now()

	time.Sleep(time.Duration(3) * time.Second)

	t2 = time.Now()

	tDifference := t2.Sub(t1)
	fmt.Printf("\nThe sleep was %s long", tDifference.String())

	//Alternative
	fmt.Printf("\nPerforming example once more using the Since method.")
	t1 = time.Now()

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("\nThe sleep was %s long.", time.Since(t1).String())
}
