package cmd

import (
	"testing"
)

// func TestMakeMenu(t *testing.T) {
// 	t.Log("Creating menu... (expected result: \n[1] a\n[2] b\n )")
// 	menuItems := []string{"a", "b"}
// 	myMenu := makeMenu(menuItems)
//
// 	if myMenu != "\n[1] a\n[2] b\n" {
// 		t.Errorf("Expected result of '\n[1] a\n[2] b\n' but it was %s instead.", myMenu)
// 	}
// }

func TestFieldMap(t *testing.T) {
	t.Log("Testing field mappings (expected result: 'myMap[EventTwitter]' = 'Twitter')")
	myMap := fieldMap()

	if myMap["EventTwitter"] != "Twitter" {
		t.Errorf("Expected result of 'myMap[EventTwitter]' = 'Twitter' but it was %s instead.", myMap["EventTwitter"])
	}
}

func TestEventStruct(t *testing.T) {
	t.Log("Testing creation of a new Event struct")
	myEvent := eventStruct("chicago", "2017")

	if myEvent.Name != "2017-chicago" {
		t.Errorf("Expected result of '2017-chicago' but it was %s instead", myEvent.Name)
	}
}
