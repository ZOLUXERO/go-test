Ejemplo de canales en go, los canales en go son pipes que conectan `goroutines` concurrentes, lo puede utilizar para sincronizacion y comunicacion entre varios `goroutines` en diferentes paquetes.
```go
-- main.go
const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber // Asigne el valor al canal que paso como argumento
}

func main() {
	// crear un canal que solo puede contener enteros, cuando termine de correr el canal cierrelo inmediatamente con `defer`
	intChan := make(chan int)
	defer close(intChan)

	// GoRoutine
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
```
```go
-- helpers.go
package helpers

import "math/rand"

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}
```
