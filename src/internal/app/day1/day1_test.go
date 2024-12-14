package day1_test

import (
	"testing"

	"github.com/bhvalent/adventOfCode2024/src/internal/app/day1"
)

func TestGetTotalDistance_ShouldReturnCorrectAnswer_WithTestData(t *testing.T) {
	result := day1.GetTotalDistance("input_test_data.txt")
	expected := 11
	if result != expected {
		t.Errorf("GetTotalDistance() = %d; Expected: %d", result, expected)
	}
}

func TestGetTotalDistance_ShouldReturnCorrectAnswer_WithRealData(t *testing.T) {
	result := day1.GetTotalDistance("input_data.txt")
	expected := 2057374
	if result != expected {
		t.Errorf("GetTotalDistance() = %d; Expected: %d", result, expected)
	}
}
