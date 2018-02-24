// Package bob simulates making a remark to a surly teenager
// and getting his response
package bob

import "strings"
import "regexp"

// Hey accepts a remark as a string and return a response as a string
func Hey(rawRemark string) string {
  remark := removeWhitespace(rawRemark)

  if isSilence(remark) {
    return "Fine. Be that way!"
  }

  if isShouting(remark) && isQuestion(remark) {
    return "Calm down, I know what I'm doing!"
  }

  if isShouting(remark) {
    return "Whoa, chill out!"
  }

  if isQuestion(remark) {
    return "Sure."
  }

  return "Whatever."
}

func isShouting(remark string) bool {
  return containsLetters(remark) && strings.ToUpper(remark) == remark
}

func isSilence(remark string) bool {
  return len(remark) == 0
}

func isQuestion(remark string) bool {
  return strings.HasSuffix(remark, "?")
}

func removeWhitespace(remark string) string {
  return regexp.MustCompile(`\s`).ReplaceAllString(remark, "")
}

func containsLetters(remark string) bool {
  return regexp.MustCompile(`[a-zA-Z]+`).MatchString(remark)
}
