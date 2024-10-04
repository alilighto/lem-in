package src

func GetShortestPaths(AllPaths [][]string) (ShortestPaths [][]string) {
	PathsLenght := len(AllPaths)
	for i := 0; i < PathsLenght; i++ {
		AppendIt := true
		for j := 0; j < len(ShortestPaths); j++ {
			if MatchAnyRoom(AllPaths[i], ShortestPaths[j]) {
				AppendIt = false
				break
			}
		}
		if AppendIt {
			ShortestPaths = append(ShortestPaths, AllPaths[i])
		}
	}
	return
}

func MatchAnyRoom(RoomsOne, RoomsTwo []string) bool {
	for i := 1; i < len(RoomsOne)-1; i++ {
		for j := 1; j < len(RoomsTwo)-1; j++ {
			if RoomsOne[i] == RoomsTwo[j] && i != 0 && i != len(RoomsOne)-1 {
				return true
			}
		}
	}
	return false
}