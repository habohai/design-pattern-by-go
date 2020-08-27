package shape

import "testing"

func TestFlyweightPattern(t *testing.T) {
	flyweightPatternDemo := FlyweightPatternDemo{}
	shapeFactory := new(ShapeFactory)
	for i := 0; i < 100; i++ {
		circle := shapeFactory.GetCircle(flyweightPatternDemo.GetRandomColor())
		circle.SetX(flyweightPatternDemo.GetRandomXAndY())
		circle.SetY(flyweightPatternDemo.GetRandomXAndY())
		circle.SetRadius(100)
		circle.Draw3()
		println(len(shapeFactory.circleMap))
	}
}
