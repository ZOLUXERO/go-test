For loops, iterando sobre slices
```go
func main()  {
	for i := 0; i <= 5; i++ {
		log.Println(i)	
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	// (_) Blank identifier, basicamente le dice al compiler que no
	// me interesa el primer valor que es retornado.
	for _, animal := range animals {
		log.Println(animal)
	}
}
```
Iterando sobre maps
```go
func main()  {
	animals := make(map[string]string)
	animals["dog"] = "Juanete"
	animals["cat"] = "Juaneta"
	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}
}
```
Iterando sobre strings
```go
func main()  {
	// Un string es un slice de bytes
	var firstLine = "hola amigos"
	for i, value := range firstLine {
		log.Println(i, value)
	}
}
```
Iterando sobre estructuras definidas por nosotros
```go
func main()  {
	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"Juanete", "De las Nieves", "juanete@gmail.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@gmail.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@gmail.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@gmail.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
```
