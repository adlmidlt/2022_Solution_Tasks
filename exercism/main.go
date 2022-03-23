package main

import "fmt"

/*func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if petDogIsPresent == true && archerIsAwake == false {
		return true
	}
	if petDogIsPresent == false && prisonerIsAwake == true && knightIsAwake == false && archerIsAwake == false {
		return true
	}

	return false
}

const (
	CostOfPlanningTenCars = 95000
	CostOfProducedOneCar  = 10000
	Percent               = 100
	NumberMinutesInAnHour = 60
)

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / Percent
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	numberOfCarsProducedInMinute := productionRate / NumberMinutesInAnHour
	numberOfSuccessfullyProducedCars := successRate / Percent * float64(numberOfCarsProducedInMinute)

	return int(numberOfSuccessfullyProducedCars)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	numberOfGroupsCars := uint(carsCount / 10)
	numberOfIndividualCars := uint(carsCount) - numberOfGroupsCars*10

	return numberOfGroupsCars*CostOfPlanningTenCars + numberOfIndividualCars*CostOfProducedOneCar
}

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	nameOfCustomerInCapitalLetters := strings.ToUpper(customer)
	messageOfGreeting := "Welcome to the Tech Palace, " + nameOfCustomerInCapitalLetters

	return messageOfGreeting
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	numberOfStars := ""
	for i := 0; i < numStarsPerLine; i++ {
		numberOfStars += "*"
	}
	framedGreeting := numberOfStars + "\n" + welcomeMsg + "\n" + numberOfStars

	return framedGreeting
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	removeAsterisksFromMessage := strings.ReplaceAll(oldMsg, "*", "")
	outputCleanMessage := strings.TrimSpace(removeAsterisksFromMessage)

	return outputCleanMessage
}

// Welcome greets a person by name.
func Welcome(name string) string {
	greetingPersonByName := fmt.Sprintf("Welcome to my party, %s!", name)
	return greetingPersonByName
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	messageAboutBirthdayAndAge := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)

	return messageAboutBirthdayAndAge
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	messageForGuests := fmt.Sprintf(
		"Welcome to my party, %s!\n"+
			"You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n"+
			"You will be sitting next to %s.", name, table, direction, distance, neighbor)
	return messageForGuests
}

// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
// You will be sitting next to Frank.

/*Если у Анналин есть домашняя собака, она может спасти заключенного, если лучник спит. Рыцарь боится пса и лучник не успеет собраться, прежде чем Анналин и пленник смогут сбежать.
информация о текущем прогнозк погоды в городе

Если у Анналин нет собаки, то она и заключенный должны быть очень хитрыми! Анналин может освободить заключенного,
если заключенный бодрствует, а рыцарь и лучник оба спят, но если заключенный спит, их нельзя спасти:
	заключенный будет поражен внезапным появлением Анналин и разбудит рыцаря и лучника.

Функция возвращает значение true, если заключенного можно освободить, основываясь на состоянии трех персонажей
и присутствии собаки Анналин. В противном случае возвращается ложь:

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	choiceOfCar := ""
	if option1 > option2 {
		choiceOfCar = option2
	} else {
		choiceOfCar = option1
	}

	return fmt.Sprintf("%s is clearly the better choice.", choiceOfCar)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	costVehicleResell := 0.0
	if age < 3 {
		costVehicleResell = originalPrice / 100 * 80
	} else if age >= 10 {
		costVehicleResell = originalPrice / 100 * 50
	} else if age > 3 && age < 10 {
		costVehicleResell = originalPrice / 100 * 70
	}

	return costVehicleResell
}
*/

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	numericCard := 0
	switch card {
	case "ace":
		numericCard = 11
	case "two":
		numericCard = 2
	case "three":
		numericCard = 3
	case "four":
		numericCard = 4
	case "five":
		numericCard = 5
	case "six":
		numericCard = 6
	case "seven":
		numericCard = 7
	case "eight":
		numericCard = 8
	case "nine":
		numericCard = 9
	case "ten":
		numericCard = 10
	case "jack":
		numericCard = 10
	case "queen":
		numericCard = 10
	case "king":
		numericCard = 10
	default:
		numericCard = 0

	}

	return numericCard
}

const (
	Stand            = "S"
	Hit              = "H"
	Split            = "P"
	AutomaticallyWin = "W"
)

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumOfTwoCard := ParseCard(card1) + ParseCard(card2)
	decision := ""

	switch {
	case card1 == "ace" && card2 == "ace":
		decision = Split
	case sumOfTwoCard == 21 && ParseCard(dealerCard) < 10:
		decision = AutomaticallyWin
	case sumOfTwoCard == 21 && ParseCard(dealerCard) >= 2:
		decision = Stand
	case sumOfTwoCard <= 20 && sumOfTwoCard >= 17:
		decision = Stand
	case sumOfTwoCard <= 16 && sumOfTwoCard >= 12 && ParseCard(dealerCard) < 7:
		decision = Stand
	case sumOfTwoCard <= 16 && sumOfTwoCard >= 12 && ParseCard(dealerCard) >= 7:
		decision = Hit
	case sumOfTwoCard <= 11:
		decision = Hit
	default:
		decision = "0"
	}

	return decision
}

/*
FirstTurn("ace", "ace", "jack") == "P"
FirstTurn("ace", "king", "ace") == "S"
FirstTurn("five", "queen", "ace") == "H"

Stand (S)
Hit (H)
Split (P)
Automatically win (W)

*/
/**
Если у вас есть пара тузов, вы всегда должны разделить их.

Если у вас есть блэкджек (две карты, сумма которых равна 21), а у дилера нет ни туза, ни цифры, ни десятки,
вы автоматически выигрываете. Если у дилера есть какая-либо из этих карт, вам придется стоять и ждать, пока откроется другая карта.

Если сумма ваших карт равна значению в диапазоне [17, 20], вы всегда должны стоять.

Если сумма ваших карт составляет значение в диапазоне [12, 16], вы всегда должны стоять, если у дилера нет 7 или выше,
и в этом случае вы всегда должны брать.

Если сумма ваших карт равна 11 или ниже, вы всегда должны брать.
*/

func main() {
	fmt.Println(FirstTurn("ace", "ace", "jack"))   // == "P"
	fmt.Println(FirstTurn("ace", "king", "ace"))   //  == "S"
	fmt.Println(FirstTurn("five", "queen", "ace")) // == "H"
	fmt.Println(FirstTurn("king", "ace", "queen")) // == "S"
}

/*Для приблизительной оценки предположим, что если транспортному средству меньше 3 лет,
оно стоит 80% от первоначальной цены, когда оно было совершенно новым. Если ему не менее 10 лет,
он стоит 50%. Если транспортному средству не менее 3 лет, но менее 10 лет, оно стоит 70% от первоначальной цены.*/
