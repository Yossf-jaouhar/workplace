package functions


type Info struct {
	NumberOfAnts int
	Start string
	End string
	Rooms map[string][]string
	Tunnels map[string][]string

	UniquePaths [][]string
	TheBestPaths [][]string
 	NumberOfrooms int
 	Res []string
}


