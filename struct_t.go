package main

import "fmt"

type Count int

func (count *Count) Increment() { *count++ }
func (count *Count) Decrement() { *count-- }
func (count Count) IsZero() bool { return count == 0 }

type Part struct {
	stat string
	Count
}

func (part Part) IsZero() bool {
	return part.Count.IsZero() && part.stat == ""
}
func (part Part) String() string {
	return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

func main() {
	var i Count = -1
	fmt.Printf("Start \"Count\" test :\nOrigin value of count: %d\n", i)
	i.Increment()
	fmt.Printf("Value of count after increment: %d\n", i)
	fmt.Printf("Count is zero t/f? ;%t\n\n", i.IsZero())
	fmt.Println("Start: \"Part\" test:")
	part := Part{"232",0}
	fmt.Print("Part: %v", part)
	fmt.Printf("Part is zero t/f?: %t\n",part.Count.IsZero())
}
