package frontend

import (
	"context"
	"html/template"
	"math/rand"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
)

func (c *Controller) Index(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.index.controller", zap.String("event", "got request"))
	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.IndexTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	// random message generator
	verbs := []string{"Люблю", "Кушаю", "Уважаю", "Умиляю", "Обнимаю", "Жгу", "Глажу", "Шакалю", "Ем"}
	nouns := []string{"Аську", "Димку", "питсу", "пельмешки", "", "айфон", "пиксель", "яблоко", "Альберта Эйнштейна"}

	data["text"] = verbs[rand.Intn(len(verbs))] + " " + nouns[rand.Intn(len(nouns))]
	return html, data, nil
}
