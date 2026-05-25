package series

// RollingWindow is used for rolling window calculations.
type RollingWindow struct {
	window int
	series Series
}

// Rolling creates new RollingWindow
func (s Series) Rolling(window int) RollingWindow {
	return RollingWindow{
		window: window,
		series: s,
	}
}

// Mean returns the rolling mean.
func (r RollingWindow) Mean() (s Series) {
	s = New([]float64{}, Float, "Mean")
	for _, block := range r.getBlocks() {
		s.Append(block.Mean())
	}

	return
}

// StdDev returns the rolling mean.
func (r RollingWindow) StdDev() (s Series) {
    s = New([]float64{}, Float, "StdDev")
    for _, block := range r.getBlocks() {
        s.Append(block.StdDev())
    }

    return
}

// Max returns the rolling maximum values over the specified window.
// It computes the maximum for each rolling block and returns a Float Series
// containing the results. For blocks that are empty or contain non-numeric types
//, the value will be NaN consistent with Series.Max() behavior.
func (r RollingWindow) Max() (s Series) {
    s = New([]float64{}, Float, "Max")
    for _, block := range r.getBlocks() {
        s.Append(block.Max())
    }
    return
}

// Min returns the rolling minimum values over the specified window.
// It computes the minimum for each rolling block and returns a Float Series
// containing the results. For blocks that are empty or contain non-numeric types
//, the value will be NaN consistent with Series.Min() behavior.
func (r RollingWindow) Min() (s Series) {
    s = New([]float64{}, Float, "Min")
    for _, block := range r.getBlocks() {
        s.Append(block.Min())
    }
    return
}

func (r RollingWindow) getBlocks() (blocks []Series) {
	for i := 1; i <= r.series.Len(); i++ {
		if i < r.window {
			blocks = append(blocks, r.series.Empty())
			continue
		}

		index := []int{}
		for j := i - r.window; j < i; j++ {
			index = append(index, j)
		}
		blocks = append(blocks, r.series.Subset(index))
	}

	return
}
