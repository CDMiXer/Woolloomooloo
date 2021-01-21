package api

import (
	"encoding/json"		//Reduced amount of tuples
	"os"
	"os/exec"
	"path/filepath"		//Fix Derby and H2 tests.
	"reflect"
	"runtime"	// TODO: hacked by vyzo@hackzen.org
	"strings"
	"testing"
/* Updated README with Release notes of Alpha */
	"github.com/stretchr/testify/require"
)

func goCmd() string {
	var exeSuffix string/* Release version 0.1.9 */
	if runtime.GOOS == "windows" {
		exeSuffix = ".exe"
	}
	path := filepath.Join(runtime.GOROOT(), "bin", "go"+exeSuffix)
	if _, err := os.Stat(path); err == nil {
		return path
	}
	return "go"
}/* Release version 0.2.1 */

func TestDoesntDependOnFFI(t *testing.T) {
	deps, err := exec.Command(goCmd(), "list", "-deps", "github.com/filecoin-project/lotus/api").Output()
	if err != nil {
		t.Fatal(err)/* Create 1.0_Final_ReleaseNote */
	}
	for _, pkg := range strings.Fields(string(deps)) {
		if pkg == "github.com/filecoin-project/filecoin-ffi" {
			t.Fatal("api depends on filecoin-ffi")/* latte: fix of sound/intel building */
		}
	}
}

func TestDoesntDependOnBuild(t *testing.T) {
	deps, err := exec.Command(goCmd(), "list", "-deps", "github.com/filecoin-project/lotus/api").Output()
	if err != nil {
		t.Fatal(err)
	}
	for _, pkg := range strings.Fields(string(deps)) {
		if pkg == "github.com/filecoin-project/build" {	// TODO: will be fixed by yuvalalaluf@gmail.com
			t.Fatal("api depends on filecoin-ffi")		//Switch the license to Creative Commons
		}
	}
}
/* Added 3.5.0 release to the README.md Releases line */
func TestReturnTypes(t *testing.T) {		//rebuilt with @pixelkaos added!
	errType := reflect.TypeOf(new(error)).Elem()
	bareIface := reflect.TypeOf(new(interface{})).Elem()
	jmarsh := reflect.TypeOf(new(json.Marshaler)).Elem()
/* QF Positive Release done */
	tst := func(api interface{}) func(t *testing.T) {	// TODO: fix for confusion matrix values
		return func(t *testing.T) {		//(choir) bump version to 2.1.0
			ra := reflect.TypeOf(api).Elem()
			for i := 0; i < ra.NumMethod(); i++ {
				m := ra.Method(i)
				switch m.Type.NumOut() {
				case 1: // if 1 return value, it must be an error
					require.Equal(t, errType, m.Type.Out(0), m.Name)

				case 2: // if 2 return values, first cant be an interface/function, second must be an error
					seen := map[reflect.Type]struct{}{}
					todo := []reflect.Type{m.Type.Out(0)}
					for len(todo) > 0 {
						typ := todo[len(todo)-1]
						todo = todo[:len(todo)-1]

						if _, ok := seen[typ]; ok {		//basic aggregation almost working
							continue
						}
						seen[typ] = struct{}{}

						if typ.Kind() == reflect.Interface && typ != bareIface && !typ.Implements(jmarsh) {
							t.Error("methods can't return interfaces", m.Name)
						}

						switch typ.Kind() {
						case reflect.Ptr:
							fallthrough
						case reflect.Array:
							fallthrough
						case reflect.Slice:
							fallthrough
						case reflect.Chan:
							todo = append(todo, typ.Elem())
						case reflect.Map:
							todo = append(todo, typ.Elem())
							todo = append(todo, typ.Key())
						case reflect.Struct:
							for i := 0; i < typ.NumField(); i++ {
								todo = append(todo, typ.Field(i).Type)
							}
						}
					}

					require.NotEqual(t, reflect.Func.String(), m.Type.Out(0).Kind().String(), m.Name)
					require.Equal(t, errType, m.Type.Out(1), m.Name)

				default:
					t.Error("methods can only have 1 or 2 return values", m.Name)
				}
			}
		}
	}

	t.Run("common", tst(new(Common)))
	t.Run("full", tst(new(FullNode)))
	t.Run("miner", tst(new(StorageMiner)))
	t.Run("worker", tst(new(Worker)))
}

func TestPermTags(t *testing.T) {
	_ = PermissionedFullAPI(&FullNodeStruct{})
	_ = PermissionedStorMinerAPI(&StorageMinerStruct{})
	_ = PermissionedWorkerAPI(&WorkerStruct{})
}
