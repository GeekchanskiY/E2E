package myfin

import (
	"context"
	"time"
)

func RunPeriodicInspection(
	ctx context.Context,
	scraper *Scraper,
) error {
	lastUpdateTime, err := getLastUpdateTime(ctx, scraper)
	if err != nil {
		return err
	}

	if lastUpdateTime == nil {

	}
	err = runInspection(ctx)
	if err != nil {
		return err
	}
	return nil
}

func getLastUpdateTime(ctx context.Context, s *Scraper) (*time.Time, error) {
	return s.currencyStatesRepo.GetLastUpdate(ctx)
}

func runInspection(ctx context.Context) error {
	return nil
}
