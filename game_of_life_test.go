package gameoflife

import "testing"
import "fmt"
import "time"

func TestSampleReplaceScreen(t *testing.T) {

	fmt.Println()
	fmt.Println()
	for _, i := range []int{1, 2, 3} {
		fmt.Printf("\033[0;0H")
		fmt.Println()
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("On %d/10 Off: %d", i, i)
		fmt.Println()
		fmt.Printf("On %d/12 Off: %d", i, i)
	}
}
