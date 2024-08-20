package orderedmap

import "testing"

func TestOrderedMap(t *testing.T) {
	om1 := NewOrderdMap([]OrderdMapElement[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}, false)
	if om1.Get(1) != "a" {
		t.Error("Get(1) should return 'a'")
	}
	if om1.Get(2) != "b" {
		t.Error("Get(2) should return 'b'")
	}
	if om1.Len() != 3 {
		t.Error("Len() should return 3")
	}

	om2 := NewOrderdMap([]OrderdMapElement[string, string]{{"1", "a"}, {"2", "b"}, {"3", "c"}}, false)
	if om2.Get("1") != "a" {
		t.Error("Get('1') should return 'a'")
	}
	om2.Set("2", "bb")
	if om2.Get("2") != "bb" {
		t.Error("Get('2') should return 'bb'")
	}

	om3 := NewOrderdMap([]OrderdMapElement[float32, []string]{{1.1, []string{"a"}}, {2.2, []string{"b"}}, {3.3, []string{"c"}}}, false)
	if om3.Get(1.1)[0] != "a" {
		t.Error("Get(1.1) should return 'a'")
	}
	count := 0
	for k, v := range om3.Range() {
		if om3.Keys[count] != k || om3.Map[om3.Keys[count]][0] != v[0] {
			t.Error("Range() should return the correct value", k, v)
		}
		count++
	}
	if om3.Get(2.2)[0] != "b" {
		t.Error("Get(2.2) should return 'b'")
	}
}
