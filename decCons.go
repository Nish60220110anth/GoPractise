package main

import "os"

var defaultErrorLog = os.Stdout

const indexErrStr = "index out of range error"

//Only non recoverable error's are printed in the errLog

type StudentName string

const StudentNameValidLen = 4

const (
	NISHANTH StudentName = "Nishanth"
	ARUN     StudentName = "Arun Prashath"
	MAHA     StudentName = "Maha"
	PREM     StudentName = "Prem Suriya"
)

type Parent [2]string

const ParentValidLen = 2

// Array is constant by defult , so we can't use const again , so we need to use var
var (
	SHIVA_PARVATHI Parent = [...]string{"Shivan", "Parvathi"}
	VISHNU_LAKSHMI Parent = [...]string{"Vishnu", "Lakshmi"}
)

type Color string

const ColorValidLen = 4

const (
	RED    Color = "Red"
	BLUE   Color = "Blue"
	GREEN  Color = "Green"
	YELLOW Color = "Yellow"
)

type ClassStandard int

const ClassStandardValidLen = 4

const (
	_ = ClassStandard(iota)
	BTECH
	BSC
	MBA
	MBBS
)

type InfoOfHumans struct {
	Name     StudentName
	Age      uint8
	Class    ClassStandard
	friends  [2]StudentName
	Parents  Parent
	LikeGo   bool
	FavColor Color
}
