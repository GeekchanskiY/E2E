package myfin

import (
	"context"
	"strconv"
	"time"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

const currencyURL = "https://myfin.by/currency/usd"

type Scraper struct {
	ctx    context.Context
	logger *zap.Logger
}

func New(log *zap.Logger) *Scraper {

	return &Scraper{
		ctx:    context.Background(),
		logger: log,
	}
}

func GetCurrencies(s *Scraper) ([]*Currency, error) {
	var currencies []*Currency

	c := colly.NewCollector(
		colly.AllowedDomains("myfin.by"),
	)

	c.OnHTML("table.currencies-courses", func(e *colly.HTMLElement) {
		e.ForEach("tr.currencies-courses__row-main", func(_ int, el *colly.HTMLElement) {
			newCurrencyByn := &Currency{}
			newCurrencyEur := &Currency{}
			newCurrencyRub := &Currency{}
			el.ForEach("td", func(col int, el *colly.HTMLElement) {
				// 0 - bank name
				// 1/2 - sell/buy usd
				// 3/4 - sell/buy eur/usd
				// 5/6 - sell/buy usd/rub
				switch col {
				case 0:
					newCurrencyByn.BankName = el.Text
					newCurrencyEur.BankName = el.Text
					newCurrencyRub.BankName = el.Text
				case 1:
					if el.Text == "" || newCurrencyByn == nil {
						newCurrencyByn = nil
						return
					}
					sellUsd, err := strconv.ParseFloat(el.Text, 64)
					if err != nil {
						s.logger.Error("Error parsing sell usd", zap.String("error", err.Error()))
					}
					newCurrencyByn.SellUsd = sellUsd
				case 2:
					if el.Text == "" || newCurrencyByn == nil {
						newCurrencyByn = nil
						return
					}
					buyUsd, err := strconv.ParseFloat(el.Text, 64)
					if err != nil {
						s.logger.Error("Error parsing sell usd", zap.String("error", err.Error()))
					}
					newCurrencyByn.BuyUsd = buyUsd
				case 4:
					if el.Text == "" || newCurrencyEur == nil {
						newCurrencyEur = nil
						return
					}
					sellUsd, err := strconv.ParseFloat(el.Text, 64)
					if err != nil {
						s.logger.Error("Error parsing sell usd", zap.String("error", err.Error()))
					}
					newCurrencyEur.SellUsd = sellUsd
				case 5:
					if el.Text == "" || newCurrencyRub == nil {
						newCurrencyRub = nil
						return
					}

					sellUsd, err := strconv.ParseFloat(el.Text, 64)
					if err != nil {
						return
					}
					newCurrencyRub.SellUsd = sellUsd
				case 6:
					if el.Text == "" || newCurrencyRub == nil {
						newCurrencyRub = nil
						return
					}

					buyUsd, err := strconv.ParseFloat(el.Text, 64)
					if err != nil {
						s.logger.Error("Error parsing sell usd", zap.String("error", err.Error()))
					}
					newCurrencyRub.BuyUsd = buyUsd
				}
			})
			if newCurrencyByn != nil {
				newCurrencyByn.Time = time.Now()
				currencies = append(currencies, newCurrencyByn)
			}
			if newCurrencyEur != nil {
				newCurrencyEur.Time = time.Now()
				currencies = append(currencies, newCurrencyEur)
			}
			if newCurrencyRub != nil {
				newCurrencyRub.Time = time.Now()
				currencies = append(currencies, newCurrencyRub)
			}
		})
	})

	err := c.Visit(currencyURL)
	if err != nil {
		return nil, err
	}
	s.logger.Info("Scraped currencies", zap.Int("count", len(currencies)))

	return currencies, nil
}
