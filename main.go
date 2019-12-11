// This is the file to run the grinchmas api

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/oliviaclyde/grinchmas/grinchmas"
	"github.com/oliviaclyde/grinchmas/storage"
)

const (
	addEventCmd       = "Add Horrible Event"
	generateExcuseCmd = "Generate Excuse to Get Out of Your Awful Event"
	viewEventListCmd  = "View Your List of Events"
)

func main() {

	// Add in storage load function
	storage.Load()

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

			// This will save each event we add as a JSON object
			err = storage.Save()
			if err != nil {
				fmt.Println("Prompt failed %v\n", err)
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
	listOfEvents := grinchmas.ListEvents()

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

	newEvent := grinchmas.Event{
		Description: description,
	}

	grinchmas.AddEvent(newEvent)

	fmt.Println("Event added", description)

	return nil

}

func generateExcusePrompt() error {
	listOfEvents := grinchmas.ListEvents()

	if len(listOfEvents) == 0 {
		fmt.Println("You need to add a horrible event you don't want to attend first")
		return nil

	}

	MissedEvent := grinchmas.PickEvent()

	Who, Reason := grinchmas.GetExcuse()
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