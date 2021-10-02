package czml

// Packet describes the graphical properties of a single object in a scene
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Packet
type Packet struct {
	Id                  string                  `json:"id,omitempty"`
	Delete              *bool                   `json:"delete,omitempty"`
	Name                string                  `json:"name,omitempty"`
	Parent              string                  `json:"parent,omitempty"`
	Description         string                  `json:"description,omitempty"`
	Clock               *Clock                  `json:"clock,omitempty"`
	Version             string                  `json:"version,omitempty"`
	Availability        *TimeIntervalCollection `json:"availability,omitempty"`
	Properties          *CustomProperties       `json:"properties,omitempty"`
	Position            *Position               `json:"position,omitempty"`
	Orientation         *Orientation            `json:"orientation,omitempty"`
	ViewFrom            *ViewFrom               `json:"viewFrom,omitempty"`
	Billboard           *Billboard              `json:"billboard,omitempty"`
	Box                 *Box                    `json:"box,omitempty"`
	Corridor            *Corridor               `json:"corridor,omitempty"`
	Cylinder            *Cylinder               `json:"cylinder,omitempty"`
	Ellipse             *Ellipse                `json:"ellipse,omitempty"`
	Ellipsoid           *Ellipsoid              `json:"ellipsoid,omitempty"`
	Label               *Label                  `json:"label,omitempty"`
	Model               *Model                  `json:"model,omitempty"`
	Path                *Path                   `json:"path,omitempty"`
	Point               *Point                  `json:"point,omitempty"`
	Polygon             *Polygon                `json:"polygon,omitempty"`
	Polyline            *Polyline               `json:"polyline,omitempty"`
	PolylineVolume      *PolylineVolume         `json:"polylineVolume,omitempty"`
	Rectangle           *Rectangle              `json:"rectangle,omitempty"`
	Tileset             *Tileset                `json:"tileset,omitempty"`
	Wall                *Wall                   `json:"wall,omitempty"`
	ConicSensor         *ConicSensor            `json:"agi_conicSensor,omitempty"`
	CustomPatternSensor *CustomPatternSensor    `json:"agi_customPatternSensor,omitempty"`
	RectangularSensor   *RectangularSensor      `json:"agi_rectangularSensor,omitempty"`
	Fan                 *Fan                    `json:"agi_fan,omitempty"`
	Vector              *Vector                 `json:"agi_vector,omitempty"`
}
