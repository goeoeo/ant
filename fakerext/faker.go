package fakerext

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/phpdi/ant/util"
	"math/rand"
	"reflect"
)

func init() {
	_ = faker.AddProvider("Mobile", Mobile)
	_ = faker.AddProvider("CommunityName", CommunityName)
	_ = faker.AddProvider("WxAvatar", WxAvatar)
	_ = faker.AddProvider("GoodsImg", GoodsImg)
	_ = faker.AddProvider("Price", Price)

}

func Price(v reflect.Value) (interface{}, error) {

	return rand.Intn(1000)*100, nil
}


func Mobile(v reflect.Value) (interface{}, error) {

	return fmt.Sprintf("1%d%s", rand.Intn(6)+3, util.RandomCreateBytes(9, []byte("1234567890")...)), nil
}

func CommunityName(v reflect.Value) (interface{}, error) {
	return randomString([]string{"光大名筑", "观筑庭园", "观河锦苑", "国典华园", "花样年华", "国展家园", "甘露家园", "广通苑", "冠云庄园", "冠城园", "硅谷先锋", "国展", "新座甘露家园", "果园西小区", "冠雅苑", "九台庄园", "名流花园"}), nil
}

func WxAvatar(v reflect.Value) (interface{}, error) {
	return randomString([]string{
		"https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=2821036723,3583437323&fm=26&gp=0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793775487&di=457120721ca3356af2daa5d5b7155b50&imgtype=0&src=http%3A%2F%2Fimg4.imgtn.bdimg.com%2Fit%2Fu%3D2875089794%2C2078586948%26fm%3D214%26gp%3D0.jpg",
		"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3816071403,244674382&fm=26&gp=0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793809397&di=dee10f124639412068767aae17c41ddc&imgtype=0&src=http%3A%2F%2Fimg1.imgtn.bdimg.com%2Fit%2Fu%3D1338403846%2C1469349235%26fm%3D214%26gp%3D0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793755205&di=9a4c3ad6638c5cdb81a537cd0af3bb75&imgtype=0&src=http%3A%2F%2Fpic.rmb.bdstatic.com%2Fddb69eb2bee4d2e9bb15a2aab1dd4e7c.jpeg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793755205&di=bb8fbd9fb6b33040f614dbebf0a17f52&imgtype=0&src=http%3A%2F%2Fb-ssl.duitang.com%2Fuploads%2Fitem%2F201711%2F28%2F20171128011630_JAvTc.thumb.700_0.jpeg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793755204&di=44fb0ab35185adc94fec7e222359f01e&imgtype=0&src=http%3A%2F%2Fimg.bq233.com%2Fkanqq%2Fpic%2Fupload%2F2018%2F0925%2F1537862023130293.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793845261&di=e09d562ee0ce47e0b8593979856904b5&imgtype=0&src=http%3A%2F%2Fpic4.zhimg.com%2Fv2-17d57d79a1181d1a6da3ffe899e58d03_b.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793755203&di=737d8ba624186c1f48ebd2058e32ecc5&imgtype=0&src=http%3A%2F%2Fwechat.shwilling.com%2Fuploads%2Fueditor%2Fimage%2F20160821%2F1471760931761147.jpg",
		"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=2315655086,265578671&fm=26&gp=0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793755202&di=9a669d1f7ba33f6af8325057a714cd14&imgtype=0&src=http%3A%2F%2Fimg.duoziwang.com%2F2016%2F12%2F17%2F22090992876.jpg",
		"https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2165449282,2094365905&fm=26&gp=0.jpg",
		"https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=2325214326,1485162395&fm=26&gp=0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793755201&di=6e0746ccbd97b596330439a9d24a0fb2&imgtype=0&src=http%3A%2F%2Fpic4.zhimg.com%2F50%2Fv2-f41af535c044d503346cc4be802b7724_hd.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793897544&di=65a1514959eac9e63a61731f6a93e355&imgtype=0&src=http%3A%2F%2Fimg3.imgtn.bdimg.com%2Fit%2Fu%3D2564818988%2C1089303898%26fm%3D214%26gp%3D0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793918569&di=4464d3ee69eef85011affbaa24c50181&imgtype=0&src=http%3A%2F%2Fimg0.imgtn.bdimg.com%2Fit%2Fu%3D2834303150%2C3850588710%26fm%3D214%26gp%3D0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594793931360&di=70ea8a737b6beeb036f3e8411dbeffc0&imgtype=0&src=http%3A%2F%2Fimg.mp.itc.cn%2Fupload%2F20161114%2F055b277895dc4809994c3f6061edeeca_th.jpeg",
	}), nil
}

func GoodsImg(v reflect.Value) (interface{}, error) {
	return randomString([]string{
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803413257&di=910341ea2e9c458d62583ceae1493136&imgtype=0&src=http%3A%2F%2Fwww.esys.cn%2Fuploads%2Fallimg%2F191001%2F1525213457_0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803413256&di=41597a8a12dc8a196d4f5901c9a02b91&imgtype=0&src=http%3A%2F%2Fwww.szthks.com%2Flocalimg%2F687474703a2f2f6777312e616c6963646e2e636f6d2f62616f2f75706c6f616465642f69372f5431595670794672787458585858585858585f2121302d6974656d5f7069632e6a7067.jpg",
		"https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=368471429,3580886966&fm=26&gp=0.jpg",
		"https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2910748489,1515877704&fm=26&gp=0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502092&di=9cf2d8e81167662a2bfdeec04a5fab08&imgtype=0&src=http%3A%2F%2Fwww2.flightclub.cn%2Fnews%2Fuploads%2Fallimg%2F161017%2F3-16101FT927.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502090&di=b65927f151c92b2a7d9d4078338fc6c1&imgtype=0&src=http%3A%2F%2Fwww2.flightclub.cn%2Fnews%2Fuploads%2Fallimg%2F170110%2F1R11J260-9_resized.jpg",
		"https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1918200844,94046446&fm=26&gp=0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502079&di=c1b03e155f5aba982b21631f29f3579a&imgtype=0&src=http%3A%2F%2Fimg10.360buyimg.com%2Fn0%2Fjfs%2Ft2668%2F172%2F1330788370%2F265854%2F76435fc4%2F573bd165Ne0fd3c50.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502078&di=ace760e3562a43e6dce530059c8f6b27&imgtype=0&src=http%3A%2F%2Fd.ifengimg.com%2Fw600%2Fp0.ifengimg.com%2Fpmop%2F2018%2F0713%2F3A517C9A7DA8783242B98CF0F93618526E80F55A_size187_w1080_h1080.jpeg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502076&di=a72e1fae8cd1767e433da880309e8fc5&imgtype=0&src=http%3A%2F%2Fwww2.flightclub.cn%2Fnews%2Fuploads%2Fallimg%2F170825%2F6-1FR51J914.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502074&di=9a6280412df537d02d61a583bbf46e2d&imgtype=0&src=http%3A%2F%2Fimg.yzcdn.cn%2Fupload_files%2F2018%2F01%2F09%2FFkMKAWBqBe9_GsvHNMzFmXOtzZEH.jpg%2521730x0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803570813&di=e79c334cf5d0954c185ffc61ab98e7d6&imgtype=0&src=http%3A%2F%2Fimg2.imgtn.bdimg.com%2Fit%2Fu%3D2719177227%2C1808577862%26fm%3D214%26gp%3D0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502062&di=9f5dd7b9e4f8568a85dcae22683d435b&imgtype=0&src=http%3A%2F%2Fm.360buyimg.com%2Fn12%2Fjfs%2Ft529%2F332%2F921366202%2F354337%2Fcc4bffe2%2F549a6652Nf0af29d3.jpg%2521q70.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502059&di=5676dd0d150a00baf56dbb3e4f938086&imgtype=0&src=http%3A%2F%2Fwww.i-size.com%2Fsystem%2Fupload%2F201312%2Fc5af24f844271316327679c2df18990b.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502056&di=90d6c67af3d9efd4c6a0c1587d36f050&imgtype=0&src=http%3A%2F%2Fwww.flightclub.cn%2Fnews%2Fuploads%2Fallimg%2F180112%2F12-1P1121K228.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803502026&di=a006af996d114c985bf26e3ab38d224a&imgtype=0&src=http%3A%2F%2Finews.gtimg.com%2Fnewsapp_bt%2F0%2F11929327998%2F641.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803602126&di=b267c692b1a6e46a4769063b17f3c7c2&imgtype=0&src=http%3A%2F%2Fwww2.flightclub.cn%2Fnews%2Fuploads%2Fallimg%2Fc140521%2F1400629321162Z-53Q0_resized.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803635139&di=84b9f6b8bf4642c912e3dd6e26fe1c7a&imgtype=0&src=http%3A%2F%2Fimage.suning.cn%2Fuimg%2Fsop%2Fcommodity%2F209391578252638084342831_x.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803602122&di=551144a9e55f53fd4d06274927523d2f&imgtype=0&src=http%3A%2F%2Fwww.sneakers.com.cn%2Fdata%2Fattachment%2Fforum%2F201609%2F12%2F152348dpzxp77np5s87z55.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803649185&di=d6360ba316f85ef45cc8c62147a56242&imgtype=0&src=http%3A%2F%2Fimg3.imgtn.bdimg.com%2Fit%2Fu%3D3680710608%2C149078428%26fm%3D214%26gp%3D0.jpg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803602121&di=47c797ccf256a6caa4e3b577828e8cdd&imgtype=0&src=http%3A%2F%2Fp0.ifengimg.com%2Fpmop%2F2018%2F0917%2F29CF2B8F01C52B5663F926967C816C001BB120A3_size135_w1080_h771.jpeg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803602114&di=6c1974853b5a31ed2fa8ee4185b780c1&imgtype=0&src=http%3A%2F%2Fimages.dunkhome.com%2Fstatic_files%2F2017-08-18%2FmnncdxfEEJzktsoZk4hn_image_wh_840x610.jpeg",
		"https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1594803602111&di=a4923700d2906184e24618fe6133c51d&imgtype=0&src=http%3A%2F%2Fi0.hdslb.com%2Fbfs%2Farticle%2F6ed52107fc80766c44d1c8a7d7dc0914174602d2.jpg",
	}), nil
}
func randomString(s []string) string {
	return s[rand.Intn(len(s))]
}
