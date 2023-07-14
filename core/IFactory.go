package core

import (
	"github.com/fatih/color"
	"os"
	"seeyonerExp/vulners"
)

type IFactory struct {
}

func (i *IFactory) NewFactory(name int) Seeyoner {
	switch name {
	case 0:
		return nil
	case 1:
		return &vulners.Sy01{}
	case 2:
		return &vulners.Sy02{}
	case 3:
		return &vulners.Sy03{}
	case 4:
		return &vulners.Sy04{}
	case 5:
		return &vulners.Sy05{}
	case 6:
		return &vulners.Sy06{}
	case 7:
		return &vulners.Sy07{}
	case 8:
		return &vulners.Sy08{}
	case 9:
		return &vulners.Sy09{}
	case 10:
		return &vulners.Sy10{}
	case 11:
		return &vulners.Sy11{}
	case 12:
		return &vulners.Sy12{}
	default:
		color.Red("[x]不存在的漏洞编号！可使用list命令查看")
		os.Exit(1)
		return nil
	}
}
