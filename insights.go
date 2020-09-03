package facebookgraph

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type InsightsResponse struct {
	Data   []Insight `mapstructure:"data"`
	Paging Paging    `mapstructure:"paging"`
}

type InsightValue struct {
	Value   string `mapstructure:"value"`
	EndTime string `mapstructure:"end_time"`
}
type Insight struct {
	Name        string         `mapstructure:"name"`
	Period      string         `mapstructure:"period"`
	Values      []InsightValue `mapstructure:"values"`
	Title       string         `mapstructure:"title"`
	Description string         `mapstructure:"description"`
	ID          string         `mapstructure:"id"`
}

// Insights return Instagram insights for a user
//
func (fg *FacebookGraph) Insights(objectID string, metrics []string, period *string, since *int64, until *int64) (*[]Insight, error) {
	params := make(map[string]interface{})
	params["metric"] = strings.Join(metrics, ",")
	if period != nil {
		params["period"] = *period
	}
	if since != nil {
		params["since"] = *since
	}
	if until != nil {
		params["until"] = *until
	}

	result, err := fg.session.Get(fmt.Sprintf("/%s/insights", objectID), params)
	if err != nil {
		return nil, err
	}

	response := InsightsResponse{}
	err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}