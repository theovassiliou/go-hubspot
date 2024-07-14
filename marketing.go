package hubspot

import "fmt"

const (
	marketingBasePath = "marketing"
)

type Marketing struct {
	Email         MarketingEmailService
	Transactional TransactionalService
	Event         MarketingEventService
}

func newMarketing(c *Client) *Marketing {
	return &Marketing{
		Email: NewMarketingEmail(c),
		Transactional: &TransactionalServiceOp{
			client:            c,
			transactionalPath: fmt.Sprintf("%s/%s/%s", marketingBasePath, c.apiVersion, transactionalBasePath),
		},
		Event: &MarketingEventServiceOp{
			marketingEventPath: fmt.Sprintf("%s/%s/%s", marketingBasePath, c.apiVersion, marketingEventBasePath),
			client:             c,
		},
	}
}
