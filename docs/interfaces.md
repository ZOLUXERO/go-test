Interfaces en go con implemetacion de receivers
```go
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main()  {
	dog := Dog{
		Name: "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Chuck",
		Color: "Grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "Legs")
}

// Para utilizar receivers la mejor practica es usar pointers
func (d *Dog) Says() string {
	return "Woof"
}

// Para utilizar receivers la mejor practica es usar pointers
func (d *Dog) NumberOfLegs() int {
	return 4
}

// Para utilizar receivers la mejor practica es usar pointers
func (d *Gorilla) Says() string {
	return "Ugh"
}

// Para utilizar receivers la mejor practica es usar pointers
func (d *Gorilla) NumberOfLegs() int {
	return 2
}
```
