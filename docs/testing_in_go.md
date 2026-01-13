Los comandos para correr tests en go son los siguientes:
```bash
go test -v
```
Si quiere informacion mas detallada de que tiene coverage y que no puede usar:
```bash
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

Para testear en go debe crear un arhivo que se llame igual al que desea testear con el sufijo `_test`, en este caso queremos testear `main.go` asi que el archivo de tests se deberia llamar `main_test.go`.
Esto es un ejemplo!, la funcion main no se deberia testear!!!.
Archivo con logica que queremos testear.
```go
package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result of division is", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by 0")
	}

	result = x / y
	return result, nil
}
```
Ejemplo de test simplificado y el que se deberia usar.
```go
-- main_test.go
package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

```
Ejemplo de test manual y que debemos evitar.
```go
package main

import "testing"
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
```
