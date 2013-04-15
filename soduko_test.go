package soduko_test

import (
	"soduko"
	"testing"
)

func TestSoduko1x1(t *testing.T) {
	soduko := &soduko.Soduko{
		Board: []int{
			0,
		},
		Size: &soduko.SquareSize{X: 1, Y: 1},
	}

	solution := []int{
		1,
	}

	err := soduko.Solve()

	if err != nil {
		t.Errorf("Unexpected error condition: %v, \nwant\n %v", err.Error(), solution)
	}

	if soduko.Board[0] != solution[0] {
		t.Errorf("Unexpected solution\n %v, \nwant\n %v", soduko.Board, solution)
	}
}

func TestSodukoRegular(t *testing.T) {
	soduko := &soduko.Soduko{
		Board: []int{
			0, 3, 5, 0, 0, 2, 0, 4, 0,
			0, 9, 0, 1, 0, 0, 0, 3, 7,
			0, 0, 0, 0, 3, 0, 0, 0, 0,
			3, 5, 9, 0, 7, 0, 4, 0, 0,
			0, 0, 0, 3, 0, 0, 0, 1, 5,
			0, 0, 0, 8, 0, 5, 0, 6, 0,
			0, 0, 4, 2, 0, 0, 0, 0, 0,
			2, 0, 7, 4, 0, 0, 1, 0, 8,
			0, 8, 0, 0, 1, 6, 9, 0, 4,
		},
		Size: &soduko.SquareSize{X: 3, Y: 3},
	}

	solution := []int{
		7, 3, 5, 9, 6, 2, 8, 4, 1,
		6, 9, 2, 1, 8, 4, 5, 3, 7,
		1, 4, 8, 5, 3, 7, 2, 9, 6,
		3, 5, 9, 6, 7, 1, 4, 8, 2,
		8, 2, 6, 3, 4, 9, 7, 1, 5,
		4, 7, 1, 8, 2, 5, 3, 6, 9,
		9, 1, 4, 2, 5, 8, 6, 7, 3,
		2, 6, 7, 4, 9, 3, 1, 5, 8,
		5, 8, 3, 7, 1, 6, 9, 2, 4,
	}

	err := soduko.Solve()

	if err != nil {
		t.Errorf("Unexpected error condition: %v, \nwant\n %v", err.Error(), solution)
	}
	wrong := false
	for i := 0; i < len(soduko.Board); i++ {
		if soduko.Board[i] != solution[i] {
			wrong = true
			break
		}
	}

	if wrong {
		t.Errorf("Unexpected solution\n %v, \nwant\n %v", soduko.Board, solution)
	}

}

func TestSodukoIrregular(t *testing.T) {
	soduko := &soduko.Soduko{
		Board: []int{
			0, 4, 0, 0, 0, 3,
			0, 0, 0, 5, 0, 2,
			0, 1, 0, 2, 0, 0,
			0, 0, 2, 0, 3, 0,
			4, 0, 6, 0, 0, 0,
			1, 0, 0, 0, 5, 0,
		},
		Size: &soduko.SquareSize{X: 3, Y: 2},
	}

	solution := []int{
		2, 4, 5, 6, 1, 3,
		6, 3, 1, 5, 4, 2,
		3, 1, 4, 2, 6, 5,
		5, 6, 2, 1, 3, 4,
		4, 5, 6, 3, 2, 1,
		1, 2, 3, 4, 5, 6,
	}

	err := soduko.Solve()

	if err != nil {
		t.Errorf("Unexpected error condition: %v, \nwant\n %v", err.Error(), solution)
	}
	wrong := false
	for i := 0; i < len(soduko.Board); i++ {
		if soduko.Board[i] != solution[i] {
			wrong = true
			break
		}
	}

	if wrong {
		t.Errorf("Unexpected solution\n %v, \nwant\n %v", soduko.Board, solution)
	}

}

func ExampleSoduko11() {
	soduko := &soduko.Soduko{
		Board: []int{
			0,
		},
		Size: &soduko.SquareSize{X: 1, Y: 1},
	}
	soduko.Solve()
	soduko.ToStdOut()

	//Output: | 1 |
	//
	// 

}
