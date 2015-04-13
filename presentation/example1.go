/* Every file is in a package */
package main

import "fmt"

/**
 * Create a new type called Meetup that
 * contains 3 exported fields and one
 * non-exported field
 */
type Meetup struct {
	Name     string
	Date     string
	Location string
	open     bool
}

func main() {
	/* Create an empty slice of meetups */
	meetups := make([]Meetup, 0)
	meetups = append(meetups, Meetup{
		"Ruby Loco Hack Night",
		"2015-04-13",
		"Phish Me",
		false,
	})

	meetups = append(meetups, Meetup{
		"Ruby Loco Lunch",
		"2015-04-24",
		"Alamo Draft House",
		true,
	})

	fmt.Printf("%v\n", meetups)
}
