package czml

import (
	"encoding/json"
	"errors"

	"github.com/cconcannon/czml/types"
)

// Czml is a .czml file, which is a valid JSON document that contains an array of packets
type Document struct {
	Packets []types.Packet
}

// Unmarshal accepts raw byte data (i.e. a .czml file) and returns
// the data Unmarshaled into the CZML structured interface
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Marshal accepts a name for the document and a Packet array,
// and returns the Marshaled czml data
func Marshal(d Document) ([]byte, error) {
	if d.Packets[0].Id != "document" {
		return nil, errors.New("the first packet is not the \"document\" packet")
	}
	return json.Marshal(d.Packets)
}

// MarshalIndent accepts a name for the document, a Packet array,
// a prefix string, and an indentation string, and returns the
// Marshaled and Indented czml data
func MarshalIndent(d Document, prefix, indent string) ([]byte, error) {
	if d.Packets[0].Id != "document" {
		return nil, errors.New("the first packet is not the \"document\" packet")
	}
	return json.MarshalIndent(&d.Packets, prefix, indent)
}

// CreateDocument initializes .czml file data with the "document" packet as the first packet
func CreateDocument(name string) Document {
	var packet types.Packet
	packet.Id = "document"
	packet.Name = name
	packet.Version = "1.0"

	return Document{[]types.Packet{packet}}
}

// AddPacket adds a Packet to the .czml document data
func (d *Document) AddPacket(p types.Packet) error {
	d.Packets = append(d.Packets, p)
	return nil
}
