package object

// DroneAdd is request for add new drone
type DroneAdd struct {
	Quadrant int    `json:"quadrant"`
	X        string `json:"x"`
	Y        string `json:"y"`
}
