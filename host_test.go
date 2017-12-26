package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleStorage = `
Filesystem     1M-blocks  Used Available Use% Mounted on
/dev/root          29894 18969      9686  67% /
devtmpfs             427     0       427   0% /dev
tmpfs                432     0       432   0% /dev/shm
tmpfs                432    33       399   8% /run
tmpfs                  5     1         5   1% /run/lock
tmpfs                432     0       432   0% /sys/fs/cgroup
/dev/mmcblk0p1        63    21        42  33% /boot
tmpfs                 87     0        87   0% /run/user/1000
`

var exampleMemory = `
             total       used       free     shared    buffers     cached
Mem:           862        836         25          9          3        759
-/+ buffers/cache:         73        788
Swap:           99         35         64
`

func Test_parseStorageInfo(t *testing.T) {
	info := parseStorageInfo("foobar")
	assert.Equal(t, false, info.Parsed)
	assert.Equal(t, 0, info.Total)
	assert.Equal(t, 0, info.Used)
	assert.Equal(t, 0, info.Available)
	assert.Equal(t, 0, info.UsedPercent)

	info = parseStorageInfo(exampleStorage)
	assert.Equal(t, true, info.Parsed)
	assert.Equal(t, 29894, info.Total)
	assert.Equal(t, 18969, info.Used)
	assert.Equal(t, 9686, info.Available)
	assert.Equal(t, 67, info.UsedPercent)
}

func Test_parseMemoryInfo(t *testing.T) {
	info := parseMemoryInfo("foobar")
	assert.Equal(t, false, info.Parsed)
	assert.Equal(t, 0, info.Total)
	assert.Equal(t, 0, info.Used)
	assert.Equal(t, 0, info.Available)
	assert.Equal(t, 0, info.UsedPercent)

	info = parseMemoryInfo(exampleMemory)
	assert.Equal(t, true, info.Parsed)
	assert.Equal(t, 862, info.Total)
	assert.Equal(t, 836, info.Used)
	assert.Equal(t, 25, info.Available)
	assert.Equal(t, 96, info.UsedPercent)
}
