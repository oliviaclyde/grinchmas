package grinchmasconnect

import (
	"database/sql"
	"math/rand"
	"time"
)

type GrinchmasService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *GrinchmasService {
	return &GrinchmasService{
		db: db,
	}
}

const (
	insertEventQuery = "INSERT INTO events (event_description) VALUES (?)"

	selectEventQuery = "SELECT id, event_description FROM events"
)

func (a *GrinchmasService) AddEvent(newEvent Event) error {
	_, err := a.db.Exec(insertEventQuery, newEvent.Description)
	if err != nil {
		return err
	}

	return nil
}

func (a *GrinchmasService) ListEvent() ([]Event, error) {
	rows, err := a.db.Query(selectEventQuery)
	if err != nil {
		return nil, err
	}

	var events []Event
	for rows.Next() {
		var event Event

		err := rows.Scan(
			&event.ID,
			&event.Description,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetExcuse() (string, string) {

	Actor := []string{"dog", "boss", "insurance company", "in-laws", "stalker"}
	Action := []string{"forgot to pay them the annual premium", "have to submit to a polygraph", "need to be alone in a dark alley", "need to scoop their poop", "need to give them grandchildren"}

	getActorNumber := rand.Intn(len(Actor))
	getActionNumber := rand.Intn(len(Action))

	return Actor[getActorNumber], Action[getActionNumber]
}

func (a *GrinchmasService) PickEvent() Event {

	events := a.ListEvent()

	eventNumber := rand.Intn(len(events))

	return events[eventNumber]

}
