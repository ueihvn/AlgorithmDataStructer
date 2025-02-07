package arrays

import "fmt"

const (
	timeEnterPassword = 1
	timeDelayRetry    = 5
)

func consolePasswords() {
	var n int
	fmt.Scanf("%d", &n)
	var k int
	fmt.Scanf("%d", &k)
	var temp string
	pass := make([]string, 0, n+1)
	for i := 0; i <= n; i++ {
		fmt.Scanf("%s", &temp)
		pass = append(pass, temp)
	}
	fmt.Println(passwords(n, k, pass))
}

func passwords(n, k int, pass []string) (int, int) {
	// move to count smaller and same instead of whole array
	lenOfRightPass := len(pass[n])
	// ignore the right password and the password
	var smaller, equal int
	for i := 0; i < n; i++ {
		if len(pass[i]) < lenOfRightPass {
			smaller++
		} else if len(pass[i]) == lenOfRightPass && pass[i] != pass[n] {
			equal++
		}
	}
	// remove the correct password from worst
	return calculateWrongAttemptTime(k, smaller) + timeEnterPassword, calculateWrongAttemptTime(k, smaller+equal) + timeEnterPassword
}

func calculateWrongAttemptTime(k, pass int) int {
	return pass*timeEnterPassword + (pass/k)*timeDelayRetry
}

func testPasswords() {
	type args struct {
		n    int
		k    int
		pass []string
	}
	type test struct {
		args  args
		best  int
		worst int
	}

	tests := []test{
		{
			args: args{
				n:    5,
				k:    5,
				pass: []string{"cba", "abc", "bb1", "abC", "ABC", "abc"},
			},
			best:  1,
			worst: 5,
		},
		{
			args: args{
				n:    5,
				k:    2,
				pass: []string{"cba", "abc", "bb1", "abC", "ABC", "abc"},
			},
			best:  1,
			worst: 15,
		},
		{
			args: args{
				n:    4,
				k:    100,
				pass: []string{"11", "22", "1", "2", "22"},
			},
			best:  3,
			worst: 4,
		},
		{
			args: args{
				n: 20,
				k: 5,
				pass: []string{
					"pine",
					"kqdhmw",
					"ufkszbursb",
					"l",
					"htalezfiosdepsgmiu",
					"v",
					"fkzfpno",
					"lrscyyhev",
					"ffaihnj",
					"omvcpnncreznp",
					"vnmydarmeqa",
					"bzjoonknqchdp",
					"qmc",
					"wvtnfkggzfwdwubw",
					"thhnwjyavvphw",
					"bidxkeuhykdbvirebpn",
					"rwuggu",
					"vjslcqestszouquyfb",
					"jrnknmtcjtdm",
					"xoixkrdwzzz",
					"xoixkrdwzzz",
				},
			},
			best:  21,
			worst: 22,
		},
	}
	for _, v := range tests {
		if best, worst := passwords(v.args.n, v.args.k, v.args.pass); v.best != best || v.worst != worst {
			fmt.Printf("args: %+v - best: %d  worst: %d / expected best: %d worst: %d\n", v.args, best, worst, v.best, v.worst)
		}
	}

}
