package adapter

import (
	"fmt"
	"testing"
)

func TestAdaptee_SpecificExecute(t *testing.T) {
	tAdaptee := NewAdaptee("HuaWei")
	tAdapter := NewAdapter(tAdaptee)
	tAdapter.Request()

	fmt.Println("=============================")

	tAdaptee = NewAdaptee("XiaoMi")
	tAdapter = NewAdapter(tAdaptee)
	tAdapter.Request()
}
