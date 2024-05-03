package main

import (
	"fmt"
    "bufio"
    "os"
	"regexp"
	"strings"
)

// validateCreditCards validates a list of credit card numbers
func validateCreditCards(cc string) {
	// Define regular expression to validate the structure of a credit card number
	validStructure := regexp.MustCompile(`^[456]\d{3}(-?\d{4}){3}$`)
    if validStructure.MatchString(cc) && !containsFourRepeats(cc) {
        fmt.Println("Valid")
    } else {
        fmt.Println("Invalid")
    }
}
// containsFourRepeats checks if a credit card number contains four repeated digits consecutively
func containsFourRepeats(cc string) bool {
	// Remove hyphens from the credit card number
	cc = strings.ReplaceAll(cc, "-", "")
	for i := 0; i < len(cc)-3; i++ {
		// Check if there are four consecutive repeated digits
		if cc[i] == cc[i+1] && cc[i] == cc[i+2] && cc[i] == cc[i+3] {
			return true
		}
	}
	return false
}

func main() {
	var numCards int
	fmt.Scanln(&numCards)

	scanner := bufio.NewScanner(os.Stdin)
	//var creditCards []string
	for scanner.Scan() {
		line := scanner.Text()
        // Validate the credit card numbers
        validateCreditCards(line)
	}

}
