package main

import "testing"

func TestCountCalories(t *testing.T) {
	arr := []string{"1000", "2000", "3000", "", "4000"}
	ans := CountCalories(arr)
	if ans != 24000 {
		t.Errorf("CountCalories(arr) = %d; want 24000", ans)
	}
}
