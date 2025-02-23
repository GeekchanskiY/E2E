package myfin

import (
	"context"
	"time"

	"go.uber.org/zap"

	"finworker/internal/models"
)

const (
	timeZoneDifference  = 3 * time.Hour // bad practice, but I dont see other way to resolve this now
	inspectionFrequency = 12 * time.Hour
)

func RunPeriodicScraping(scraper *Scraper) error {
	ctx := context.Background()

	errChan := make(chan error)
	lastUpdateTime, err := scraper.currencyStatesRepo.GetLastUpdate(ctx)
	if err != nil {
		return err
	}

	// running periodic scraping
	ticker := time.NewTicker(inspectionFrequency)
	defer ticker.Stop()

	// running initial scraping if required
	if lastUpdateTime.IsZero() || lastUpdateTime.Before(time.Now().Add(inspectionFrequency)) {
		go runScraping(ctx, errChan, scraper)
	} else {
		ticker.Reset(inspectionFrequency - time.Now().Add(timeZoneDifference).Sub(lastUpdateTime))
	}

	select {
	case <-ctx.Done():
		scraper.logger.Info("scraping finished")
		return nil
	case <-ticker.C:
		go runScraping(ctx, errChan, scraper)
		ticker.Reset(inspectionFrequency)
	case err := <-errChan:
		scraper.logger.Warn("scraping failed", zap.Error(err))
		return err
	}

	return nil
}

func runScraping(ctx context.Context, errChan chan error, scraper *Scraper) {
	currencies, err := scraper.GetCurrencies()
	if err != nil {
		errChan <- err

		return
	}

	for _, currency := range currencies {
		switch currency.BankName {
		case "INSNC by Alfa Bank":
			bank, err := scraper.banksRepo.GetByName("alfabank")
			if err != nil {
				errChan <- err

				return
			}
			_, err = scraper.currencyStatesRepo.Create(ctx, &models.CurrencyState{
				BankId:       bank.Id,
				CurrencyName: currency.Name,
				SourceName:   "INSNC", // alfabank app name
				SellUsd:      currency.SellUsd,
				BuyUsd:       currency.BuyUsd,
				Time:         currency.Time,
			})
			if err != nil {
				errChan <- err

				return
			}
			break
		case "Альфа Банк":
			bank, err := scraper.banksRepo.GetByName("alfabank")
			if err != nil {
				errChan <- err

				return
			}
			_, err = scraper.currencyStatesRepo.Create(ctx, &models.CurrencyState{
				BankId:       bank.Id,
				CurrencyName: currency.Name,
				SourceName:   "BANK", // alfabank app name
				SellUsd:      currency.SellUsd,
				BuyUsd:       currency.BuyUsd,
				Time:         currency.Time,
			})
			if err != nil {
				errChan <- err

				return
			}
			break
		case "Приорбанк":
			bank, err := scraper.banksRepo.GetByName("priorbank")
			if err != nil {
				errChan <- err

				return
			}
			_, err = scraper.currencyStatesRepo.Create(ctx, &models.CurrencyState{
				BankId:       bank.Id,
				CurrencyName: currency.Name,
				SourceName:   "BANK", // alfabank app name
				SellUsd:      currency.SellUsd,
				BuyUsd:       currency.BuyUsd,
				Time:         currency.Time,
			})
			if err != nil {
				errChan <- err

				return
			}
			break
		default:
			scraper.logger.Debug("bank skipping", zap.String("bank", currency.BankName))

		}
	}

	return
}
