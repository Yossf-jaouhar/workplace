package functions

func (a *Info) LogicAnts() {
	a.UniquePath()
}

func (a *Info) UniquePath() {
	num := a.NumberOfAnts
	for _, path := range a.AllPaths {
		if len(path) == num+2 {
		}
	}
}
