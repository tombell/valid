package valid_test

import (
	"cmp"
	"regexp"
	"testing"

	"github.com/tombell/valid"
)

func TestEmpty(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{
			name:     "empty value",
			value:    "",
			expected: true,
		},
		{
			name:     "non-empty value",
			value:    "hello world",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Empty(tt.value); got != tt.expected {
				t.Errorf("Empty() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{
			name:     "empty value",
			value:    "",
			expected: false,
		},
		{
			name:     "non-empty value",
			value:    "hello world",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.NotEmpty(tt.value); got != tt.expected {
				t.Errorf("NotEmpty() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		prefix   string
		expected bool
	}{
		{
			name:     "starts with prefix",
			value:    "hello world",
			prefix:   "hello",
			expected: true,
		},
		{
			name:     "does not start with prefix",
			value:    "hello world",
			prefix:   "Hello",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.StartsWith(tt.value, tt.prefix); got != tt.expected {
				t.Errorf("StartsWith() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		suffix   string
		expected bool
	}{
		{
			name:     "ends with suffix",
			value:    "hello world",
			suffix:   "world",
			expected: true,
		},
		{
			name:     "does not end witn suffix",
			value:    "hello world",
			suffix:   "World",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.EndsWith(tt.value, tt.suffix); got != tt.expected {
				t.Errorf("EndsWith() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		max      int
		expected bool
	}{
		{
			name:     "less characters than max",
			value:    "asdf",
			max:      5,
			expected: true,
		},
		{
			name:     "same number of characters as max",
			value:    "asdf",
			max:      4,
			expected: true,
		},
		{
			name:     "more characters than max",
			value:    "asdfasdfasdf",
			max:      6,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.MaxLength(tt.value, tt.max); got != tt.expected {
				t.Errorf("MaxLength() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMinLength(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		min      int
		expected bool
	}{
		{
			name:     "less characters than min",
			value:    "asdf",
			min:      5,
			expected: false,
		},
		{
			name:     "same number of characters as min",
			value:    "asdf",
			min:      4,
			expected: true,
		},
		{
			name:     "more characters than min",
			value:    "asdfasdfasdf",
			min:      6,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.MinLength(tt.value, tt.min); got != tt.expected {
				t.Errorf("MinLength() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestRangeLength(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		min      int
		max      int
		expected bool
	}{
		{
			name:     "length between min and max",
			value:    "123",
			min:      1,
			max:      4,
			expected: true,
		},
		{
			name:     "length outside range",
			value:    "boom!",
			min:      5,
			max:      9,
			expected: true,
		},
		{
			name:     "length equals min",
			value:    "hello",
			min:      5,
			max:      8,
			expected: true,
		},
		{
			name:     "length equals max",
			value:    "some text",
			min:      3,
			max:      9,
			expected: true,
		},
		{
			name:     "utf8 text in range",
			value:    "やばい",
			min:      3,
			max:      6,
			expected: true,
		},
		{
			name:     "utf8 text outside range",
			value:    "やばい",
			min:      1,
			max:      2,
			expected: false,
		},
		{
			name:     "utf8 text equal min",
			value:    "やばい",
			min:      3,
			max:      10,
			expected: true,
		},
		{
			name:     "utf8 text equal max",
			value:    "やばい",
			min:      2,
			max:      3,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.RangeLength(tt.value, tt.min, tt.max); got != tt.expected {
				t.Errorf("RangeLength() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMatches(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		rx       *regexp.Regexp
		expected bool
	}{
		{
			name:     "numbers regex",
			value:    "5432",
			rx:       regexp.MustCompile(`[0-9]+`),
			expected: true,
		},
		{
			name:     "uppercase letters regex",
			value:    "HELLO WORLD",
			rx:       regexp.MustCompile(`^\P{L}*\p{Lu}\P{Ll}*$`),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Matches(tt.value, tt.rx); got != tt.expected {
				t.Errorf("Matches() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{
			name:     "valid number",
			value:    "1234567890",
			expected: true,
		},
		{
			name:     "invalid number",
			value:    "!p4ssw0rd",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.IsNumber(tt.value); got != tt.expected {
				t.Errorf("IsNumber() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type test[T cmp.Ordered] struct {
		name     string
		value    T
		max      T
		expected bool
	}

	tests := []test[int]{
		{
			name:     "less than max",
			value:    0,
			max:      5,
			expected: true,
		},
		{
			name:     "equal to max",
			value:    8,
			max:      8,
			expected: true,
		},
		{
			name:     "greater than max",
			value:    5,
			max:      3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Max(tt.value, tt.max); got != tt.expected {
				t.Errorf("Max() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type test[T cmp.Ordered] struct {
		name     string
		value    T
		min      T
		expected bool
	}

	tests := []test[int]{
		{
			name:     "less than min",
			value:    -1,
			min:      5,
			expected: false,
		},
		{
			name:     "equal to min",
			value:    91,
			min:      91,
			expected: true,
		},
		{
			name:     "greater than min",
			value:    12,
			min:      1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Min(tt.value, tt.min); got != tt.expected {
				t.Errorf("Min() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestRange(t *testing.T) {
	type test[T cmp.Ordered] struct {
		name  string
		value T
		min   T
		max   T
		want  bool
	}

	tests := []test[int]{
		{
			name:  "value between min and max",
			value: 5,
			min:   2,
			max:   7,
			want:  true,
		},
		{
			name:  "value outside range",
			value: 7,
			min:   11,
			max:   18,
			want:  false,
		},
		{
			name:  "value equals min",
			value: 7,
			min:   7,
			max:   30,
			want:  true,
		},
		{
			name:  "value equals max",
			value: -1,
			min:   -10,
			max:   -1,
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Range(tt.value, tt.min, tt.max); got != tt.want {
				t.Errorf("Range() = %v, expected %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type test[T cmp.Ordered] struct {
		name   string
		values []T
		want   bool
	}

	numberTests := []test[int]{
		{
			name:   "unique numbers",
			values: []int{30, 910, 0, -1},
			want:   true,
		},
		{
			name:   "non-unique numbers",
			values: []int{20, 7, 1, 70, 1},
			want:   false,
		},
	}

	stringTests := []test[string]{
		{
			name:   "unique strings",
			values: []string{"one", "two", "three"},
			want:   true,
		},
		{
			name:   "non-unique strings",
			values: []string{"one", "two", "three", "three"},
			want:   false,
		},
	}

	for _, tt := range numberTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Unique(tt.values); got != tt.want {
				t.Errorf("Unique() = %v, expected %v", got, tt.want)
			}
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.Unique(tt.values); got != tt.want {
				t.Errorf("Unique() = %v, expected %v", got, tt.want)
			}
		})
	}
}

func TestIn(t *testing.T) {
	type test[T comparable] struct {
		name  string
		value T
		list  []T
		want  bool
	}

	numberTests := []test[int]{
		{
			name:  "number exists in list",
			value: 5,
			list:  []int{4, 5, 6, 7, 8},
			want:  true,
		},
		{
			name:  "number does not exist in list",
			value: -5,
			list:  []int{0, 2, 0, 4},
			want:  false,
		},
	}
	stringTests := []test[string]{
		{
			name:  "string exists in list",
			value: "orange",
			list:  []string{"apple", "banana", "orange"},
			want:  true,
		},
		{
			name:  "string does not exist in list",
			value: "chicken",
			list:  []string{"apple", "banana", "orange"},
			want:  false,
		},
	}

	for _, tt := range numberTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.In(tt.value, tt.list...); got != tt.want {
				t.Errorf("In() = %v, expected %v", got, tt.want)
			}
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.In(tt.value, tt.list...); got != tt.want {
				t.Errorf("In() = %v, expected %v", got, tt.want)
			}
		})
	}
}

func TestIsDate(t *testing.T) {
	type test struct {
		name  string
		value string
		want  bool
	}

	tests := []test{
		{
			name:  "empty string",
			value: "",
			want:  false,
		},
		{
			name:  "invalid date string",
			value: "202aana",
			want:  false,
		},
		{
			name:  "valid date string",
			value: "2024-03-18T13:34:15Z",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid.IsDate(tt.value); got != tt.want {
				t.Errorf("IsDate() = %v, expected %v", got, tt.want)
			}
		})
	}
}
