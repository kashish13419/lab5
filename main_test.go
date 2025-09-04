package main

import "testing"

func TestAddition(t *testing.T) {
	tests := []struct{a,b,exp int}{
		{1,2,3},
		{5,7,12},
		{-1,1,0},
	}
	for _, tt := range tests {
		if got := tt.a + tt.b; got != tt.exp {
			t.Fatalf("expected %d got %d", tt.exp, got)
		}
	}
}
