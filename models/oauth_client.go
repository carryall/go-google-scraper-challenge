package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type OAuthClient struct {
	ClientID     string `orm:"column(id);pk" json:"client_id,omitempty"`
	ClientSecret string `orm:"column(secret)" json:"client_secret,omitempty"`
	Domain       string
	Data         string `orm:"type(jsonb)"`
}

func init() {
	orm.RegisterModel(new(OAuthClient))
}

func (client *OAuthClient) TableName() string {
	return "oauth2_clients"
}

// FindClientByID retrieves OAuthClient by ID and returns error if OAuthClient with given ID doesn't exist.
func FindClientByID(id string) (client *OAuthClient, err error) {
	ormer := orm.NewOrm()
	client = &OAuthClient{ClientID: id}

	err = ormer.QueryTable(OAuthClient{}).Filter("ClientID", id).RelatedSel().One(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}
