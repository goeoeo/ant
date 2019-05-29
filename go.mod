module ant

require github.com/astaxie/beego v1.11.1

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85
	golang.org/x/net => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
)
