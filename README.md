# CPU Usage

This is a simple Golang package (in development) giving the actual CPU usage based on a 1 sec average.

### Installation
```go
go get github.com/hilt0n/cpuusage
```

### Basic usage
```go
package main

import (
    "github.com/hilt0n/cpuusage"
    "time"
)

func main() {
    cpu := cpuusage.CPUUsage{}
    
    cpu.Start()
    
    for {
        time.Sleep(1 * time.Second)
        cpu.Print()    
    } 
    
	return
}
```
