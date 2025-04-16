package base

import (
	"context"
	"html/template"
	"math/rand/v2"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
)

func (c *controller) Index(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.index.controller", zap.String("event", "got request"))
	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.IndexTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	// random server message
	randomText := []string{
		"Билд в студию",
		"Это не баг, это фича",
		"Сдал релиз иди в Ливиз",
		"Когда я вижу, как ты верстаешь, Малыш, ты меня пугаешь",
		"Хорошо написанная программа — это программа, написанная 2 раза",
		"Ничто так не ограничивает полет мысли начинающего программиста, как компилятор",
		"Ничто так не ограничивает полет мысли зрелого программиста, как Project Manager",
		"Хорошая жена может провожать мужа-программиста на работу словами «Чистого кода тебе!»",
		"Тяжела и неказиста жизнь без парня-программиста",
		"Программисты — это устройства, преобразующие кофеин в код",
		"Зачем я работаю...",
		"Sex, drugs & rock'n'roll? — Bugs, hex & source control!",
		"Улучшение работающего продукта приводит к его ухудшению",
		"Хорошо задокументированный баг, автоматически становится фичей!",
		"Критичный баг, найденный тестировщиком в последний день, является багом в работе самого тестировщика",
		"Если что-то может сломаться, оно должно сломаться именно сейчас",
		"Быстро откаченное выложенным не считается",
		"Семь бед — один reset",
		"Не было печали — апдейтов накачали",
		"Первый файл com’ом",
		"Какой error не мечтает стать general’ом…",
		"Первый тост за локалхост",
		"Знаю отличную шутку про UDP, но не факт, что она до вас дойдет",
		"Восстановление системы после критического сбоя, у админа похмелье",
		"Глаза болят, а руки делают",
		"double ять!",
		"Поддерживаю устройства категории plug'n'pray",
		"#define true false",
		"Ой, что-то пошло не так. А нет, всё так",
		"НЛО прилетело и опубликовало эту надпись здесь",
		"Не соврал, а ударил пизде-джитсу",
	}

	data["text"] = randomText[rand.IntN(len(randomText))]
	return html, data, nil
}
