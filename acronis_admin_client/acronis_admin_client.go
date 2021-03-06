package acronis_admin_client

import (
	"google.golang.org/api/admin/directory/v1"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"context"
)

type AdminClient struct {
	s *admin.Service
	domain string
}

func Init(domain string, admin_email string) (*AdminClient, error) {
	client := AdminClient{}

	ctx := context.Background()

	b, err := ioutil.ReadFile("./Acronis-backup-project-8b80e5be7c37.json")
	if err != nil {
		return nil, err
	}

	data, err := google.JWTConfigFromJSON(b, admin.AdminDirectoryUserReadonlyScope)
	if err != nil {
		return nil, err
	}

	data.Subject = admin_email

	client.s, err = admin.New(data.Client(ctx))
	if err != nil {
		return nil, err
	}

	client.domain = domain

	return &client, nil
}

func (c *AdminClient) GetListOfUsers() (*admin.Users, error) {
	usersListCall := c.s.Users.List()
	usersListCall = usersListCall.Domain(c.domain)
	usersListCall = usersListCall.Projection("full")

	res, err := usersListCall.Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}