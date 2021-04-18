package models

import "errors"

// AdviceArgs contains all the parameters used in the GiveMeAdvice method
type AdviceRequestArgs struct {
	Topic  string `json:"topic"`
	Amount *int   `json:"amount"`
}

// AdviceResponse struct represents an advice response
type AdviceResponse struct {
	AdviceList []string `json:"adviceList"`
}

func (adr *AdviceRequestArgs) Validate() error {

	if isBlankString(&adr.Topic) {
		return errors.New("please provide a topic")
	}

	if len(adr.Topic) < 3 {
		return errors.New("please provide a keyword with at least 3 characters as a topic")
	}

	if isNegativeInt(adr.Amount) {
		return errors.New("can't perform requests for negative amount")
	}

	if isZeroInt(adr.Amount) {
		return errors.New("can't perform requests for zero amount")
	}

	return nil

}

// If input variable is empty string, returns true
func isBlankString(s *string) bool {
	return s != nil && *s == ""
}

// If input variable is negative, returns true
func isNegativeInt(n *int) bool {
	return n != nil && *n < 0
}

// If input variable is zero, returns true
func isZeroInt(n *int) bool {
	return n != nil && *n == 0
}
