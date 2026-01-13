TLDR: basicamente son tipos que nosotros podemos crear.
Las estructuras son agrupamientos de variables de diferentes tipos bajo un solo nombre, este concepto es muy parecido a como lo maneja `C`. Estas estructuras puede tener funciones asociadas a ellas como lo ven en el ejemplo.
```go
var s = "seven"

type User struct {
	FirstNAme string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

type myStuct struct {
	FirstNAme string
}

// Funcion asociada de la estructura myStruct
func (m *myStuct) printFirstName() string {
	return m.FirstNAme
}

func main()  {
	user := User {
		FirstNAme: "Juanete",
		LastName: "De las nieves",
	}

	log.Println(user.FirstNAme, user.LastName, user.BirthDate)

	var myVar myStuct
	myVar.FirstNAme = "john"

	myVar2 := myStuct{
		FirstNAme: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
```
