package obj

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Library holds information about a Library in the Google Books Hashcode problem
type Library struct {
	Index       int
	Signup      int
	BooksPerDay int
	NBooks      int
	Books       []bool
}

// BooksProblem holds information about Google Books Hashcode problem
type BooksProblem struct {
	Days       int
	NBooks     int
	NLibraries int
	BookScores []int
	Libraries  []Library
}

// Parse fills BooksProblem with information from string
func (p *BooksProblem) Parse(f string) (err error) {
	lines := strings.Split(f, "\n")
	head := strings.Split(lines[0], " ")

	p.NBooks, err = strconv.Atoi(head[0])
	if err != nil {
		return fmt.Errorf("Failed to convert NBooks string to integer: %v", err)
	}
	p.NLibraries, err = strconv.Atoi(head[1])
	if err != nil {
		return fmt.Errorf("Failed to convert NLibraries string to integer: %v", err)
	}
	p.Days, err = strconv.Atoi(head[2])
	if err != nil {
		return fmt.Errorf("Failed to convert Days string to integer: %v", err)
	}

	p.BookScores = make([]int, p.NBooks)
	scores := strings.Split(lines[1], " ")
	for i := range p.BookScores {
		p.BookScores[i], err = strconv.Atoi(scores[i])
		if err != nil {
			return fmt.Errorf("Failed to convert BookScores string to integer: %v", err)
		}
	}

	p.Libraries = make([]Library, p.NLibraries)
	for i := range p.Libraries {
		lib := Library{}
		lib.Index = i
		lib.Books = make([]bool, p.NBooks)
		line := 2 + 2*i
		stats := strings.Split(lines[line], " ")
		lib.NBooks, err = strconv.Atoi(stats[0])
		if err != nil {
			return fmt.Errorf("Failed to convert NBooks string to integer: %v", err)
		}
		lib.Signup, err = strconv.Atoi(stats[1])
		if err != nil {
			return fmt.Errorf("Failed to convert Signup string to integer: %v", err)
		}
		lib.BooksPerDay, err = strconv.Atoi(stats[2])
		if err != nil {
			return fmt.Errorf("Failed to convert BooksPerDay string to integer: %v", err)
		}
		books := strings.Split(lines[line+1], " ")
		for j := 0; j < lib.NBooks; j++ {
			book, err := strconv.Atoi(books[j])
			if err != nil {
				return fmt.Errorf("Failed to convert book string to integer: %v", err)
			}
			lib.Books[book] = true
		}
		p.Libraries[i] = lib
	}
	return nil
}

// BooksSolution holds a solution to BooksProblem
type BooksSolution struct {
	LibraryOrder  []int
	AssignedBooks map[int][]int
}

// Parse fills BooksSolution with information from string
func (s *BooksSolution) Parse(f string) (err error) {
	lines := strings.Split(f, "\n")
	NLibraries, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("Failed to parse NLibraries string: %v", err)
	}
	order := make([]int, NLibraries)
	assignment := map[int][]int{}
	for i := 0; i < NLibraries; i++ {
		line := 1 + 2*i
		info := strings.Split(lines[line], " ")
		idx, err := strconv.Atoi(info[0])
		if err != nil {
			return fmt.Errorf("Failed to parse idx string: %v", err)
		}
		nbooks, err := strconv.Atoi(info[1])
		if err != nil {
			return fmt.Errorf("Failed to parse nbooks string: %v", err)
		}

		order[i] = idx
		assignment[idx] = make([]int, nbooks)
		books := strings.Split(lines[line+1], " ")
		for j := 0; j < nbooks; j++ {
			assignment[idx][j], err = strconv.Atoi(books[j])
			if err != nil {
				return fmt.Errorf("Failed to parse assignment[idx][j] string: %v", err)
			}
		}
	}
	s.LibraryOrder = order
	s.AssignedBooks = assignment
	return nil
}

// BooksScore evaluates a solution for the problem and returns the score
func BooksScore(problem *BooksProblem, solution *BooksSolution) (int, error) {
	score := 0
	endOfSignup := 0                        // day when next sign up process ends
	signedUpLibraries := -1                 // count number of signed up libraries
	scanned := make([]bool, problem.NBooks) // track which books have already been scanned
	libStartDay := map[int]int{}            // track when a library starts scanning

	for day := 0; day < problem.Days; day++ {
		// finish / start sign up process of library
		if endOfSignup == day {
			// mark day of library scan start
			if signedUpLibraries != -1 {
				libStartDay[solution.LibraryOrder[signedUpLibraries]] = day
			}

			signedUpLibraries++

			// start sign up process for next library
			if signedUpLibraries != len(solution.LibraryOrder) {
				endOfSignup = day + problem.Libraries[solution.LibraryOrder[signedUpLibraries]].Signup
			}
		}
		// scan books signed up libraries
		for i := 0; i < signedUpLibraries; i++ {
			lib := &problem.Libraries[solution.LibraryOrder[i]]
			start := (day - libStartDay[lib.Index]) * lib.BooksPerDay
			for j := 0; j < lib.BooksPerDay; j++ {
				if start+j > len(solution.AssignedBooks[lib.Index])-1 {
					break
				}
				book := solution.AssignedBooks[lib.Index][start+j]
				if !lib.Books[book] {
					return -1, errors.New("The assigned book is not in the library")
				}
				if !scanned[book] {
					scanned[book] = true
					score += problem.BookScores[book]
				}
			}
		}
	}
	return score, nil
}
