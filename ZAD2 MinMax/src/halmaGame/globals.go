package halmaGame

// SETTINGS
var _heuristic Heuristic

// STATISTICS
var _turnCounter = 0
var _allVisitedNodesCounter int64 = 0
var _allPrunedNodesCounter int64 = 0

var _visitedNodesCounter int64
var _prunedNodesCounter int64

var _visitedon1Level int64
var _visitedon2Level int64
var _visitedon3Level int64

// FUNCTIONS
func initGlobals() {
	_turnCounter = 0
	_allVisitedNodesCounter = 0
	_allPrunedNodesCounter = 0
	_visitedNodesCounter = 0
	_prunedNodesCounter = 0
	_visitedon1Level = 0
	_visitedon2Level = 0
	_visitedon3Level = 0
}
