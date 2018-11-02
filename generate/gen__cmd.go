package main

import (
	"fmt"
	"io/ioutil"
	"jpm/gt/util"
	"regexp"
	"strings"
	"time"
)

func printDef(path, funcName, def string) string {
	re := regexp.MustCompile(`(?m)^`)
	//fmt.Println("<<<" + def + ">>>")
	def = re.ReplaceAllString(strings.TrimRight(def, " \t\n\r"), "\t")
	opts := def[strings.Index(def, "CmdOpts")+7:]
	//fmt.Println("1<<" + opts + ">>>")
	args := def[strings.Index(def, "CmdArgs")+7:]
	//fmt.Println("2<<" + opts + ">>>")
	opts = strings.TrimSpace(opts[:strings.Index(opts, "CmdArgs")])
	fmt.Println("opts:<<<" + opts + ">>>")
	fmt.Println("args:<<<" + args + ">>>")
	//defstr:=fmt.Sprintf()
	//json.Marshal(cmddef.CmdDef,defstr)
	//c:=cmddef.CmdDef{CmdName: funcName,CmdArgs: args,CmdOpts: opts}

	x := `
// Cmd%[1]sOpts defined in %[2]s
var %[1]sCmdDef = NewCmdDef(
	"%[1]s",
%[3]s,
)

// Cmd%[1]sOpts is the list of options available for the command 
type Cmd%[1]sOpts struct {
%[4]s
}

// %[1]sCaller helper command to call function
func %[1]sCaller(args []string) int {
	cmdDef := cmddef.Commands["%[1]s"]
	o := Cmd%[1]sOpts{}
	xv := reflect.ValueOf(&o).Elem()     // Dereference into addressable value
	xt := xv.Type()                      // Now get the type of the value object
	idx1 := cmddef.checkOpts(args, cmdDef, xv, xt) // idx1 is the offset of the first argument after the options
	arg1 := cmddef.IntArg(cmd, args[idx1], 0)
	return %[1]s(arg1, o)	
}
`
	out := fmt.Sprintf(x, funcName, path, def, opts)
	return out
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dir := "./cmds"
	files, _ := ioutil.ReadDir(dir)
	inComment := false
	inDef := false
	def := ""
	funcName := ""
	output := ""
	for _, f := range files {
		path := dir + "/" + f.Name()
		fmt.Printf("[generate] Reading %v\n", path)
		for l := range util.Readlines(path) {
			t := strings.TrimSpace(l)
			switch {
			case util.Startswith(t, "//"):
				continue
			case inComment:
				if strings.Contains(l, "*/") {
					inComment = false
				}
				continue
			case util.Startswith(t, "/*") && util.Endswith(t, " CmdDef"):
				fmt.Println(l)
				inDef = true
				x := strings.Index(t, " CmdDef")
				funcName = t[2:x]
			case util.Startswith(t, "/*"):
				if !strings.Contains(l, "*/") {
					inComment = true
				}
				continue
			case inDef:
				if util.Endswith(t, "*/") {
					inDef = false
					out := printDef(path, funcName, def)
					output += out
					fmt.Print(out)
				} else {
					def += l + "\n"
				}
			}
		}
	}

	//fmt.Println("Calling GC")
	// run garbage collection
	//runtime.GC()
	time.Sleep(1 * time.Second)
}

/*
                                 Example Command
// FooCmdDef is ...
var FooCmdDef = NewCmdDef(
	"foo",
	CmdOpts{
		"Name": {shortOpt: "n", desc: "name of city", datatype: "string", defaultVal: "metropolis"},
		"Age":  {shortOpt: "a", desc: "age of city", datatype: "int", defaultVal: "0"},
	},
	CmdArgs{{"bar", "int"}})

// Foo is ...
func Foo(bar int, o CmdFooOpts) int {
	fmt.Printf("bar=%d age=%d name=%s\n", bar, o.Age, o.Name)
	return o.Age
}
*/
