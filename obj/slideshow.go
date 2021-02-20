package obj

import (
	"fmt"
	"strconv"
	"strings"
)

// Photo represents one photo in the Photo Slideshow problem
type Photo struct {
	Horizontal bool
	Tags       map[string]bool
}

// SlideshowProblem is the collection of photos in the Photo Slideshow problem
type SlideshowProblem []Photo

// Parse converts a slideshow problem string into SlideshowProblem
func (s *SlideshowProblem) Parse(f string) error {
	lines := strings.Split(f, "\n")
	N, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("failed to convert string to integer: %v", err)
	}
	problem := make(SlideshowProblem, N)
	for i := range problem {
		stats := strings.Split(lines[1+i], " ")
		if stats[0] != "H" && stats[0] != "V" {
			return fmt.Errorf("unknown orientation tag. Expected H or V, got %v", stats[0])
		}
		horizontal := stats[0] == "H"
		tags := map[string]bool{}
		for _, tag := range stats[2:] {
			tags[tag] = true
		}
		problem[i] = Photo{
			Horizontal: horizontal,
			Tags:       tags,
		}
	}
	(*s) = problem
	return nil
}

// Slide holds indices for photos in this slide. A horizontal Slide has only Photo1 and Photo2 is -1
type Slide struct {
	Photo1 int
	Photo2 int
}

// Valid checks if a slide is valid
func (s *Slide) Valid(problem *SlideshowProblem) bool {
	if s.Photo1 == -1 {
		return false
	}
	// a horizontal photo must be alone on the slide
	if (*problem)[s.Photo1].Horizontal {
		return s.Photo2 == -1
	}
	if s.Photo2 == -1 {
		return false
	}
	return !(*problem)[s.Photo1].Horizontal && !(*problem)[s.Photo2].Horizontal
}

// Tags returns merged tags of slide
func (s *Slide) Tags(problem *SlideshowProblem) map[string]bool {
	tags := map[string]bool{}
	for tag := range (*problem)[s.Photo1].Tags {
		tags[tag] = true
	}
	if s.Photo2 == -1 {
		return tags
	}
	for tag := range (*problem)[s.Photo2].Tags {
		tags[tag] = true
	}
	return tags
}

// SlideshowSolution represents the order of slides
type SlideshowSolution []Slide

// Parse converts a string into SlideshowSolution
func (s *SlideshowSolution) Parse(f string) error {
	lines := strings.Split(f, "\n")
	N, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("failed to convert string to integer: %v", err)
	}
	show := make(SlideshowSolution, N)
	used := map[int]bool{}
	for i := range show {
		stats := strings.Split(lines[1+i], " ")
		first, err := strconv.Atoi(stats[0])
		if err != nil {
			return fmt.Errorf("failed to convert string to integer: %v", err)
		}
		if _, exists := used[first]; exists {
			return fmt.Errorf("picture %v appears more than 1 time", first)
		}
		used[first] = true
		second := -1
		if len(stats) > 1 {
			second, _ = strconv.Atoi(stats[1])
		}
		if _, exists := used[second]; second != -1 && exists {
			return fmt.Errorf("picture %v appears more than 1 time", second)
		}
		used[second] = true
		slide := Slide{
			Photo1: first,
			Photo2: second,
		}
		show[i] = slide
	}
	(*s) = show
	return nil
}

// SlideshowScore evaluates the score of a problem/solution combination
func SlideshowScore(problem *SlideshowProblem, solution *SlideshowSolution) (int, error) {
	score := 0
	for i := 1; i < len(*solution); i++ {
		// check validity of slide
		if !(*solution)[i-1].Valid(problem) || !(*solution)[i].Valid(problem) {
			return -1, fmt.Errorf("slide %v or %v is invalid", i-1, i)
		}
		// evaluate each slide transition
		prevTags := (*solution)[i-1].Tags(problem)
		currTags := (*solution)[i].Tags(problem)

		// count tags
		distinctPrev := 0 // in prev but not in current
		common := 0       // in both
		distinctCurr := 0 // in current but not in prev
		for tag := range prevTags {
			_, present := currTags[tag]
			if present {
				common++
			} else {
				distinctPrev++
			}
		}
		for tag := range currTags {
			_, present := prevTags[tag]
			if !present {
				distinctCurr++
			}
		}
		// score is increased by min(distinctPrev, common, distinctCurr)
		if distinctPrev < common && distinctPrev < distinctCurr {
			score += distinctPrev
		} else if distinctCurr < common {
			score += distinctCurr
		} else {
			score += common
		}
	}
	return score, nil
}
