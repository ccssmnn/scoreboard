package obj

import "testing"

func TestTrafficlightsProblemParse(t *testing.T) {
	problem := `6 4 5 2 1000
2 0 rue-de-londres 1
0 1 rue-d-amsterdam 1
3 1 rue-d-athenes 1
2 3 rue-de-rome 2
1 2 rue-de-moscou 3
4 rue-de-londres rue-d-amsterdam rue-de-moscou rue-de-rome
3 rue-d-athenes rue-de-moscou rue-de-londres
`
	tp := TrafficlightsProblem{}
	err := tp.Parse(problem)
	if err != nil {
		t.Errorf("failed to parse problem: %v", err)
	}
	if len(tp.Routes) != 2 {
		t.Errorf("unexpected number of routes: %v", len(tp.Routes))
	}
	if tp.Routes[0][0] != "rue-de-londres" {
		t.Errorf("unexpected first street in route: %v", tp.Routes[0][0])
	}
	if length := tp.Streets["rue-de-londres"].Length; length != 1 {
		t.Errorf("unexpected length for 'rue-de-londres': %v", length)
	}
}

func TestTrafficlightsSolutionParse(t *testing.T) {
	solution := `4
3
1
rue-de-rome 1
0
1
rue-de-londres 1
1
2
rue-d-amsterdam 1
rue-d-athenes 1
2
1
rue-de-moscou 1`
	ts := TrafficlightsSolution{}
	err := ts.Parse(solution)
	if err != nil {
		t.Errorf("failed to parse solution: %v", ts)
	}
	if len(ts) != 4 {
		t.Errorf("unexpected number of schedules: %v", len(ts))
	}
	if len(ts[0]) != 1 {
		t.Errorf("unexpected length of schedule 0: %v", len(ts[0]))
	}
}
