package utils

import (
  "testing"
)

func TestGetKeys(t *testing.T) {
  m := make(map[string]bool)
  m["foo"] = true
  m["bar"] = true
  
  keys := GetKeys(m)
  if len(keys) != 2 {
    t.Error("Unexpected length of keys array", len(keys))
  }
}
