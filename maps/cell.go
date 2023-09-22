package maps

// Represents the different values of a tile in the map
type Cell int

const (
	Void          = -1
	FlatGround    = 1
	Stairs        = 2
	CrowdedStairs = 3
	CrowdedRoom   = 4
)
