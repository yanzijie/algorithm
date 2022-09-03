package main

import "github.com/yanzijie/algorithm/link"

func main() {
	TestOneWayLinkList()
}

func TestOneWayLinkList() {
	list := link.InitOneWayLinkList()
	list.InsertOnHead(1)
	list.InsertOnHead(2)
	list.InsertOnHead(3)
	list.InsertOnHead(4)
	list.PrintLink()
}
