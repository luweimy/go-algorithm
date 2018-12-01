package main

import (
	"bytes"
	"container/list"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 前缀表达式
// 前缀表达式定义：
// 1. 一个数是一个前缀表达式，值为该数
// 2. "运算符[空格]前缀表达式[空格]前缀表达式"是前缀表达式，值为结果
func PrefixExpr(expr *[]string) float64 {
	if len(*expr) == 0 {
		return 0
	}
	// 从表达式pop出第一个元素（可能是操作符，也可能是数字）
	var first = (*expr)[0]
	*expr = (*expr)[1:]
	switch first {
	case "+":
		return PrefixExpr(expr) + PrefixExpr(expr)
	case "-":
		return PrefixExpr(expr) - PrefixExpr(expr)
	case "*":
		return PrefixExpr(expr) * PrefixExpr(expr)
	case "/":
		return PrefixExpr(expr) / PrefixExpr(expr)
	default:
		return mustParseFloat(first)
	}
	return 0
}

func mustParseFloat(v string) float64 {
	f, err := strconv.ParseFloat(v, 0)
	if err != nil {
		panic(err)
	}
	return f
}

// 中缀表达式
// 中缀表达式也是由多个子中缀表达式组合而成
// 假设：表达式无空格，以简化处理
func InfixExpr(expr string) float64 {
	var nu = list.New()
	var op = list.New()
	for i := 0; i < len(expr); i++ {
		fmt.Println("expr", expr[i:])
		switch expr[i] {
		case '(':
			op.PushBack(expr[i])
		case ')':
			// 遇到右括号，则op出栈并计算，直接出栈左括号停止
			for op.Len() > 0 {
				var opr = op.Remove(op.Back())
				if opr.(uint8) == '(' {
					break
				}
				nu.PushBack(calc(nu.Remove(nu.Back()), opr, nu.Remove(nu.Back())))
			}
		case '+', '-':
			// 当前为+-操作，则计算上一步操作（为了保障*/的优先级）
			// 若上一步op是左括号，则略过
			if op.Len() > 0 {
				switch op.Back().Value.(uint8) {
				case '+', '-', '*', '/':
					nu.PushBack(calc(nu.Remove(nu.Back()), op.Remove(op.Back()), nu.Remove(nu.Back())))
				}
			}
			op.PushBack(expr[i])
		case '*', '/':
			op.PushBack(expr[i])
		default:
			// 解析出一个数字，并入栈
			j := i
			for ; j < len(expr); j++ {
				var stop = false
				switch expr[j] {
				case '+', '-', '*', '/', '(', ')':
					stop = true
				}
				if stop {
					break
				}
			}
			nu.PushBack(mustParseFloat(expr[i:j]))
			i = j - 1 // 下次循环会i++
		}
	}
	for op.Len() > 0 {
		nu.PushBack(calc(nu.Remove(nu.Back()), op.Remove(op.Back()), nu.Remove(nu.Back())))
	}
	return nu.Back().Value.(float64)
}

func calc(n1, op, n2 interface{}) float64 {
	switch op.(uint8) {
	case '+':
		return n1.(float64) + n2.(float64)
	case '-':
		return n1.(float64) - n2.(float64)
	case '*':
		return n1.(float64) * n2.(float64)
	case '/':
		return n2.(float64) / n1.(float64) // 注意操作数顺序
	default:
		panic(-1)
	}
	return 0
}

func dump(s string, l *list.List) {
	var sb = bytes.Buffer{}
	var el = l.Front()
	for el != nil {
		switch el.Value.(type) {
		case uint8:
			sb.WriteString(fmt.Sprintf("%c ", el.Value))
		case float64:
			sb.WriteString(fmt.Sprintf("%f ", el.Value))
		default:
			panic(fmt.Sprintln("unsupport type", reflect.TypeOf(el.Value)))
		}
		el = el.Next()
	}
	fmt.Printf("%s(%d): %s\n", s, l.Len(), sb.String())
}

func main() {
	// 前缀表达式求值测试
	{
		// (11+12)*(24+35) = 1357
		var expr = strings.Fields("* + 11.0 12.0 + 24.0 35.0")
		fmt.Println(PrefixExpr(&expr))
	}
	// 中缀表达式求值测试
	{
		fmt.Println(InfixExpr("2*35/7"))            // result: 10
		fmt.Println(InfixExpr("(2+3)*(5+7)+9/3"))   // result: 63
		fmt.Println(InfixExpr("((2+3)*(5+7)+9/3)")) // result: 63
		fmt.Println(InfixExpr("(0)"))               // result: 63

	}
}
