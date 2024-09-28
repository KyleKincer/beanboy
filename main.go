package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Please supply at least one argument.")
		return
	}

	displayBean()

	if args[0] == "brew" {
		brew()
	}
}

func displayBean() {
	fmt.Println(`
    Welcome to beanboy! 🤠

                                                                  
                              ▒▒▓▓▒▒                              
                            ▒▒░░░░░░▓▓                            
                            ▓▓░░░░░░▒▒░░                          
            ▒▒▓▓▒▒▓▓▓▓      ▓▓░░░░░░  ▓▓      ▓▓██▓▓              
            ▓▓░░░░░░░░▓▓▒▒  ▓▓░░░░░░  ██  ░░▓▓░░░░░░▓▓            
            ▒▒▒▒▓▓░░░░░░▒▒██▓▓▒▒░░▒▒▒▒▓▓▓▓░░░░░░░░░░▓▓            
            ██░░░░░░  ░░░░░░  ░░░░░░░░  ░░░░░░░░░░▒▒░░░░          
            ░░▒▒░░▓▓▓▓▒▒▒▒▒▒▓▓░░░░▓▓▒▒░░  ▒▒▒▒░░▓▓  ░░░░          
              ▓▓▒▒░░░░░░▓▓░░░░░░██░░░░▓▓▒▒░░░░▓▓░░░░▓▓            
                ▒▒▓▓░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒██░░            
                    ▒▒▓▓██▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒░░                
              ▒▒▒▒      ██▒▒▓▓▒▒▒▒▒▒▓▓▓▓▓▓    ░░▓▓▒▒▒▒░░          
            ░░░░▒▒░░    ▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓    ▒▒▓▓▓▓██            
        ░░▒▒  ▒▒▒▒▒▒    ▓▓▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓    ▒▒▒▒▒▒▒▒            
          ░░▒▒▒▒▓▓▓▓    ▓▓▒▒▓▓        ▓▓▓▓    ▓▓▓▓░░              
            ░░▒▒▒▒▓▓▓▓░░▓▓▒▒▒▒██▒▒▒▒▓▓▒▒▓▓  ▒▒▓▓                  
                    ▓▓▒▒▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓██▒▒                    
                      ▒▒▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒                      
                        ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒██                      
                        ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░                    
                      ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒                    
                      ░░▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓                    
                        ▓▓▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▒▒▒▒                    
                        ▓▓▒▒▒▒▒▒▒▒▒▒▓▓▒▒▒▒▓▓░░                    
                          ██▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓                      
                            ▓▓▓▓▓▓▒▒▒▒▓▓██                        
                        ░░░░▒▒░░▒▒▓▓▓▓▓▓░░░░                      
                ░░▒▒▒▒▒▒▒▒▒▒▓▓▓▓▒▒▒▒▒▒▓▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒            
                  ░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒              
                            ░░░░░░░░░░░░░░          ░░            
                                          ░░  ░░░░░░  ░░  ░░░░░░  
                                              ░░░░                

    `)
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
				time:   45,
				amount: int(b.GrindWeight * 2),
			},
			{
				action: Pour,
				time:   15,
				amount: int(b.GrindWeight * float32(brewRatio) / 2),
			},
			{
				action: Wait,
				time:   45,
			},
			{
				action: Pour,
				time:   15,
				amount: int(b.GrindWeight * float32(brewRatio) / 2),
			},
			{
				action: Wait,
				time:   45,
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

func (s Step) String() string {
	o := s.action.String()
	if s.amount > 0 {
		o += " (" + strconv.Itoa(s.amount) + "g)"
	}
	o += " for " + strconv.Itoa(s.time) + " seconds."
	return o
}

func (r Recipe) String() string {
	var s string
	for i, step := range r.steps {
		s += step.String()
		if i < len(r.steps)-1 {
			s += "\n"
		}
	}
	return s
}

func (r Recipe) execute() {
	for _, step := range r.steps {
		fmt.Print(step.action)
		if step.amount > 0 {
			fmt.Printf(" (%dg)", step.amount)
		}
		fmt.Println(" for", step.time, "seconds.")
		ticker := time.NewTicker(time.Second * 1)

		seconds := step.time
		for range ticker.C {
			fmt.Printf("\r%d seconds", seconds)
			seconds--
			if seconds < 0 {
				ticker.Stop()
				break
			}
		}
		fmt.Println()
	}
	fmt.Println("Done! Enjoy your coffee! ☕️")
}

func (r Recipe) record() {
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

	recipe := brew.recommendedRecipe()

	fmt.Println("Recommended recipe is as follows:")
	fmt.Println(recipe)
	fmt.Print("Continue? (y/n)")

	var r string
	_, err = fmt.Scan(&r)
	if err != nil {
		fmt.Print("Error parsing input:", err)
	}

	if r != "y" && r != "Y" {
		return
	}

	recipe.execute()
}
