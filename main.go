package main

import (
	_ "scripts/config"
	_ "scripts/db"
	_ "scripts/library"
	_ "scripts/tasks"
)

func main() {
	//阻止main函数退出
	defer func() { <-make(chan bool) }()
}
