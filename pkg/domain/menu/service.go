package menu

type Service struct {
	imageParser ImageParser
}

func NewService(parser ImageParser) Service {
	return Service{
		parser,
	}
}

func (s Service) ParseMenuImage(url string) (*Menu, error) {
	return s.imageParser.Parse(url)
}
