package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Please supply at least one argument.")
	}

	if args[0] == "brew" {
		brew()
	}
}

type Brew struct {
	Roaster     string
	GrindWeight float32
	Temperature Temperature
}

type Temperature int

const (
	Hot Temperature = iota
	Iced
)

func newBrew() Brew {
	return Brew{
		Temperature: Hot,
	}
}

func (b Brew) recommendedRecipe() Recipe {
	brewRatio := 17
	return Recipe{
		steps: []Step{
			{
				action: Bloom,
				time:   60,
				amount: int(b.GrindWeight * 2),
			},
			{
				action: Pour,
				time:   30,
				amount: int(b.GrindWeight * float32(brewRatio) / 2),
			},
			{
				action: Wait,
				time:   30,
			},
			{
				action: Pour,
				time:   30,
				amount: int(b.GrindWeight * float32(brewRatio) / 2),
			},
		},
	}
}

type Recipe struct {
	steps []Step
}

type Step struct {
	action ActionType
	time   int
	amount int
}

type ActionType int

const (
	Wait ActionType = iota
	Bloom
	Pour
	Swirl
)

func (at ActionType) String() string {
	return [...]string{"Wait", "Bloom", "Pour", "Swirl"}[at]
}

func (r Recipe) String() string {
	fmt.Println(r.steps)
	var s string
	for _, step := range r.steps {
		s += step.action.String()
		s += " for "
		s += string(step.time)

		s += " seconds."
		s += "\n"
	}
	return s
}

func brew() {
	brew := newBrew()
	fmt.Println("What is the brand/roaster of the coffee?")
	_, err := fmt.Scan(&brew.Roaster)
	if err != nil {
		fmt.Println("Error parsing roaster:", err)
		return
	}

	fmt.Println("What is the weight of the ground coffee in grams?")
	_, err = fmt.Scan(&brew.GrindWeight)
	if err != nil {
		fmt.Println("Error parsing weight:", err)
		return
	}

	fmt.Println(brew.recommendedRecipe())
}
