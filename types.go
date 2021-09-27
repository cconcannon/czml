package czml

// Packet describes the graphical properties of a single object in a scene
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Packet
type Packet struct {
	Id                  *string                 `json:"id,omitempty"`
	Delete              *bool                   `json:"delete,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	Parent              *string                 `json:"parent,omitempty"`
	Description         *string                 `json:"description,omitempty"`
	Clock               *Clock                  `json:"clock,omitempty"`
	Version             *string                 `json:"version,omitempty"`
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

// Clock defines a simulated clock when a document is loaded
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Clock
type Clock struct {
	Interval    *string `json:"interval,omitempty"`
	CurrentTime *string `json:"currentTime,omitempty"`
	Multiplier  *int    `json:"multiplier,omitempty"`
	Range       *string `json:"range,omitempty"`
	Step        *string `json:"step,omitempty"`
}

// TimeInterval Collection can be a single string or an array of strings
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/TimeIntervalCollection
type TimeIntervalCollection string

// CustomProperties represents a key-value mapping
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CustomProperties
type CustomProperties map[string]string

// Position defines a position
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Position
type Position struct {
	// TODO
}

type Orientation struct {
	// TODO
}

type ViewFrom struct {
	// TODO
}

type Billboard struct {
	// TODO
}

type Box struct {
	// TODO
}

type Corridor struct {
	// TODO
}

type Cylinder struct {
	// TODO
}

type Ellipse struct {
	// TODO
}

type Ellipsoid struct {
	// TODO
}

type Label struct {
	// TODO
}

type Model struct {
	// TODO
}

type Path struct {
	// TODO
}

type Point struct {
	// TODO
}

type Polygon struct {
	// TODO
}

type Polyline struct {
	// TODO
}

type PolylineVolume struct {
	// TODO
}

type Rectangle struct {
	// TODO
}

type Tileset struct {
	// TODO
}

type Wall struct {
	// TODO
}

type ConicSensor struct {
	// TODO
}

type CustomPatternSensor struct {
	// TODO
}

type RectangularSensor struct {
	// TODO
}

type Fan struct {
	// TODO
}

type Vector struct {
	// TODO
}
