package exercise

import (
	"fmt"
	"slices"
)

type Item struct {
	Name    string
	Type    string
	Message string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s!\n", p.Name, item.Name)
}
func (p *Player) DropItem(itemName string) {

	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = slices.Delete(p.Inventory, i, i)
			fmt.Printf("%s dropped %s!\n", p.Name, itemName)
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory\n", p.Name, itemName)
}
func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s used %s and %s\n", p.Name, itemName, item.Message)
		} else {
			fmt.Printf("%s used %s\n", p.Name, itemName)
		}
	}
}

func Exercise() {
	// pf := fmt.Printf
	pln := fmt.Println
	items := []Item{
		{Name: "Sword", Type: "Weapon", Message: "to slash enemy"},
		{Name: "Shield", Type: "Weapon", Message: "to protect from attach"},
		{Name: "Darts", Type: "Weapon", Message: "to kill a few enemies"},
		{Name: "Potion", Type: "Drink", Message: "feel rejuvenated and healed"},
	}
	pln("==================Exercise=====================")
	roy := Player{Name: "Roy"}
	roy.PickUpItem(items[0])
	roy.PickUpItem(items[1])
	roy.PickUpItem(items[2])
	roy.UseItem(items[0].Name)

}
