package serializers

import (
	"go-google-scraper-challenge/lib/models"
)

type ResultResponse struct {
	ID        int64             `jsonapi:"primary,result"`
	Keyword   string            `jsonapi:"attr,keyword"`
	UserID    int64             `jsonapi:"attr,user_id"`
	Status    string            `jsonapi:"attr,status"`
	PageCache string            `jsonapi:"attr,page_cache"`
	User      *UserResponse     `jsonapi:"relation,user,omitempty"`
	AdLinks   []*AdLinkResponse `jsonapi:"relation,ad_links,omitempty"`
	Links     []*LinkResponse   `jsonapi:"relation,links,omitempty"`
}

type ResultsJSONResponse struct {
	Data []struct {
		ID         string `json:"id"`
		Attributes struct {
			Keyword   string `json:"keyword"`
			UserID    int64  `json:"user_id"`
			Status    string `json:"status"`
			PageCache string `json:"page_cache"`
		} `json:"attributes"`
		Relationships struct {
			User struct {
				Data RelationshipData `json:"data"`
			} `json:"user"`
			AdLinks struct {
				Data []RelationshipData `json:"data"`
			} `json:"ad_links"`
			Links struct {
				Data []RelationshipData `json:"data"`
			} `json:"links"`
		}
	} `json:"data"`
	Included []struct {
		ID         string                 `json:"id"`
		Type       string                 `json:"type"`
		Attributes map[string]interface{} `json:"attributes"`
	} `json:"included"`
}

type ResultSerializer struct {
	Result *models.Result
}

func (s ResultSerializer) Response() (response *ResultResponse) {
	response = &ResultResponse{
		ID:        s.Result.ID,
		Keyword:   s.Result.Keyword,
		UserID:    s.Result.UserID,
		Status:    s.Result.Status,
		PageCache: s.Result.PageCache,
	}

	if s.Result.User != nil {
		response.User = UserSerializer{User: s.Result.User}.Response()
	}

	if s.Result.AdLinks != nil {
		adLinks := []*AdLinkResponse{}
		for _, adLink := range s.Result.AdLinks {
			adLinks = append(adLinks, AdLinkSerializer{AdLink: adLink}.Response())
		}

		response.AdLinks = adLinks
	}

	if s.Result.Links != nil {
		Links := []*LinkResponse{}
		for _, Link := range s.Result.Links {
			Links = append(Links, LinkSerializer{Link: Link}.Response())
		}

		response.Links = Links
	}

	return response
}
