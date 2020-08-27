package shape

import (
	"math/rand"
	"time"
)

type FlyweightPatternDemo struct {
	colors []string
}

func (f *FlyweightPatternDemo) GetRandomColor() string {
	rand.Seed(time.Now().UnixNano())
	if f.colors == nil {
		f.colors = []string{"Red", "Green", "Blue", "White", "Black"}
	}

	num := rand.Intn(len(f.colors))
	return f.colors[num]
}

func (f *FlyweightPatternDemo) GetRandomXAndY() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10) * 100
	return num
}
