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
	ID                                      *HsStr  `json:"id,omitempty"`
	Name                                    *HsStr  `json:"name,omitempty"`
	Industry                                *HsStr  `json:"industry,omitempty"`
	Domain                                  *HsStr  `json:"domain,omitempty"`
	Phone                                   *HsStr  `json:"phone,omitempty"`
	City                                    *HsStr  `json:"city,omitempty"`
	State                                   *HsStr  `json:"state,omitempty"`
	HsCreateDate                            *HsTime `json:"hs_createdate,omitempty"`
	HsLastModifiedDate                      *HsTime `json:"hs_lastmodifieddate,omitempty"`
	HsObjectID                              *HsStr  `json:"hs_object_id,omitempty"`
	HubspotOwnerAssignedDate                *HsTime `json:"hubspot_owner_assigneddate,omitempty"`
	HubspotOwnerID                          *HsStr  `json:"hubspot_owner_id,omitempty"`
	AboutUs                                 *HsStr  `json:"about_us,omitempty"`
	DaysToClose                             *HsStr  `json:"days_to_close,omitempty"`
	HsAnalyticsNumPageViews                 *HsStr  `json:"hs_analytics_num_page_views,omitempty"`
	HsAnalyticsNumVisits                    *HsStr  `json:"hs_analytics_num_visits,omitempty"`
	HsCreatedByUserID                       *HsStr  `json:"hs_created_by_user_id,omitempty"`
	NumContactedNotes                       *HsStr  `json:"num_contacted_notes,omitempty"`
	AnnualRevenue                           *HsStr  `json:"annualrevenue,omitempty"`
	HsNumChildCompanies                     *HsStr  `json:"hs_num_child_companies,omitempty"`
	HsNumOpenDeals                          *HsStr  `json:"hs_num_open_deals,omitempty"`
	HsParentCompanyID                       *HsStr  `json:"hs_parent_company_id,omitempty"`
	HsTotalDealValue                        *HsStr  `json:"hs_total_deal_value,omitempty"`
	NumAssociatedContacts                   *HsStr  `json:"num_associated_contacts,omitempty"`
	NumAssociatedDeals                      *HsStr  `json:"num_associated_deals,omitempty"`
	NumberOfEmployees                       *HsStr  `json:"numberofemployees,omitempty"`
	RecentDealAmount                        *HsStr  `json:"recent_deal_amount,omitempty"`
	TotalRevenue                            *HsStr  `json:"total_revenue,omitempty"`
	NumConversionEvents                     *HsStr  `json:"num_conversion_events,omitempty"`
	FacebookFans                            *HsStr  `json:"facebookfans,omitempty"`
	TwitterFollowers                        *HsStr  `json:"twitterfollowers,omitempty"`
	HsNumBlockers                           *HsStr  `json:"hs_num_blockers,omitempty"`
	HsNumContactsWithBuyingRoles            *HsStr  `json:"hs_num_contacts_with_buying_roles,omitempty"`
	HsNumDecisionMakers                     *HsStr  `json:"hs_num_decision_makers,omitempty"`
	IsPublic                                *HsBool `json:"is_public,omitempty"`
	HsIsTargetAccount                       *HsBool `json:"hs_is_target_account,omitempty"`
	HsAnalyticsFirstTimestamp               *HsTime `json:"hs_analytics_first_timestamp,omitempty"`
	HsAnalyticsFirstVisitTimestamp          *HsTime `json:"hs_analytics_first_visit_timestamp,omitempty"`
	HsAnalyticsLastTimestamp                *HsTime `json:"hs_analytics_last_timestamp,omitempty"`
	HsAnalyticsLastVisitTimestamp           *HsTime `json:"hs_analytics_last_visit_timestamp,omitempty"`
	EngagementsLastMeetingBooked            *HsTime `json:"engagements_last_meeting_booked,omitempty"`
	FirstContactCreatedate                  *HsTime `json:"first_contact_createdate,omitempty"`
	FirstDealCreatedDate                    *HsTime `json:"first_deal_created_date,omitempty"`
	HsLastBookedMeetingDate                 *HsTime `json:"hs_last_booked_meeting_date,omitempty"`
	HsLastLoggedCallDate                    *HsTime `json:"hs_last_logged_call_date,omitempty"`
	HsLastOpenTaskDate                      *HsTime `json:"hs_last_open_task_date,omitempty"`
	HsLastSalesActivityTimestamp            *HsTime `json:"hs_last_sales_activity_timestamp,omitempty"`
	NotesLastContacted                      *HsTime `json:"notes_last_contacted,omitempty"`
	NotesLastUpdated                        *HsTime `json:"notes_last_updated,omitempty"`
	NotesNextActivityDate                   *HsTime `json:"notes_next_activity_date,omitempty"`
	CloseDate                               *HsTime `json:"closedate,omitempty"`
	CreateDate                              *HsTime `json:"createdate,omitempty"`
	HsAnalyticsLatestSourceTimestamp        *HsTime `json:"hs_analytics_latest_source_timestamp,omitempty"`
	RecentDealCloseDate                     *HsTime `json:"recent_deal_close_date,omitempty"`
	FirstConversionDate                     *HsTime `json:"first_conversion_date,omitempty"`
	RecentConversionDate                    *HsTime `json:"recent_conversion_date,omitempty"`
	HsAnalyticsSource                       *HsStr  `json:"hs_analytics_source,omitempty"`
	HsAnalyticsLatestSource                 *HsStr  `json:"hs_analytics_latest_source,omitempty"`
	HsMergedObjectIds                       *HsStr  `json:"hs_merged_object_ids,omitempty"`
	HsLeadStatus                            *HsStr  `json:"hs_lead_status,omitempty"`
	HsObjectSourceLabel                     *HsStr  `json:"hs_object_source_label,omitempty"`
	HubspotTeamId                           *HsStr  `json:"hubspot_team_id,omitempty"`
	LifecycleStage                          *HsStr  `json:"lifecyclestage,omitempty"`
	Type                                    *HsStr  `json:"type,omitempty"`
	WebTechnologies                         *HsStr  `json:"web_technologies,omitempty"`
	HsIdealCustomerProfile                  *HsStr  `json:"hs_ideal_customer_profile,omitempty"`
	HsAnalyticsFirstTouchConvertingCampaign *HsStr  `json:"hs_analytics_first_touch_converting_campaign,omitempty"`
	HsAnalyticsLastTouchConvertingCampaign  *HsStr  `json:"hs_analytics_last_touch_converting_campaign,omitempty"`
	HsAnalyticsSourceData1                  *HsStr  `json:"hs_analytics_source_data_1,omitempty"`
	HsAnalyticsSourceData2                  *HsStr  `json:"hs_analytics_source_data_2,omitempty"`
	EngagementsLastMeetingBookedCampaign    *HsStr  `json:"engagements_last_meeting_booked_campaign,omitempty"`
	EngagementsLastMeetingBookedMedium      *HsStr  `json:"engagements_last_meeting_booked_medium,omitempty"`
	EngagementsLastMeetingBookedSource      *HsStr  `json:"engagements_last_meeting_booked_source,omitempty"`
	HsAnalyticsLatestSourceData1            *HsStr  `json:"hs_analytics_latest_source_data_1,omitempty"`
	HsAnalyticsLatestSourceData2            *HsStr  `json:"hs_analytics_latest_source_data_2,omitempty"`
	Address                                 *HsStr  `json:"address,omitempty"`
	Address2                                *HsStr  `json:"address2,omitempty"`
	Country                                 *HsStr  `json:"country,omitempty"`
	Description                             *HsStr  `json:"description,omitempty"`
	FoundedYear                             *HsStr  `json:"founded_year,omitempty"`
	HsObjectSourceDetail1                   *HsStr  `json:"hs_object_source_detail_1,omitempty"`
	HsObjectSourceDetail2                   *HsStr  `json:"hs_object_source_detail_2,omitempty"`
	HsObjectSourceDetail3                   *HsStr  `json:"hs_object_source_detail_3,omitempty"`
	Timezone                                *HsStr  `json:"timezone,omitempty"`
	TotalMoneyRaised                        *HsStr  `json:"total_money_raised,omitempty"`
	Website                                 *HsStr  `json:"website,omitempty"`
	Zip                                     *HsStr  `json:"zip,omitempty"`
	FirstConversionEventName                *HsStr  `json:"first_conversion_event_name,omitempty"`
	RecentConversionEventName               *HsStr  `json:"recent_conversion_event_name,omitempty"`
	FacebookCompanyPage                     *HsStr  `json:"facebook_company_page,omitempty"`
	GoogleplusPage                          *HsStr  `json:"googleplus_page,omitempty"`
	LinkedinCompanyPage                     *HsStr  `json:"linkedin_company_page,omitempty"`
	LinkedinBio                             *HsStr  `json:"linkedinbio,omitempty"`
	TwitterBio                              *HsStr  `json:"twitterbio,omitempty"`
	TwitterHandle                           *HsStr  `json:"twitterhandle,omitempty"`
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
	"about_us",
	"days_to_close",
	"hs_analytics_num_page_views",
	"hs_analytics_num_visits",
	"hs_created_by_user_id",
	"hs_updated_by_user_id",
	"num_contacted_notes",
	"annualrevenue",
	"hs_num_child_companies",
	"hs_num_open_deals",
	"hs_parent_company_id",
	"hs_total_deal_value",
	"num_associated_contacts",
	"num_associated_deals",
	"numberofemployees",
	"recent_deal_amount",
	"total_revenue",
	"num_conversion_events",
	"facebookfans",
	"twitterfollowers",
	"hs_num_blockers",
	"hs_num_contacts_with_buying_roles",
	"hs_num_decision_makers",
	"is_public",
	"hs_is_target_account",
	"hs_analytics_first_timestamp",
	"hs_analytics_first_visit_timestamp",
	"hs_analytics_last_timestamp",
	"hs_analytics_last_visit_timestamp",
	"engagements_last_meeting_booked",
	"first_contact_createdate",
	"first_deal_created_date",
	"hs_last_booked_meeting_date",
	"hs_last_logged_call_date",
	"hs_last_open_task_date",
	"hs_last_sales_activity_timestamp",
	"notes_last_contacted",
	"notes_last_updated",
	"notes_next_activity_date",
	"closedate",
	"createdate",
	"hs_analytics_latest_source_timestamp",
	"recent_deal_close_date",
	"first_conversion_date",
	"recent_conversion_date",
	"hs_analytics_source",
	"hs_analytics_latest_source",
	"hs_merged_object_ids",
	"hs_lead_status",
	"hs_object_source_label",
	"hubspot_team_id",
	"lifecyclestage",
	"type",
	"web_technologies",
	"hs_ideal_customer_profile",
	"hs_analytics_first_touch_converting_campaign",
	"hs_analytics_last_touch_converting_campaign",
	"hs_analytics_source_data_1",
	"hs_analytics_source_data_2",
	"engagements_last_meeting_booked_campaign",
	"engagements_last_meeting_booked_medium",
	"engagements_last_meeting_booked_source",
	"hs_analytics_latest_source_data_1",
	"hs_analytics_latest_source_data_2",
	"address",
	"address2",
	"country",
	"description",
	"founded_year",
	"hs_object_source_detail_1",
	"hs_object_source_detail_2",
	"hs_object_source_detail_3",
	"timezone",
	"total_money_raised",
	"website",
	"zip",
	"first_conversion_event_name",
	"recent_conversion_event_name",
	"facebook_company_page",
	"googleplus_page",
	"linkedin_company_page",
	"linkedinbio",
	"twitterbio",
	"twitterhandle"}

var AllCompanyFields = []string{
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
	"about_us",
	"days_to_close",
	"hs_analytics_num_page_views",
	"hs_analytics_num_visits",
	"hs_created_by_user_id",
	"hs_updated_by_user_id",
	"num_contacted_notes",
	"annualrevenue",
	"hs_num_child_companies",
	"hs_num_open_deals",
	"hs_parent_company_id",
	"hs_total_deal_value",
	"num_associated_contacts",
	"num_associated_deals",
	"numberofemployees",
	"recent_deal_amount",
	"total_revenue",
	"num_conversion_events",
	"facebookfans",
	"twitterfollowers",
	"hs_num_blockers",
	"hs_num_contacts_with_buying_roles",
	"hs_num_decision_makers",
	"is_public",
	"hs_is_target_account",
	"hs_analytics_first_timestamp",
	"hs_analytics_first_visit_timestamp",
	"hs_analytics_last_timestamp",
	"hs_analytics_last_visit_timestamp",
	"engagements_last_meeting_booked",
	"first_contact_createdate",
	"first_deal_created_date",
	"hs_last_booked_meeting_date",
	"hs_last_logged_call_date",
	"hs_last_open_task_date",
	"hs_last_sales_activity_timestamp",
	"notes_last_contacted",
	"notes_last_updated",
	"notes_next_activity_date",
	"closedate",
	"createdate",
	"hs_analytics_latest_source_timestamp",
	"recent_deal_close_date",
	"first_conversion_date",
	"recent_conversion_date",
	"hs_analytics_source",
	"hs_analytics_latest_source",
	"hs_merged_object_ids",
	"hs_lead_status",
	"hs_object_source_label",
	"hubspot_team_id",
	"lifecyclestage",
	"type",
	"web_technologies",
	"hs_ideal_customer_profile",
	"hs_analytics_first_touch_converting_campaign",
	"hs_analytics_last_touch_converting_campaign",
	"hs_analytics_source_data_1",
	"hs_analytics_source_data_2",
	"engagements_last_meeting_booked_campaign",
	"engagements_last_meeting_booked_medium",
	"engagements_last_meeting_booked_source",
	"hs_analytics_latest_source_data_1",
	"hs_analytics_latest_source_data_2",
	"address",
	"address2",
	"country",
	"description",
	"founded_year",
	"hs_object_source_detail_1",
	"hs_object_source_detail_2",
	"hs_object_source_detail_3",
	"timezone",
	"total_money_raised",
	"website",
	"zip",
	"first_conversion_event_name",
	"recent_conversion_event_name",
	"facebook_company_page",
	"googleplus_page",
	"linkedin_company_page",
	"linkedinbio",
	"twitterbio",
	"twitterhandle",
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
