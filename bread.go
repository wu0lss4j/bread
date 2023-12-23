// 21.12.2023 aqui vou eu experimentar golang
//
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("*************************************************")
	fmt.Println("Welcome to the best sourdough recipe in the World")
	fmt.Println("This will be good for ~1.7 kg bread sourdough")
	fmt.Println("This assumes you have already made your starter")
	fmt.Println("Starter is 50% water and 50% Rogen flour")
	fmt.Println("*************************************************")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nType how much starter you have in grams (default 97g)?")
	var starter float64
	scanner.Scan()
	starter, _ = strconv.ParseFloat(scanner.Text(), 64)

	switch starter {
	case 0.0:
		starter = 97.0
	}

	fmt.Println("Type desired water/solid ratio (default = 62%)?")
	var ratio float64
	scanner.Scan()
	ratio, _ = strconv.ParseFloat(scanner.Text(), 64)

	switch ratio {
	case 0.0:
		ratio = 0.62
	default:
		ratio = ratio / 100
	}

	salt := 0.02 * (1000 + 0.5*starter)
	solids := 1000 + 0.5*starter + salt
	water := ratio*solids - 0.5*starter

	fmt.Printf("Step 1: add %7.1f g starter\n", starter)
	fmt.Printf("Step 2: add %7.1f g salt\n", salt)
	fmt.Printf("Step 3: add %7.1f g water\n", water)
	fmt.Printf("Step 4: add %7.1f g Vollkorn Dinkel flour\n", 100.0)
	fmt.Printf("Step 5: add %7.1f g T1065 Dinkel flour\n", 250.0)
	fmt.Printf("Step 6: add %7.1f g T630 Dinkel flour\n", 650.0)
	fmt.Printf("\n ...now go follow the rest of the instructions!")
}
