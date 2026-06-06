package main

type Point struct {
	X, Y int
}

type DirectionManager struct {
	currentPosition Point
	historyStack    []Point
	redoStack       []Point
}

func NewDirectionManager() *DirectionManager {
	return &DirectionManager{
		currentPosition: Point{0, 0},
		historyStack:    []Point{{0, 0}},
	}
}

// Move position
func (dm *DirectionManager) move(direction string) (int, int) {

	// Move Up
	if direction == "U" {
		dm.currentPosition.Y++
	} else if direction == "D" { // Move Down
		dm.currentPosition.Y--
	} else if direction == "L" { // Move Left
		dm.currentPosition.X--
	} else if direction == "R" { // Move Right
		dm.currentPosition.X++
	} else {
		// Invalid direction
		return dm.currentPosition.X, dm.currentPosition.Y
	}

	// Save new position in history
	dm.historyStack = append(dm.historyStack, dm.currentPosition)

	// Clear redo stack after new move
	dm.redoStack = nil

	return dm.currentPosition.X, dm.currentPosition.Y
}

// Undo last move
func (dm *DirectionManager) undoMove() (int, int, bool) {

	// Cannot undo starting point
	if len(dm.historyStack) <= 1 {
		return dm.currentPosition.X, dm.currentPosition.Y, false
	}

	// Take current position
	last := dm.historyStack[len(dm.historyStack)-1]

	// Remove from history
	dm.historyStack = dm.historyStack[:len(dm.historyStack)-1]

	// Save in redo stack
	dm.redoStack = append(dm.redoStack, last)

	// New current position
	dm.currentPosition = dm.historyStack[len(dm.historyStack)-1]

	return dm.currentPosition.X, dm.currentPosition.Y, true
}

// Redo last undone move
func (dm *DirectionManager) redoMove() (int, int, bool) {

	// Nothing to redo
	if len(dm.redoStack) == 0 {
		return dm.currentPosition.X, dm.currentPosition.Y, false
	}

	// Get last redo position
	last := dm.redoStack[len(dm.redoStack)-1]

	// Remove from redo stack
	dm.redoStack = dm.redoStack[:len(dm.redoStack)-1]

	// Add back to history
	dm.historyStack = append(dm.historyStack, last)

	// Update current position
	dm.currentPosition = last

	return dm.currentPosition.X, dm.currentPosition.Y, true
}

// Get all visited points
func (dm *DirectionManager) getPath() []Point {

	// Return copy of history
	path := make([]Point, len(dm.historyStack))
	copy(path, dm.historyStack)

	return path
}
