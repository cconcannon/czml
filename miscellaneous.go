package czml

// PixelOffset is a pixel offset in viewport coordinates. A pixel offset is the number of pixels up
// and to the right to place an element relative to an origin.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PixelOffset
type PixelOffset struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// ArcType is the type of arc
// Valid values are `NONE`, `GEODESIC`, and `RHUMB`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ArcType
type ArcType struct {
	ArcType   string          `json:"arcType,omitempty"`
	Reference *ReferenceValue `json:"reference,omitempty"`
}

// ShadowMode specifies whether or not an object casts or receives shadows from each light source
// when shadows are enabled.
// Valid values are `DISABLED`, `ENABLED`, `CAST_ONLY`, and `RECEIVE_ONLY`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ShadowMode
type ShadowMode string

// DistanceDisplayCondition indicates the visibility of an object based on the distance to the
// camera. The value is specified as two values [NearDistance, FarDistance], in meters.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/DistanceDisplayCondition
type DistanceDisplayCondition []float64

// ClassificationType specifies whether a classification affects terrain, 3D tiles, or both.
// Valid values are `TERRAIN`, `CESIUM_3D_TILE`, and `BOTH`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ClassificationTypeValue
type ClassificationType string

// ReferenceValue represents a reference to another property. References can be used to specify
// that two properties on different objects are in fact, the same property
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ReferenceValue
type ReferenceValue string

// ViewFrom is a suggested initial camera position offset when tracking this object, specified as
// a Cartesian position. Typically defined in the East (x), North (y), Up (z) reference frame
// relative to the object's position, but may use another frame depending on the object's velocity.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ViewFrom
type ViewFrom struct {
	Cartesian *Cartesian3Value `json:"cartesian,omitempty"`
	Reference ReferenceValue   `json:"reference,omitempty"`
}

// VerticalOrigin holds the vertical origin of an element, which can optionally vary over time. It
// controls whether the element is bottom-, center-, or top-aligned with the position.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/VerticalOrigin
type VerticalOrigin struct {
	VerticalOrigin *VerticalOriginValue `json:"verticalOrigin,omitempty"`
	Reference      ReferenceValue       `json:"reference,omitempty"`
}

// VerticalOriginValue is the vertical location of an origin relative to an object's position.
// Valid values are `BASELINE`, `BOTTOM`, `CENTER`, and `TOP`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/VerticalOriginValue
type VerticalOriginValue string

// VelocityRefernceValue is the normalized velocity vector of a position property. The reference
// must be to a `position` property
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/VelocityReferenceValue
type VelocityReferenceValue string

// Uri holds a URI value. The URI can optionally vary with time.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Uri
type Uri struct {
	Uri       *UriValue      `json:"uri,omitempty"`
	Interval  string         `json:"interval,omitempty"`
	Reference ReferenceValue `json:"reference,omitempty"`
}

// UriValue is a URI value
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UriValue
type UriValue string

// UnitSphericalValue is a unit spherical value specified as [Clock, Cone] angles. The clock angle
// is measured in the XY plane from the positive X axis toward the positive Y axis. The cone angle
// is the angle from the positive Z axis toward the negative Z axis. If the array has two elements,
// the value is constant. If it has three or more elements, they are time-tagged samples arranged
// as [Time, Clock, Cone, Time, Clock, Cone, ...], where Time is an ISO 8601 date and time string
// or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitSphericalValue
type UnitSphericalValue []interface{}

// UnitQuaternionValue is set of 4-dimensional coordinates used to represent rotation in
// 3-dimensional space, specified as [X, Y, Z, W]. If the array has four elements, the value is
// constant. If it has five or more elements, they are time-tagged samples arranged as
// [Time, X, Y, Z, W, Time, X, Y, Z, W, ...], where Time is an ISO 8601 date and time string or
// seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitQuaternionValue
type UnitQuaternionValue []float64

// Clock defines a simulated clock when a document is loaded
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Clock
type Clock struct {
	Interval    string   `json:"interval,omitempty"`
	CurrentTime string   `json:"currentTime,omitempty"`
	Multiplier  *float64 `json:"multiplier,omitempty"`
	Range       string   `json:"range,omitempty"`
	Step        string   `json:"step,omitempty"`
}

// NearFarScalar holds a numeric value which will be linearly *interpolated between two values based
// on an object's distance from the camera, in eye coordinates. The computed value will *interpolate
// between the near value and the far value while the camera distance falls between the near
// distance and the far distance, and will be clamped to the near or far value while the distance is
// less than the near distance or greater than the far distance, respectively.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/NearFarScalar
type NearFarScaler struct {
	NearFarScalar *NearFarScalarValue `json:"nearFarScalar,omitempty"`
	Reference     ReferenceValue      `json:"reference,omitempty"`
}

// NearFarScalarValue is near-far scalar value specified as four values [NearDistance, NearValue,
// FarDistance, FarValue]. If the array has four elements, the value is constant. If it has five or
// more elements, they are time-tagged samples arranged as [Time, NearDistance, NearValue,
// FarDistance, FarValue, Time, NearDistance, NearValue, FarDistance, FarValue, ...], where Time is
// an ISO 8601 date and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/NearFarScalarValue
type NearFarScalarValue []interface{}

// TimeIntervalCollection can be a single string or an array of strings
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/TimeIntervalCollection
type TimeIntervalCollection string

// HeightReference holds the height reference of an object, which indicates if the object's position
// is relative to terrain or not.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/HeightReference
type HeightReference struct {
	HeightReference *HeightReferenceValue `json:"heightReference,omitempty"`
	Reference       ReferenceValue        `json:"reference,omitempty"`
}

// HeightReferenceValue is height reference of an object, which indicates if the object's position
// is relative to terrain or not.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/HeightReferenceValue
type HeightReferenceValue string

// HorizontalOrigin is the horizontal origin of an element, which can optionally vary over time. It
// controls whether the element is left-, center-, or right-aligned with the position.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/HorizontalOrigin
type HorizontalOrigin struct {
	HorizontalOrigin *HorizontalOriginValue `json:"horizontalOrigin,omitempty"`
	Reference        ReferenceValue         `json:"reference,omitempty"`
}

// HorizontalOriginValue is the horizontal location of an origin relative to an object's position.
// Eligible values are `LEFT`, `RIGHT`, and `CENTER`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/HorizontalOriginValue
type HorizontalOriginValue string

// EyeOffset is an offset in eye coordinates which can optionally vary over time. Eye coordinates
// are a left-handed coordinate system where the X-axis points toward the viewer's right, the Y-axis
// points up, and the Z-axis points *into the screen.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/EyeOffset
type EyeOffset struct {
	Cartesian *Cartesian3Value `json:"cartesian,omitempty"`
	Reference ReferenceValue   `json:"reference,omitempty"`
}

// CustomProperties represents a key-value mapping
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CustomProperties
type CustomProperties map[string]struct{}

// ReferenceListOfListsValue is a list of lists of references to other properties
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ReferenceListOfListsValue
type ReferenceListOfListsValue []ReferenceValue

// DoubleList is a list of floating-point numbers
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/DoubleList
type DoubleList []float64

// DirectionList is a list of directions
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/DirectionList
type DirectionList struct {
	Spherical     *SphericalListValue      `json:"spherical,omitempty"`
	UnitSpherical *UnitSphericalListValue  `json:"unitSpherical,omitempty"`
	Cartesian     *Cartesian3ListValue     `json:"cartesian,omitempty"`
	UnitCartesian *UnitCartesian3ListValue `json:"unitCartesian,omitempty"`
}

// SphericalListValue is a list of spherical values [Clock, Cone, Magnitude, Clock, Cone,
// Magnitude, ...], with angles in radians and magnitude in meters. The clock angle is measured
// in the XY plane from the positive X axis toward the positive Y axis. The cone angle is the angle
// from the positive Z axis toward the negative Z axis.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/SphericalListValue
type SphericalListValue []interface{}

// UnitSphericalListValue is a list of unit spherical values specified as
// [Clock, Cone, Clock, Cone, ...] angles.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitSphericalListValue
type UnitSphericalListValue []interface{}

// ReferenceListValue is a list of references to other properties
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ReferenceListValue
type ReferenceListValue []string

// AlignedAxis is an aligned axis represented by a unit vector which can optionally vary over time
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/AlignedAxis
type AlignedAxis struct {
	UnitCartesian     *UnitCartesian3Value    `json:"unitCartesian,omitempty"`
	UnitSpherical     *UnitSphericalValue     `json:"unitSpherical,omitempty"`
	Reference         ReferenceValue          `json:"reference,omitempty"`
	VelocityReference *VelocityReferenceValue `json:"velocityReference,omitempty"`
}
