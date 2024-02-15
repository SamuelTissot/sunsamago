package sunsama

import (
	"context"
	"net/http"
	"time"

	"sunsamago/internal/session"

	"github.com/shurcooL/graphql"
)

const (
	sessionCookieName = "sunsamaSession"
	apiUrl            = "https://api.sunsama.com/graphql"
)

type cookieTransport struct {
	parent    http.RoundTripper
	sessionId string
}

func (c cookieTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	request.AddCookie(
		&http.Cookie{
			Name:  sessionCookieName,
			Value: c.sessionId,
		},
	)

	return c.parent.RoundTrip(request)
}

type Client struct {
	gqlClient *graphql.Client
	groupeID  string
	userID    string
	timezone  string
}

func NewClient(sessionID string, timezone string) (*Client, error) {
	httpClient := &http.Client{}
	httpClient.Transport = cookieTransport{
		parent:    http.DefaultTransport,
		sessionId: sessionID,
	}

	groupID, err := session.GroupID(sessionID)
	if err != nil {
		return nil, err
	}

	userID, err := session.UserID(sessionID)
	if err != nil {
		return nil, err
	}

	return &Client{
		gqlClient: graphql.NewClient(apiUrl, httpClient),
		groupeID:  groupID,
		userID:    userID,
		timezone:  timezone,
	}, nil
}

func (c *Client) streamsByGroupID() (streams, error) {
	var q struct {
		StreamsByGroupID []stream `graphql:"streamsByGroupId(groupId: $groupId)"`
	}

	variables := map[string]any{
		"groupId": graphql.String(c.groupeID),
	}

	err := c.gqlClient.Query(context.Background(), &q, variables)
	if err != nil {
		return nil, err
	}

	return q.StreamsByGroupID, nil
}

func (c *Client) TaskByDay(day time.Time) ([]task, error) {
	var q struct {
		TasksByDay []task `graphql:"tasksByDay(day: $day,timezone: $timezone,userId: $userId,groupId: $groupId)"`
	}

	variables := map[string]any{
		"day":      graphql.String(day.Format(time.DateOnly)),
		"timezone": graphql.String(c.timezone),
		"userId":   graphql.String(c.userID),
		"groupId":  graphql.String(c.groupeID),
	}

	err := c.gqlClient.Query(context.Background(), &q, variables)
	if err != nil {
		return nil, err
	}

	return q.TasksByDay, nil
}
