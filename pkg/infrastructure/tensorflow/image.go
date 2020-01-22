package tensorflow

import "github.com/gileslloyd/menu-parser/pkg/domain/menu"

type MlParser struct {

}

func (p MlParser) Parse(url string) *menu.Menu {
	menu := menu.NewMenu(
		"hbfdiff-dfgd-dgfgf-sfgd",
		"Hot Pocket: £4.99",
	)

	return menu
}

