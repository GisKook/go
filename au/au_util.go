package au

import (
	"github.com/giskook/go/util"
	"strings"
)

func GetMachineNo() string {
	base_board_sn := strings.TrimSpace(util.GetBaseBoardSn())
	hw_addr := strings.TrimSpace(util.GetHWaddr())

	machine_code := "echo " + base_board_sn + hw_addr + "|tr '[a-z]' '[A-Z]'|tr -d ':'|tr -d '-'|tr '[0-9]' [A-J]"
	return util.UtilRunCmd(machine_code)
}
