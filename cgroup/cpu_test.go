package cgroup

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const cpuPath = "testdata/sys/fs/cgroup/cpu/docker/b29faf21b7eff959f64b4192c34d5d67a707fe8561e9eaa608cb27693fba4242"

func TestCpuStats(t *testing.T) {
	cpu := CPUSubsystem{}
	if err := cpuStat(cpuPath, &cpu); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, uint64(769021), cpu.Stats.Periods)
	assert.Equal(t, uint64(1046), cpu.Stats.ThrottledPeriods)
	assert.Equal(t, uint64(352597023453), cpu.Stats.ThrottledTime)
}

func TestCpuCFS(t *testing.T) {
	cpu := CPUSubsystem{}
	if err := cpuCFS(cpuPath, &cpu); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, uint64(100000), cpu.CFS.PeriodMicrosecs)
	assert.Equal(t, uint64(0), cpu.CFS.QuotaMicrosecs) // -1 is changed to 0.
	assert.Equal(t, uint64(1024), cpu.CFS.Shares)
}

func TestCpuRT(t *testing.T) {
	cpu := CPUSubsystem{}
	if err := cpuRT(cpuPath, &cpu); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, uint64(1000000), cpu.RT.PeriodMicrosecs)
	assert.Equal(t, uint64(0), cpu.RT.RuntimeMicrosecs)
}

func TestCpuSubsystemGet(t *testing.T) {
	cpu := CPUSubsystem{}
	if err := cpu.Get(cpuPath); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, cpuPath, cpu.Path)
	assert.Equal(t, uint64(769021), cpu.Stats.Periods)
	assert.Equal(t, uint64(100000), cpu.CFS.PeriodMicrosecs)
	assert.Equal(t, uint64(1000000), cpu.RT.PeriodMicrosecs)
}

func TestCpuSubsystemJSON(t *testing.T) {
	cpu := CPUSubsystem{}
	if err := cpu.Get(cpuPath); err != nil {
		t.Fatal(err)
	}

	json, err := json.MarshalIndent(cpu, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(json))
}
