package segment

import (
	"fmt"
	de "github.com/bbiskup/edify/edifact/spec/dataelement"
	"github.com/bbiskup/edify/edifact/util"
)

// A data element which is part of a segment specification
// (EDSD)
type SegDataElemSpec struct {
	DataElemSpec de.DataElemSpec
	Count        int
	IsMandatory  bool
}

func (e *SegDataElemSpec) String() string {
	mandatoryStr := util.CustBoolStr(e.IsMandatory, "mand.", "cond.")
	return fmt.Sprintf("SegDataElem %s %dx %s", e.DataElemSpec.Id(), e.Count, mandatoryStr)
}

func NewSegDataElemSpec(
	dataElemSpec de.DataElemSpec, count int, isMandatory bool) *SegDataElemSpec {
	return &SegDataElemSpec{
		DataElemSpec: dataElemSpec,
		Count:        count,
		IsMandatory:  isMandatory,
	}
}
