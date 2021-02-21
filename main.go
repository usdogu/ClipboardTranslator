package main

import (
	"context"

	"github.com/bregydoc/gtranslate"
	"github.com/gen2brain/dlgs"
	"golang.design/x/clipboard"
	"golang.org/x/text/language"
)

func main() {

	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		result, err := translate(string(data))
		if err != nil {
			dlgs.Error("Error!", err.Error())
		}
		dlgs.Info("Result", result)
	}
}

func translate(text string) (string, error) {
	result, err := gtranslate.Translate(text, language.English, language.Turkish)
	if err != nil {
		return "", err
	}
	return result, nil
}
