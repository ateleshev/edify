package validation

import (
	"fmt"
	"github.com/bbiskup/edify/edifact/msg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var authorSegSeqSpec = []struct {
	descr       string
	segmentIDs  []string
	expectError bool
	errorKind   SegSeqErrorKind
}{
	{"Optional segment in incorrect position",
		[]string{
			"UNH",
			"DTM", // Should come after BGM
			"BGM", "UNT",
		}, true, missingMandatorySegment},

	{"Missing mandatory group 4",
		[]string{
			"UNH",
			"BGM",
			"DTM", // optional
			"UNT",
		}, true, missingGroup},

	/*{"minimal message (only mandatory segments)",
	[]string{
		"UNH", "BGM",
		// Group 1
		"LIN",
		"UNT",
	}, false},*/
	/*

		{
			"Mostly mandatory",
			[]string{
				"UNH", "BGM",
				"DTM", "BUS", // both conditional
				// Group 4
				"LIN",
				"UNT",
			}, false},

		{
			"Mostly mandatory; one conditional group",
			[]string{
				"UNH", "BGM",
				"DTM", "BUS",
				// Group 1
				"LIN",
				// Group 2
				"FII", "CTA", "COM",

				"UNT",
			}, false},

		{
			"Some repeat counts > 1",
			[]string{
				"UNH", "BGM",
				"DTM", "BUS",
				// Group 4
				"LIN", "LIN", "LIN", "LIN",
				// Group 7
				"FII", "CTA", "COM", "COM", "COM",
				"FII", "CTA", "COM", "COM", "COM",

				"UNT",
			}, false},

		{
			"No segments at all",
			[]string{}, true},

		{
			"Missing mandatory segments",
			[]string{"UNH"}, true},

		{
			"First mandatory segment repeated too often",
			[]string{
				"UNH", "UNH", "BGM",
				// Group 1
				"LIN",
				"UNT",
			}, true},

		{
			"group 7 repeated too often",
			[]string{
				"UNH", "BGM",
				"DTM", "BUS",
				// Group 4
				"LIN", "LIN", "LIN", "LIN",
				// Group 7
				"FII", "CTA", "COM", "COM", "COM",
				"FII", "CTA", "COM", "COM", "COM",
				"FII", "CTA", "COM", "COM", "COM",

				"UNT",
			}, true},*/
}

func TestSegSeqValidator1(t *testing.T) {
	msgSpec := getMessageSpec("AUTHOR_D.14B")
	validator, err := NewSegSeqValidator(msgSpec)
	require.Nil(t, err)
	require.NotNil(t, validator)

	for _, spec := range authorSegSeqSpec {
		fmt.Printf("spec: %#v\n", spec)
		segments := mapToSegments(spec.segmentIDs)
		message := msg.NewMessage("AUTHOR", segments)
		err := validator.Validate(message)

		if spec.expectError {
			assert.NotNil(t, err)
			err, ok := err.(SegSeqError)
			assert.True(t, ok)
			assert.Equal(t, spec.errorKind, err.kind)
			fmt.Printf("Expected error was: %s", err)
		} else {
			assert.Nil(t, err)
		}
	}
}
