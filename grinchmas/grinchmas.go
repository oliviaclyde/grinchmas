// API : package that holds the logic

package grinchmas

import (
	"math/rand"
	"time"
)

var (
	events []Event
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Event struct {
	Description string
}

// Adding function to set events to array for use with storage
func SetEvents(e []Event) {
	events = e
}

// Add the event to our list of events
func AddEvent(event Event) {
	events = append(events, event)
}

func ListEvents() []Event {
	return events
}

func GetExcuse() (string, string) {

	Actor := []string{"dog", "boss", "insurance company", "in-laws", "stalker"}
	Action := []string{"forgot to pay them the annual premium", "have to submit to a polygraph", "need to be alone in a dark alley", "need to scoop their poop", "need to give them grandchildren"}

	getActorNumber := rand.Intn(len(Actor))
	getActionNumber := rand.Intn(len(Action))

	return Actor[getActorNumber], Action[getActionNumber]
}

// func PickActor() string {
// 	Actor := []string{"dog", "boss", "insurance company", "in-laws", "stalker"}

// 	getActorNumber := rand.Intn(len(Actor))

// 	return Actor[getActorNumber]
// }

func PickEvent() Event {

	eventNumber := rand.Intn(len(events))

	return events[eventNumber]

}

// func santaList (bool) bool {
// 	if str(bool) == True {
// 		fmt.Println("You're on the nice list.")
// 	} else {
// 		fmt.Println("NAUGHTY LIST!!!!!!")
// 	}
// }

// func howManyEvents {
// 	return len(events)
// }

// type favoriteNumber struct {
// 	Number	int
// }

// // Reduce number to less than 10
// // favorite number % 10 will always be < 10
// func reduceNumber {
// 	if reduceNumber > 10 {
// 		reduceNumber % 10
// 	}
// 	return reduceNumber
// }

// Reducing the length of 10 horrible gifts down to just a few
// Use (x % len(array)) and it will always be less than 10
