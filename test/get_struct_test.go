package test

import (
	"testing"
)

func TestGetValue(t *testing.T) {
	pe :=Person{Age: 18,Name: "eric"}
	t.Run("success", func(t *testing.T) {
		GetValue(pe)
	})
}
