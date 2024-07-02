package main

import "testing"

func TestRunWriter(t *testing.T) {
	prices := RunWriter()

	expected := []map[string]float64{
		{"inst1": 2.1, "inst2": 3.1, "inst3": 4.1, "inst4": 5.1},
		{"inst1": 3.1, "inst2": 4.1, "inst3": 5.1, "inst4": 6.1},
		{"inst1": 4.1, "inst2": 5.1, "inst3": 6.1, "inst4": 7.1},
		{"inst1": 5.1, "inst2": 6.1, "inst3": 7.1, "inst4": 8.1},
	}

	var results []map[string]float64
	for price := range prices {
		results = append(results, price)
	}

	if len(results) != len(expected) {
		t.Fatalf("expected %d prices, got %d", len(expected), len(results))
	}

	for i, result := range results {
		for k, v := range expected[i] {
			if result[k] != v {
				t.Fatalf("expected %v for key %s, got %v", v, k, result[k])
			}
		}
	}
}
