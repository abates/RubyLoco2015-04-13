/* Every file is in a package */
package main

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

	/* Create a channel for communication between
	 * thread
	 */
	c := make(chan Meetup)

	/* Start the AppendLoop function in its own thread */
	go func() {
		for {
			meetups = append(meetups, <-c)
		}
	}()

	/* Add two items to the channel */
	c <- Meetup{
		"Ruby Loco Hack Night",
		"2015-04-13",
		"Phish Me",
		false,
	}

	c <- Meetup{
		"Ruby Loco Lunch",
		"2015-04-24",
		"Alamo Draft House",
		true,
	}
}
