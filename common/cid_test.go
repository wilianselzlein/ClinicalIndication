package common

import (
	"testing"
)

func TestCidSimple(t *testing.T) {
	res, err := ExtractCid("123 c 123 123")
	want := "c123"
	if res[0] != want || err != nil  {
		t.Errorf("ExtractCidSimple = %s, %v; want %s", res, err, want)
	}

	res, err = ExtractCid("123 c123 123")
	if res[0] != want || err != nil  {
		t.Errorf("ExtractCidSimple = %s, %v; want %s", res, err, want)
	}
}

func testRange(cids *map[string]string, t *testing.T) {
	for k, v := range *cids {
		res, err := ExtractCid(k)
		if (len(res) > 0 && res[0] != v) || err != nil  {
			t.Errorf("ExtractCid(%s) != %s, %v; want %s", k, res, err, v)
		}
	}
}

func TestRangeCidsStartWithBlancSpace(t *testing.T) {
	cids := map[string]string{
		" c111": "c111", 
		" c111 a": "c111",
		" c11 1 a": "c111",
		" c 111": "c111",
		" c 111 a": "c111",
		" c 11 1": "c111",
		" c 11 1 a": "c111",
		" c 1 1 1": "c111",
		" c 1 1 1 |": "c111",
		" c 11": "c11",
		" c 11 a": "c11",
		" c11": "c11",
		" c11 |": "c11",
	}
	testRange(&cids, t)
}

func TestRangeCidsStartWithoutBlancSpace(t *testing.T) {
	cids := map[string]string{
		"c111": "c111", 
		"c111 a": "c111",
		"c11 1 a": "c111",
		"c 111": "c111",
		"c 111 a": "c111",
		"c 11 1": "c111",
		"c 11 1 a": "c111",
		"c 1 1 1": "c111",
		"c 1 1 1 |": "c111",
		"c 11": "c11",
		"c 11 a": "c11",
		"c11": "c11",
		"c11 |": "c11",
	}

	testRange(&cids, t)
}


func TestRangeCidsEmpty(t *testing.T) {
	cids := map[string]string{
		"AC111": "", 
		"AC111 a": "", 
		"ac11": "", 
		"ac11 |": "", 
		"ac11 1 a": "", 
		" ac11 1 a": "", 
		"ad 11": "", 
		"ad 11 a": "", 
		"ac 111": "", 
		"ac 111 a": "", 
		"ac 11 1": "", 
		"ac 11 1 a": "", 
		"ad 11 1": "", 
		"ad 11 1 a": "", 
		"ac 1 1 1": "", 
		"ac 1 1 1 |": "", 
		" d 1": "", 
		" d 1 a": "", 
		" C11145": "", 
		" ": "", 
	}

	testRange(&cids, t)
}