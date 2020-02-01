package main

const (
	m = 252533
	q = 33554393
)

/* j
	  0   1  2  3  4  5 ...

i  0  0   2  5  9 14 20 ...
   1  1   4  8 13 19 26 ...
   2  3   7 12 18 25 ...
   3  6  11 17 24 ...
   4  10 16 23 ...
	  ...

Denoting I_{i,j} the linear index in this table, it holds:;
	I_{0, j} = I_{0, j-1} + j+1 (as the first row it's just the previous plus the number of entries in the diagonal)
from which, using I_{0, 0} = 0
	I_{0, j} = sum_{k = 0}^j j+1 = j + (j * (j + 1)) / 2
Furthermore
	I_{i, j} = I_{i, 0} + I_{i, j-1} + j+1 (as the i-th row "starts" from I_{i, 0})

Now we just need to find an expression for I_{i, 0}, which along the lines of I_{0, j}, reads
	I_{i, 0} = I_{i-1, 0} + i
from which, using I_{0, 0} = 0
	I_{i, 0} = (i * (i + 1)) / 2

Finally:
	I_{i, j} = [(i * (i + 1)) / 2] + [j + (j * (j + 1)) / 2]
*/
func GetCumSum(n int) int {
	return (n * (n + 1)) / 2
}
func getRangedCumSum(from, to int) int {
	if to < from {
		return 0
	}
	return GetCumSum(to) - GetCumSum(from)
}
func GetLinearizedIndex(i, j int) int {
	I_i0 := GetCumSum(i)
	I_0j := getRangedCumSum(i+1, i+j+1)
	return I_i0 + I_0j
}

func GetNextCode(n int) int {
	return (n * m) % q
}
