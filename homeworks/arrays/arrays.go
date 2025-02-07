package arrays

import "fmt"

const (
	YES = "YES"
	NO  = "NO"
)

func consoleArrays() {
	var nA int
	fmt.Scanf("%d", &nA)
	var nB int
	fmt.Scanf("%d", &nB)
	var k int
	fmt.Scanf("%d", &k)
	var m int
	fmt.Scanf("%d", &m)
	numsA := make([]int, 0, nA)
	var temp int
	for i := 0; i < nA; i++ {
		fmt.Scanf("%d", &temp)
		numsA = append(numsA, temp)
	}
	numsB := make([]int, 0, nB)
	for i := 0; i < nB; i++ {
		fmt.Scanf("%d", &temp)
		numsB = append(numsB, temp)
	}
	fmt.Println(arrays(nA, nB, k, m, numsA, numsB))
}

func arrays(nA, nB, k, m int, numsA, numsB []int) string {
	numK := numsA[k-1]
	if numsB[nB-m] <= numK {
		return NO
	}
	return YES
}

func testArrays() {
	type args struct {
		nA    int
		nB    int
		k     int
		m     int
		numsA []int
		numsB []int
	}

	type test struct {
		args args
		want string
	}

	tests := []test{
		{
			args: args{
				nA:    3,
				nB:    3,
				k:     2,
				m:     1,
				numsA: []int{1, 2, 3},
				numsB: []int{3, 4, 5},
			},
			want: YES,
		},
		{
			args: args{
				nA:    3,
				nB:    3,
				k:     3,
				m:     3,
				numsA: []int{1, 2, 3},
				numsB: []int{3, 4, 5},
			},
			want: NO,
		},
		{
			args: args{
				nA:    5,
				nB:    2,
				k:     3,
				m:     1,
				numsA: []int{1, 1, 1, 1, 1},
				numsB: []int{2, 2},
			},
			want: YES,
		},
		{
			args: args{
				nA:    4,
				nB:    4,
				k:     2,
				m:     2,
				numsA: []int{3, 4, 5, 6},
				numsB: []int{3, 4, 5, 6},
			},
			want: YES,
		},
	}
	for _, v := range tests {
		if out := arrays(v.args.nA, v.args.nB, v.args.k, v.args.m, v.args.numsA, v.args.numsB); out != v.want {
			fmt.Printf("args: %+v - out: %s / expected: %s\n", v.args, out, v.want)
		}
	}
}
