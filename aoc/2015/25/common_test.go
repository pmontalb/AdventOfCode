package main

import "testing"

func TestFirstRow(t *testing.T) {
	i := 0
	for j := 0; j < 100; j++ {
		act := GetLinearizedIndex(i, j)
		exp := GetCumSum(j+1) - 1
		if act != exp {
			t.Errorf("got(%d) want(%d)", act, exp)
		}
	}
}

func TestSecondRow(t *testing.T) {
	i := 1
	for j := 0; j < 100; j++ {
		act := GetLinearizedIndex(i, j)
		exp := 0
		if j > 0 {
			exp = GetLinearizedIndex(i, j-1) + j + 2
		} else {
			exp = GetLinearizedIndex(i, j-1)
		}
		if act != exp {
			t.Errorf("got(%d) want(%d)", act, exp)
		}
	}
}

func TestThirdRow(t *testing.T) {
	i := 2
	for j := 0; j < 100; j++ {
		act := GetLinearizedIndex(i, j)
		exp := 0
		if j > 0 {
			exp = GetLinearizedIndex(i, j-1) + j + 3
		} else {
			exp = GetLinearizedIndex(i, j-1)
		}
		if act != exp {
			t.Errorf("got(%d) want(%d)", act, exp)
		}
	}
}
