package businesslock

import (
	"errors"
	"sync"
	"sync/atomic"
	"unsafe"
)

//基础锁
const (
	LockedFlag   int32 = 1
	UnlockedFlag int32 = 0
)

type Mutex struct {
	mutex sync.Mutex
}

func (m *Mutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.mutex)), UnlockedFlag, LockedFlag) {
		return true
	}
	return false
}

func (m *Mutex) IsLocked() bool {
	if atomic.LoadInt32((*int32)(unsafe.Pointer(&m.mutex))) == LockedFlag {
		return true
	}
	return false
}

//业务流程锁
type BusinessLock struct {
	lockMap map[string]map[int]*Mutex //mid=>lockType=>Mutex
}

func NewBusinessLock() *BusinessLock {
	bl := &BusinessLock{}
	bl.lockMap = make(map[string]map[int]*Mutex)
	return bl

}

//加锁
func (this *BusinessLock) Lock(mid string, lockType int) error {

	if _, ok := this.lockMap[mid]; !ok {
		this.lockMap[mid] = make(map[int]*Mutex)
	}

	if _, ok := this.lockMap[mid][lockType]; !ok {
		this.lockMap[mid][lockType] = new(Mutex)
	}

	if this.lockMap[mid][lockType].TryLock() {
		return nil
	}

	return errors.New("当前数据有其他用户操作，请稍后再试！")
}

//解锁
func (this *BusinessLock) UnLock(mid string, lockType int) {
	if _, ok := this.lockMap[mid]; ok {
		if mu, ok := this.lockMap[mid][lockType]; ok {
			if mu.IsLocked() {
				mu.mutex.Unlock()
			}
		}
	}
}
