package horizon

import (
	"testing"
)

func assertOffsetCoordinates(a, b []offsetCoordinates, t *testing.T) {
	if len(a) != len(b) {
		t.Errorf("arrays differ in length: %d != %d\na: %v\nb: %v\n", len(a), len(b), a, b)
		t.Fail()
	}

	for i := range a {
		if a[i] != b[i] {
			t.Errorf("arrays differ at element %d: %v != %v\n", i, a[i], b[i])
			t.Fail()
		}
	}
}

func TestArcSimplest(t *testing.T) {
	c := arc(1)

	expected := []offsetCoordinates{{1, 0}}
	assertOffsetCoordinates(c, expected, t)
}

func TestArcXS(t *testing.T) {
	c := arc(2)

	expected := []offsetCoordinates{{2, 0}, {2, 1}}
	assertOffsetCoordinates(c, expected, t)
}

func TestArcS(t *testing.T) {
	c := arc(10)

	expected := []offsetCoordinates{{10, 0}, {10, 1}, {10, 2}, {10, 3}, {9, 4}, {9, 5}, {8, 6}, {7, 7}}
	assertOffsetCoordinates(c, expected, t)
}

func TestL(t *testing.T) {
	a := arc(10000)
	b := arc_old(10000)

	assertOffsetCoordinates(a, b, t)
}

func BenchmarkArc(b *testing.B) {
	r := 1000
	for i := 0; i < b.N; i++ {
		arc(r)
	}
}

func arc_old(r int) []offsetCoordinates {
	arc := make([]offsetCoordinates, 1)
	x, y := r, 0

	arc[0] = offsetCoordinates{r, 0}

	for x > y {
		if radiusError(x-1, y+1, r) < radiusError(x, y+1, r) {
			x -= 1
		}
		y += 1

		if x < y {
			break
		}

		arc = append(arc, offsetCoordinates{latitude: x, longitude: y})
	}

	return arc
}
