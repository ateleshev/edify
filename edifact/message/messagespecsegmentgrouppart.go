package message

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/util"
)

// Segment group specification in message specification
type MessageSpecSegmentGroupPart struct {
	MessageSpecPartBase
	Name     string
	children []MessageSpecPart
}

func (p *MessageSpecSegmentGroupPart) String() string {
	mandatoryStr := util.CustBoolStr(p.IsMandatory(), "mand.", "cond.")
	return fmt.Sprintf("Segment group %s %d %s (%d children)", p.Name, p.MaxCount(), mandatoryStr, p.Count())
}

func (p *MessageSpecSegmentGroupPart) IsGroup() bool {
	return true
}

func (p *MessageSpecSegmentGroupPart) Count() int {
	return len(p.children)
}

func (p *MessageSpecSegmentGroupPart) Children() []MessageSpecPart {
	return p.children
}

func (p *MessageSpecSegmentGroupPart) Append(messageSpecPart MessageSpecPart) {
	p.children = append(p.children, messageSpecPart)
}

func NewMessageSpecSegmentGroupPart(name string, children []MessageSpecPart, maxCount int, isMandatory bool) *MessageSpecSegmentGroupPart {
	return &MessageSpecSegmentGroupPart{
		MessageSpecPartBase{
			maxCount:    maxCount,
			isMandatory: isMandatory,
		},
		name,
		children,
	}
}
