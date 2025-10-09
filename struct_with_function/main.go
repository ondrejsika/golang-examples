package main

type Dog struct {
	Name         string
	BarkFunction func() string
}

func main() {
	dogs := []Dog{
		{
			Name: "Dela",
		},
		{
			Name: "Rex",
			BarkFunction: func() string {
				return "Woof!"
			},
		},
	}

	for _, dog := range dogs {
		if dog.BarkFunction != nil {
			println(dog.Name + " says: " + dog.BarkFunction())
		} else {
			println(dog.Name + " is silent.")
		}
	}
}
