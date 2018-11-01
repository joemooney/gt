package main

import (
	"fmt"
	"jpm/gt/cmds"
)

// 
//go:generate go run generate/gen__cmd.go

// Each source file can have an init funciton
// this is called after all varialb declartions in the package
// have evaluated their initializers which is after all imported
// packages have been initialized
func init() {

}

func main() {
	// Parse command line options
	//flag.Parse()

	// print active options
	//fmt.Println("hello world long:", optDetailedOutput)
	//fmt.Println("hello world user:", optUser)

	fmt.Println("hello my flags")

	//fs := flag.NewFlagSet("flags for foo",flag.PanicOnError)
	//fs := flag.NewFlagSet("ctx foo", flag.ContinueOnError)

	//fs.IntVar(&optAge, "age", defaultAge, "Age")

	//foo := fs.Parse(args)
	//fmt.Println(foo)
	//args := []string{"--zip", "85260", "--age", "99"}
	args := []string{"--age", "99", "77"}

	//opts := o.NewCmdHero(args)
	//fmt.Printf("%v is prime? %v\n",23,IsPrime(23))

	cmds.FooCaller(args)

	cmds.DumpDef()

	//var def1 CommandDef
	//def1.help = "foo"

	/*
		if opts.Age > 49 {
			log.Fatal(fmt.Sprintf("You must be under 50 %s, %d is too old", opts.Name, opts.Age))
		} else if opts.Age > 29 {
			fmt.Println("older than 29")
			os.Exit(1)
		} else {
			fmt.Println("young buck")
			//exit 1
			os.Exit(0)
		}
	*/

}
