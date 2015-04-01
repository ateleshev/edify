package dataelement

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/codes"
	"github.com/bbiskup/edify/edifact/util"
)

// DataElement specification
type SimpleDataElementSpec struct {
	Num        int32
	Name       string
	Descr      string
	Repr       *Repr
	CodesSpecs *codes.CodesSpec
}

func (s *SimpleDataElementSpec) String() string {
	return fmt.Sprintf("SimpleDataElementSpec: %d '%s' [%s]", s.Num, s.Name, s.Repr)
}

func NewSimpleDataElementSpec(num int32, name string, descr string, repr *Repr, codes *codes.CodesSpec) (*SimpleDataElementSpec, error) {
	err := util.CheckNotNil(num, name, descr, repr, codes)
	if err != nil {
		return nil, err
	}
	return &SimpleDataElementSpec{
		Num:        num,
		Name:       name,
		Descr:      descr,
		Repr:       repr,
		CodesSpecs: codes,
	}, nil
}
