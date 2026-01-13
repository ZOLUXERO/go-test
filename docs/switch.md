Ejemplo estructura de decision switch en go
```go
func main()  {
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("Cat is set to cat")	
	case "dog":
		log.Println("Dog is set to dog")	
	case "fish":
		log.Println("Fish is set to fish")	
	default:
		log.Println("This is default")	
	}
}
```
