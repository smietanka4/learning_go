package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var imie string
	var wiek int

	fmt.Println("Podaj swoje imie: ")
	fmt.Scanln(&imie)

	fmt.Println("Podaj swój wiek: ")
	fmt.Scanln(&wiek)

	if wiek >= 18 {
		fmt.Printf("Cześć %s, jesteś pełnoletni i masz %d lat", imie, wiek)
	} else {
		missing_age := 18 - wiek
		fmt.Printf("Cześć %s, niestety nie jesteś pełnoletni i brakuje Ci %d", imie, missing_age)
	}

	second_way_of_input()
}

func second_way_of_input() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nNapisz jakieś zdanie: ")
	scanner.Scan()

	tekst := scanner.Text()
	fmt.Printf("Wpisałeś %s\n", tekst)
}