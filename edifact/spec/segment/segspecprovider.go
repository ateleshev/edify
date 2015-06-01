package segment

import (
	"fmt"
	"log"
)

// Provides segment spec by Id
type SegSpecProvider interface {
	Get(id string) *SegSpec
	Len() int
}

type SegSpecMap map[string]*SegSpec

// Regular implementation of SegSpecProvider for production
type SegSpecProviderImpl struct {
	segSpecs SegSpecMap
}

func (p *SegSpecProviderImpl) Get(id string) *SegSpec {
	result := p.segSpecs[id]
	if result == nil {
		// e.g. UNH, UNT are not defined in UNCE specs, because they
		// are not part of the release cycle. Instead, they are defined
		// in part 1 of ISO9735 (file testdata/r1241.txt)
		log.Printf("######################## Missing segment spec: '%s'", id)
		return NewSegSpec(
			id, fmt.Sprintf("missing-%s", id),
			"dummy_function", nil)
	} else {
		return result
	}
}

func (p *SegSpecProviderImpl) Len() int {
	return len(p.segSpecs)
}