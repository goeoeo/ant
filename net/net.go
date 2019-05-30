package net

import "os/exec"

//检测网络状态
func NetWorkStatus(ip string) bool {
	cmd := exec.Command("ping", ip, "-c", "1", "-W", "5")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
