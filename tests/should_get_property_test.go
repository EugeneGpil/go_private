package should_get_property_test

import (
	"testing"

	"github.com/EugeneGpil/go_private"
	"github.com/EugeneGpil/tester"
)

func Test_should_get_property(t *testing.T) {
	tester.SetTester(t)

	theStruct := go_private.NewTheStruct(go_private.NewTheStructDto{
		TheProperty: "test",
	})

	tester.AssertSame("test", theStruct.GetProperty())
}
