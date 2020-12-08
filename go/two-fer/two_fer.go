package twofer

import (
  "fmt"
  "strings"
)
// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
  name = strings.TrimSpace(name)
  if name != "" {
    return fmt.Sprintf("One for %s, one for me.", name)
  } 
  return "One for you, one for me."
}
