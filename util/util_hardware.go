package util

import (
	"log"
	"os/exec"
)

func util_run_cmd(cmd string) string {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Printf("Failed to excute command %s\n", cmd)
	}
	return string(out)
}

func UtilRunCmd(cmd string) string {
	return util_run_cmd(cmd)
}

func GetHWaddr() string {
	return util_run_cmd("/sbin/ifconfig -a |grep HWaddr|egrep -i \"eth0|em1\" |awk '{print $NF}'")
}

func GetBaseBoardSn() string {
	return util_run_cmd("dmidecode -s baseboard-serial-number")
}
