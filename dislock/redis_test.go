package dislock

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/phpdi/ant/redisutil"
	"log"
	"testing"
	"time"
)

var dislockRedis *DisLockRedis
var c *redis.Pool

func init() {

	r := redisutil.Node{Address: "127.0.0.1:6379", Password: "123456"}
	c = r.NewRedisPool()
	dislockRedis = NewDisLockRedis("a", c)

}

func TestDisLockRedis_TryLock(t *testing.T) {
	//err:=dislockRedis.TryLock()
	//if err != nil {
	//	t.Error(err)
	//}
	for i := 0; i < 3; i++ {
		//c:=redisutil.Node{Address:"127.0.0.1:6379",Password:"123456"}
		go worker(i)
	}

	<-make(chan int)

	//time.Sleep(30*time.Second)
}

func worker(workerid int) {

	for {
		d := NewDisLockRedis("a", c)
		time.Sleep(2 * time.Second)

		if err := d.TryLock(); err != nil {
			fmt.Println("worker:", workerid, err)
			continue
		}
		log.Println("worker:", workerid, "上锁成功")
		time.Sleep(10 * time.Second)

		if err := d.Unlock(); err != nil {
			log.Println("worker:", workerid, err)
			continue
		}
	}

}
