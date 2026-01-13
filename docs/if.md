Ejemplo de estructura de decision if en go
```go
func main()  {
	var isTrue bool
	isTrue = false
	if isTrue {
		log.Println("isTrue is", isTrue)	
	} else {
		log.Println("isTrue is", isTrue)	
	}

	myNum := 100
	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")	
	} else if myNum < 100 && isTrue {
		log.Println("1")	
	} else if myNum == 101 || isTrue {
		log.Println("2")	
	} else {
		log.Println("3")	
	}
}
```
