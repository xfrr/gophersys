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

func (a App) deleteGophers(id string) error {
	// TODO: implement this
	// _, err := a.cmdbus.Dispatch(context.Background(), queries.DeleteGopherCommand{ID: id})
	// if err != nil {
	// 	a.logger.Error().Err(err).Msg("failed to dispatch DeleteGopherCommand")
	// 	return err
	// }

	return nil
}

func (a App) createGopher(name, color string) error {
	// TODO: implement this
	// _, err := a.cmdbus.Dispatch(context.Background(), queries.CreateGopherCommand{Name: name, Color: color})
	// if err != nil {
	// 	a.logger.Error().Err(err).Msg("failed to dispatch CreateGopherCommand")
	// 	return err
	// }

	return nil
}

func (a App) updateGopher(id, name, color string) error {
	// TODO: implement this
	// _, err := a.cmdbus.Dispatch(context.Background(), queries.UpdateGopherCommand{ID: id, Name: name, Color: color})
	// if err != nil {
	// 	a.logger.Error().Err(err).Msg("failed to dispatch UpdateGopherCommand")
	// 	return err
	// }

	return nil
}
