package functions


type Info struct {
	NumberOfAnts string
	Start string
	End string
	Rooms map[string][]string
	Tunnels map[string][]string
	AllPaths [][]string
} 	