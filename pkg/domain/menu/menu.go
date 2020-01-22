package menu

type Menu struct {
	TraderID string
	Content string
}

func NewMenu(traderID string, content string) *Menu {
	m := Menu{
		TraderID: traderID,
		Content:  content,
	}

	return &m
}
