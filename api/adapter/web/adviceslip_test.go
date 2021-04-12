package adviceslip

import (
	"testing"
)

var topicTestCase string

func TestSearchAdviceSlips(t *testing.T) {

	topicTestCase = "cars"

	advices, err := SearchAdviceSlips(topicTestCase)
	if err != nil {
		t.Fatal(err)
	}

	if len(*advices) < 1 {
		t.Fatal("No advices")
	}
}
