package types

import (
	"errors"
	"fmt"
)

func (W *World) PlotCount() int {
	return len(W.plots)
}

func (W *World) GetPlot(plot_id uint64) (*Plot, error) {
	W.plots_lock.RLock()
	defer W.plots_lock.RUnlock()
	if val, ok := W.plots[NewPlotKey(plot_id)]; ok {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("plot %d could not be found", plot_id))
}

func (W *World) SearchPlot(x, z int64) (*Plot, error) {
	if (x <= 2 && x >= -2) || (z <= 2 && z >= -2) {
		return nil, errors.New(fmt.Sprintf("plot at %d %d disabled", x, z))
	}
	if val, ok := W.cache.CheckPlot(x, z); ok {
		W.plots_lock.RLock()
		defer W.plots_lock.RUnlock()
		return W.GetPlot(val)
	}
	return nil, errors.New(fmt.Sprintf("plot at %d %d not found", x, z))
}
