package calculator

import "testing"

// Table-driven test
// 🔍 SET BREAKPOINT HERE — Inside test cases
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		// 🔍 SET CONDITIONAL BREAKPOINT: tt.name == "negative numbers"
		t.Run(tt.name, func(t *testing.T) {
			// 👀 Watch tt.a, tt.b, tt.expected in the Variables panel
			result := Add(tt.a, tt.b)

			// 🔍 SET BREAKPOINT HERE — Inspect result before assertion
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Subtests
func TestDivide(t *testing.T) {
	t.Run("normal division", func(t *testing.T) {
		// 🔍 SET BREAKPOINT HERE
		result, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// 👀 Watch the result
		if result != 5 {
			t.Errorf("Divide(10, 2) = %d; want 5", result)
		}
	})

	t.Run("division by zero", func(t *testing.T) {
		// 🔍 SET BREAKPOINT HERE
		_, err := Divide(10, 0)

		// 👀 Inspect the error
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("division with remainder", func(t *testing.T) {
		// ⚠️ THIS TEST WILL FAIL (integer division loses remainder)
		// 🔍 SET BREAKPOINT HERE
		result, err := Divide(10, 3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// This will fail: 10/3 = 3 (integer division), not 3.33
		if result != 3 {
			t.Errorf("Divide(10, 3) = %d; want 3", result)
		}
	})
}

// Test with a bug
func TestFindMax(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"multiple values", []int{1, 5, 3, 9, 2}, 9},
		{"single value", []int{42}, 42},
		{"negative values", []int{-5, -1, -10}, -1},
		{"all same", []int{7, 7, 7}, 7},
		// ⚠️ THIS TEST WILL PASS but there's still a bug
		{"empty slice", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 🔍 SET BREAKPOINT HERE — Step into FindMax
			result := FindMax(tt.nums)

			// 🔍 SET BREAKPOINT HERE — Before assertion
			if result != tt.expected {
				t.Errorf("FindMax(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}

// Benchmark (can also be debugged)
func BenchmarkAdd(b *testing.B) {
	// 🔍 SET BREAKPOINT HERE — Will hit b.N times
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

// Test helper function
func assertEqual(t *testing.T, got, want int) {
	t.Helper() // Marks this as a helper function
	// 🔍 SET BREAKPOINT HERE
	if got != want {
		// 👀 When this fails, the stack trace points to the caller, not here
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestWithHelper(t *testing.T) {
	result := Add(2, 3)
	// 🔍 SET BREAKPOINT HERE
	assertEqual(t, result, 5)
}
