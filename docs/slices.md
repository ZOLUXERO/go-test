Un slice es una estructura de datos que es parecida y funciona como un array, pero **NO es un array**, generalemte los slices una tamaño predefinido, pero se le pueden agregar items con `append()`.
```go
func main()  {
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)
}
```
Shorthand de slices
```go
func main()  {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)
	log.Println(numbers[6:9])

    names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
```
