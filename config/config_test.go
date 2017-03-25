package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

func TestFoo(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("Current test filename: " + filename)
	foo := filepath.Dir(filename)
	fmt.Printf("%v\n", foo)
	raw, _ := ioutil.ReadFile(foo + "/test_config.json")
	fmt.Printf("%v\n", raw)
}
