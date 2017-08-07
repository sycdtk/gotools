package rpn

import (
	"testing"
)

func TestParse(t *testing.T) {
	s := "@MUL(@DIV(@IF(@AND(@OR(@EQ(DEPACT,11),@EQ(DEPACT,12)),@OR(@EQ(DEPSTA,0),@EQ(DEPSTA,2))),1,0),IPNDEP),100)"
	if Parse(s) == "DEPACT 11 @EQ DEPACT 12 @EQ @OR DEPSTA 0 @EQ DEPSTA 2 @EQ @OR @AND 1 0 @IF IPNDEP @DIV 100 @MUL" {
		t.Log(Parse(s))
	} else {
		t.Fatal(Parse(s))
	}

}
