package stackqueue

import (
	"fmt"
	"reflect"
)

func TransformTheExpression(t int, expressions []string) []string {
	mapOperator := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"^": true,
	}
	closeBracket := ")"
	res := make([]string, 0, t)
	for _, expression := range expressions {
		var rpn string
		stack := NewStack[string]()
		for i := range expression {
			char := string(expression[i])
			if char == "(" {
				continue
			}
			if _, ok := mapOperator[char]; ok {
				stack.Push(char)
				continue
			}
			if char == closeBracket {
				if !stack.IsEmpty() {
					rpn += stack.Top()
					stack.Pop()
				}
				continue
			}
			rpn += char
		}
		res = append(res, rpn)
	}
	return res
}

func ConsoleTransformTheExpression() {
	var t int
	fmt.Scanln(&t)
	expressions := make([]string, 0, t)
	var temp string
	for i := 0; i < t; i++ {
		fmt.Scanln(&temp)
		expressions = append(expressions, temp)
	}
	res := TransformTheExpression(t, expressions)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestTransformTheExpression() {
	type args struct {
		t           int
		expressions []string
	}

	cases := []struct {
		args args
		want []string
	}{
		{
			args: args{
				t: 3,
				expressions: []string{
					"(a+(b*c))",
					"((a+b)*(z+x))",
					"((a+t)*((b+(a+c))^(c+d)))",
				},
			},
			want: []string{
				"abc*+",
				"ab+zx+*",
				"at+bac++cd+^*",
			},
		},
	}

	for _, v := range cases {
		if got := TransformTheExpression(v.args.t, v.args.expressions); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
