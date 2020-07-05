package dislock

import (
	"context"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"log"
)

var ErrLockAlreadyUse = errors.New("锁被占用")

//分布式锁 ,通过etcd的TXN事务实现
type DisLockEtcd struct {
	key       string //锁名称
	ttl       int64  //锁超时时间
	isLocked  bool   //上锁成功标识
	cancelFun context.CancelFunc

	kv      clientv3.KV
	lease   clientv3.Lease
	leaseId clientv3.LeaseID
}

func NewDisLock(key string, client *clientv3.Client) *DisLockEtcd {

	this := &DisLockEtcd{
		key:   key,
		ttl:   30,
		kv:    clientv3.NewKV(client),
		lease: clientv3.NewLease(client),
	}

	return this
}

//尝试上锁
func (this *DisLockEtcd) TryLock() error {
	var (
		leaseGranResp *clientv3.LeaseGrantResponse
		leaseKeepResp <-chan *clientv3.LeaseKeepAliveResponse
		err           error
		txnResp       *clientv3.TxnResponse
		txn           clientv3.Txn
		cancelCtx     context.Context
		cancelFunc    context.CancelFunc
	)
	defer func() {
		if err != nil {
			cancelFunc() //取消自动续租
			if leaseGranResp != nil {
				//释放租约资源
				if _, err = this.lease.Revoke(context.TODO(), leaseGranResp.ID); err != nil {
					log.Println("释放租约资源,错误:", err.Error())
				}

			}
		}
	}()

	//1.创建租约（5秒）
	if leaseGranResp, err = this.lease.Grant(context.TODO(), this.ttl); err != nil {
		return err
	}

	//context用于取消自动续租
	cancelCtx, cancelFunc = context.WithCancel(context.TODO())

	//2.自动续租
	if leaseKeepResp, err = this.lease.KeepAlive(cancelCtx, leaseGranResp.ID); err != nil {
		return err
	}

	//处理续租应答
	go func() {
		var (
			keepResp *clientv3.LeaseKeepAliveResponse
		)

		for {
			select {
			case keepResp = <-leaseKeepResp: //自动续租应答
				if keepResp == nil {
					goto END
				}

			}
		}

	END:
	}()
	//3.创建事务
	txn = this.kv.Txn(context.TODO())

	//4.事务抢锁
	txn.If(clientv3.Compare(clientv3.CreateRevision(this.key), "=", 0)).
		Then(clientv3.OpPut(this.key, "xxx", clientv3.WithLease(leaseGranResp.ID))).
		Else(clientv3.OpGet(this.key))

	//提交事务
	if txnResp, err = txn.Commit(); err != nil {
		return err
	}
	//5.成功返回，失败释放租约
	if !txnResp.Succeeded {
		//被占用
		return ErrLockAlreadyUse
	}

	//抢锁成功
	this.leaseId = leaseGranResp.ID
	this.cancelFun = cancelFunc
	this.isLocked = true

	return nil
}

//释放锁
func (this *DisLockEtcd) Unlock() (err error) {
	if this.isLocked {
		this.cancelFun() //取消程序自动续租的协程

		if _, err := this.lease.Revoke(context.TODO(), this.leaseId); err != nil {
			//释放租约资源
			return err
		}
	}

	return nil
}
