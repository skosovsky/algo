package main

import "testing"

func Test_findPN(t *testing.T) {
	tests := []struct {
		k  int
		m  int
		k2 int
		p2 int
		n2 int
		p1 int
		n1 int
	}{
		{89, 20, 41, 1, 11, 2, 3},
		{11, 1, 1, 1, 1, 0, 1},
		{3, 2, 2, 2, 1, -1, -1},
		{16, 2, 15, 2, 1, 2, 0},
		{17, 2, 15, 2, 1, 2, 0},
		{18, 2, 15, 2, 1, 2, 0},
		{19, 2, 15, 2, 1, 2, 0},
		{20, 2, 15, 2, 1, 2, 0},
		{21, 2, 15, 2, 1, 0, 0},
		{22, 2, 15, 2, 1, 0, 0},
		{23, 2, 15, 2, 1, 0, 0},
		{24, 2, 15, 2, 1, 0, 0},
		{25, 2, 15, 2, 1, 0, 0},
		{26, 2, 15, 2, 1, 0, 0},
		{27, 2, 15, 2, 1, 0, 0},
		{28, 2, 15, 2, 1, 0, 0},
		{29, 2, 15, 2, 1, 3, 0},
		{5, 20, 2, 1, 1, 1, 0},
		{20, 20, 2, 1, 1, 1, 0},
		{21, 20, 2, 1, 1, 1, 0},
		{753, 10, 1000, 1, 1, 1, 1},
		{10, 3, 50, 1, 50, -1, -1},
		{25, 3, 1, 1, 1, 0, 0},
		{25, 3, 1, 1, 1, 0, 0},
		{24, 3, 1, 1, 1, 0, 0},
		{23, 3, 1, 1, 1, 0, 0},
		{22, 3, 1, 1, 1, 0, 0},
		{21, 3, 1, 1, 1, 0, 0},
		{20, 3, 1, 1, 1, 0, 0},
		{19, 3, 1, 1, 1, 0, 0},
		{18, 3, 1, 1, 1, 0, 0},
		{17, 3, 1, 1, 1, 0, 0},
		{16, 3, 1, 1, 1, 0, 0},
		{15, 3, 1, 1, 1, 0, 0},
		{19, 3, 8, 1, 1, 1, 0},
		{19, 3, 7, 1, 1, 1, 0},
		{19, 3, 6, 1, 1, 0, 0},
		{18, 3, 6, 1, 1, 1, 0},
		{6, 3, 18, 1, 1, 1, 1},
		{3, 1, 9, 7, 3, -1, -1},
		{3, 1, 2, 1, 1, 0, 1},
		{2, 1, 1, 1, 1, 0, 1},
		{3, 2, 2, 1, 1, 1, 0},
		{2, 3, 1, 1, 1, 1, 0},
		{842887, 10, 163729, 24, 8, 123, 0},
		{20, 10, 4, 1, 5, -1, -1},
		{20, 10, 5, 1, 5, 2, 10},
		{11, 2, 4, 1, 2, 0, 2},
	}
	for _, tt := range tests {
		res1, res2 := findPN(tt.k, tt.m, tt.k2, tt.p2, tt.n2)
		if res1 != tt.p1 && res2 != tt.n1 {
			t.Errorf("apartOnFloor(%v, %v, %v, %v, %v) = %v, %v, expected %v, %v",
				tt.k, tt.m, tt.k2, tt.p2, tt.n2, res1, res2, tt.p1, tt.n1)
		}
	}
}
