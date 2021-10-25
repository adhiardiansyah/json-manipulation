package main

import (
	"fmt"
	"time"
)

type Inventory struct {
	InventoryID int        `json:"inventory_id"`
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Tags        []string   `json:"tags"`
	PurchasedAt time.Time  `json:"purchased_at"`
	Placement   *Placement `json:"placement"`
}

type Placement struct {
	RoomID int    `json:"room_id"`
	Name   string `json:"name"`
}

var inventory []Inventory

func findInMeetingRoom() {
	for _, item := range inventory {
		if item.Placement.Name == "Meeting Room" {
			fmt.Println(item.Name)
		}
	}
}

func findAllElectronic() {
	for _, item := range inventory {
		if item.Type == "electronic" {
			fmt.Println(item.Name)
		}
	}
}

func findAllFurniture() {
	for _, item := range inventory {
		if item.Type == "furniture" {
			fmt.Println(item.Name)
		}
	}
}

func findPurchasedOn() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02"
	value = "2020-01-16"
	date, _ = time.Parse(layoutFormat, value)
	for _, item := range inventory {
		if item.PurchasedAt.Format("16-Jan-20") == date.Format("16-Jan-20") {
			fmt.Println(item.Name)
		}
	}
}

func findBrownColor() {
	for _, item := range inventory {
		if item.Tags[2] == "brown" {
			fmt.Println(item.Name)
		}
	}
}

func main() {
	inventory = append(inventory, Inventory{InventoryID: 9382, Name: "Brown Chair", Type: "furniture", Tags: []string{"chair", "furniture", "brown"}, PurchasedAt: time.Unix(1579190471, 0).UTC(), Placement: &Placement{RoomID: 3, Name: "Meeting Room"}})
	inventory = append(inventory, Inventory{InventoryID: 9380, Name: "Big Desk", Type: "furniture", Tags: []string{"desk", "furniture", "brown"}, PurchasedAt: time.Unix(1579190642, 0).UTC(), Placement: &Placement{RoomID: 3, Name: "Meeting Room"}})
	inventory = append(inventory, Inventory{InventoryID: 2932, Name: "LG Monitor 50 inch", Type: "electronic", Tags: []string{"monitor", "", ""}, PurchasedAt: time.Unix(1579017842, 0).UTC(), Placement: &Placement{RoomID: 3, Name: "Lavender"}})
	inventory = append(inventory, Inventory{InventoryID: 232, Name: "Sharp Pendingin Ruangan 2PK", Type: "electronic", Tags: []string{"ac", "", ""}, PurchasedAt: time.Unix(1578931442, 0).UTC(), Placement: &Placement{RoomID: 5, Name: "Dhanapala"}})
	inventory = append(inventory, Inventory{InventoryID: 9382, Name: "Alat Makan", Type: "tableware", Tags: []string{"spoon", "fork", "tableware"}, PurchasedAt: time.Unix(1578672242, 0).UTC(), Placement: &Placement{RoomID: 10, Name: "Rajawali"}})

	fmt.Println("1. Find items in the Meeting Room.")
	findInMeetingRoom()
	fmt.Println()
	fmt.Println("2. Find all electronic devices.")
	findAllElectronic()
	fmt.Println()
	fmt.Println("3. Find all the furniture.")
	findAllFurniture()
	fmt.Println()
	fmt.Println("4. Find all items were purchased on 16 Januari 2020.")
	findPurchasedOn()
	fmt.Println()
	fmt.Println("5. Find all items with brown color.")
	findBrownColor()
}
