package obj

import "testing"

func TestBooksProblemParse(t *testing.T) {
	// a_example.txt problem
	problem := `6 2 7
1 2 3 6 5 4
5 2 2
0 1 2 3 4
4 3 1
0 2 3 5
`
	bp := BooksProblem{}
	err := bp.Parse(problem)
	if err != nil {
		t.Errorf("Failed to parse problem: %v", err)
	}
	if bp.Days != 7 {
		t.Errorf("Unexpected number of days: %v", bp.Days)
	}
	if bp.NBooks != 6 {
		t.Errorf("Unexpected number of books: %v", bp.NBooks)
	}
	if bp.BookScores[3] != 6 {
		t.Errorf("Wrong content in BookScores: %v", bp.BookScores)
	}
	if bp.Libraries[1].Books[2] != true {
		t.Errorf("Wrong book assignment in library: %v", bp.Libraries[1].Books)
	}
}

func TestBooksSolutionParse(t *testing.T) {
	solution := `2
1 3
5 2 3
0 5
0 1 2 3 4`
	bs := BooksSolution{}
	err := bs.Parse(solution)
	if err != nil {
		t.Errorf("Unexpected error while parsing solution: %v", err)
	}
	if len(bs.LibraryOrder) != 2 {
		t.Error("Wrong number of libraries in solution")
	}
	if bs.LibraryOrder[0] != 1 {
		t.Errorf("Wrong first library. Expected %v, got %v", 1, bs.LibraryOrder[0])
	}
	if bs.AssignedBooks[0][2] != 2 {
		t.Error("Wrong book in assigned books")
	}
}

func TestBooksScore(t *testing.T) {
	problem := `6 2 7
1 2 3 6 5 4
5 2 2
0 1 2 3 4
4 3 1
0 2 3 5
`
	bp := BooksProblem{}
	err := bp.Parse(problem)
	if err != nil {
		t.Error("error while parsing problem")
	}
	solution := `2
1 3
5 2 3
0 5
0 1 2 3 4`
	bs := BooksSolution{}
	err = bs.Parse(solution)
	if err != nil {
		t.Error("error while parsing solution")
	}
	score, err := BooksScore(&bp, &bs)
	if err != nil {
		t.Errorf("error while evaluating books score: %v", err)
	}
	if score <= 0 {
		t.Error("score must be greater than 0")
	}
	if score != 16 {
		t.Error("Wrong total score")
	}
}
