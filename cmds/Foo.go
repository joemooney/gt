package cmds

import (
	"fmt"
)

/*Foo CmdDef
CmdOpts: {
	Name: {ShortOpt: "n", Desc: "name of city", DataType: "string", DefaultVal: "metropolis"},
	Age:  {ShortOpt: "a", Desc: "age of city", DataType: "int", DefaultVal: "0"}
}
CmdArgs: [{bar, int}]
*/
func Foo(bar int, o CmdFooOpts) int {
	fmt.Printf("bar=%d age=%d name=%s\n", bar, o.Age, o.Name)
	return o.Age
}
