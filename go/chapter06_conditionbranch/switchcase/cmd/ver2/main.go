package main

import (
	"fmt"
	manager "switchcase/cmd/ver2/magicManager"
	magic "switchcase/magicutil"
)

func main() {
	fmt.Println("hello")
	member := magic.Member{Level: 1, Agility: 2, MagicAttack: 3, Vitality: 5}
	var magicType magic.MagicType = magic.Fire
	fmt.Printf("%d", member.Level)
	fmt.Printf("%d", member.Agility)
	fmt.Printf("%d", member.MagicAttack)
	fmt.Printf("%d", member.Vitality)

	magicManager := manager.MagicManager{}

	name := magicManager.GetName(magicType)
	fmt.Printf("%s", name)
}
