package prototype

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "出去浪"
	proto := ConcretePrototype{name: name}
	newProto := proto.Clone()
	actualName := newProto.Name()

	if !cmp.Equal(name, actualName) {
		t.Error("name: ", name, "actualName: ", actualName)
	}
}
