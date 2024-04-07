package generator

import "testing"

func TestReadFileLines(t *testing.T) {
	var filepath = "tests/testdata/lines.txt"
	var expectedLines = []string{"line 1", "line 2", "line_3", "line_4"}
	lines, err := readFileLines(filepath)

	if err != nil {
		t.Fatalf("expected no error; error = %v", err)
	}

	if lines == nil {
		t.Fatal("expected non-nil")
	}

	if len(lines) != len(expectedLines) {
		t.Fatalf("expected lines = %d; got = %d", len(expectedLines), len(lines))
	}

	for i, expectedLine := range expectedLines {
		if lines[i] != expectedLine {
			t.Errorf("expected line = %s; got = %s", expectedLine, lines[i])
		}
	}
}
