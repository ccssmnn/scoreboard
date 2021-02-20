package obj

import "testing"

func TestRidesProblemParse(t *testing.T) {
	problem := `3 4 2 3 2 10
0 0 1 3 2 9
1 2 1 0 0 9
2 0 2 2 0 9
`
	rp := RidesProblem{}
	err := rp.Parse(problem)
	if err != nil {
		t.Errorf("unexpected error when parsing problem: %v", err)
	}
	if rp.Rows != 3 {
		t.Error("wrong number of rows")
	}
	if rp.Cols != 4 {
		t.Error("wrong number of cols")
	}
	if rp.NVehicles != 2 {
		t.Error("wrong number of vehicles")
	}
	if len(rp.Rides) != 3 {
		t.Error("wrong number of rides")
	}
	if rp.Bonus != 2 {
		t.Error("wrong bonus")
	}
	if rp.T != 10 {
		t.Error("wrong T")
	}
	if rp.Rides[1].EarliestStart != 0 {
		t.Error("wrong earliest start for ride 1")
	}
	if rp.Rides[0].From[0] != 0 {
		t.Error("wrong start for ride 0")
	}
}

func TestRidesSolutionParse(t *testing.T) {
	solution := `1 0
2 2 1
`
	rs := RidesSolution{}
	err := rs.Parse(solution)
	if err != nil {
		t.Errorf("unexpected error when parsing solution: %v", err)
	}
	if len(rs) != 2 {
		t.Error("wrong number of schedules")
	}
	if len(rs[0]) != 1 {
		t.Error("wrong number of rides for vehicle 0")
	}
	if rs[1][0] != 2 {
		t.Error("wrong first ride for vehicle 1")
	}
}

func TestRidesScore(t *testing.T) {
	problem := `3 4 2 3 2 10
0 0 1 3 2 9
1 2 1 0 0 9
2 0 2 2 0 9
`
	rp := RidesProblem{}
	err := rp.Parse(problem)
	if err != nil {
		t.Errorf("unexpected error when parsing problem: %v", err)
	}
	solution := `1 0
2 2 1
`
	rs := RidesSolution{}
	err = rs.Parse(solution)
	if err != nil {
		t.Errorf("unexpected error when parsing solution: %v", err)
	}
	score, err := RidesScore(&rp, &rs)
	if err != nil {
		t.Errorf("unexpected error when computing score: %v", err)
	}
	if score != 10 {
		t.Errorf("wrong score, expected 10 got %v", score)
	}
}
