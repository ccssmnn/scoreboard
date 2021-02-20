package obj

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Intersection is a coordinate in the Self-driving Rides Problem
type Intersection [2]int

func distance(a, b Intersection) int {
	x := a[0] - b[0]
	if x < 0 {
		x *= -1
	}
	y := a[1] - b[1]
	if y < 0 {
		y *= -1
	}
	return x + y
}

// Ride represents a pre booked ride in the Self-driving Rides Problem
type Ride struct {
	From          Intersection
	To            Intersection
	EarliestStart int
	LatestFinish  int
}

// RidesProblem represents a problem described in Google Hashcode Self-driving Rides
type RidesProblem struct {
	Rows      int
	Cols      int
	NVehicles int
	T         int
	Bonus     int
	Rides     []Ride
}

// Parse populates problem with information from string
func (r *RidesProblem) Parse(f string) error {
	lines := strings.Split(f, "\n")
	stats := strings.Split(lines[0], " ")
	R, err := strconv.Atoi(stats[0])
	if err != nil {
		return fmt.Errorf("failed to parse R: %v", err)
	}
	C, err := strconv.Atoi(stats[1])
	if err != nil {
		return fmt.Errorf("failed to parse C: %v", err)
	}
	F, err := strconv.Atoi(stats[2])
	if err != nil {
		return fmt.Errorf("failed to parse F: %v", err)
	}
	N, err := strconv.Atoi(stats[3])
	if err != nil {
		return fmt.Errorf("failed to parse N: %v", err)
	}
	B, err := strconv.Atoi(stats[4])
	if err != nil {
		return fmt.Errorf("failed to parse B: %v", err)
	}
	T, err := strconv.Atoi(stats[5])
	if err != nil {
		return fmt.Errorf("failed to parse T: %v", err)
	}
	problem := RidesProblem{
		T:         T,
		Rows:      R,
		Cols:      C,
		Bonus:     B,
		NVehicles: F,
		Rides:     make([]Ride, N),
	}
	for i := range problem.Rides {
		stats := strings.Split(lines[1+i], " ")
		a, err := strconv.Atoi(stats[0])
		if err != nil {
			return fmt.Errorf("failed to parse a for ride %v: %v", i, err)
		}
		b, err := strconv.Atoi(stats[1])
		if err != nil {
			return fmt.Errorf("failed to parse b for ride %v: %v", i, err)
		}
		x, err := strconv.Atoi(stats[2])
		if err != nil {
			return fmt.Errorf("failed to parse x for ride %v: %v", i, err)
		}
		y, err := strconv.Atoi(stats[3])
		if err != nil {
			return fmt.Errorf("failed to parse y for ride %v: %v", i, err)
		}
		s, err := strconv.Atoi(stats[4])
		if err != nil {
			return fmt.Errorf("failed to parse s for ride %v: %v", i, err)
		}
		f, err := strconv.Atoi(stats[5])
		if err != nil {
			return fmt.Errorf("failed to parse f for ride %v: %v", i, err)
		}
		problem.Rides[i] = Ride{
			From:          Intersection{a, b},
			To:            Intersection{x, y},
			EarliestStart: s,
			LatestFinish:  f,
		}
	}
	(*r) = problem
	return nil
}

// RidesSolution represents assignment of rides to a car in fleet
type RidesSolution map[int][]int

// Parse populates solution with information from string
func (r *RidesSolution) Parse(f string) error {
	lines := strings.Split(f, "\n")
	solution := map[int][]int{}
	assigned := map[int]bool{} // check if a ride is assigned twice
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	for i, line := range lines {
		stats := strings.Split(line, " ")
		rides := make([]int, len(stats)-1)
		for j, stat := range stats[1:] {
			ride, err := strconv.Atoi(stat)
			if err != nil {
				return fmt.Errorf("failed to parse ride: %v", err)
			}
			if _, ok := assigned[ride]; ok {
				return fmt.Errorf("multiple assignments for ride %v", ride)
			}
			assigned[ride] = true
			rides[j] = ride
		}
		solution[i] = rides
	}
	(*r) = solution
	return nil
}

// RidesScore runs simulation for assigned rides and returns score
func RidesScore(problem *RidesProblem, solution *RidesSolution) (int, error) {
	score := 0
	if len(*solution) > problem.NVehicles {
		return score, errors.New("too many vehicles assigned")
	}
	for vehid, schedule := range *solution {
		t := 0                    // reset time
		pos := Intersection{0, 0} // reset position
		for _, ridx := range schedule {
			// check if ride is part of the problem
			if ridx >= len(problem.Rides) {
				return score, fmt.Errorf("illegal ride assignment in for vehicle %v. Ride %v is not available", vehid, ridx)
			}
			ride := problem.Rides[ridx]
			// drive to ride start
			t += distance(pos, ride.From)
			if t > problem.T {
				break
			}
			pos = ride.From
			// make car wait if early and give bonus
			if t <= ride.EarliestStart {
				score += problem.Bonus
				t = ride.EarliestStart
			}
			// drive to ride end
			t += distance(pos, ride.To)
			if t > problem.T {
				break
			}
			pos = ride.To
			// give bonus if ride is on time
			if t <= ride.LatestFinish {
				score += distance(ride.From, ride.To)
			}
		}
	}
	return score, nil
}
