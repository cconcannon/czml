package types

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

// Clock defines a simulated clock when a document is loaded
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Clock
type Clock struct {
	Interval    string   `json:"interval,omitempty"`
	CurrentTime string   `json:"currentTime,omitempty"`
	Multiplier  *float64 `json:"multiplier,omitempty"`
	Range       string   `json:"range,omitempty"`
	Step        string   `json:"step,omitempty"`
}

// TimeInterval Collection can be a single string or an array of strings
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/TimeIntervalCollection
type TimeIntervalCollection string

// CustomProperties represents a key-value mapping
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CustomProperties
type CustomProperties map[string]struct{}

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

// Cartesian3Value is the position specified as a three-dimensional
// Cartesian value [X, Y, Z] in meters relative to the referenceFrame
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3Value
type Cartesian3Value []float64

// CartographicRadiansValue is a geodetic, WGS84 position specified as [Longitude, Latitude, Height]
// where Longitude and Latitude are in radians and Height is in meters. If the array has three
// elements, the value is constant. If it has four or more elements, they are time-tagged samples
// arranged as [Time, Longitude, Latitude, Height, Time, Longitude, Latitude, Height, ...], where
// Time is an ISO 8601 date and time string or seconds since epoch
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRadiansValue
type CartographicRadiansValue []float64

// CartographicDegreesValue is a geodetic, WGS84 position specified as [Longitude, Latitude, Height]
// where Longitude and Latitude are in degrees and Height is in meters. If the array has three
// elements, the value is constant. If it has four or more elements, they are time-tagged samples
// arranged as [Time, Longitude, Latitude, Height, Time, Longitude, Latitude, Height, ...], where
// Time is an ISO 8601 date and time string or seconds since epoch
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicDegreesValue
type CartographicDegreesValue []float64

// Cartesian3VelocityValue is a three-dimensional Cartesian value and its derivative specified as
// [X, Y, Z, dX, dY, dZ]. If the array has six elements, the value is constant. If it has seven or
// more elements, they are time-tagged samples arranged as [Time, X, Y, Z, dX, dY, dZ, Time, X,...],
// where Time is an ISO 8601 date and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3VelocityValue
type Cartesian3VelocityValue []float64

// Orientation is a rotation that takes a vector expresxsed in the "body" axes of the object and
// transforms it to the Earth fixed axes.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Orientation
type Orientation struct {
	UnitQuaternion    *UnitQuaternionValue    `json:"unitQuaternion,omitempty"`
	Reference         ReferenceValue          `json:"reference,omitempty"`
	VelocityReference *VelocityReferenceValue `json:"velocityReference,omitempty"`
}

// UnitQuaternionValue is set of 4-dimensional coordinates used to represent rotation in
// 3-dimensional space, specified as [X, Y, Z, W]. If the array has four elements, the value is
// constant. If it has five or more elements, they are time-tagged samples arranged as
// [Time, X, Y, Z, W, Time, X, Y, Z, W, ...], where Time is an ISO 8601 date and time string or
// seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitQuaternionValue
type UnitQuaternionValue []float64

// VelocityRefernceValue is the normalized velocity vector of a position property. The reference
// must be to a `position` property
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/VelocityReferenceValue
type VelocityReferenceValue string

// ViewFrom is a suggested initial camera position offset when tracking this object, specified as
// a Cartesian position. Typically defined in the East (x), North (y), Up (z) reference frame
// relative to the object's position, but may use another frame depending on the object's velocity.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ViewFrom
type ViewFrom struct {
	Cartesian *Cartesian3Value `json:"cartesian,omitempty"`
	Reference ReferenceValue   `json:"reference,omitempty"`
}

// Billboard is a billboard, or viewport-aligned image. The billboard is positioned in the scene
// by the position property. A billboard is sometimes called a marker.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Billboard
type Billboard struct {
	Show                       *bool                     `json:"show,omitempty"`
	Image                      *Uri                      `json:"image,omitempty"`
	Scale                      *float64                  `json:"scale,omitempty"`
	PixelOffset                *PixelOffset              `json:"pixelOffset,omitempty"`
	EyeOffset                  *EyeOffset                `json:"eyeOffset,omitempty"`
	HorizontalOrigin           *HorizontalOrigin         `json:"horizontalOrigin,omitempty"`
	VerticalOrigin             *VerticalOrigin           `json:"verticalOrigin,omitempty"`
	HeightReference            *HeightReference          `json:"heightReference,omitempty"`
	Color                      *Color                    `json:"color,omitempty"`
	Rotation                   *float64                  `json:"rotation,omitempty"`
	AlignedAxis                *AlignedAxis              `json:"alignedAxis,omitempty"`
	SizeInMeters               *bool                     `json:"sizeInMeters,omitempty"`
	Width                      *float64                  `json:"width,omitempty"`
	Height                     *float64                  `json:"height,omitempty"`
	ScaleByDistance            *NearFarScaler            `json:"scaleByDistance,omitempty"`
	TranslucencyByDistance     *NearFarScaler            `json:"translucencyByDistance,omitempty"`
	PixelOffsetScaleByDistance *NearFarScaler            `json:"pixelOffsetScaleByDistance,omitempty"`
	ImageSubRegion             *BoundingRectangle        `json:"imageSubRegion,omitempty"`
	DistanceDisplayCondition   *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	DisableDepthTestDistance   *float64                  `json:"disableDepthTestDistance,omitempty"`
}

// BoundingRectangle holds a bounding rectangle specified by a corner, width and height.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/BoundingRectangle
type BoundingRectangle struct {
	BoundingRectangle *BoundingRectangleValue `json:"boundingRectangle,omitempty"`
	Reference         ReferenceValue          `json:"reference,omitempty"`
}

// BoundingRectangleValue is a  near-far scalar value specified as four values
// [X, Y, Width, Height]. If the array has four elements, the value is constant. If it has five or
// more elements, they are time-tagged samples arranged as
// [Time, X, Y, Width, Height, Time, X, Y, Width, Height, ...], where Time is an ISO 8601 date and
// time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/BoundingRectangleValue
type BoundingRectangleValue []interface{}

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

// AlignedAxis is an aligned axis represented by a unit vector which can optionally vary over time
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/AlignedAxis
type AlignedAxis struct {
	UnitCartesian     *UnitCartesian3Value    `json:"unitCartesian,omitempty"`
	UnitSpherical     *UnitSphericalValue     `json:"unitSpherical,omitempty"`
	Reference         ReferenceValue          `json:"reference,omitempty"`
	VelocityReference *VelocityReferenceValue `json:"velocityReference,omitempty"`
}

// UnitSphericalValue is a unit spherical value specified as [Clock, Cone] angles. The clock angle
// is measured in the XY plane from the positive X axis toward the positive Y axis. The cone angle
// is the angle from the positive Z axis toward the negative Z axis. If the array has two elements,
// the value is constant. If it has three or more elements, they are time-tagged samples arranged
// as [Time, Clock, Cone, Time, Clock, Cone, ...], where Time is an ISO 8601 date and time string
// or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitSphericalValue
type UnitSphericalValue []interface{}

// UnitCartesian3Value is a three-dimensional unit magnitude Cartesian value specified as [X, Y, Z].
// If the array has three elements, the value is constant. If it has four or more elements, they are
// time-tagged samples arranged as [Time, X, Y, Z, Time, X, Y, Z, ...], where Time is an ISO 8601
// date and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitCartesian3Value
type UnitCartesian3Value []interface{}

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

// EyeOffset is an offset in eye coordinates which can optionally vary over time. Eye coordinates
// are a left-handed coordinate system where the X-axis points toward the viewer's right, the Y-axis
// points up, and the Z-axis points *into the screen.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/EyeOffset
type EyeOffset struct {
	Cartesian *Cartesian3Value `json:"cartesian,omitempty"`
	Reference ReferenceValue   `json:"reference,omitempty"`
}

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

// PixelOffset is a pixel offset in viewport coordinates. A pixel offset is the number of pixels up
// and to the right to place an element relative to an origin.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PixelOffset
type PixelOffset struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// Cartesian2 is two-dimensional Cartesian value specified as [X, Y]. If the array has two elements,
// the value is constant. If it has three or more elements, they are time-tagged samples arranged as
// [Time, X, Y, Time, X, Y, ...], where Time is an ISO 8601 date and time string or seconds since
// epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian2Value
type Cartesian2Value []float64

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

// Box is a closed rectangular cuboid
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Box
type Box struct {
	Show                     *bool                     `json:"show,omitempty"`
	Dimensions               *BoxDimensions            `json:"dimensions"`
	HeightReference          *HeightReference          `json:"heightReference"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// Material is a definition of how a surface is colored or shaded
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Material
type Material struct {
	SolidColor   *SolidColorMaterial   `json:"solidColor,omitempty"`
	Image        *ImageMaterial        `json:"image,omitempty"`
	Grid         *GridMaterial         `json:"grid,omitempty"`
	Stripe       *StripeMaterial       `json:"stripe,omitempty"`
	Checkerboard *CheckerboardMaterial `json:"checkerboard,omitempty"`
}

// BoxDimensions is the width, depth, and height of a box
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/BoxDimensions
type BoxDimensions struct {
	Cartesian *Cartesian3Value `json:"cartesian,omitempty"`
	Reference ReferenceValue   `json:"reference,omitempty"`
}

// Corridor is a shape defined by a centerline and width that conforms to the curvature of the
// globe. It can be placed on the surface or at altitude and can optionally be extruded *into a
// volume.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Corridor
type Corridor struct {
	Show                     *bool                     `json:"show,omitempty"`
	Positions                *PositionList             `json:"positions,omitempty"`
	Width                    *float64                  `json:"width,omitempty"`
	Height                   *float64                  `json:"height,omitempty"`
	HeightReference          *HeightReference          `json:"heightReference,omitempty"`
	ExtrudedHeight           *float64                  `json:"extrudedHeight,omitempty"`
	ExtrudedHeightReference  *HeightReference          `json:"extrudedHeightReference,omitempty"`
	CornerType               *CornerType               `json:"cornerType,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	ClassificationType       ClassificationType        `json:"classificationType,omitempty"`
	ZIndex                   *int                      `json:"zIndex,omitempty"`
}

// CornerType holds the style of a corner
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CornerType
type CornerType struct {
	CornerType CornerTypeValue `json:"cornerType,omitempty"`
	Reference  ReferenceValue  `json:"reference,omitempty"`
}

// CornerTypeValue is the style of a corner
// valid values are `ROUNDED`, `MITERED`, `BEVELED`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CornerTypeValue
type CornerTypeValue string

// Cylinder is a cylinder, truncated cone, or cone defined by a length, top radius, and bottom
// radius.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cylinder
type Cylinder struct {
	Show                     *bool                     `json:"show,omitempty"`
	Length                   *float64                  `json:"length"`
	TopRadius                *float64                  `json:"topRadius"`
	BottomRadius             *float64                  `json:"bottomRadius"`
	HeightReference          *HeightReference          `json:"heightReference,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	NumberOfVerticalLines    *int                      `json:"numberOfVerticalLines,omitempty"`
	Slices                   *int                      `json:"slices,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// Ellipse is a closed curve on or above the surface of the Earth.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Ellipse
type Ellipse struct {
	Show                     *bool                     `json:"show,omitempty"`
	SemiMajorAxis            *float64                  `json:"semiMajorAxis"`
	SemiMinorAxis            *float64                  `json:"semiMinorAxis"`
	Height                   *float64                  `json:"height,omitempty"`
	HeightReference          *HeightReference          `json:"heightReference,omitempty"`
	ExtrudedHeight           *float64                  `json:"extrudedHeight,omitempty"`
	ExtrudedHeightReference  *HeightReference          `json:"extrudedHeightReference,omitempty"`
	Rotation                 *float64                  `json:"rotation,omitempty"`
	StRotation               *float64                  `json:"stRotation,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	NumberOfVerticalLines    *int                      `json:"numberOfVerticalLines,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	ClassificationType       ClassificationType        `json:"classificationType,omitempty"`
	ZIndex                   *int                      `json:"zIndex,omitempty"`
}

// Ellipsoid is a closed quadric surface that is a three-dimensional analogue of an ellipse.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Ellipsoid
type Ellipsoid struct {
	Show                     *bool                     `json:"show,omitempty"`
	Radii                    *EllipsoidRadii           `json:"radii"`
	InnerRadii               *EllipsoidRadii           `json:"innerRadii,omitempty"`
	MinimumClock             *float64                  `json:"minimumClock,omitempty"`
	MaximumClock             *float64                  `json:"maximumClock,omitempty"`
	MinimumCone              *float64                  `json:"minimumCone,omitempty"`
	MaximumCone              *float64                  `json:"maximumCone,omitempty"`
	HeightReference          *HeightReference          `json:"heightReference,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	StackPartitions          *int                      `json:"stackPartitions,omitempty"`
	SlicePartitions          *int                      `json:"slicePartitions,omitempty"`
	Subdivisions             *int                      `json:"subdivisions,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// EllipsoidRadii is the radii of an ellipsoid
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/EllipsoidRadii
type EllipsoidRadii struct {
	Cartesian Cartesian3Value `json:"cartesian,omitempty"`
	Reference ReferenceValue  `json:"reference,omitempty"`
}

// Label is a string of text
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Label
type Label struct {
	Show                       *bool                     `json:"show,omitempty"`
	Text                       string                    `json:"text,omitempty"`
	Font                       *Font                     `json:"font,omitempty"`
	Style                      *LabelStyle               `json:"style,omitempty"`
	Scale                      *float64                  `json:"scale,omitempty"`
	ShowBackground             *bool                     `json:"showBackground,omitempty"`
	BackgroundColor            *Color                    `json:"backgroundColor,omitempty"`
	BackgroundPadding          *BackgroundPadding        `json:"backgroundPadding,omitempty"`
	PixelOffset                *PixelOffset              `json:"pixelOffset,omitempty"`
	EyeOffset                  *EyeOffset                `json:"eyeOffset,omitempty"`
	HorizontalOrigin           *HorizontalOrigin         `json:"horizontalOrigin,omitempty"`
	VerticalOrigin             *VerticalOrigin           `json:"verticalOrigin,omitempty"`
	HeightReference            *HeightReference          `json:"heightReference,omitempty"`
	FillColor                  *Color                    `json:"fillColor,omitempty"`
	OutlineColor               *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth               *float64                  `json:"outlineWidth,omitempty"`
	TranslucencyByDistance     *NearFarScaler            `json:"translucencyByDistance,omitempty"`
	PixelOffsetScaleByDistance *NearFarScaler            `json:"pixelOffsetScaleByDistance,omitempty"`
	ScaleByDistance            *NearFarScaler            `json:"scaleByDistance,omitempty"`
	DistanceDisplayCondition   *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	DisableDepthTestDistance   *float64                  `json:"disableDepthTestDistance,omitempty"`
}

// BackgroundPadding describes the amount of horizontal and vertical padding, in pixels, between a
// label's text and its background.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/BackgroundPadding
type BackgroundPadding struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// LabelStyle describes the style of a label
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LabelStyle
type LabelStyle struct {
	LabelStyle LabelStyleValue `json:"labelStyle,omitempty"`
	Reference  ReferenceValue  `json:"reference,omitempty"`
}

// LabelStyleValue is the style of a label
// valid values are `FILL`, `OUTLINE`, and `FILL_AND_OUTLINE`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LabelStyleValue
type LabelStyleValue string

// Font describes a font used to draw text. Fonts are specified using the same syntax as the CSS
// "font" property.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Font
type Font struct {
	Font      FontValue      `json:"font,omitempty"`
	Reference ReferenceValue `json:"reference,omitempty"`
}

// FontValue is a font, specified using the same syntax as the CSS "font" property.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/FontValue
type FontValue string

// Model describes a 3D model
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Model
type Model struct {
	Show                      *bool                     `json:"show,omitempty"`
	Gltf                      *Uri                      `json:"uri"`
	Scale                     *float64                  `json:"scale,omitempty"`
	MinimumPixelSize          *float64                  `json:"minimumPixelSize,omitempty"`
	MaximumScale              *float64                  `json:"maximumScale,omitempty"`
	MinimumCone               *float64                  `json:"minimumCone,omitempty"`
	IncrementallyLoadTextures *bool                     `json:"incrementallyLoadTextures,omitempty"`
	RunAnimations             *bool                     `json:"runAnimations,omitempty"`
	Shadows                   ShadowMode                `json:"shadows,omitempty"`
	HeightReference           *HeightReference          `json:"heightReference,omitempty"`
	SilhouetteColor           *Color                    `json:"silhouetteColor,omitempty"`
	SilhouetteSize            *float64                  `json:"silhouetteSize,omitempty"`
	Color                     *Color                    `json:"color,omitempty"`
	ColorBlendMode            *ColorBlendMode           `json:"colorBlendMode,omitempty"`
	ColorBlendAmount          *float64                  `json:"colorBlendAmount,omitempty"`
	DistanceDisplayCondition  *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	NodeTransformations       *NodeTransformations      `json:"nodeTransformations,omitempty"`
	Articulations             *Articulations            `json:"articulations,omitempty"`
}

// Articulations is a mapping of keys to articulation values, where the keys are the name of the
// articulation, a single space, and the name of the stage.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Articulations
type Articulations map[string]interface{}

// NodeTransformations is a mapping of node names to node transformations
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/NodeTransformations
type NodeTransformations map[string]interface{}

// ColorBlendMode contains the mode of blending between a target color and an entity's source color
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ColorBlendMode
type ColorBlendMode struct {
	ColorBlendMode ColorBlendModeValue `json:"colorBlendMode,omitempty"`
	Reference      ReferenceValue      `json:"reference,omitempty"`
}

// ColorBlendModeValue is the mode of blending between a target color and an entity's source color.
// Valid values are `HIGHLIGHT`, `REPLACE`, and `MIX`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ColorBlendModeValue
type ColorBlendModeValue string

// Path is a polyline defined by the motion of an object over time. The possible vertices of the
// path are specified by the position property. Note that because clients cannot render a truly
// infinite path, the path must be limited, either by defining availability for this object, or by
// using the leadTime and trailTime properties.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Path
type Path struct {
	Show                     *bool                     `json:"show,omitempty"`
	LeadTime                 *float64                  `json:"leadTime,omitempty"`
	TrailTime                *float64                  `json:"trailTime,omitempty"`
	Width                    *float64                  `json:"width,omitempty"`
	Resolution               *float64                  `json:"resolution,omitempty"`
	Material                 *PolylineMaterial         `json:"material,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// Point is a viewport-aligned circle.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Point
type Point struct {
	Show                     *bool                     `json:"show,omitempty"`
	PixelSize                *float64                  `json:"pixelSize,omitempty"`
	HeightReference          *HeightReference          `json:"heightReference"`
	Color                    *Color                    `json:"color,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	ScaleByDistance          *NearFarScaler            `json:"scaleByDistance,omitempty"`
	TranslucencyByDistance   *NearFarScaler            `json:"translucencyByDistance,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	DisableDepthTestDistance *float64                  `json:"disableDepthTestDistance,omitempty"`
}

// Polygon is a closed figure on the surface of the Earth.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Polygon
type Polygon struct {
	Show                     *bool                     `json:"show,omitempty"`
	Positions                *PositionList             `json:"positions,omitempty"`
	Holes                    *PositionListOfLists      `json:"holes,omitempty"`
	ArcType                  *ArcType                  `json:"arcType,omitempty"`
	Height                   *float64                  `json:"height,omitempty"`
	HeightReference          *HeightReference          `json:"heightReference,omitempty"`
	ExtrudedHeight           *float64                  `json:"extrudedHeight,omitempty"`
	ExtrudedHeightReference  *HeightReference          `json:"extrudedHeightReference,omitempty"`
	StRotation               *float64                  `json:"stRotation,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	PerPositionHeight        *bool                     `json:"perPositionHeight,omitempty"`
	CloseTop                 *bool                     `json:"closeTop,omitempty"`
	CloseBottom              *bool                     `json:"closeBottom,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	ClassificationType       ClassificationType        `json:"classificationType,omitempty"`
	ZIndex                   *int                      `json:"zIndex,omitempty"`
}

// PositionListOfLists is a list of lists of positions
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PositionListOfLists
type PositionListOfLists struct {
	Cartesian           *Cartesian3ListOfListsValue          `json:"cartesian,omitempty"`
	CartographicRadians *CartographicRadiansListOfListsValue `json:"cartographicRadians,omitempty"`
	CartographicDegrees *CartographicDegreesListOfListsValue `json:"cartographicDegrees,omitempty"`
	References          *ReferenceListOfListsValue           `json:"references,omitempty"`
}

// Cartesian3ListOfListsValue is a list of lists of three-dimensional Cartesian values specified
// as [X, Y, Z, X, Y, Z, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3ListOfListsValue
type Cartesian3ListOfListsValue []Cartesian3Value

// CartographicRadiansListOfListsValue is a list of lists of geodetic, WGS84 positions specified as
// [Longitude, Latitude, Height, Longitude, Latitude, Height, ...], where Longitude and Latitude are
// in radians and Height is in meters.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRadiansListOfListsValue
type CartographicRadiansListOfListsValue []CartographicRadiansValue

// CartographicDegreesListOfListsValue is a list of lists of geodetic, WGS84 positions specified as
// [Longitude, Latitude, Height, Longitude, Latitude, Height, ...], where Longitude and Latitude are
// in degrees and Height is in meters.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicDegreesListOfListsValue
type CartographicDegreesListOfListsValue []CartographicDegreesValue

// ReferenceListOfListsValue is a list of lists of references to other properties
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ReferenceListOfListsValue
type ReferenceListOfListsValue []ReferenceValue

// Polyline is a line in the scene composed of multiple segments.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Polyline
type Polyline struct {
	Show                     *bool                     `json:"show,omitempty"`
	Positions                *PositionList             `json:"positions"`
	ArcType                  *ArcType                  `json:"arcType,omitempty"`
	Width                    *float64                  `json:"width,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Material                 *PolylineMaterial         `json:"material,omitempty"`
	FollowSurface            *bool                     `json:"followSurface,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DepthFailMaterial        *PolylineMaterial         `json:"depthFailMaterial,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	ClampToGround            *bool                     `json:"clampToGround,omitempty"`
	ClassificationType       ClassificationType        `json:"classificationType,omitempty"`
	ZIndex                   *int                      `json:"zIndex,omitempty"`
}

// PolylineVolume is a polyline with a volume, defined as a 2D shape extruded along a polyline
// that conforms to the curvature of the globe.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineVolume
type PolylineVolume struct {
	Show                     *bool                     `json:"show,omitempty"`
	Positions                *PositionList             `json:"positions"`
	Shape                    *Shape                    `json:"shape"`
	CornerType               *CornerType               `json:"cornerType,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *PolylineMaterial         `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// Shape is a list of two-dimensional positions defining a shape.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Shape
type Shape struct {
	Cartesian2 *Cartesian2ListValue `json:"cartesian2"`
}

// Cartesian2ListValue is a list of two-dimensional Cartesian values specified as
// [X, Y, X, Y, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian2ListValue
type Cartesian2ListValue *[]float64

// Rectangle is a cartographic rectangle, which conforms to the curvature of the globe and can be
// placed on the surface or at altitude and can optionally be extruded into a volume.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Rectangle
type Rectangle struct {
	Show                     *bool                     `json:"show,omitempty"`
	Coordinates              *RectangleCoordinates     `json:"coordinates"`
	Height                   *float64                  `json:"height,omitempty"`
	HeightReference          *HeightReference          `json:"heightReference,omitempty"`
	ExtrudedHeight           *float64                  `json:"extrudedHeight,omitempty"`
	ExtrudedHeightReference  *HeightReference          `json:"extrudedHeightReference,omitempty"`
	Rotation                 *float64                  `json:"rotation,omitempty"`
	StRotation               *float64                  `json:"stRotation,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
	ClassificationType       ClassificationType        `json:"classificationType,omitempty"`
	ZIndex                   *int                      `json:"zIndex,omitempty"`
}

// RectangleCoordinates is a set of coordinates describing a cartographic rectangle on the surface
// of the ellipsoid.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/RectangleCoordinates
type RectangleCoordinates struct {
	Wsen        *CartographicRectangleRadiansValue `json:"wsen,omitempty"`
	WsenDegrees *CartographicRectangleDegreesValue `json:"wsenDegrees,omitempty"`
	Reference   ReferenceValue                     `json:"reference"`
}

// CartographicRectangleRadiansValue is a two-dimensional region specified as [WestLongitude,
// SouthLatitude, EastLongitude, NorthLatitude], with values in radians. If the array has four
// elements, the value is constant. If it has five or more elements, they are time-tagged samples
// arranged as [Time, WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, Time,
// WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, ...], where Time is an ISO 8601 date
// and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRectangleRadiansValue
type CartographicRectangleRadiansValue []interface{}

// CartographicRectangleDegreesValue is a two-dimensional region specified as [WestLongitude,
// SouthLatitude, EastLongitude, NorthLatitude], with values in degrees. If the array has four
// elements, the value is constant. If it has five or more elements, they are time-tagged samples
// arranged as [Time, WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, Time,
// WestLongitude, SouthLatitude, EastLongitude, NorthLatitude, ...], where Time is an ISO 8601 date
// and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRectangleDegreesValue
type CartographicRectangleDegreesValue []interface{}

// Tileset is a 3D Tiles tileset
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Tileset
type Tileset struct {
	Show                    *bool    `json:"show,omitempty"`
	Uri                     *Uri     `json:"uri"`
	MaximumScreenSpaceError *float64 `json:"maximumScreenSpaceError,omitempty"`
}

// Wall is a  two-dimensional wall defined as a line strip and optional maximum and minimum heights,
// which conforms to the curvature of the globe and can be placed along the surface or at altitude.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Wall
type Wall struct {
	Show                     *bool                     `json:"show,omitempty"`
	Positions                *PositionList             `json:"positions"`
	MinimumHeights           *DoubleList               `json:"minimumHeights,omitempty"`
	MaximumHeights           *DoubleList               `json:"maximumHeights,omitempty"`
	Granularity              *float64                  `json:"granularity,omitempty"`
	Fill                     *bool                     `json:"fill,omitempty"`
	Material                 *Material                 `json:"material,omitempty"`
	Outline                  *bool                     `json:"outline,omitempty"`
	OutlineColor             *Color                    `json:"outlineColor,omitempty"`
	OutlineWidth             *float64                  `json:"outlineWidth,omitempty"`
	Shadows                  ShadowMode                `json:"shadows,omitempty"`
	DistanceDisplayCondition *DistanceDisplayCondition `json:"distanceDisplayCondition,omitempty"`
}

// DoubleList is a list of floating-point numbers
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/DoubleList
type DoubleList []float64

// ConicSensor is a conical sensor volume taking into account occlusion of an ellipsoid, i.e.,
// the globe.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ConicSensor
type ConicSensor struct {
	Show                            *bool                         `json:"show,omitempty"`
	InnerHalfAngle                  *float64                      `json:"innerHalfAngle,omitempty"`
	OuterHalfAngle                  *float64                      `json:"outerHalfAngle,omitempty"`
	MinimumClockAngle               *float64                      `json:"minimumClockAngle,omitempty"`
	MaximumClockAngle               *float64                      `json:"maximumClockAngle,omitempty"`
	Radius                          *float64                      `json:"radius,omitempty"`
	ShowIntersection                *bool                         `json:"showIntersection,omitempty"`
	IntersectionColor               *Color                        `json:"intersectionColor,omitempty"`
	IntersectionWidth               *float64                      `json:"intersectionWidth,omitempty"`
	ShowLateralSurfaces             *bool                         `json:"showLateralSurfaces,omitempty"`
	LateralSurfaceMaterial          *Material                     `json:"lateralSurfaceMaterial,omitempty"`
	ShowEllipsoidSurfaces           *bool                         `json:"showEllipsoidSurfaces,omitempty"`
	EllipsoidSurfaceMaterial        *Material                     `json:"ellipsoidSurfaceMaterial,omitempty"`
	ShowEllipsoidHorizonSurfaces    *bool                         `json:"showEllipsoidHorizonSurfaces,omitempty"`
	EllipsoidHorizonSurfaceMaterial *Material                     `json:"ellipsoidHorizonSurfaceMaterial,omitempty"`
	ShowDomeSurfaces                *bool                         `json:"showDomeSurfaces,omitempty"`
	DomeSurfaceMaterial             *Material                     `json:"domeSurfaceMaterial,omitempty"`
	PortionToDisplay                *SensorVolumePortionToDisplay `json:"portionToDisplay"`
	EnvironmentConstraint           *bool                         `json:"environmentConstraint,omitempty"`
	ShowEnvironmentOcclusion        *bool                         `json:"showEnvironmentOcclusion,omitempty"`
	EnvironmentOcclusionMaterial    *Material                     `json:"environmentOcclusionMaterial,omitempty"`
	ShowEnvironmentIntersection     *bool                         `json:"showEnvironmentIntersection,omitempty"`
	EnvironmentIntersectionColor    *Color                        `json:"environmentIntersectionColor,omitempty"`
	EnvironmentIntersectionWidth    *float64                      `json:"environmentIntersectionWidth,omitempty"`
	ShowThroughEllipsoid            *bool                         `json:"showThroughEllipsoid,omitempty"`
	ShowViewshed                    *bool                         `json:"showViewshed,omitempty"`
	ViewshedVisibleColor            *Color                        `json:"viewshedVisibleColor,omitempty"`
	ViewshedOccludedColor           *Color                        `json:"viewshedOccludedColor,omitempty"`
	ViewshedResolution              *int                          `json:"viewshedResolution,omitempty"`
}

// SensorVolumePortionToDisplay is the part of a sensor that should be displayed
// Valid values are `COMPLETE`, `BELOW_ELLIPSOID_HORIZON`, and `ABOVE_ELLIPSOID_HORIZON`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/SensorVolumePortionToDisplay
type SensorVolumePortionToDisplay string

// CustomPatternSensor is a custom sensor volume taking into account occlusion of an ellipsoid,
// i.e., the globe.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CustomPatternSensor
type CustomPatternSensor struct {
	Show                            *bool                         `json:"show,omitempty"`
	Directions                      *DirectionList                `json:"directions"`
	Radius                          *float64                      `json:"radius,omitempty"`
	ShowIntersection                *bool                         `json:"showIntersection,omitempty"`
	IntersectionColor               *Color                        `json:"intersectionColor,omitempty"`
	IntersectionWidth               *float64                      `json:"intersectionWidth,omitempty"`
	ShowLateralSurfaces             *bool                         `json:"showLateralSurfaces,omitempty"`
	LateralSurfaceMaterial          *Material                     `json:"lateralSurfaceMaterial,omitempty"`
	ShowEllipsoidSurfaces           *bool                         `json:"showEllipsoidSurfaces,omitempty"`
	EllipsoidSurfaceMaterial        *Material                     `json:"ellipsoidSurfaceMaterial,omitempty"`
	ShowEllipsoidHorizonSurfaces    *bool                         `json:"showEllipsoidHorizonSurfaces,omitempty"`
	EllipsoidHorizonSurfaceMaterial *Material                     `json:"ellipsoidHorizonSurfaceMaterial,omitempty"`
	ShowDomeSurfaces                *bool                         `json:"showDomeSurfaces,omitempty"`
	DomeSurfaceMaterial             *Material                     `json:"domeSurfaceMaterial,omitempty"`
	PortionToDisplay                *SensorVolumePortionToDisplay `json:"portionToDisplay"`
	EnvironmentConstraint           *bool                         `json:"environmentConstraint,omitempty"`
	ShowEnvironmentOcclusion        *bool                         `json:"showEnvironmentOcclusion,omitempty"`
	EnvironmentOcclusionMaterial    *Material                     `json:"environmentOcclusionMaterial,omitempty"`
	ShowEnvironmentIntersection     *bool                         `json:"showEnvironmentIntersection,omitempty"`
	EnvironmentIntersectionColor    *Color                        `json:"environmentIntersectionColor,omitempty"`
	EnvironmentIntersectionWidth    *float64                      `json:"environmentIntersectionWidth,omitempty"`
	ShowThroughEllipsoid            *bool                         `json:"showThroughEllipsoid,omitempty"`
	ShowViewshed                    *bool                         `json:"showViewshed,omitempty"`
	ViewshedVisibleColor            *Color                        `json:"viewshedVisibleColor,omitempty"`
	ViewshedOccludedColor           *Color                        `json:"viewshedOccludedColor,omitempty"`
	ViewshedResolution              *int                          `json:"viewshedResolution,omitempty"`
}

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

// UnitCartesian3ListValue is a list of three-dimensional unit magnitude Cartesian values,
// specified as [X, Y, Z, X, Y, Z, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/UnitCartesian3ListValue
type UnitCartesian3ListValue []interface{}

// RectangularSensor is a rectangular pyramid sensor volume taking into account occlusion of an
// ellipsoid, i.e., the globe.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/RectangularSensor
type RectangularSensor struct {
	Show                            *bool                         `json:"show,omitempty"`
	XHalfAngle                      *float64                      `json:"xHalfAngle,omitempty"`
	YHalfAngle                      *float64                      `json:"yHalfAngle,omitempty"`
	Radius                          *float64                      `json:"radius,omitempty"`
	ShowIntersection                *bool                         `json:"showIntersection,omitempty"`
	IntersectionColor               *Color                        `json:"intersectionColor,omitempty"`
	IntersectionWidth               *float64                      `json:"intersectionWidth,omitempty"`
	ShowLateralSurfaces             *bool                         `json:"showLateralSurfaces,omitempty"`
	LateralSurfaceMaterial          *Material                     `json:"lateralSurfaceMaterial,omitempty"`
	ShowEllipsoidSurfaces           *bool                         `json:"showEllipsoidSurfaces,omitempty"`
	EllipsoidSurfaceMaterial        *Material                     `json:"ellipsoidSurfaceMaterial,omitempty"`
	ShowEllipsoidHorizonSurfaces    *bool                         `json:"showEllipsoidHorizonSurfaces,omitempty"`
	EllipsoidHorizonSurfaceMaterial *Material                     `json:"ellipsoidHorizonSurfaceMaterial,omitempty"`
	ShowDomeSurfaces                *bool                         `json:"showDomeSurfaces,omitempty"`
	DomeSurfaceMaterial             *Material                     `json:"domeSurfaceMaterial,omitempty"`
	PortionToDisplay                *SensorVolumePortionToDisplay `json:"portionToDisplay"`
	EnvironmentConstraint           *bool                         `json:"environmentConstraint,omitempty"`
	ShowEnvironmentOcclusion        *bool                         `json:"showEnvironmentOcclusion,omitempty"`
	EnvironmentOcclusionMaterial    *Material                     `json:"environmentOcclusionMaterial,omitempty"`
	ShowEnvironmentIntersection     *bool                         `json:"showEnvironmentIntersection,omitempty"`
	EnvironmentIntersectionColor    *Color                        `json:"environmentIntersectionColor,omitempty"`
	EnvironmentIntersectionWidth    *float64                      `json:"environmentIntersectionWidth,omitempty"`
	ShowThroughEllipsoid            *bool                         `json:"showThroughEllipsoid,omitempty"`
	ShowViewshed                    *bool                         `json:"showViewshed,omitempty"`
	ViewshedVisibleColor            *Color                        `json:"viewshedVisibleColor,omitempty"`
	ViewshedOccludedColor           *Color                        `json:"viewshedOccludedColor,omitempty"`
	ViewshedResolution              *int                          `json:"viewshedResolution,omitempty"`
}

// Fan starts at a point or apex and extends in a specified list of directions from the apex. Each
// pair of directions forms a face of the fan extending to the specified radius.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Fan
type Fan struct {
	Show               *bool          `json:"show,omitempty"`
	Directions         *DirectionList `json:"directions"`
	Radius             *float64       `json:"radius,omitempty"`
	PerDirectionRadius *bool          `json:"perDirectionRadius,omitempty"`
	Material           *Material      `json:"material,omitempty"`
	Fill               *bool          `json:"fill,omitempty"`
	Outline            *bool          `json:"outline,omitempty"`
	OutlineColor       *Color         `json:"outlineColor,omitempty"`
	OutlineWidth       *float64       `json:"outlineWidth,omitempty"`
	NumberOfRings      *int           `json:"numberOfRings,omitempty"`
}

// Vector defines a graphical vector that originates at the position property and extends in the
// provided direction for the provided length.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Vector
type Vector struct {
	Show                  *bool      `json:"show,omitempty"`
	Color                 *Color     `json:"color,omitempty"`
	Direction             *Direction `json:"direction"`
	Length                *float64   `json:"length,omitempty"`
	MinimumLengthInPixels *float64   `json:"minimumLengthInPixels,omitempty"`
}

// Spherical is a spherical value [Clock, Cone, Magnitude], with angles in radians and magnitude in
// meters. The clock angle is measured in the XY plane from the positive X axis toward the positive
// Y axis. The cone angle is the angle from the positive Z axis toward the negative Z axis. If the
// array has three elements, the value is constant. If it has four or more elements, they are
// time-tagged samples arranged as [Time, Clock, Cone, Magnitude, Time, Clock, Cone, Magnitude, ...]
// where Time is an ISO 8601 date and time string or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/SphericalValue
type SphericalValue []interface{}

// Direction is a unit vector, in world coordinates, that defines a direction
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Direction
type Direction struct {
	Spherical           *SphericalValue     `json:"spherical,omitempty"`
	UnitSpherical       *UnitSphericalValue `json:"unitSpherical,omitempty"`
	Cartesian           *Cartesian3Value    `json:"cartesian,omitempty"`
	UnitCartesian3Value `json:"unitCartesian,omitempty"`
	Reference           ReferenceValue `json:"reference,omitempty"`
}

type PositionList struct {
	ReferenceFrame      string                        `json:"referenceFrame,omitempty"`
	Cartesian           *Cartesian3ListValue          `json:"cartesian,omitempty"`
	CartographicRadians *CartographicRadiansListValue `json:"cartographicRadians,omitempty"`
	CartographicDegrees []float64                     `json:"cartographicDegrees,omitempty"`
	References          *ReferenceListValue           `json:"references,omitempty"`
}

type ArcType struct {
	// TODO
}

type PolylineMaterial struct {
	SolidColor      *SolidColorMaterial      `json:"solidColor,omitempty"`
	PolylineOutline *PolylineOutlineMaterial `json:"polylineOutline,omitempty"`
	PolylineArrow   *PolylineArrowMaterial   `json:"polylineArrow,omitempty"`
	PolylineDash    *PolylineDashMaterial    `json:"polylineDash,omitempty"`
	PolylineGlow    *PolylineGlowMaterial    `json:"polylineGlow,omitempty"`
	Image           *ImageMaterial           `json:"image,omitempty"`
	Grid            *GridMaterial            `json:"grid,omitempty"`
	Stripe          *StripeMaterial          `json:"stripe,omitempty"`
	Checkerboard    *CheckerboardMaterial    `json:"checkerboard,omitempty"`
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

// Cartesian3ListValue is a list of three-dimensional Cartesian values specified as
// [X, Y, Z, X, Y, Z, ...].
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Cartesian3ListValue
type Cartesian3ListValue []float64

// CartographicRadiansListValue is a list of geodetic, WGS84 positions specified as
// [Longitude, Latitude, Height, Longitude, Latitude, Height, ...], where Longitude and Latitude
// are in radians and Height is in meters.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CartographicRadiansListValue
type CartographicRadiansListValue []float64

// ReferenceListValue is a list of references to other properties
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ReferenceListValue
type ReferenceListValue []string

// SolidColorMaterial is a material that fills the surface with a solid color
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/SolidColorMaterial
type SolidColorMaterial struct {
	Color *Color `json:"color,omitempty"`
}

// PolylineOutlineMaterial is a material that fills the surface of a line with an outlined color.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineOutlineMaterial
type PolylineOutlineMaterial struct {
	Color        *Color   `json:"color,omitempty"`
	OutlineColor *Color   `json:"outlineColor,omitempty"`
	OutlineWidth *float64 `json:"outlineWidth,omitempty"`
}

// PolylineArrowMaterial is a material that fills the surface of a line with an arrow.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineArrowMaterial
type PolylineArrowMaterial struct {
	Color *Color `json:"color,omitempty"`
}

// PolylineDashMaterial is a material that fills the surface of a line with a pattern of dashes.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineDashMaterial
type PolylineDashMaterial struct {
	Color       *Color   `json:"color,omitempty"`
	GapColor    *Color   `json:"gapColor,omitempty"`
	DashLength  *float64 `json:"dashLength,omitempty"`
	DashPattern *int     `json:"dashPattern,omitempty"`
}

// PolylineGlowMaterial is a material that fills the surface of a line with a glowing color.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/PolylineGlowMaterial
type PolylineGlowMaterial struct {
	Color      *Color   `json:"color,omitempty"`
	GlowPower  *float64 `json:"glowPower,omitempty"`
	TaperPower *float64 `json:"taperPower,omitempty"`
}

// ImageMaterial is a material that fills the surface with an image
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ImageMaterial
type ImageMaterial struct {
	Image       *Uri    `json:"image,omitempty"`
	Repeat      *Repeat `json:"repeat,omitempty"`
	Color       *Color  `json:"color,omitempty"`
	Transparent *bool   `json:"transparent,omitempty"`
}

// Repeat is the number of times an image repeats along each axis.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Repeat
type Repeat struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// GridMaterial is a material that fills the surface with a two-dimensional grid.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/GridMaterial
type GridMaterial struct {
	Color         *Color         `json:"color,omitempty"`
	CellAlpha     *float64       `json:"cellAlpha,omitempty"`
	LineCount     *LineCount     `json:"lineCount,omitempty"`
	LineThickness *LineThickness `json:"lineThickness,omitempty"`
	LineOffset    *LineOffset    `json:"lineOffset,omitempty"`
}

// LineCount is the number of grid lines along each axis
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LineCount
type LineCount struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// LineThickness is the thickness of grid lines along each axis, in pixels
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LineThickness
type LineThickness struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// LineOffset is the offset of grid lines along each axis, as a percentage from 0 to 1
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/LineOffset
type LineOffset struct {
	Cartesian2 *Cartesian2Value `json:"cartesian2,omitempty"`
	Reference  ReferenceValue   `json:"reference,omitempty"`
}

// StripeMaterial is a material that fills the surface with alternating colors
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/StripeMaterial
type StripeMaterial struct {
	Orientation StripeOrientation `json:"orientation,omitempty"`
	EvenColor   *Color            `json:"evenColor,omitempty"`
	OddColor    *Color            `json:"oddColor,omitempty"`
	Offset      *float64          `json:"offset,omitempty"`
	Repeat      *float64          `json:"repeat,omitempty"`
}

// StripeOrientation describes the orientation of stripes in a stripe material
// Valid values are `HORIZONTAL` and `VERTICAL`
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/StripeOrientation
type StripeOrientation string

// CheckerboardMaterial is a material that fills the surface with a checkerboard pattern.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CheckerboardMaterial
type CheckerboardMaterial struct {
	EvenColor *Color  `json:"evenColor,omitempty"`
	OddColor  *Color  `json:"oddColor,omitempty"`
	Repeat    *Repeat `json:"repeat,omitempty"`
}

// Color describes a color. The color can optionally vary over time
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Color
type Color struct {
	Rgba      RgbaValue      `json:"rgba,omitempty"`
	Rgbaf     RgbafValue     `json:"rgbaf,omitempty"`
	Reference ReferenceValue `json:"reference,omitempty"`
}

// RgbaValue is a color specified as an array of color components [Red, Green, Blue, Alpha] where
// each component is in the range 0-255. If the array has four elements, the color is constant. If
// it has five or more elements, they are time-tagged samples arranged as [Time, Red, Green, Blue,
// Alpha, Time, Red, Green, Blue, Alpha, ...], where Time is an ISO 8601 date and time string or
// seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/RgbaValue
type RgbaValue []int

// RgbafValue is a color specified as an array of color components [Red, Green, Blue, Alpha] where
// each component is in the range 0.0-1.0. If the array has four elements, the color is constant.
// If it has five or more elements, they are time-tagged samples arranged as [Time, Red, Green,
// Blue, Alpha, Time, Red, Green, Blue, Alpha, ...], where Time is an ISO 8601 date and time string
// or seconds since epoch.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/RgbafValue
type RgbafValue []float32

// ReferenceValue represents a reference to another property. References can be used to specify
// that two properties on different objects are in fact, the same property
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/ReferenceValue
type ReferenceValue string
