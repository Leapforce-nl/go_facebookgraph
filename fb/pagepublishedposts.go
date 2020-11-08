package facebookgraph

import (
	"fmt"

	fb2 "github.com/huandu/facebook/v2"
	api "github.com/leapforce-libraries/go_facebookgraph/api"
	models "github.com/leapforce-libraries/go_facebookgraph/models"
	utils "github.com/leapforce-libraries/go_utilities"
)

const pagePublishedPostsLimit int = 50 //limit 100 icm comments does not work...

type PagePublishedPostsResponse struct {
	Data   []PagePost `mapstructure:"data"`
	Paging Paging     `mapstructure:"paging"`
}

type PagePostFrom struct {
	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

type PagePostShares struct {
	Count int64 `mapstructure:"count"`
}

type PagePost struct {
	ID           string             `mapstructure:"id"`
	Attachments  models.Attachments `mapstructure:"attachments"`
	CreatedTime  string             `mapstructure:"created_time"`
	From         PagePostFrom       `mapstructure:"from"`
	FullPicture  string             `mapstructure:"full_picture"`
	Message      string             `mapstructure:"message"`
	PermalinkURL string             `mapstructure:"permalink_url"`
	Shares       PagePostShares     `mapstructure:"shares"`
	StatusType   string             `mapstructure:"status_type"`
	UpdatedTime  string             `mapstructure:"updated_time"`
}

// PagePublishedPosts return Instagram medias for a user
//
func (fb *Facebook) PagePublishedPosts(pageID string, accessToken string, after string) (*PagePublishedPostsResponse, error) {
	path := fmt.Sprintf("/%s/published_posts", pageID)

	params := fb2.Params{
		"limit":        pagePublishedPostsLimit,
		"after":        after,
		"access_token": accessToken,
		"fields":       utils.GetTaggedTagNames("mapstructure", PagePost{}),
	}

	result, err := api.GetWithRetry(fb.session, path, params)
	if err != nil {
		return nil, err
	}

	response := PagePublishedPostsResponse{}
	err = result.DecodeField("", &response)
	//err = mapstructure.Decode(result, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
