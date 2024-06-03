package qp

import (
	"testing"
)

func sliceEqual(a []string, b ...string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func mapEqualFlat(m map[string][]string, b ...string) bool {
	if len(m) != len(b)/2 {
		return false
	}
	for i := 0; i < len(b); i += 2 {
		if m[b[i]][0] != b[i+1] {
			return false
		}
	}
	return true
}

func TestCasing(t *testing.T) {
	attrs, noise := Find("keY:VaLuE KeY:\"VaLuE\" noiSe")
	if !mapEqualFlat(attrs, "key", "VaLuE") {
		t.Fail()
	}
	if !sliceEqual(noise, "noiSe") {
		t.Fail()
	}
}

func TestUnbalancedQuotes(t *testing.T) {
	attrs, noise := Find("a:\"b")
	if len(attrs) != 0 {
		t.Fail()
	}
	if !sliceEqual(noise, "a:\"b") {
		t.Fail()
	}
}

func TestUnbalancedQuotesNoise(t *testing.T) {
	attrs, noise := Find("a:\"b c d")
	if len(attrs) != 0 {
		t.Fail()
	}
	if !sliceEqual(noise, "a:\"b", "c", "d") {
		t.Fail()
	}
}

func TestBalancedQuotes(t *testing.T) {
	attrs, noise := Find("a:\"b\"")
	if !mapEqualFlat(attrs, "a", "b") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestBalancedQuotesSpace(t *testing.T) {
	attrs, noise := Find("a:\"b c\"")
	if !mapEqualFlat(attrs, "a", "b c") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestMultipleAttributes(t *testing.T) {
	attrs, noise := Find("a:b c:d")
	if !mapEqualFlat(attrs, "a", "b", "c", "d") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestMultipleAttributesQuote(t *testing.T) {
	attrs, noise := Find("a:\"b\" c:\"d\"")
	if !mapEqualFlat(attrs, "a", "b", "c", "d") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestMultipleAttributesQuoteSpace(t *testing.T) {
	attrs, noise := Find("a:\"b c\" d:\"e f\"")
	if !mapEqualFlat(attrs, "a", "b c", "d", "e f") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestEmptyAttributeStrings(t *testing.T) {
	attrs, noise := Find("a b")
	if len(attrs) != 0 {
		t.Fail()
	}
	if !sliceEqual(noise, "a", "b") {
		t.Fail()
	}
}

func TestMixedOrder(t *testing.T) {
	attrs, noise := Find("a b c:d e f g:h")
	if !mapEqualFlat(attrs, "c", "d", "g", "h") {
		t.Fail()
	}
	if !sliceEqual(noise, "a", "b", "e", "f") {
		t.Fail()
	}
}

func TestQuotedParentheses(t *testing.T) {
	attrs, noise := Find("a:\"b (c)\"")
	if !mapEqualFlat(attrs, "a", "b (c)") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestUnquotedHyphen(t *testing.T) {
	attrs, noise := Find("a:b-c d:e")
	if !mapEqualFlat(attrs, "a", "b-c", "d", "e") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestQuotedHyphen(t *testing.T) {
	attrs, noise := Find("a:\"b-c\"")
	if !mapEqualFlat(attrs, "a", "b-c") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestQuotedUnderscore(t *testing.T) {
	attrs, noise := Find("a:\"b_c\"")
	if !mapEqualFlat(attrs, "a", "b_c") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestQuotedPeriod(t *testing.T) {
	attrs, noise := Find("a:\"b.c\"")
	if !mapEqualFlat(attrs, "a", "b.c") {
		t.Fail()
	}
	if len(noise) != 0 {
		t.Fail()
	}
}

func TestNoiseSymbols(t *testing.T) {
	attrs, noise := Find("a-b c?d e_f xy:!z")
	if len(attrs) != 0 {
		t.Fail()
	}
	if !sliceEqual(noise, "a-b", "c?d", "e_f", "xy:!z") {
		t.Fail()
	}
}
