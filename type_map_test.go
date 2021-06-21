package reflect2_tests

import (
	"github.com/modern-go/reflect2"
	"testing"
)

type MyStruct struct {
}

var typByPtr = reflect2.TypeOfPtr((*MyStruct)(nil)).Elem()

func TestTypeByName(t *testing.T) {
	typByName := reflect2.TypeByName("reflect2_tests.MyStruct")
	if typByName != typByPtr {
		t.Fail()
	}
}

func TestTypeByPackageName(t *testing.T) {
	typByPkg := reflect2.TypeByPackageName(
		"github.com/modern-go/reflect2-tests",
		"MyStruct")
	if typByPkg != typByPtr {
		t.Fail()
	}
}

func TestTypesInPackage(t *testing.T) {
	typInPkg := reflect2.TypesInPackage("github.com/modern-go/reflect2-tests")
	for key, typ := range typInPkg {
		if key != "MyStruct" || typ != typByPtr {
			t.Fail()
		}
	}
}
