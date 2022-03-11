package main

type TestStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (ts *TestStruct) Shoot() bool {
	if ts.On && ts.Ammo > 0 {
		ts.Ammo--
		return true
	} else {
		return false
	}
}

func (ts *TestStruct) RideBike() bool {
	if ts.On && ts.Power > 0 {
		ts.Power--
		return true
	} else {
		return false
	}
}

func main() {
	//testStruct := &TestStruct{}
}
