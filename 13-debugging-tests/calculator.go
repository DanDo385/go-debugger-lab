package calculator

// Simple calculator functions for testing

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &ErrDivisionByZero{}
	}
	// ⚠️ BUG: Integer division loses remainder
	return a / b, nil
}

type ErrDivisionByZero struct{}

func (e *ErrDivisionByZero) Error() string {
	return "division by zero"
}

// FindMax returns the maximum value in a slice
func FindMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	// ⚠️ BUG: Off-by-one error (should start at i := 1)
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
