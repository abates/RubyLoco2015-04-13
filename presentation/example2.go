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

/**
 * The Stringer interface specifies a single String()
 * function that takes no arguments and returns a
 * single string value
 */
func (m Meetup) String() string {
	return "" +
		"    Name: " + m.Name + "\n" +
		"    Date: " + m.Date + "\n" +
		"Location: " + m.Date + "\n"
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

	/* The %v format specifier will call String() on the
	 * object if it implements the Stringer interface
	 */
	fmt.Printf("%v\n", meetups)
}
