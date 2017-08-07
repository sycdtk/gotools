package rpn

import (
	"strings"

	"github.com/sycdtk/gotools/stack"
)

const (
	space            = " " //空格占位
	comma            = "," //逗号分隔符，分隔函数参数
	leftParenthesis  = "(" //左小括号，需要区分函数括号还是运算符括号
	rightParenthesis = ")" //右小括号，需要区分函数括号还是运算符括号
)

var functionNames = []string{"@LOG", "@LOG10", "@EXP", "@SQRT", "@ABS", "@INTERP",
	"@INTERPBAD", "@CAV", "@EFILT", "@DIFF", "@INTEG",
	"@SHIFT", "@VALIM", "@MKBAD", "@LOCLP", "@HICLP",
	"@MAX", "@MIN", "@IF", "@EQ", "@NE",
	"@GT", "@GE", "@LT", "@LE", "@OR",
	"@AND", "@NOT", "@BADVAL", "@PCT2", "@PFPCT",
	"@REFPCT", "@SELECT", "@MAPRANGE", "@SPLINE",
	"@ADD", "@SUB", "@MUL", "@DIV", "@POW"} //函数名称

//表达式解析函数
//输入表达式串，输出逆波兰表达式结果
func Parse(exp string) (rpn string) {

	s := stack.NewStack()

	exp = prepare(exp) //表达式预处理，关键字拆解

	for _, foo := range strings.Fields(exp) { //处理函数部分

		if isFN(foo) { //是函数名则压栈
			stack.Push(s, foo)
		} else {
			switch foo {
			case leftParenthesis: //"(" 左括号则压栈
				stack.Push(s, foo)
			case rightParenthesis: //")" 右括号则出栈，按函数方式处理
				for {
					fn := stack.Pop(s)
					if fn == leftParenthesis { //忽略左括号
					} else {
						rpn += space + fn
						if isFN(fn) { //若为函数表达式名称，跳出循环
							break
						}
					}
				}
			case comma: //","分隔符不做处理
			default: //追加参数
				if len(rpn) > 0 {
					rpn += space
				}
				rpn += foo
			}
		}
	}

	//	if empty(s){//为空栈则正常
	//	}

	return
}

//表达式预处理，对函数名称或符号名称关键字进行拆分
//例如输入：@EQ(SREIND,1)/100
//输出为：@EQ ( SREIND,1 )  / 100
func prepare(exp string) string {

	//处理函数名称
	for _, fn := range functionNames {
		exp = strings.Replace(exp, fn, space+fn+space, -1)
	}

	//处理左右括号及逗号
	for _, s := range []string{comma, leftParenthesis, rightParenthesis} {
		exp = strings.Replace(exp, s, space+s+space, -1)
	}

	return exp
}

//判断输入表达式是否为函数名称，是返回true，否则返回false
func isFN(fn string) bool {
	for _, functionName := range functionNames {
		if fn == functionName {
			return true
		}
	}
	return false
}

//判断输入表达式是否为左小括号，是返回true，否则返回false
func isLP(lp string) bool {
	return lp == leftParenthesis
}
