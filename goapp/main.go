package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["google"] = 10
	m["microsoft"] = 20
	m["ubuntu"] = 30

	fmt.Println(m)
	fmt.Println("after deleting")

	if _, ok := m["microsoft"]; ok {
		delete(m, "microsoft")
	}
	fmt.Println(m)

	fmt.Println("after modifying")
	if _, ok := m["ubuntu"]; ok {
		m["ubuntu"] = (12)
	}
	fmt.Println(m)
}
