package src

func (antcolonie *antcolonie) DFS(Path []string, Room string, Visited map[string]bool, res [][]string) [][]string {
	Path = append(Path, Room)
	if Room == antcolonie.EndRoom {
		P := make([]string, len(Path))
		copy(P, Path)
		res = append(res, P)
		return res
	}
	Visited[Room] = true
	for _, GoToRoom := range antcolonie.Rooms[Room] {
		if !Visited[GoToRoom] {
			res = antcolonie.DFS(Path, GoToRoom, Visited, res)
		}
	}
	Visited[Room] = false
	return res
}
