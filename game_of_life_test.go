package gameoflife

import "testing"
import "fmt"
import "time"

func TestSampleReplaceScreen(t *testing.T) {

	// reen = "\033[01;32m"
	red := "\033[01;31m"
	// yello = "\033[01;33m"
	// reset = "\033[00m"

	fmt.Println()
	fmt.Println()
	for _, i := range []int{1, 2, 3} {
		fmt.Printf("\033[0;0H")
		fmt.Println()
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("On %d/10 Off: %d", i, i)
		fmt.Println(red)
		row := "4"
		colums := "3"
		fmt.Printf("\033["+row+";"+colums+"H")
		fmt.Printf("On %d/12 Off: %d", i, i)
	}
}

func TestRuleOneHasZeroLifeNeighboursShouldBeDead(t *testing.T) {
	nextState := getNextCellState(Life, 0)
	assertDead(t, nextState)
}

func TestRuleTwoHasTwoLifeNeighboursShouldBeLife(t *testing.T) {
	nextState := getNextCellState(Life, 2)
	assertLife(t, nextState)
}

func TestRuleTwoHasOneLifeNeighboursShouldBeDead(t *testing.T) {
	nextState := getNextCellState(Life, 1)
	assertDead(t, nextState)
}

func TestRuleTwoHasThreeLifeNeighboursShouldBeLife(t *testing.T) {
	nextState := getNextCellState(Life, 3)
	assertLife(t, nextState)
}

func TestRuleThreeHasMoreThanThreeLifeNeighboursShouldBeDead(t *testing.T) {
	nextState := getNextCellState(Life, 4)
	assertDead(t, nextState)
}

func TestRuleFourDeadCellHasEqualThreeLifeNeighboursShouldBeLife(t *testing.T) {
	nextState := getNextCellState(Dead, 3)
	assertLife(t, nextState)
}

func TestRuleFourDeadCellHasTwoLifeNeighboursShouldBeDead(t *testing.T) {
	nextState := getNextCellState(Dead, 2)
	assertDead(t, nextState)
}

func TestRuleFourDeadCellHasFourLifeNeighboursShouldBeDead(t *testing.T) {
	nextState := getNextCellState(Dead, 4)
	assertDead(t, nextState)
}

func assertLife(t *testing.T, nextState bool) {
	if nextState != Life  {
		t.Error("cell should be Life")
	}
}

func assertDead(t *testing.T, nextState bool) {
	if nextState != Dead  {
		t.Error("cell should be Dead")
	}
}
