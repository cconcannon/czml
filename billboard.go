package czml

// Billboard is a billboard, or viewport-aligned image. The billboard is positioned in the scene
// by the position property. A billboard is sometimes called a marker.
// https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/Billboard
type Billboard struct {
	Show                       *bool                     `json:"show,omitempty"`
	Image                      string                    `json:"image"`
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
