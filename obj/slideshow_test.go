package obj

import (
	"testing"
)

func TestSlideshowProblemParse(t *testing.T) {
	problem := `4
H 3 cat beach sun
V 2 selfie smile
V 2 garden selfie
H 2 garden cat
`
	sp := SlideshowProblem{}
	err := sp.Parse(problem)
	if err != nil {
		t.Errorf("failed to parse problem: %v", err)
	}
	if len(sp) != 4 {
		t.Errorf("unexpected number of photos: %v", len(sp))
	}
	if !sp[0].Horizontal {
		t.Error("photo 0 should be horizontal")
	}
	if sp[1].Horizontal {
		t.Error("photo 1 should be vertical")
	}
	if len(sp[2].Tags) != 2 {
		t.Error("photo 2 should have 2 tags")
	}
}

func TestSlideshowSolutionParse(t *testing.T) {
	solution := `3
0
3
1 2
`
	ss := SlideshowSolution{}
	err := ss.Parse(solution)
	if err != nil {
		t.Errorf("failed to parse Solution: %v", err)
	}
	if len(ss) != 3 {
		t.Errorf("unexpected number of slides: %v", len(ss))
	}
	if !(ss[0].Photo1 == 0) && !(ss[0].Photo2 == -1) {
		t.Error("slide 0 should contain photo 0")
	}
	if !(ss[2].Photo1 == 1) && !(ss[2].Photo2 == 2) {
		t.Error("slide 2 should contain photo 1 and 2")
	}
	solution = `3
	0
	0
	0
	`
	err = ss.Parse(solution)
	if err == nil {
		t.Error("parsing solution with reappearing photos should fail")
	}
}

func TestSlide(t *testing.T) {
	problem := `4
H 3 cat beach sun
V 2 selfie smile
V 2 garden selfie
H 2 garden cat
`
	sp := SlideshowProblem{}
	err := sp.Parse(problem)
	if err != nil {
		t.Errorf("failed to parse problem: %v", err)
	}
	slide := Slide{
		Photo1: -1,
		Photo2: -1,
	}
	if slide.Valid(&sp) {
		t.Error("Photo1 == -1 should be invalid")
	}
	slide = Slide{
		Photo1: 0,
		Photo2: 3,
	}
	if slide.Valid(&sp) {
		t.Error("two horizontal photos should be invalid")
	}
	slide = Slide{
		Photo1: 1,
		Photo2: -1,
	}
	if slide.Valid(&sp) {
		t.Error("only one vertical photo should be invalid")
	}
	slide = Slide{
		Photo1: 0,
		Photo2: -1,
	}
	if !slide.Valid(&sp) {
		t.Error("one horizontal photo should be valid")
	}
	slide = Slide{
		Photo1: 1,
		Photo2: 2,
	}
	if !slide.Valid(&sp) {
		t.Error("two vertical photos should be valid")
	}
}

func TestSlideshowScore(t *testing.T) {
	solution := `3
0
3
1 2
`
	ss := SlideshowSolution{}
	err := ss.Parse(solution)
	if err != nil {
		t.Errorf("failed to parse Solution: %v", err)
	}
	problem := `4
H 3 cat beach sun
V 2 selfie smile
V 2 garden selfie
H 2 garden cat
`
	sp := SlideshowProblem{}
	err = sp.Parse(problem)
	if err != nil {
		t.Errorf("failed to parse problem: %v", err)
	}
	score, err := SlideshowScore(&sp, &ss)
	if err != nil {
		t.Errorf("unexpected error when computing score: %v", err)
	}
	if score != 2 {
		t.Errorf("wrong score. Expected 2, got %v", score)
	}
}
