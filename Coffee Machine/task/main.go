package main

import "fmt"

var water, milk, coffee, cups, money int = 400, 540, 120, 9, 550

func main() {
	//makeCoffee()
	//calcIngredient(cups)
	//maxCupFinder()
	//printState()
	machineAction()
}

func machineAction() {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	var action string
	fmt.Scan(&action)
	switch action {
	case "remaining":
		printState()
	case "buy":
		buy()
	case "fill":
		fill()
	case "take":
		take()
	case "exit":
		return
	default:
		fmt.Println("Invalid action")
	}
	machineAction()
}

func fill() {
	water += fillItem("Write how many ml of water you want to add:")
	milk += fillItem("Write how many ml of milk you want to add:")
	coffee += fillItem("Write how many grams of coffee beans you want to add:")
	cups += fillItem("Write how many disposable coffee cups you want to add:")
}

func fillItem(message string) int {
	var item int
	fmt.Println(message)
	fmt.Scan(&item)
	return item
}

func take() {
	fmt.Printf("I gave you $%d\n", money)
	money = 0
}

func buy() {
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	var drink string
	fmt.Scan(&drink)
	switch drink {
	case "1":
		buyDrink(250, 16, 0, 4)
	case "2":
		buyDrink(350, 20, 75, 7)
	case "3":
		buyDrink(200, 12, 100, 6)
	case "back":
		return
	default:
		fmt.Println("Invalid Drink; Please select again")
		buy()
	}
}
func buyDrink(waterP, coffeeP, milkP, moneyP int) {
	if waterP > water {
		fmt.Println("Sorry, not enough water!")
	} else if coffeeP > coffee {
		fmt.Println("Sorry, not enough coffee!")
	} else if milkP > milk {
		fmt.Println("Sorry, not enough milk!")
	} else if cups == 0 {
		fmt.Println("Sorry, not enough cup!")
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
		water -= waterP
		coffee -= coffeeP
		milk -= milkP
		cups -= 1
		money += moneyP
	}
}

func printState() {
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d of water\n", water)
	fmt.Printf("%d of milk\n", milk)
	fmt.Printf("%d of coffee beans\n", coffee)
	fmt.Printf("%d of disposable cups\n", cups)
	fmt.Printf("%d of money\n\n", money)
}

func maxCupFinder() {
	var water, milk, coffee, cups int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&coffee)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)
	maxCups := findMaximumCups(water, milk, coffee)
	if cups == maxCups {
		fmt.Println("Yes, I can make that amount of coffee")
	} else if maxCups < cups {
		fmt.Printf("No, I can make only %d cups of coffee\n", maxCups)
	} else {
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", maxCups-cups)
	}
}

func findMaximumCups(water, milk, coffee int) int {
	maxWaterCup := water / 200
	maxMilkCup := milk / 50
	maxCoffeeCup := coffee / 15

	return findMin(maxWaterCup, maxMilkCup, maxCoffeeCup)
}
func findMin(maxCups ...int) int {
	var min = maxCups[0]
	for _, max := range maxCups {
		if max < min {
			min = max
		}
	}
	return min
}

func calcIngredient(cups int) {
	fmt.Printf("For %d cups of coffee you will need:\n", cups)
	fmt.Printf("%d ml of water\n", 200*cups)
	fmt.Printf("%d ml of milk\n", 50*cups)
	fmt.Printf("%d g of coffee beans\n", 15*cups)
}

func makeCoffee() {
	fmt.Println("Starting to make a coffee")
	fmt.Println("Grinding coffee beans")
	fmt.Println("Boiling water")
	fmt.Println("Mixing boiled water with crushed coffee beans")
	pourIntoCup("coffee")
	pourIntoCup("some milk")
	fmt.Println("Coffee is ready!")
}

func pourIntoCup(item string) {
	fmt.Printf("Pouring %s into the cup\n", item)
}
