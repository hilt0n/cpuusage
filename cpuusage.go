package cpuusage

import (
    "errors"
    "fmt"
    "os"
    proc "github.com/c9s/goprocinfo/linux"
    "sync"
    "time"
)

type CPUUsage struct {
    CPU		float64
    lock	sync.Mutex
    // TODO add a channel or similar to link the worker 
    // goroutine with a given object
}

func (cpu *CPUUsage) Start() error {
    // test if file exists
    if _, err := os.Stat("/proc/stat"); os.IsNotExist(err) {
        return errors.New("/proc/stat file doesn't exist") 
    }
    
    go func () {
        stat, err := proc.ReadStat("/proc/stat")
        if err != nil {
            return
        }
        for {
            time.Sleep(time.Millisecond * 1000)
            
            statn, err := proc.ReadStat("/proc/stat")
            if err != nil {
                return
            }
            
            deltaU := (float64)(statn.CPUStatAll.User - stat.CPUStatAll.User)
            deltaN := (float64)(statn.CPUStatAll.Nice - stat.CPUStatAll.Nice)
            deltaS := (float64)(statn.CPUStatAll.System - stat.CPUStatAll.System)
            deltaI := (float64)(statn.CPUStatAll.Idle - stat.CPUStatAll.Idle)
         
            val := 100 * (deltaU + deltaN + deltaS) / (deltaU + deltaN + deltaS + deltaI)
            
            cpu.lock.Lock()
            cpu.CPU = val
            cpu.lock.Unlock()
            
            stat = statn   
        }
    }()
    
    return nil
}

// TODO Add a stop function to stop the worker

// TODO Add a notification system based on callback or channel

func (cpu *CPUUsage) Print() error {
    // test if file exists
    cpu.lock.Lock()
    val := cpu.CPU
    cpu.lock.Unlock()
    
    fmt.Printf("CPU val: %.1f%%\n", val)
        
    return nil
}