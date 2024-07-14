package hubspot

import (
	"fmt"
	"time"
)

const (
	marketingEventBasePath = "marketing-events/events"
)

// marketingEventService is an interface for the marketingEvent/marketingEvent service of the HubSpot API.
// Reference: https://developers.hubspot.com/docs/api/marketing/marketing-events
type MarketingEventService interface {
	CreateUpdateEvent(event *MarketingEvent) (*CreateEventSingleResponse, error)
}

var _ MarketingEventService = (*MarketingEventServiceOp)(nil)

// marketingEventServiceOp provides the default implementation of MarketingEvent.
type MarketingEventServiceOp struct {
	marketingEventPath string
	client             *Client
}

type MarketingEvent struct {
	EventName         HsStr       `json:"eventName"`
	ExternalAccountId HsStr       `json:"externalAccountId"`
	ExternalEventId   HsStr       `json:"externalEventId"`
	EventType         *HsStr      `json:"eventType,omitempty"`
	StartDate         *HsTime     `json:"startDateTime,omitempty"`
	EndDate           *HsTime     `json:"endDateTime,omitempty"`
	EventOrganizier   *HsStr      `json:"eventOrganizer,omitempty"`
	EventDescription  *HsStr      `json:"eventDescription,omitempty"`
	EventUrl          *HsStr      `json:"eventUrl,omitempty"`
	EventCancelled    bool        `json:"eventCancelled,omitempty"`
	EventCompplted    bool        `json:"eventCompleted,omitempty"`
	CustomProperties  interface{} `json:"customProperties,omitempty"`
	CreatedAt         *HsTime     `json:"createdAt,omitempty"`
	UpdatedAt         *HsTime     `json:"updatedAt,omitempty"`
	Id                *HsStr      `json:"id,omitempty"`
}

type CreateEventSingleResponse struct {
	Status      string           `json:"status"`
	StartedAt   *time.Time       `json:"startedAt"`
	CompletedAt *time.Time       `json:"completedAt"`
	Results     []MarketingEvent `json:"results"`
}

func (s *MarketingEventServiceOp) CreateUpdateEvent(event *MarketingEvent) (*CreateEventSingleResponse, error) {
	inputs := struct {
		Input []MarketingEvent `json:"input"`
	}{
		Input: []MarketingEvent{*event},
	}
	externalEventId := event.ExternalEventId
	resource := &CreateEventSingleResponse{}
	if err := s.client.Post(fmt.Sprintf("%s/upsert/%s", s.marketingEventPath, externalEventId), inputs, resource); err != nil {
		return nil, err
	}
	return resource, nil
}
