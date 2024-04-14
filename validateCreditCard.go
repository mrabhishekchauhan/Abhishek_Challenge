package main

import (
	"fmt"
	"regexp"
	"strings"
)

// validateCreditCards validates a list of credit card numbers
func validateCreditCards(creditCards []string) {
	// Define regular expression to validate the structure of a credit card number
	validStructure := regexp.MustCompile(`[456]\d{3}(-?\d{4}){3}$`)
	// Create a list of filters containing the valid structure
	filters := []string{validStructure.String()}

	for _, cc := range creditCards {
		valid := true
		for _, f := range filters {
			// Check if the credit card number matches the valid structure
			// and does not contain four repeated digits consecutively
			if !regexp.MustCompile(f).MatchString(cc) || containsFourRepeats(cc) {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
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

	creditCards := make([]string, numCards)
	for i := 0; i < numCards; i++ {
		var card string
		fmt.Scanln(&card)
		// Trim leading and trailing whitespace from the credit card number
		creditCards[i] = strings.TrimSpace(card)
	}

	// Validate the credit card numbers
	validateCreditCards(creditCards)
}
