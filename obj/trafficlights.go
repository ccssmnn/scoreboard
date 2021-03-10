package obj

import (
	"strconv"
	"strings"
)

type Street struct {
	StartNode int
	EndNode   int
	Length    int
}

type Node struct {
	IncomingStreets map[string]bool
	OutgoingStreets map[string]bool
}

type Route []string

type TrafficlightsProblem struct {
	Name    string
	Time    int
	Points  int
	Nodes   map[int]Node
	Streets map[string]Street
	Routes  []Route
}

func (p *TrafficlightsProblem) Parse(s string) error {
	lines := strings.Split(s, "\n")
	header := strings.Split(lines[0], " ")
	T, _ := strconv.Atoi(header[0])
	nIntersections, _ := strconv.Atoi(header[1])
	nStreets, _ := strconv.Atoi(header[2])
	nCars, _ := strconv.Atoi(header[3])
	points, _ := strconv.Atoi(header[4])
	p.Time = T
	p.Points = points
	p.Nodes = map[int]Node{}
	for i := 0; i < nIntersections; i++ {
		p.Nodes[i] = Node{
			IncomingStreets: map[string]bool{},
			OutgoingStreets: map[string]bool{},
		}
	}
	p.Streets = map[string]Street{}
	p.Routes = make([]Route, 0, nCars)
	for i := 1; i < nStreets+1; i++ {
		stats := strings.Split(lines[i], " ")
		start, _ := strconv.Atoi(stats[0])
		end, _ := strconv.Atoi(stats[1])
		name := stats[2]
		length, _ := strconv.Atoi(stats[3])
		p.Streets[name] = Street{
			StartNode: start,
			EndNode:   end,
			Length:    length,
		}
		p.Nodes[start].OutgoingStreets[name] = true
		p.Nodes[end].IncomingStreets[name] = true
	}
	for i := nStreets + 1; i < nStreets+1+nCars; i++ {
		stats := strings.Split(lines[i], " ")
		N, _ := strconv.Atoi(stats[0])
		r := make(Route, N)
		for j := range r {
			r[j] = stats[1+j]
		}
		p.Routes = append(p.Routes, r)
	}
	return nil
}

type Listing struct {
	Name     string
	Duration int
}

type Schedule []Listing

type TrafficlightsSolution map[int]Schedule

func (ts *TrafficlightsSolution) Parse(s string) error {
	lines := strings.Split(s, "\n")
	nSchedules, _ := strconv.Atoi(lines[0])
	baseIdx := 1
	for i := 0; i < nSchedules; i++ {
		node, _ := strconv.Atoi(lines[baseIdx])
		nListings, _ := strconv.Atoi(lines[baseIdx+1])
		sched := make(Schedule, nListings)
		for j := range sched {
			meta := strings.Split(lines[baseIdx+2+j], " ")
			listing := Listing{}
			listing.Name = meta[0]
			listing.Duration, _ = strconv.Atoi(meta[1])
			sched[j] = listing
		}
		(*ts)[node] = sched
		baseIdx += 2 + len(sched)
	}
	return nil
}

func TrafficSolutionScore(problem *TrafficlightsProblem, solution *TrafficlightsSolution) (int, error) {
	score := 0
	
	return score, nil
}
