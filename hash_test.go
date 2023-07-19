package main

import (
	"fmt"
	"testing"
)

func TestHashFile(t *testing.T) {
	hash, err := hashFile("testData/hashTestFile")
	if err != nil {
		fmt.Println("encountered error" + err.Error())
		t.Fail()
	}
	expected := "7dca1c31120b34ea5696bec8cc40374c7b8cbe7d2bd00619ce630bf1266a3945"
	if hash != expected {
		fmt.Printf("Expected hash %s, got %s", expected, hash)
		t.Fail()
	}
}

func TestGetHash(t *testing.T) {
	result := getHashes([]string{"testData/hashTestFile"})
	if result[0][0] != "testData/hashTestFile" ||
		result[0][1] != "7dca1c31120b34ea5696bec8cc40374c7b8cbe7d2bd00619ce630bf1266a3945" ||
		len(result) != 1 {
		t.Fail()
	}
}

func TestFindPathHash(t *testing.T) {
	hashes := getHashes([]string{"testData/hashTestFile"})
	path, hash, found := findPathHash(hashes, "testData/hashTestFile", true)
	if path != "testData/hashTestFile" || hash != "7dca1c31120b34ea5696bec8cc40374c7b8cbe7d2bd00619ce630bf1266a3945" || !found {
		fmt.Println("a")
		t.Fail()
	}

	_, _, found = findPathHash(hashes, "testData/hashTestFi", true)
	if found {
		fmt.Println("b")
		t.Fail()
	}
	_, _, found = findPathHash(hashes, "7dca1c31120b34ea5696bec8cc40374c7b8cbe7d2bd00619ce630bf1266a3945", false)
	if !found {

		fmt.Println("c")
		t.Fail()
	}
	_, _, found = findPathHash(hashes, "7dca1c31120b34ea5696bec8cc40374c3945", false)
	if found {

		fmt.Println("d")
		t.Fail()
	}

}
