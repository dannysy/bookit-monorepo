package object

import (
	"testing"

	"github.com/casdoor/casdoor/object"
)

func TestDumpToFile(t *testing.T) {
	object.InitConfig()

	err := object.DumpToFile("../etc/init_data.json")
	if err != nil {
		panic(err)
	}
}
