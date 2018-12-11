package main

import "flag"
import "fmt"
import "os"

// java [-options] your_main_class_fully_names [args...]
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string   //zhongshu-comment 对应[-options]
	class       string   //zhongshu-comment 对应your_main_class_fully_names
	args        []string //zhongshu-comment 对应[args...]
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage

	/*
		1、flag这个package是顺序解析参数的，先解析flag，后解析args

		2、举例：假如传入这些参数：-version -help q w e "g g" -hehe
			./ch01 -version -help q w e -hehe
			flag会顺序解析-version -help q w e "g g" -hehe这一串参数，
			第1个-version，无需参数值，就会赋值给&cmd.versionFlag，
			第2个-help，无需参数值，就会赋值给&cmd.helpFlag，
			第3个q，q之前没有带横杠，那就意味着从q开始以及后面都是args，不是flag(flag是带横杠的)
	*/
	//flag.BoolVar是不需要参数值，只需要变量名即可，eg：java -version，只有version这个变量名
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")

	//flag.StringVar需要变量名和参数值，例如：java -cp /a/b/c/dd.jar
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	fmt.Println("_____________________")
	fmt.Println(args)
	fmt.Println("_____________________")
	if len(args) > 0 {
		cmd.class = args[0] //zhongshu-comment 第一个非选项参数给出主类的完全限定名
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}
