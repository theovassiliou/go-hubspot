package hubspot

const (
	ownerBasePath = "owners"
)

// Owner is an interface of owner endpoints of the HubSpot API.
// HubSpot uses owners to assign specific users to contacts, companies, deals, tickets, or engagements.
// Reference: https://developers.hubspot.com/docs/api/crm/owners
type OwnerService interface {
	Get(ownerID string, owner interface{}, option *RequestQueryOption) (ResponseResourceNonObject, error)
	GetAll(owner interface{}, option *RequestQueryOption) (*ResponseResourceAll, error)
}

// OwnerServiceOp handles communication with the product related methods of the HubSpot API.
type OwnerServiceOp struct {
	ownerPath string
	client    *Client
}

type Owner struct {
	ID        *HsStr  `json:"id,omitempty"`
	Email     *HsStr  `json:"email,omitempty"`
	FirstName *HsStr  `json:"firstName,omitempty"`
	LastName  *HsStr  `json:"lastName,omitempty"`
	UserID    int     `json:"userId,omitempty"`
	CreatedAt *HsTime `json:"createdAt,omitempty"`
	UpdatedAt *HsTime `json:"updatedAt,omitempty"`
	Archived  bool    `json:"archived,omitempty"`
}

var defaultOwnerFields = []string{
	"id",
	"email",
	"firstName",
	"lastName",
	"userId",
	"createdAt",
	"updatedAt",
	"archived",
}

// Get gets an owner.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *OwnerServiceOp) Get(ownerID string, owner interface{}, option *RequestQueryOption) (ResponseResourceNonObject, error) {
	if err := s.client.Get(s.ownerPath+"/"+ownerID, owner, option.setupProperties(defaultOwnerFields)); err != nil {
		return nil, err
	}
	return owner, nil
}

// Get gets all owners.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *OwnerServiceOp) GetAll(owner interface{}, option *RequestQueryOption) (*ResponseResourceAll, error) {
	result := []interface{}{}
	result = append(result, owner)
	resource := &ResponseResourceAll{Results: result}
	if err := s.client.Get(s.ownerPath, resource, option.setupProperties(defaultOwnerFields)); err != nil {
		return nil, err
	}
	return resource, nil
}
