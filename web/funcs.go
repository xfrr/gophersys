package web

import (
	"context"

	"github.com/xfrr/gophersys/internal/queries"
)

func (a App) getGophers() []queries.GopherView {
	res, err := a.queryBus.Dispatch(context.Background(), queries.GetGophersQuery{})
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to dispatch GetGophersQuery")
		return nil
	}

	result, ok := res.(queries.GetGophersQueryResult)
	if !ok {
		a.logger.Error().Msg("failed to cast result to GetGophersQueryResult")
		return nil
	}

	return result.Gophers
}
