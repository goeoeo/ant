package services

import (
	"ant"
	"ant/stringutil"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLwagentServicePost(t *testing.T) {
	//var res []entity.Client

	lwagent := new(lwagentService)
	clientListParams := ant.StockHsas{Name: "test"}
	err := lwagent.Post("127.0.0.1", ClientList, clientListParams)
	if err != nil {
		t.Error(err)
	}

}

func TestDes(t *testing.T) {
	lwagent := new(lwagentService)
	//解密map[code:gc7jOoVVeCUlsISvOWa9uU28QBTLU399 time:1554113158 data:P3B3MUaNMeKXSIF1Erm4gTQOdTvQMZIa2AumFM-w_bEq0OmIcPLxKDM=]
	//key u0vODMj1c2XX4Kb2
	aesKey := lwagent.getAesKey("wmXhAHdGaTSoCzj4Cfcn8ex8OAN1pwUs", "1554355904")

	fmt.Println("aesKey:", aesKey)
	str, ok := stringutil.AESDecrypt([]byte(aesKey), "slgfYU_nX4r0NgsxeHoR8lgSnABUuLP8SUTmsbaTpkdW4R4STnKx4ABnJ_avw8ysW0CVlxwdHS5_pIXxL4-eCEFTwU_s7Z-UmiyacPzZ80_1JgbMGF4W6r1TpKHJV7mR-NExz8ci3euLJFZZFKK4e8ASudlohG9_Z14b8sULr2yfuJd3xcAM39uMxIkb5Y0S8zBy2jfj5cIdraSUj7nOun9QdqxMNIDhFUNtvDhZtamiZ51LJT0MfxCO16DLowTACyW7g0rnr0K6mHrYniBXbt7jMv-EjJJT0Cli05sp2s9VB0wMxYgjkrf_88IdxPitb29lfVenZ-rBoEb0GzdRH865DbOrCah7c21manpkXvjAqpsIwUN83ufc75LbDH5nvKffwN9rhJ-xS5r_x-esxYDqfHuqO0_UciWlz96V4AcnjdBB9QTX6obVowJC5rBkmyWGWV7OXYFMMCpuw_Kx8Eb2S229uN3qQ3S6wazHzdnGyT5FVy13pBMCi5MgqeZkgSdT3l4Dc_y-t1LK26nzvsublMJpPovllQtqtyAutW3q8PEj0v9_NJbdcQ==")
	if !ok {
		t.Error("解密失败")
	}

	fmt.Println("str:", str)

}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(10)
	fmt.Println(rnd)
}
