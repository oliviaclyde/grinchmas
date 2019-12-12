package cli

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"gitlab.com/parallellearning/lessons/lesson-10/02-db-applied/code-activities/03-arcades-db/unsolved/arcades"
)

type CLI struct {
	totalSpent     int
	arcadesService *arcades.ArcadesService
}

func New(arcadesService *arcades.ArcadesService) *CLI {
	return &CLI{
		totalSpent:     0,
		arcadesService: arcadesService,
	}
}

const (
	addArcadeCmd      = "Add Arcade"
	addGameCabinetCmd = "Add Game Cabinet to Arcade"
	visitArcadeCmd    = "Visit Arcade"
)

func (c *CLI) MainMenu() error {

	prompt := promptui.Select{
		Label: "Select Action",
		Items: []string{
			addArcadeCmd,
			addGameCabinetCmd,
			visitArcadeCmd,
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return err
	}

	switch result {
	case addArcadeCmd:
		err := c.addArcadePrompt()
		if err != nil {
			return err
		}

	case addGameCabinetCmd:
		err := c.addGameCabinet()
		if err != nil {
			return err
		}

	case visitArcadeCmd:
		err := c.visitArcade()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CLI) addArcadePrompt() error {
	namePrompt := promptui.Prompt{
		Label: "Name",
	}
	name, err := namePrompt.Run()
	if err != nil {
		return err
	}

	entryPrice, err := promptInt("Entry Price")
	if err != nil {
		return err
	}

	err = c.arcadesService.AddArcade(name, entryPrice)
	if err != nil {
		return err
	}

	fmt.Println("Added arcade", name)

	return nil
}

func (c *CLI) visitArcade() error {
	chosenArcade, err := c.chooseArcadePromptHelper()
	if err != nil {
		return err
	}

	c.totalSpent = c.totalSpent + chosenArcade.EntryPrice
	err = c.chooseGameAtArcade(chosenArcade)
	if err != nil {
		return err
	}

	return nil
}

func (c *CLI) addGameCabinet() error {

	// TODO: implement this functionality
	// Show the list of arcades
	chosenArcade, err := c.chooseArcadePromptHelper()
	if err != nil {
		return err
	}

	chosenGame, err := c.chooseGamePromptHelper()
	if err != nil {
		return err
	}

	// How much does the game cost
	pricePerGame, err := promptInt("Price per Game")
	if err != nil {
		return err
	}

	// Save to database
	err = c.arcadesService.AddGameCabinetToArcade(chosenGame.ID, chosenArcade.ID, pricePerGame)
	if err != nil {
		return err
	}

	fmt.Println("Added game cabinet", chosenArcade, chosenGame, pricePerGame)

	return nil
}

func (c *CLI) chooseArcadePromptHelper() (arcades.Arcade, error) {
	availableArcades, err := c.arcadesService.ListArcades()
	if err != nil {
		return arcades.Arcade{}, err
	}

	if len(availableArcades) == 0 {
		return arcades.Arcade{}, fmt.Errorf("No arcades to select!")
	}

	var options []string
	for _, arcade := range availableArcades {
		options = append(options, fmt.Sprintf("%s ($%d)", arcade.Name, arcade.EntryPrice))
	}

	selectArcadePrompt := promptui.Select{
		Label: "Select Arcade",
		Items: options,
	}

	chosenIndex, _, err := selectArcadePrompt.Run()
	if err != nil {
		return arcades.Arcade{}, err
	}

	chosenArcade := availableArcades[chosenIndex]

	return chosenArcade, nil
}

func (c *CLI) chooseGamePromptHelper() (arcades.Game, error) {
	availableGames, err := c.arcadesService.ListGames()
	if err != nil {
		return arcades.Game{}, err
	}

	if len(availableGames) == 0 {
		return arcades.Game{}, fmt.Errorf("No games to select!")
	}

	var options []string
	for _, game := range availableGames {
		options = append(options, fmt.Sprintf("%s (%s)", game.GameTitle, game.Rating))
	}

	selectArcadePrompt := promptui.Select{
		Label: "Select Game to Add",
		Items: options,
	}

	chosenIndex, _, err := selectArcadePrompt.Run()
	if err != nil {
		return arcades.Game{}, err
	}

	chosenGame := availableGames[chosenIndex]

	return chosenGame, nil
}

func (c *CLI) chooseGameAtArcade(arcade arcades.Arcade) error {
	gameCabinets, err := c.arcadesService.ListGamesAtArcade(arcade.ID)
	if err != nil {
		return err
	}

	if len(gameCabinets) == 0 {
		fmt.Println("No games at arcade to select!")
		return nil
	}

	var gameOptions []string
	for _, gameCabinet := range gameCabinets {
		game, err := c.arcadesService.GetGame(gameCabinet.GameID)
		if err != nil {
			return err
		}

		gameOptionTxt := fmt.Sprintf("%s (%s) - $%d", game.GameTitle, game.Rating, gameCabinet.PricePerPlay)

		gameOptions = append(gameOptions, gameOptionTxt)
	}

	gameOptions = append(gameOptions, "Exit Arcade")

	for {

		selectGamePrompt := promptui.Select{
			Label: "Select Game",
			Items: gameOptions,
		}

		chosenIndex, _, err := selectGamePrompt.Run()
		if err != nil {
			return err
		}

		if chosenIndex >= len(gameCabinets) {
			fmt.Println("Bye!")
			return nil
		}

		chosenGameCabinet := gameCabinets[chosenIndex]

		c.totalSpent = c.totalSpent + chosenGameCabinet.PricePerPlay

		fmt.Printf("You've spent $%d so far\n", c.totalSpent)
		time.Sleep(500 * time.Millisecond)
	}
}

// promptInt is a helper function to prompt the user for a number input. It validates as the user types.
func promptInt(label string) (int, error) {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Validate: validate,
		Label:    label,
	}
	inputStr, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		return 0, err
	}

	return input, nil
}
