package encryptreq

import (
	"fmt"
	"github.com/phpdi/ant"
	"math/rand"
	"testing"
	"time"
)

var encryptReq EncryptReq

func init() {
	encryptReq = EncryptReq{}

}

func TestEncryptReq_PostCheckCode(t *testing.T) {
	req := ant.StockParse{
		Id:   2,
		Code: "aaa",
		Name: "bbb",
	}

	ack := ant.StockParse{}

	encryptReq.PostCheckCode("http://localhost:8080/", req, &ack)
	fmt.Println(ack)
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(10)
	fmt.Println(rnd)
}
