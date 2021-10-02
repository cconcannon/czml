package czml

// Position defines a position
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Position
type Position struct {
	Epoch               string                    `json:"epoch,omitempty"`
	ReferenceFrame      string                    `json:"referenceFrame,omitempty"`
	Cartesian           *Cartesian3Value          `json:"cartesian,omitempty"`
	CartographicRadians *CartographicRadiansValue `json:"cartographicRadians,omitempty"`
	CartographicDegrees *CartographicDegreesValue `json:"cartographicDegrees,omitempty"`
	CartesianVelocity   *Cartesian3VelocityValue  `json:"cartesianVelocity,omitempty"`
	Reference           ReferenceValue            `json:"reference,omitempty"`
}

// PositionList defines a list of positions
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PositionList
type PositionList struct {
	ReferenceFrame      string                        `json:"referenceFrame,omitempty"`
	Cartesian           *Cartesian3ListValue          `json:"cartesian,omitempty"`
	CartographicRadians *CartographicRadiansListValue `json:"cartographicRadians,omitempty"`
	CartographicDegrees []float64                     `json:"cartographicDegrees,omitempty"`
	References          *ReferenceListValue           `json:"references,omitempty"`
}

// PositionListOfLists is a list of lists of positions
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PositionListOfLists
type PositionListOfLists struct {
	Cartesian           *Cartesian3ListOfListsValue          `json:"cartesian,omitempty"`
	CartographicRadians *CartographicRadiansListOfListsValue `json:"cartographicRadians,omitempty"`
	CartographicDegrees *CartographicDegreesListOfListsValue `json:"cartographicDegrees,omitempty"`
	References          *ReferenceListOfListsValue           `json:"references,omitempty"`
}
