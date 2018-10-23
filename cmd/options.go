package cmddef

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Goals:
// access options via cmdOpts.Age where --age=9 is an option
// Inherit options, common options can be separated and included as needed
// Detect unsupplied options
// Hide options from user if only meant to be provided programaticallay
// Have default and if not then detect unsupplied values

// CmdOpt An option taking a string as a value --foo="abc"
type CmdOpt struct {
	ShortOpt   string
	Desc       string
	Help       string // long format help
	DataType   string
	DefaultVal string
}

// CmdOpts ...
type CmdOpts map[string]CmdOpt

// CmdArgs ...
type CmdArgs []CmdArgDef

// CmdArgDef ...
type CmdArgDef struct {
	Var      string
	DataType string
}

// CmdDef could be generated from FooCmdDef
type CmdDef struct {
	CmdName string
	CmdOpts CmdOpts
	CmdArgs CmdArgs
}

// Commands list of registered commands
var Commands = make(map[string]CmdDef)

// AddCmdDef add command to list of registered commands
func AddCmdDef(cmd CmdDef) CmdDef {
	Commands[cmd.CmdName] = cmd
	fmt.Printf("Defined %v\n", cmd)
	return cmd
}

// NewCmdDef add command to list of registered commands
func NewCmdDef(CmdName string, CmdOpts CmdOpts, CmdArgs CmdArgs) CmdDef {
	cmd := CmdDef{CmdName, CmdOpts, CmdArgs}
	Commands[CmdName] = cmd
	fmt.Printf("Defined %v\n", cmd)
	return cmd
}

//IntArg converts string to int arg and stores value
func IntArg(cmdDef CmdDef, val string, idx int) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("Invalid integer value " + val + " for arg " + cmdDef.CmdArgs[idx].Var + " in command " + cmdDef.CmdName)
	}
	return i
}

//BoolArg converts string to bool arg and stores value
func BoolArg(cmdDef CmdDef, val string, idx int) bool {
	b, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatal("Invalid boolean value " + val + " for arg " + cmdDef.CmdArgs[idx].Var + " in command " + cmdDef.CmdName)
	}
	return b
}

//CheckOpts converts an args array into options
func CheckOpts(args []string, cmdDef CmdDef, xv reflect.Value, xt reflect.Type) int {

	//fs := flag.NewFlagSet("flags for foo", flag.PanicOnError)
	fs := flag.NewFlagSet("flags for foo", flag.ContinueOnError)

	//xv := reflect.ValueOf(&outOpts).Elem() // Dereference into addressable value
	//xt := xv.Type()                        // Now get the type of the value object

	fmt.Println("here")
	for i := 0; i < xt.NumField(); i++ {
		f := xt.Field(i) // get the field

		//fmt.Printf("%d: %s %s = %v\n", i, xt.Field(i).Name, f.Type(), f.Interface())

		//if f.Name == "cmdDef" { // skip the definition struct
		//	continue
		//}

		name := strings.ToLower(f.Name)        // get the name
		Name := f.Name                         // get the exported name
		addr := xv.Field(i).Addr().Interface() //
		//desc := optdef.Desc

		optdef, ok := cmdDef.CmdOpts[Name]
		if !ok {
			log.Fatal("Command " + cmdDef.CmdName + " missing definition for option: " + Name)
			os.Exit(1)
		}

		switch ptr := addr.(type) {
		case *int:
			i, err := strconv.Atoi(optdef.DefaultVal)
			if err != nil {
				log.Fatal("Invalid integer default value for option")
			}
			fs.IntVar(ptr, name, i, optdef.Help)
		case *string:
			fs.StringVar(ptr, name, optdef.DefaultVal, optdef.Help)
		case *bool:
			b, err := strconv.ParseBool(optdef.DefaultVal)
			if err != nil {
				log.Fatal("Invalid bool default value for option")
			}
			fs.BoolVar(ptr, name, b, optdef.Help)
		}
	}
	err := fs.Parse(args)
	if err != nil {
		log.Fatal(err.Error())
	}
	providedArgs := fs.NArg() // number of args remaining
	if providedArgs != len(cmdDef.CmdArgs) {
		log.Fatal("Error command " + cmdDef.CmdName + " expects " + strconv.Itoa(len(cmdDef.CmdArgs)) + " arguments, you provided " + strconv.Itoa(providedArgs))
	}
	return len(args) - fs.NArg()
}

/*

                            Example Command

// FooCmdDef is ...
var FooCmdDef = NewCmdDef(
	"foo",
	CmdOpts{
		"Name": {ShortOpt: "n", Desc: "name of city", DataType: "string", DefaultVal: "metropolis"},
		"Age":  {ShortOpt: "a", Desc: "age of city", DataType: "int", DefaultVal: "0"},
	},
	CmdArgs{{"bar", "int"}})

// Foo is ...
func Foo(bar int, o CmdFooOpts) int {
	fmt.Printf("bar=%d age=%d name=%s\n", bar, o.Age, o.Name)
	return o.Age
}

// CmdFooOpts could be generated from FooCmdDef
type CmdFooOpts struct {
	Name string
	Age  int
}

// FooCaller is ...
func FooCaller(args []string) int {
	cmd := Commands["foo"]
	o := CmdFooOpts{}
	xv := reflect.ValueOf(&o).Elem()     // Dereference into addressable value
	xt := xv.Type()                      // Now get the type of the value object
	idx1 := checkOpts(args, cmd, xv, xt) // idx1 is the offset of the first argument after the options
	arg1 := intArg(cmd, args[idx1], 1)
	return Foo(arg1, o)
}
*/
