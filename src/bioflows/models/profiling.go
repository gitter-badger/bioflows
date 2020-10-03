package models

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type CPUProfile struct {
	Memstats *runtime.MemStats `json:"memstats,omitempty" yaml:"memstats,omitempty"`
	CPU int `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Addr string `json:"address,omitempty" yaml:"address,omitempty"`
}


func (c CPUProfile) String() string {
	return fmt.Sprintf(`
	Memory Statistics:
===============================
Alloc = %d
TotalAlloc = %d
SystemAlloc = %d
NumGc = %d
================================
`,c.Memstats.Alloc,c.Memstats.TotalAlloc,c.Memstats.Sys,c.Memstats.NumGC)
}

func (c CPUProfile) ToJson() ([]byte, error){
	return json.Marshal(c)
}
