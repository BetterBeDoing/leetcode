package leetcode

import "strings"

func isPalindrome(s string) bool {
	// Remove non-alphanumeric characters and convert to lowercase
	s = strings.ToLower(s)
	s = removeNonAlphaNumeric(s)

	// Check if the string is a palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

// Helper function to remove non-alphanumeric characters from a string
func removeNonAlphaNumeric(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
