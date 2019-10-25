package io

import (
	"testing"
)

func TestContainer(t *testing.T) {
	path := "test/test.json"
	pathDir := "test"
	pathWrite := "test/test_write.json"

	json := testReadAll(path, t)
	testWriteFile(pathWrite, json, t)
	testReadDir(pathDir, t)

}
func testReadAll(path string, t *testing.T) string {
	result, err := ReadAll(path)
	if err != nil || result == "" {
		t.Fatal("read file error")
	}
	return result
}
func testWriteFile(path string, json string, t *testing.T) {
	err := WriteFile(path, json)
	if err != nil {
		t.Log(err)
		t.Fatal("write file error")
	}
}
func testReadDir(path string, t *testing.T) {
	_, err := ReadDir(path)
	if err != nil {
		t.Fatal("read dirctory error")
	}
}
