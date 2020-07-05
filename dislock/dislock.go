package dislock

type DisLock interface {
	TryLock() error //上锁
	Unlock() error  //释放锁
}
