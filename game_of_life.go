package gameoflife

var Life = true;
var Dead = false;

func getNextCellState(currentState bool, totalLifeNeighbours int) bool {
	if totalLifeNeighbours == 2 {
		return currentState
	}

	if totalLifeNeighbours == 3 {
		return Life
	}

	return Dead
}
