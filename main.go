package main

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/datadog-agent/pkg/gohai/cpu"
	"github.com/DataDog/datadog-agent/pkg/gohai/filesystem"
	"github.com/DataDog/datadog-agent/pkg/gohai/memory"
	"github.com/DataDog/datadog-agent/pkg/gohai/network"
	"github.com/DataDog/datadog-agent/pkg/gohai/platform"
)

type CPUInfo struct {
	CPUCores             string `json:"cpu_cores"`
	CPULogicalProcessors string `json:"cpu_logical_processors"`
	Family               string `json:"family"`
	Mhz                  string `json:"mhz"`
	Model                string `json:"model"`
	ModelName            string `json:"model_name"`
	Stepping             string `json:"stepping"`
	VendorID             string `json:"vendor_id"`
}

func main() {
	cput := cpu.Cpu{}
	cpuInfo, err := cput.Collect()
	if err != nil {
		fmt.Println("Error in getting CPU info")
	}
	bytes, err := json.Marshal(cpuInfo)
	localCPU := CPUInfo{}
	json.Unmarshal(bytes, &localCPU)
	fmt.Println(localCPU.CPUCores)
	// fmt.Println(cput.Collect())
	fs := filesystem.FileSystem{}
	fmt.Println("File System")
	fmt.Println(fs.Collect())
	mem := memory.Memory{}
	fmt.Println("Memory")
	fmt.Println(mem.Collect())
	nwk := network.Network{}
	fmt.Println("Network")
	fmt.Println(nwk.Collect())
	plt := platform.Platform{}
	fmt.Println("Platform")
	fmt.Println(plt.Collect())
	// ps := processes.Processes{}
	// fmt.Println("Processes")
	// fmt.Println(ps.Collect())

}
