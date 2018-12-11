package main

import (
	"fmt"
)

/*
step 1
cd /Users/zhongshu/kowloon/program_project/mine/go/src/jvmgo-book/v1/code/go/src/jvmgo/ch01
go install

step 2
cd /Users/zhongshu/kowloon/program_project/mine/go/bin
ch01 -cp /a/b/c/dd.jar  com.test.Hehe gg

ch01是go二进制可执行程序
*/

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s \n"+
		"class:%s \n"+
		"args:%v\n",
		cmd.cpOption,
		cmd.class,
		cmd.args)
}
