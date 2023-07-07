package hubspot

import (
	"fmt"
)

const (
	companyBasePath = "companies"
)

// Owner is an interface of owner endpoints of the HubSpot API.
// HubSpot uses owners to assign specific users to contacts, companies, deals, tickets, or engagements.
// Reference: https://developers.hubspot.com/docs/api/crm/owners
type CompanyService interface {
	Get(companyID string, owner interface{}, option *RequestQueryOption) (*ResponseResource, error)
	GetAll(company interface{}, option *RequestQueryOption) (*ResponseResourceMulti, error)
	Search(company interface{}, option *RequestSearchOption) (*ResponseResourceMulti, error)
	Create(company interface{}) (*ResponseResource, error)
	Delete(companyID string) error
}

// OwnerServiceOp handles communication with the product related methods of the HubSpot API.
type CompanyServiceOp struct {
	companyPath string
	client      *Client
}

type Company struct {
	ID                       *HsStr  `json:"id,omitempty"`
	Name                     *HsStr  `json:"name,omitempty"`
	Industry                 *HsStr  `json:"industry,omitempty"`
	Domain                   *HsStr  `json:"domain,omitempty"`
	Phone                    *HsStr  `json:"phone,omitempty"`
	City                     *HsStr  `json:"city,omitempty"`
	State                    *HsStr  `json:"state,omitempty"`
	HsCreateDate             *HsTime `json:"hs_createdate,omitempty"`
	HsLastModifiedDate       *HsTime `json:"hs_lastmodifieddate,omitempty"`
	HsObjectID               *HsStr  `json:"hs_object_id,omitempty"`
	HubspotOwnerAssignedDate *HsTime `json:"hubspot_owner_assigneddate,omitempty"`
	HubspotOwnerID           *HsStr  `json:"hubspot_owner_id,omitempty"`
}

var defaultCompanyFields = []string{
	"id",
	"name",
	"industry",
	"domain",
	"phone",
	"city",
	"state",
	"hs_createdate",
	"hs_lastmodifieddate",
	"hs_object_id",
	"hubspot_owner_assigneddate",
	"hubspot_owner_id",
}

// Get gets a company.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *CompanyServiceOp) Get(companyID string, company interface{}, option *RequestQueryOption) (*ResponseResource, error) {
	resource := &ResponseResource{Properties: company}
	path := s.companyPath + "/" + companyID
	if len(option.Associations) != 0 {
		path += "/associations/" + option.Associations[0]
		result := []interface{}{}
		result = append(result, &AssociationResult{})
		resource = &ResponseResource{Results: result}
	}
	fmt.Printf("Result path %s\n", path)
	//if err := s.client.Get(path, company, option.setupProperties(defaultCompanyFields)); err != nil {
	if err := s.client.Get(path, resource, option.setupProperties(defaultCompanyFields)); err != nil {
		return nil, err
	}
	return resource, nil
}

// Create creates a new company.
// In order to bind the created content, a structure must be specified as an argument.
// When using custom fields, please embed hubspot.Contact in your own structure.
func (s *CompanyServiceOp) Create(company interface{}) (*ResponseResource, error) {
	req := &RequestPayload{Properties: company}
	resource := &ResponseResource{Properties: company}
	if err := s.client.Post(s.companyPath, req, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Get gets all companies.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *CompanyServiceOp) GetAll(company interface{}, option *RequestQueryOption) (*ResponseResourceMulti, error) {
	//result := []interface{}{}
	//result = append(result, company)
	//resource := &ResponseResourceAll{Results: result}
	resource := &ResponseResourceMulti{}
	if len(option.Properties) == 0 {
		option = option.setupProperties(defaultCompanyFields)
	}
	//if err := s.client.Get(s.companyPath, resource, option.setupProperties(defaultCompanyFields)); err != nil {
	if err := s.client.Get(s.companyPath, resource, option); err != nil {
		return nil, err
	}
	return resource, nil
}

// Search finds a company.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *CompanyServiceOp) Search(company interface{}, option *RequestSearchOption) (*ResponseResourceMulti, error) {
	resources := []ResponseResource{}
	resources = append(resources, ResponseResource{Properties: company})
	resource := &ResponseResourceMulti{Results: resources}
	if err := s.client.Post(s.companyPath+"/search", option, resource); err != nil {
		return nil, err
	}
	return resource, nil
}

// Delete deletes a company.
// A HubSpot internal Company ID must be specified.
func (s *CompanyServiceOp) Delete(companyID string) error {
	path := s.companyPath + "/" + companyID
	if err := s.client.Delete(path, nil); err != nil {
		return err
	}
	return nil
}
