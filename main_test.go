package main

import (
	"reflect"
	"testing"
)

func TestFun(t *testing.T) {
	out := generateParenthesis(2)

	if reflect.DeepEqual(out, []string{"(())", "()()"}) != true {
		//if out != []string{"(())", "()()"} {
		t.Errorf("got %#v, want %#v", out, []string{"(())", "()()"})
	}

	out2 := generateParenthesis(3)
	if reflect.DeepEqual(out2, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}) != true {
		t.Errorf("got %#v, want %#v", out2, []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
	}
}
