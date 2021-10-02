package czml

// Orientation is a rotation that takes a vector expresxsed in the "body" axes of the object and
// transforms it to the Earth fixed axes.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Orientation
type Orientation struct {
	UnitQuaternion    *UnitQuaternionValue    `json:"unitQuaternion,omitempty"`
	Reference         ReferenceValue          `json:"reference,omitempty"`
	VelocityReference *VelocityReferenceValue `json:"velocityReference,omitempty"`
}
