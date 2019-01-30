package object

// Drone is response with drone info
type Drone struct {
	ID       string `json:"id"`
	Quadrant int    `json:"quadrant"`
	X        string `json:"x"`
	Y        string `json:"y"`
}
