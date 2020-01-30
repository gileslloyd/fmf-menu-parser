package cloudvision

import (
	"context"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/gileslloyd/menu-parser/internal/domain/menu"
)

type MlParser struct {
}

func NewMLParser() menu.ImageParser {
	return MlParser{}
}

func (p MlParser) Parse(url string) (*menu.Menu, error) {
	menuContent, err := detectTextURI(url)

	if err == nil {
		parsedMenu := menu.NewMenu(
			"hbfdiff-dfgd-dgfgf-sfgd",
			menuContent,
		)

		return parsedMenu, nil
	}

	return nil, err
}

func detectTextURI(file string) (content string, err error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return "", err
	}

	image := vision.NewImageFromURI(file)
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		return "", err
	}

	if len(annotations) == 0 {
		return "No text found.", nil
	}

	menuContent := ""
	for _, annotation := range annotations {
		menuContent += annotation.Description
	}

	return menuContent, nil
}
