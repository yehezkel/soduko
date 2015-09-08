/*
This package find the solution for soduko boards of square sizes of NxM
where N,M > 0

Take into account that the square size is not the dimensions of the board but
the dimensions of a single square. For an standard soduko of board 9x9
the sizes are 3x3

look at the  test file for usage examples

the algorithm is similiar to the brute force one with some improvements
to save execution time

The idea is to build a SdkCell (Soduko Cell) for each board position
Every cell contain a reference to 3 unique sets representing
the items on the same row, column an square

In order to find the solution a board of cells is created together
with the relations between their respective unique sets

Then a recursive function will pass over the cells setting the values for each
of them
*/
package soduko

import (
	"errors"
	"fmt"
)

//structure to define the sizes of the soduko squares
type SquareSize struct {
	X, Y int
}

/*
Calculate the amount of items on the soduko board
according to the square sizes
*/
func (s *SquareSize) ItemsOnBoard() int {
	return s.X * s.X * s.Y * s.Y
}

/*
Structure that represent the Soduko game
its formed by a slice of int representing the default values
and a pointer to an SquareSize structure
*/
type Soduko struct {
	Board []int
	Size  *SquareSize
}

/*
Find the solution for the Soduko
if no squareSize is set on the soduko an standard size of 3x3 is assumed
If a solution is found the board values will be overriten with it
*/
func (s *Soduko) Solve() error {

	if nil == s.Size {
		s.Size = &SquareSize{3, 3}
	}

	length := len(s.Board)

	if s.Size.X < 1 || s.Size.Y < 1 || length != s.Size.ItemsOnBoard() {
		return errors.New("Wrong soduko board dimensions")
	}

	cells, err := s.buildBoard()

	if err != nil {
		return err
	}

	solved := s.solveBoard(cells, 0)
	if solved {
		for i := 0; i < length; i++ {
			if cells[i].value > 0 {
				s.Board[i] = cells[i].value
			}
		}
		return nil
	}

	return errors.New("Soduko Unbreakable")
}

/*
 Print the Soduko to the standard output
 This is kind of a String function
*/
func (s *Soduko) ToStdOut() {
	length := len(s.Board)

	if nil == s.Size {
		s.Size = &SquareSize{3, 3}
	}

	rows_length := s.Size.X * s.Size.Y

	for i := 0; i < length; i++ {
		fmt.Printf("| %v ", s.Board[i])
		if (i+1)%rows_length == 0 {
			fmt.Println("|")
		}
	}

}

/*
 Build the board of SdkCells to find the solution
*/
func (s *Soduko) buildBoard() ([]*SdkCell, error) {
	length := len(s.Board)
	square_items := s.Size.X * s.Size.Y

	unique_sets := make([]sdkUniqueSet, square_items*3)
	for i := 0; i < len(unique_sets); i++ {
		unique_sets[i] = make(sdkUniqueValues, square_items+1)
	}

	groups := []sdkGroupsInterface{
		//interesting here to know how expensives the following casts are
		SdkRows(unique_sets[:square_items]),
		SdkColumns(unique_sets[square_items : 2*square_items]),
		SdkSquares(unique_sets[2*square_items:]),
	}

	cells := make([]*SdkCell, length)
	for i := range cells {
		cells[i] = new(SdkCell)
		for _, group := range groups {
			group.RegisterCell(cells[i], i, s.Size)
		}

		if s.Board[i] > 0 {
			if !cells[i].Set(s.Board[i]) {
				return nil, errors.New("The initial values are not Soduko Complaint")
			}
			cells[i].Fix()
		}
	}

	return cells, nil
}

func (s *Soduko) solveBoard(cells []*SdkCell, pos int) bool {

	length := len(cells)
	if pos == length {
		return true
	}

	if cells[pos].value < 0 {
		return s.solveBoard(cells, pos+1)
	}

	count := s.Size.X * s.Size.Y

	for i := 0; i < count; i++ {
		if cells[pos].Set(i + 1) {
			if s.solveBoard(cells, pos+1) {
				return true
			}
		}
		cells[pos].Clean()
	}
	return false
}

type sdkUniqueSet interface {
	HasValue(value int) bool
	Set(value int) bool
	Clean(value int)
}

/*
Implementation of the sdkUniqueSet interface using an slice of bools
in order to save memory and runtime, as in order to know if an
item is on the set we need to check if the value of the key
equal to the item is true
*/
type sdkUniqueValues []bool

func (shv sdkUniqueValues) HasValue(value int) bool {

	return value > 0 && value < len(shv) && shv[value]
}

func (shv sdkUniqueValues) Set(value int) bool {

	hasValue := shv.HasValue(value)
	if !hasValue {
		shv[value] = true
	}

	return !hasValue
}

func (shv sdkUniqueValues) Clean(value int) {

	if value > 0 && value < len(shv) {
		shv[value] = false
	}
}

type sdkGroup []sdkUniqueSet

type sdkGroupsInterface interface {
	RegisterCell(cell *SdkCell, position int, size *SquareSize)
}

type SdkCell struct {
	value int
	sets  sdkGroup
}

func (cell *SdkCell) Set(value int) bool {
	allow := true
	for _, set := range cell.sets {
		if set.HasValue(value) {
			allow = false
			break
		}
	}

	if allow {
		for _, set := range cell.sets {
			set.Set(value)
		}
		cell.value = value
	}
	return allow
}

func (cell *SdkCell) Clean() {
	for _, set := range cell.sets {
		set.Clean(cell.value)
	}
	cell.value = 0
}

func (cell *SdkCell) Fix() {
	cell.value *= -1
}

type SdkRows sdkGroup

func (rows SdkRows) RegisterCell(cell *SdkCell, position int, size *SquareSize) {

	cell.sets = append(cell.sets, rows[position/(size.X*size.Y)])
}

type SdkColumns sdkGroup

func (clmns SdkColumns) RegisterCell(cell *SdkCell, position int, size *SquareSize) {
	cell.sets = append(cell.sets, clmns[position%(size.X*size.Y)])
}

type SdkSquares sdkGroup

func (squares SdkSquares) RegisterCell(cell *SdkCell, position int, size *SquareSize) {

	square_items := size.X * size.Y
	index := ((position/square_items)/size.Y)*size.Y + (position%square_items)/size.X
	cell.sets = append(cell.sets, squares[index])
}
