package main

import "fmt"

func main() {
	m := make(map[string]string)
	// computer says invalid argument m (type map[string]string) for cap.
	fmt.Println(cap(m))
}
