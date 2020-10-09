package link

import "io"

//Link represents a tag in a HTML Document
type Link struct {
	Href string
	Text string
}
// Parse will take in a HTML document and will return a
// Slice of links parsed from it
func Parse (r io.Reader) ([]Link, error) {
 return nil, nil
}