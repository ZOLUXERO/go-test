Todos los archivos `.go` siempre deben llevar `package <nombre-paquete>` al inicio de cada archivo.
Los paquetes son formas de agrupar codigo por achivo y sirven para ser llamados desde otros lugares.
```go
package helpers

type SomeType struct {
	TypeName string
	TypeNumber int
}

func PrintHelper() string  {
	return "hola helper"
}
```
Para usar estos paquetes debe importarlos de la siguiente manera.
```go
import (
	"github.com/ZOLUXERO/go-test/helpers"
)

func main()  {
	log.Println("Hello")
	var myVar helpers.SomeType
	log.Println(helpers.PrintHelper())
	myVar.TypeName = "MyVar"
	myVar.TypeNumber = 1
	log.Println(
		myVar.TypeName,
		myVar.TypeNumber,
	)
}
```
