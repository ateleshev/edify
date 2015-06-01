package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMsgSpecPartString(t *testing.T) {
	segSpec := fixtureSimpleSegSpec()
	part := NewMsgSpecSegmentPart(segSpec, 5, true, nil)
	assert.Equal(t, "Segment 5 mand. TESTSEGMENT_NAME", part.String())
	assert.Equal(t, 5, part.MaxCount())
	assert.Equal(t, true, part.IsMandatory())
	assert.Equal(t, "TESTSEGMENT_NAME", part.Name())
}
