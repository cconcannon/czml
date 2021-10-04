# github.com/cconcannon/czml

A module for interfacing with `.czml` files and data.

More info about the [Cesium Ecosystem can be found on cesium.com](https://cesium.com)

The [spec for `.czml` files seems to be maintained here.](https://github.com/AnalyticalGraphicsInc/czml-writer/wiki/CZML-Guide)

## Use

### Import the Module

```go
import (
	...

	"github.com/cconcannon/czml"
)
```

### Initialize a CZML data structure

```go
var c czml.Czml

c.InitializeDocument("name")
```

### Create & Add CZML Packets to CZML data

```go
packet := czml.CreateEmptyPacket("id", "name")

c.AddPacket(packet)
```

> **Note: Packet field types and sub-types are available, but constructors and schema-checking are not, so you must use some other means to validate your packet schema**

### Create JSON binary

```go
json, err := czml.Marshal(c)
```

## About the CZML format

- `.czml` files are valid `.json`
- a `.czml` file is an array of `Packet` objects
- the first `Packet` in a `.czml` file must be the `document` Packet

## Opinions in Implementation

Most struct fields are references, and this is by design. The core of this reasoning is that **default values of cesium properties are a mix of zero and non-zero values across numerical and boolean types**. To preserve non-zero (i.e. `1.0`, `true`) implied default values, all fields except for strings have been implemented as reference values so that non-presence does not equate to zero-value.

## Future Development

Most of the type implementations are untested. There's a lot of opportunity for further development here - add testing, implement constructors, etc.

The documentation about value formats for many of the `.czml` properties, especially those around coordinates, is relatively unclear. Many of the types declared here could probably be updated to reflect valid data.

## Contributing

Yes, please. 