package msg

type Segment struct {
	Id       string
	Elements []*DataElement
}

func (s *Segment) String() string {
	elementsStr := ""
	for _, e := range s.Elements {
		elementsStr += "\t\t" + e.String() + "\n"
	}
	return s.Id + "\n" + elementsStr
}

func (s *Segment) AddElement(element *DataElement) {
	s.Elements = append(s.Elements, element)
}

func (s *Segment) AddElements(elements []*DataElement) {
	s.Elements = elements
}

func NewSegment(id string) *Segment {
	return &Segment{id, []*DataElement{}}
}
