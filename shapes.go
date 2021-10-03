package czml

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

// BoxDimensions is the width, depth, and height of a box
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/BoxDimensions
type BoxDimensions struct {
	Cartesian *Cartesian3Value `json:"cartesian,omitempty"`
	Reference ReferenceValue   `json:"reference,omitempty"`
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

// Shape is a list of two-dimensional positions defining a shape.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Shape
type Shape struct {
	Cartesian2 *Cartesian2ListValue `json:"cartesian2"`
}

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
