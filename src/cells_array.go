package src

import (
	"fmt"
	"math/big"
	"sync"
)

const (
	MAX_BALLS_IN_CELL = 50000
)

type CellsArray struct {
	cellsLock        sync.Mutex
	indexLock        sync.RWMutex
	cells            map[string]int
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
		} else if ca.cells[ca.largestCellIndex] == value {
			largestIndex := big.Int{}
			keyIndex := big.Int{}

			largestIndex.SetString(ca.largestCellIndex, 10)
			keyIndex.SetString(key, 10)

			if largestIndex.Cmp(&keyIndex) == 1 {
				ca.largestCellIndex = key
			}
		}
	}
}

func (ca *CellsArray) UpdateCell(cellIndex string, value int) {
	index := big.Int{}
	_, ok := index.SetString(cellIndex, 10)
	if !ok {
		return
	}
	strIndex := index.String()
	ca.cellsLock.Lock()
	if _, ok := ca.cells[strIndex]; ok {
		ca.cells[strIndex] += value
		if ca.cells[strIndex] > MAX_BALLS_IN_CELL {
			ca.cells[strIndex] = 0
		}
	} else {
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

func (ca *CellsArray) DumpCells() {
	ca.cellsLock.Lock()
	defer ca.cellsLock.Unlock()

	for key, value := range ca.cells {
		fmt.Println(key, ":\t", value)
	}
}
