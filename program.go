package MVM

import (
	"fmt"
	"github.com/nilspolek/goLog"
	"time"
)

const (
	HALT = iota
	PUSHC
	ADD
	SUB
	MUL
	DIV
	MOD
	RDINT
	WRINT
	RDCHR
	WRCHR
)

var stack Stack

type Program struct {
	Pc           int
	Instructions []Instruction
}

func (p *Program) Execute() {
	var (
		err error
	)
	for ; p.Pc < len(p.Instructions); p.Pc++ {
		err = p.Instructions[p.Pc].exectue()
		if err != nil {
			goLog.Error("Error: %v", err)
		}
	}
}

type Instruction struct {
	Instruction uint8
	Payload     int
}

func (i Instruction) exectue() error {
	switch i.Instruction {
	case HALT:
		return i.halt()
	case PUSHC:
		return i.pushc(i.Payload)
	case ADD:
		return i.add()
	case SUB:
		return i.sub()
	case MUL:
		return i.mul()
	case DIV:
		return i.div()
	case MOD:
		return i.mod()
	case RDINT:
		return i.rdint()
	case WRINT:
		return i.wrint()
	case RDCHR:
		return i.rdchr()
	case WRCHR:
		return i.wrchr()
	default:
		return fmt.Errorf("unknown instruction %v", i.Instruction)
	}
}

func (i Instruction) halt() error {
	time.Sleep(3 * time.Second)
	return nil
}

func (i Instruction) pushc(num int) error {
	stack.push(num)
	return nil
}

func (i Instruction) add() error {
	rhs, err := stack.pop()
	if err != nil {
		return err
	}
	lhs, err := stack.pop()
	if err != nil {
		return err
	}
	stack.push(lhs + rhs)
	return nil
}

func (i Instruction) sub() error {
	rhs, err := stack.pop()
	if err != nil {
		return err
	}
	lhs, err := stack.pop()
	if err != nil {
		return err
	}
	stack.push(lhs - rhs)
	return nil
}

func (i Instruction) mul() error {
	rhs, err := stack.pop()
	if err != nil {
		return err
	}
	lhs, err := stack.pop()
	if err != nil {
		return err
	}
	stack.push(lhs * rhs)
	return nil
}
func (i Instruction) div() error {
	rhs, err := stack.pop()
	if err != nil {
		return err
	}
	lhs, err := stack.pop()
	if err != nil {
		return err
	}
	stack.push(lhs / rhs)
	return nil
}

func (i Instruction) mod() error {
	rhs, err := stack.pop()
	if err != nil {
		return err
	}
	lhs, err := stack.pop()
	if err != nil {
		return err
	}
	stack.push(lhs % rhs)
	return nil
}

func (i Instruction) rdint() error {
	var (
		payload int
		err     error
	)
	_, err = fmt.Scanf("%d", &i.Payload)
	stack.push(payload)
	return err
}

func (i Instruction) wrint() error {
	num, err := stack.pop()
	if err != nil {
		return err
	}
	fmt.Printf("%d", num)
	return nil
}

func (i Instruction) rdchr() error {
	var (
		payload int
		err     error
	)
	_, err = fmt.Scanf("%i", &i.Payload)
	stack.push(payload)
	return err
}

func (i Instruction) wrchr() error {
	num, err := stack.pop()
	if err != nil {
		return err
	}
	fmt.Printf("%c", rune(num))
	return nil
}
