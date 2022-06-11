package clients

import (
	"testing"
)

// Tests the formatting of serial numbers
func TestUserInput_GetSerialNumber(t *testing.T) {
	testCases := []struct{
		name string
		input UserInput
		expected string
	}{
		{name: "SerialNumberValid",input: UserInput{AwsAccount: "13423425",Username: "mark.teasdale"},expected: "arn:aws:iam::13423425:mfa/mark.teasdale"},
	}
	for _, tc := range testCases {
		actual := tc.input.GetSerialNumber()
		if *actual != tc.expected {
			t.Errorf("Test Failed expected %s got %s", *actual,tc.expected)
		}
	}
}