package rpn

import (
	"testing"
)

func TestParse(t *testing.T) {
	s := "@MUL(@DIV(@IF(@AND(@OR(@EQ(DEPACT,11),@EQ(DEPACT,12)),@OR(@EQ(DEPSTA,0),@EQ(DEPSTA,2))),1,0),IPNDEP),100)"
	if Parse(s) == "| | | | | | DEPACT 11 @EQ | DEPACT 12 @EQ @OR | | DEPSTA 0 @EQ | DEPSTA 2 @EQ @OR @AND 1 0 @IF IPNDEP @DIV 100 @MUL" {
		t.Log(Parse(s))
	} else {
		t.Error("error:", Parse(s))
	}

	s = "@SUB(@PLU(1,@PLU(2,3)),@PLU(1,2))"
	if Parse(s) == "| | 1 | 2 3 @PLU @PLU | 1 2 @PLU @SUB" {
		t.Log(Parse(s))
	} else {
		t.Error("error:", Parse(s))
	}

	s = "@IF(@NOT(@OR(true,false,@AND(false,true,false)))1,2)"
	if Parse(s) == "| | | true false | false true false @AND @OR @NOT 1 2 @IF" {
		t.Log(Parse(s))
	} else {
		t.Error("error:", Parse(s))
	}

	s = "@OR(true,false,@AND(false,true,false))"
	if Parse(s) == "| true false | false true false @AND @OR" {
		t.Log(Parse(s))
	} else {
		t.Error("error:", Parse(s))
	}

	s = "@LOG10(@LOG(5))"
	if Parse(s) == "| true false | false true false @AND @OR" {
		t.Log(Parse(s))
	} else {
		t.Error("error:", Parse(s))
	}

}
