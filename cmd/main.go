package main

import (
	"flag"
	"fmt"
	"github.com/nilspolek/MVM"
	"github.com/nilspolek/goLog"
)

const (
	versionInfo    = "Ninja Virtual Machine version 0"
	startMsg       = "My Virtual Machine started"
	stopMsg        = "My Virtual Machine stopped"
	debugEnableMsg = "My Virtual Machine debug enabled"
)

var (
	isDebug   *bool
	isVersion *bool
)

func main() {
	isDebug = flag.Bool("d", false, "debug mode")
	isVersion = flag.Bool("v", false, "version mode")
	flag.Parse()
	if isVersion != nil && *isVersion {
		fmt.Println(versionInfo)
		return
	}
	if isDebug != nil && !*isDebug {
		goLog.LoggingLevel = goLog.HIGH
	}
	goLog.Info(debugEnableMsg)
	goLog.Info(startMsg)
	program := MVM.Program{
		Pc: 0,
		Instructions: []MVM.Instruction{
			{
				Instruction: MVM.PUSHC,
				Payload:     3,
			},
			{
				Instruction: MVM.PUSHC,
				Payload:     4,
			},
			{
				Instruction: MVM.ADD,
			},
			{
				Instruction: MVM.PUSHC,
				Payload:     10,
			},
			{
				Instruction: MVM.PUSHC,
				Payload:     6,
			},
			{
				Instruction: MVM.SUB,
			},
			{
				Instruction: MVM.MUL,
			},
			{
				Instruction: MVM.WRINT,
			},
			{
				Instruction: MVM.PUSHC,
				Payload:     int('\n'),
			},
			{
				Instruction: MVM.WRCHR,
			},
			{
				Instruction: MVM.HALT,
			},
		},
	}
	program.Execute()
	goLog.Info(stopMsg)
}
