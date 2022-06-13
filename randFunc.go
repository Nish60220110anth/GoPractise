package main

import "math/rand"

func GetName(index uint8) StudentName {
	switch index {
	case 1:
		return NISHANTH
	case 2:
		return ARUN
	case 3:
		return MAHA
	case 4:
		return PREM
	default:
		panic(indexErrStr)
	}
}

func GetRndmClassStandard() ClassStandard {
	rand.Seed(2)
	num := rand.Intn(ClassStandardValidLen) + 1

	switch num {
	case 1:
		return BTECH
	case 2:
		return BSC
	case 3:
		return MBA
	case 4:
		return MBBS
	default:
		panic(indexErrStr)
	}
}

//TODO: complete this function
func FillRndmData(oneInfo *InfoOfHumans) {
	oneInfo.Age = uint8(rand.Uint32())
}

// Fills the given struct with data randomly
func FillData(info *[]InfoOfHumans) {
	for index := range *info {
		FillRndmData(&(*info)[index])
	}
}
