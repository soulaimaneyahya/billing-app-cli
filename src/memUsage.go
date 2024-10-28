/**
* Created by @soulaimaneyh on 2024-10-28
 */

package main

import (
	"fmt"
	"runtime"
)

func printMemUsage(stage string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("%s:\nAlloc = %v KB\tTotalAlloc = %v KB\tSys = %v KB\tNumGC = %v\n",
		stage, m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
