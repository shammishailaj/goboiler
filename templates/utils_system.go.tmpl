package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"os"
)

func (u *Utils) LoadAvg() (*load.AvgStat, error) {
	return load.Avg()
}

func (u *Utils) CPUInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

func (u *Utils) CPUCores() int32 {
	var cores int32 = 0
	cpuInfo, cpuInfoErr := u.CPUInfo()
	if cpuInfoErr != nil {
		u.Log.Errorf("Error getting CPU info. %s", cpuInfoErr.Error())
		return CPU_CORES_CPUINFO_ERROR
	}

	for _, info := range cpuInfo {
		cores += info.Cores
	}
	return cores
}

func (u *Utils) LoadAvgCheck() int {
	lavg, lavgErr := u.LoadAvg()

	if lavgErr != nil {
		u.Log.Errorf("Error getting load average. %s", lavgErr.Error())
		return LAVG_TREND_ERROR
	}
	if lavg.Load1 > lavg.Load5 {
		cpuCores := u.CPUCores()
		u.Log.Infof("LoadAvgCheck():: Found %d CPU Cores", cpuCores)
		if cpuCores != CPU_CORES_CPUINFO_ERROR {
			loadValueYellow := (LAVG_LOAD_LEVEL_YELLOW * float64(cpuCores)) / 100
			if lavg.Load1 >= loadValueYellow {
				return LAVG_TREND_INCREASING
			} else {
				return LAVG_TREND_NORMAL
			}
		} else {
			return LAVG_TREND_CPUINFO_ERROR
		}
	} else {
		return LAVG_TREND_NORMAL
	}
}

func (u *Utils) LoadAvgCheckCPUCores(cpuCores int32) int {
	lavg, lavgErr := u.LoadAvg()

	if lavgErr != nil {
		u.Log.Errorf("Error getting load average. %s", lavgErr.Error())
		return LAVG_TREND_ERROR
	}
	if lavg.Load1 > lavg.Load5 {
		if cpuCores != CPU_CORES_CPUINFO_ERROR {
			loadValueYellow := (LAVG_LOAD_LEVEL_YELLOW * float64(cpuCores)) / 100
			if lavg.Load1 >= loadValueYellow {
				return LAVG_TREND_INCREASING
			} else {
				return LAVG_TREND_NORMAL
			}
		} else {
			return LAVG_TREND_CPUINFO_ERROR
		}
	} else {
		return LAVG_TREND_NORMAL
	}
}

func (u *Utils) Hostname() string {
	hostname := os.Getenv("CONTAINER_HOST_NAME")
	var err error
	if hostname == "" {
		hostname, err = os.Hostname()
		if err != nil {
			hostname = err.Error()
		}
	}
	return hostname
}
