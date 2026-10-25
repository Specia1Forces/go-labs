package main

import "testing"

func TestParallelSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "several numbers",
			numbers: []int{1, 2, 3, 4, 5, 6},
			want:    21,
		},
		{
			name:    "empty slice",
			numbers: []int{},
			want:    0,
		},
		{
			name:    "one number",
			numbers: []int{42},
			want:    42,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ParallelSum(test.numbers)
			if got != test.want {
				t.Fatalf("ParallelSum() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestRunCounter(t *testing.T) {
	const want = 5000

	got := RunCounter()
	if got != want {
		t.Fatalf("RunCounter() = %d, want %d", got, want)
	}
}
