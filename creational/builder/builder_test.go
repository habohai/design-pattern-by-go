package builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_GetResult(t *testing.T) {
	var builder Builder = &CharacterBuilder{}
	var director *Director = &Director{builder: builder}
	var character *Character = director.Create("loader", "AK47")
	fmt.Println(character.GetName() + "," + character.GetArms())
}
