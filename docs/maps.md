Los maps son inmutables y funcionan muy parecido a un hashtable se podria decir con son la implementacion de hash tables en go, puede utilizar estos maps con estructuras personalizadas.
```go
type User struct {
	FirstName string
	LastName string
}

func main()  {
	myMap := make(map[string]User)
	me := User {
		FirstName: "Juanete",
		LastName: "De las Nieves",
	}
	myMap["me"] = me
	log.Println(myMap["me"].FirstName)
	var myNewVar float32
	myNewVar = 11.1
}
```
