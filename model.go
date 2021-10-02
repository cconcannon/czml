package czml

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
