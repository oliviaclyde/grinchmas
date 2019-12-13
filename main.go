package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/manifoldco/promptui"
	"github.com/oliviaclyde/grinchmas/db"
	"github.com/oliviaclyde/grinchmas/grinchmasconnect"
)

const (
	addEventCmd       = "Add Horrible Event"
	generateExcuseCmd = "Generate Excuse to Get Out of Your Awful Event"
	viewEventListCmd  = "View Your List of Events"
)

var grinchmasService *grinchmasconnect.GrinchmasService

func main() {

	db, err := db.ConnectDatabase("grinchmas_db.config")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	grinchmasService = grinchmasconnect.NewService(db)

	for {
		fmt.Println(`Welcome to GRINCHMAS!!!! 
		Ever felt like December is overscheduled and overpriced?
		Welcome to a safe space to air your grievances.
		Don't want to attend that party? No judgment here!
		Add all your events to our Grinch caclulator and we'll choose one you don't attend and give you a valid excuse!
		Try us risk-free for 30 days!`)

		prompt := promptui.Select{
			Label: "Select Action",
			Items: []string{
				addEventCmd,
				generateExcuseCmd,
				viewEventListCmd,
			},
		}

		_, choice, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch choice {
		case addEventCmd:
			err := addEventPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case generateExcuseCmd:
			err := generateExcusePrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case viewEventListCmd:
			err := viewEventListPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func viewEventListPrompt() error {
	listOfEvents, _ := grinchmasService.ListEvent()

	if len(listOfEvents) == 0 {
		fmt.Println("You need to add a horrible event you don't want to attend first")
		return nil
	}

	for _, event := range listOfEvents {
		fmt.Println(event.Description)
	}
	return nil
}

func addEventPrompt() error {
	descriptionPrompt := promptui.Prompt{
		Label: "Description",
	}
	description, err := descriptionPrompt.Run()
	if err != nil {
		return err
	}

	newEvent := grinchmasconnect.Event{
		Description: description,
	}

	grinchmasService.AddEvent(newEvent)

	fmt.Println("Event added", description)

	return nil

}

func generateExcusePrompt() error {
	listOfEvents, _ := grinchmasService.ListEvent()

	if len(listOfEvents) == 0 {
		fmt.Println("You need to add a horrible event you don't want to attend first")
		return nil

	}

	MissedEvent := grinchmasService.PickEvent()

	Who, Reason := grinchmasService.GetExcuse()
	// Who := grinchmas.PickActor()

	fmt.Println("I'm terribly sorry I won't be able to " + MissedEvent.Description + ". My " + Who + " said I " + Reason)

	os.Exit(0)
	return nil

}

// func promptForNumber(label string) (float64, error) {
// numberPrompt := promptui.Prompt{
// Label: label,
// // Could add in a validate so you can make sure a number is entered
// }
// numberStr, err := numberPrompt.Run()
// if err != nil {
// return 0, err
// }
// number, err := strconv.ParseFloat(numberStr, 64)
// if err != nil {
// return 0, err
// }
// return number, nil
// }
