package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func PKID() string {
	v7, err := uuid.NewV7()
	if err != nil {
		return strings.ReplaceAll(uuid.NewString(), "-", "")
	} else {
		return hex.EncodeToString(v7[:])
	}
}

const (
	epoch        = int64(1609459200000) // 基准时间:2021‑01‑01
	machineBits  = uint(10)
	sequenceBits = uint(12)

	maxMachineID = int64(1<<machineBits - 1)
	maxSequence  = int64(1<<sequenceBits - 1)

	machineShift   = sequenceBits
	timestampShift = sequenceBits + machineBits
)

type SnowFlake struct {
	mu        sync.Mutex
	machineID int64
	sequence  int64
	lastTime  int64
}

func NewSnowFlake(machineID int64) *SnowFlake {

	return &SnowFlake{
		machineID: machineID,
	}
}

var sf *SnowFlake = &SnowFlake{machineID: 1}

func (s *SnowFlake) NextID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli() - epoch
	if now < s.lastTime {
		return 0, errors.New("clock back")
	}
	if now == s.lastTime {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// 本毫秒号用尽，等到下一毫秒
			for now <= s.lastTime {
				now = time.Now().UnixMilli() - epoch
			}
		}
	} else {
		s.sequence = 0
	}
	s.lastTime = now

	id := (now << timestampShift) | (s.machineID << machineShift) | s.sequence
	return id, nil
}

func NextInt[T int64 | int32 | int16 | int | uint64 | uint32 | uint16 | uint]() (id T) {
	ret, _ := sf.NextID()
	return T(ret)
}

func NextIntString() (id string) {
	ret, _ := sf.NextID()
	return fmt.Sprintf("%d", ret)
}
