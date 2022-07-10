// Netflag demonstrates an integer type used as a bit field.
// Netflag 演示了用作位域的整数类型。

package main

import (
	"fmt"
	. "net"
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%d %t\n", v, IsUp(v)) // "17 true"
	TurnDown(&v)
	fmt.Printf("%d %t\n", v, IsUp(v)) // "16 false"
	SetBroadcast(&v)
	fmt.Printf("%d %t\n", v, IsUp(v))   // "18 false"
	fmt.Printf("%d %t\n", v, IsCast(v)) // "18 true"

}
