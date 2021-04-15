package adviceslip

import (
	"testing"
)

var topicTestCase string

func TestSearchAdviceSlips(t *testing.T) {

	topicTestCase = "cars"

	advice, err := SearchAdviceSlips(topicTestCase)
	if err != nil {
		t.Fatal(err)
	}

	if len(*advice) < 1 {
		t.Fatal("No advice")
	}
}
