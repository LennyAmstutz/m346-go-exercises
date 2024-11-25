package main

import "fmt"

func main() {
	modules := make(map[string]int)

	modules["Module 104"] = 104
	modules["Module 117"] = 117
	modules["Module 346"] = 346

	fmt.Println("Modul 104:", 104)
	fmt.Println("Modul 117:", 117)
	fmt.Println("Modul 346:", 346)

	delete(modules, "Module 117")
	fmt.Println("Module nach gelöschtem Module", modules)

	modules["Module 111"] = 111
	fmt.Println("Module nach Hinzugefügtem Module", modules)

	modules["Module 346"] = 222
	fmt.Println("Module nach Wächseln vom Module", modules)

	fmt.Println(modules)
}
