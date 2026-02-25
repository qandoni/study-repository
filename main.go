package main

import (
	"Study/feature1"
	"Study/feature2"
	simpleconnection "Study/feature_postgres/simple_connection"
	"fmt"
)

func main() {
	fmt.Println("Hello, Git!")
	feature1.Feature1()
	feature2.Feature2()
	simpleconnection.CheckConnection()
}
