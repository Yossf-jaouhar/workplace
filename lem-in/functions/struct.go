package functions

type Info struct {
	NumberOfAnts int
	Start        string
	End          string

	Rooms   map[string][]string
	Tunnels map[string][]string

	Paths [][]string

	Res []string


	
	AllGroups    [][][]string

	UniqueGroups [][][]string


	TheBestGroup [][]string


	SG int
	GG int
}

