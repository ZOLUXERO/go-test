Los receivers son una forma de declarar funciones que adiciona una funcion a una estructura.
```go
type myStuct struct {
	FirstNAme string
}

// Esto es un receiver, funcion asociada de la estructura myStruct
func (m *myStuct) printFirstName() string {
	return m.FirstNAme
}

func main()  {
	var myVar myStuct
	myVar.FirstNAme = "john"

	myVar2 := myStuct{
		FirstNAme: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
```
