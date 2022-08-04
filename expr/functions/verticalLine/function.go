//go:build !cairo
// +build !cairo

package verticalLine

import (
	"context"
	"fmt"
	"github.com/grafana/carbonapi/expr/interfaces"
	"github.com/grafana/carbonapi/expr/types"
	"github.com/grafana/carbonapi/pkg/parser"
)

var UnsupportedError = fmt.Errorf("must build w/ cairo support")

type verticalLine struct {
	interfaces.FunctionBase
}

func GetOrder() interfaces.Order {
	return interfaces.Any
}

func New(_ string) []interfaces.FunctionMetadata {
	res := make([]interfaces.FunctionMetadata, 0)

	f := &verticalLine{}
	functions := []string{"verticalLine"}
	for _, n := range functions {
		res = append(res, interfaces.FunctionMetadata{Name: n, F: f})
	}

	return res
}

func (f *verticalLine) Do(_ context.Context, _ parser.Expr, _, _ int64, _ map[parser.MetricRequest][]*types.MetricData) ([]*types.MetricData, error) {
	return nil, UnsupportedError
}

func (f *verticalLine) Description() map[string]types.FunctionDescription {
	return map[string]types.FunctionDescription{
		"verticalLine": {
			Description: "Draws a vertical line at the designated timestamp with optional\n  'label' and 'color'. Supported timestamp formats include both\n  relative (e.g. -3h) and absolute (e.g. 16:00_20110501) strings,\n  such as those used with ``from`` and ``until`` parameters. When\n  set, the 'label' will appear in the graph legend.",
			Function:    "verticalLine(ts, label=None, color=None)",
			Group:       "Graph",
			Module:      "graphite.render.functions",
			Name:        "verticalLine",
			Params: []types.FunctionParam{
				{
					Name:     "ts",
					Required: true,
					Type:     types.Date,
				},
				{
					Name:     "label",
					Required: false,
					Type:     types.String,
				},
				{
					Name:     "color",
					Required: false,
					Type:     types.String,
				},
			},
		},
	}
}
