package src

import (
	"sync"
	"math/big"
	"fmt"
)

const (
	maxBallsInCell = 50000
)

type CellsArray struct {
	cellsLock sync.Mutex
	indexLock sync.RWMutex
	cells map[string]int
	largestCellIndex string
}

func NewCellsArray() *CellsArray {
	ca := CellsArray{cells: make(map[string]int)}
	return &ca
}

func (ca *CellsArray) updateLargestCellIndex() {
	for key, value := range ca.cells {
		if ca.cells[ca.largestCellIndex] < value {
			ca.largestCellIndex = key
		}
	}
}

func (ca *CellsArray) UpdateCell(cellIndex string, value int) {
	index := big.Int{}
	_, ok := index.SetString(cellIndex, 10)
	strIndex := index.String()
	if !ok {
		return
	}
	ca.cellsLock.Lock()
	if _, ok := ca.cells[strIndex]; ok {
		ca.cells[strIndex] += value
		if ca.cells[strIndex] > maxBallsInCell {
			ca.cells[strIndex] = 0
		}
	} else {
		fmt.Println("Index: ", strIndex)
		ca.cells[strIndex] = value
	}
	ca.cellsLock.Unlock()

	ca.indexLock.Lock()
	defer ca.indexLock.Unlock()
	if len(ca.cells) == 1 {
		ca.largestCellIndex = strIndex
	} else {
		ca.updateLargestCellIndex()
	}
}

func (ca *CellsArray) GetLargestCellIndex() string {
	ca.indexLock.Lock()
	defer ca.indexLock.Unlock()

	return ca.largestCellIndex
}
