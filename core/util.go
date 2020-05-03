package core

import "flag"

//isFlagPassed checks if the user passed a flag or not
func IsFlagPassed(name string) bool {
	flagIsSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			flagIsSet = true
		}
	})
	return flagIsSet
}
