package main

import "fmt"

var age int = 20
var name string = "Karol"

func main() {
	fmt.Println("Hello Go!")

	city := "Gdańsk"
	height := 1.83

	fmt.Printf("Cześć jestem %s, mam %d lat i %v wzrostu. Mieszkam w %s", name, age, height, city)

	if age >= 18 {
		fmt.Println("Jesteś pełnoletni")
	} else {
		missing_age := 18 - age
		fmt.Printf("Nie jesteś pełnoletni i brakuje Ci %v", missing_age)
	}

	programming_languages := []string{"Python", "Javascript", "HTML", "CSS", "Go", "C++", "C"}

	for index, element := range programming_languages {
		fmt.Println(index+1, element)
	}

	current_language := "Go"
	switch current_language {
	case programming_languages[0]:
		fmt.Printf("Programujesz w %v", programming_languages[0])
	case programming_languages[1]:
		fmt.Printf("Programujesz w %v", programming_languages[1])
	case programming_languages[2]:
		fmt.Printf("Programujesz w %v", programming_languages[2])
	case programming_languages[3]:
		fmt.Printf("Programujesz w %v", programming_languages[3])
	case programming_languages[4]:
		fmt.Printf("Programujesz w %v", programming_languages[4])
	}
}