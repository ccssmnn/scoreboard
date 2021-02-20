package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ccssmnn/scoreboard/obj"
	"github.com/gorilla/mux"
)

var bookProblems = map[string]*obj.BooksProblem{}

var allowedBookProblems = []string{
	"a_example",
	"b_read_on",
	"c_incunabula",
	"d_touch_choices",
	"e_so_many_books",
	"f_libraries_of_the_world",
}

func stringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}
	return false
}

// HandleBooks handles a Google Books related http request
func HandleBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problem := vars["problem"]
	// check if problem is supported
	if !stringInSlice(problem, allowedBookProblems) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "requested problem is not supported")
		return
	}
	// read problem file if not done already
	if _, found := bookProblems[problem]; !found {
		bp := &obj.BooksProblem{}
		file, err := ioutil.ReadFile("static/2020/qualification/" + problem + ".txt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "failed to parse problem: %v", err)
			return
		}
		bp.Parse(string(file))
		bookProblems[problem] = bp
	}
	// read solution
	defer r.Body.Close()

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "failed to read request body: %v", err)
			return
		}
	}

	solution := obj.BooksSolution{}
	err = solution.Parse(string(bytes))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to parse solution from request body: %v", err)
		return
	}
	// compute result
	score, err := obj.BooksScore(bookProblems[problem], &solution)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed to compute score: %v", err)
		return
	}
	fmt.Fprintf(w, "%v", score)
}
