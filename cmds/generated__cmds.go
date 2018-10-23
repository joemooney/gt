package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
	cmddef "jpm/gt/cmd"
	"log"
	"reflect"
	"strings"
)

// CmdFooOpts could be generated from FooCmdDef
type CmdFooOpts struct {
	Name string
	Age  int
}

// FooCaller is ...
func FooCaller(args []string) int {
	cmdDef := cmddef.Commands["Foo"]
	o := CmdFooOpts{}
	xv := reflect.ValueOf(&o).Elem()               // Dereference into addressable value
	xt := xv.Type()                                // Now get the type of the value object
	idx1 := cmddef.CheckOpts(args, cmdDef, xv, xt) // idx1 is the offset of the first argument after the options
	arg1 := cmddef.IntArg(cmdDef, args[idx1], 1)
	return Foo(arg1, o)
}

// PrettyJson does stuff
func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

/*
//FooCmdDef is...
var FooCmdDef = cmddef.NewCmdDef(
	"Foo",
	cmddef.CmdOpts{
		"Name": {ShortOpt: "n", Desc: "name of city", DataType: "string", DefaultVal: "metropolis"},
		"Age":  {ShortOpt: "a", Desc: "age of city", DataType: "int", DefaultVal: "0"},
	},
	cmddef.CmdArgs{{Var: "bar", DataType: "int"}})
*/

// DumpDef prints out the JSON definitions for commands
func DumpDef() {
	out, err := PrettyJson(FooCmdDef)
	if err != nil {
		log.Fatal("Could not marshall cmd")
	}
	fmt.Println("<<<" + out + ">>>")
	out = strings.Replace(out, "\t", "  ", -1)
	//out = strings.Replace(out, " ", ".", -1)
	out = strings.Replace(out, "\r", "", -1)

	fmt.Println("<<<" + out + ">>>")
	out = strings.Replace(out, "\"CmdName\": ", "CmdName: ", 1)
	out = strings.Replace(out, "\"CmdOpts\": {", "CmdOpts: cmddef.CmdOpts{", 1)
	out = strings.Replace(out, "\"CmdArgs\": [\n    {", "CmdArgs: cmddef.CmdArgs{{", 1)
	out = strings.Replace(out, "\": {\n", "\": cmddef.CmdOpt{\n", -1)
	out = strings.Replace(out, "\"\n    },\n    \"", "\"},\n    \"", -1)
	out = strings.Replace(out, "\"\n    }\n  },\n", "\"}},\n", 1)
	out = strings.Replace(out, "\n    }\n  ]\n}\n", "}}})\n", 1)

	fmt.Println("<<<" + out + ">>>")
	//out = strings.TrimSpace(out)
	//re := regexp.MustCompile(`\s*\]\n`)
	//out = re.ReplaceAllString(out, "}")
	//re = regexp.MustCompile(`\s*\}\n`)
	//out = re.ReplaceAllString(out, "}")
	//re = regexp.MustCompile(`\s*\[\n`)
	//out = re.ReplaceAllString(out, "{")
	//re = regexp.MustCompile(`\s*\{\n`)
	//out = re.ReplaceAllString(out, "{")
	fmt.Println("var FooCmdDef = cmddef.AddCmdDef(cmddef.CmdDef" + out)
	//json.MarshalIndent(FooCmdDef, "", "  ")
}

var FooCmdDef = cmddef.AddCmdDef(cmddef.CmdDef{
	CmdName: "Foo",
	CmdOpts: cmddef.CmdOpts{
		"Age": cmddef.CmdOpt{
			ShortOpt:   "a",
			Desc:       "age of city",
			Help:       "",
			DataType:   "int",
			DefaultVal: "0"},
		"Name": cmddef.CmdOpt{
			ShortOpt:   "n",
			Desc:       "name of city",
			Help:       "",
			DataType:   "string",
			DefaultVal: "metropolis"}},
	CmdArgs: cmddef.CmdArgs{{
		Var:      "bar",
		DataType: "int"}}})

//YooCmdDef asldkjf
var YooCmdDef = cmddef.AddCmdDef(cmddef.CmdDef{
	CmdName: "Foo",
	CmdOpts: cmddef.CmdOpts{
		"Age": cmddef.CmdOpt{
			ShortOpt:   "a",
			Desc:       "age of city",
			Help:       "",
			DataType:   "int",
			DefaultVal: "0"},
		"Name": cmddef.CmdOpt{
			ShortOpt:   "n",
			Desc:       "name of city",
			Help:       "",
			DataType:   "string",
			DefaultVal: "metropolis"}},
	CmdArgs: cmddef.CmdArgs{{
		Var:      "bar",
		DataType: "int"}}})
