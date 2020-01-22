package menu

type ImageParser interface {
	Parse(url string) (*Menu, error)
}
