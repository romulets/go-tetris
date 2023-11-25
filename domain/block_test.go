package domain

import (
	"reflect"
	"strings"
	"testing"
)

func TestRotateOBlock(t *testing.T) {
	b := buildOBlock()
	expectedOBody := [][]bool{
		{true, true},
		{true, true},
	}

	if b.shape != oBlock {
		t.Errorf("Expected shape %c, but it was %c", oBlock, b.shape)
	}

	assertBodyIs(t, b, expectedOBody)

	b.rotate()
	assertBodyIs(t, b, expectedOBody)

	b.rotate()
	assertBodyIs(t, b, expectedOBody)

	b.rotate()
	assertBodyIs(t, b, expectedOBody)

	b.rotate()
	assertBodyIs(t, b, expectedOBody)

	b.rotate()
	assertBodyIs(t, b, expectedOBody)
}

func TestRotateIBlock(t *testing.T) {
	b := buildIBlock()
	if b.shape != iBlock {
		t.Errorf("Expected shape %c, but it was %c", iBlock, b.shape)
	}

	assertBodyIs(t, b, [][]bool{
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
		{true, true, true, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, false, false, true},
		{false, false, false, true},
		{false, false, false, true},
		{false, false, false, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, true, true, true},
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, false, false, false},
		{true, false, false, false},
		{true, false, false, false},
		{true, false, false, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, false, false, false},
		{false, false, false, false},
		{false, false, false, false},
		{true, true, true, true},
	})
}

func TestRotateJBlock(t *testing.T) {
	b := buildJBlock()
	if b.shape != jBlock {
		t.Errorf("Expected shape %c, but it was %c", jBlock, b.shape)
	}

	assertBodyIs(t, b, [][]bool{
		{false, false, false},
		{true, false, false},
		{true, true, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, false, true},
		{false, false, true},
		{false, true, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, true, true},
		{false, false, true},
		{false, false, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, true, false},
		{true, false, false},
		{true, false, false},
	})
}

func TestRotateLBlock(t *testing.T) {
	b := buildLBlock()
	if b.shape != lBlock {
		t.Errorf("Expected shape %c, but it was %c", lBlock, b.shape)
	}

	assertBodyIs(t, b, [][]bool{
		{false, false, false},
		{false, false, true},
		{true, true, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, true, true},
		{false, false, true},
		{false, false, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, true, true},
		{true, false, false},
		{false, false, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, false, false},
		{true, false, false},
		{true, true, false},
	})
}

func TestRotateSBlock(t *testing.T) {
	b := buildSBlock()
	if b.shape != sBlock {
		t.Errorf("Expected shape %c, but it was %c", sBlock, b.shape)
	}

	assertBodyIs(t, b, [][]bool{
		{false, false, false},
		{false, true, true},
		{true, true, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, true, false},
		{false, true, true},
		{false, false, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, true, true},
		{true, true, false},
		{false, false, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, false, false},
		{true, true, false},
		{false, true, false},
	})
}

func TestRotateZBlock(t *testing.T) {
	b := buildZBlock()
	if b.shape != zBlock {
		t.Errorf("Expected shape %c, but it was %c", zBlock, b.shape)
	}

	assertBodyIs(t, b, [][]bool{
		{false, false, false},
		{true, true, false},
		{false, true, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, false, true},
		{false, true, true},
		{false, true, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, true, false},
		{false, true, true},
		{false, false, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, true, false},
		{true, true, false},
		{true, false, false},
	})
}

func TestRotateTBlock(t *testing.T) {
	b := buildTBlock()
	if b.shape != tBlock {
		t.Errorf("Expected shape %c, but it was %c", tBlock, b.shape)
	}

	assertBodyIs(t, b, [][]bool{
		{false, false, false},
		{false, true, false},
		{true, true, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{false, false, true},
		{false, true, true},
		{false, false, true},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, true, true},
		{false, true, false},
		{false, false, false},
	})

	b.rotate()
	assertBodyIs(t, b, [][]bool{
		{true, false, false},
		{true, true, false},
		{true, false, false},
	})
}

func TestBuildRandomBlock(t *testing.T) {
	allShapes := []shape{lBlock, jBlock, iBlock, oBlock, sBlock, zBlock, tBlock}
	occ := make(map[shape]int, len(allShapes))

	for x := 0; x < 1000; x++ {
		b := buildRandomBlock()
		occ[b.shape] = occ[b.shape] + 1
	}

	for _, s := range allShapes {
		if occ[s] == 1000 || occ[s] == 0 {
			t.Errorf("%c occurred %d times (expected something between 1 and 999)", s, occ[s])
		} else {
			t.Logf("%c occurred %d times", s, occ[s])
		}
	}
}

func assertBodyIs(t *testing.T, b block, expectedBody [][]bool) {
	t.Helper()

	if !reflect.DeepEqual(b.body, expectedBody) {
		t.Errorf("Unexpected body \n%s wanted \n%s", fmtBody(b.body), fmtBody(expectedBody))
	}
}

func fmtBody(body [][]bool) string {
	builder := strings.Builder{}

	for _, row := range body {
		for _, pos := range row {
			if pos {
				builder.WriteString("#")
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
