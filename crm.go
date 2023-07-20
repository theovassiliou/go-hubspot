package hubspot

import "fmt"

const (
	crmBasePath = "crm"

	objectsBasePath = "objects"
)

type CRM struct {
	Company    CompanyService
	Contact    ContactService
	Deal       DealService
	Owner      OwnerService
	Pipeline   PipelineService
	Imports    CrmImportsService
	Schemas    CrmSchemasService
	Properties CrmPropertiesService
	Tickets    CrmTicketsServivce
}

func newCRM(c *Client) *CRM {
	crmPath := fmt.Sprintf("%s/%s", crmBasePath, c.apiVersion)
	return &CRM{
		Company: &CompanyServiceOp{
			companyPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, companyBasePath),
			client:      c,
		},
		Contact: &ContactServiceOp{
			contactPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, contactBasePath),
			client:      c,
		},
		Deal: &DealServiceOp{
			dealPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, dealBasePath),
			client:   c,
		},
		Owner: &OwnerServiceOp{
			ownerPath: fmt.Sprintf("%s/%s", crmPath, ownerBasePath),
			client:    c,
		},
		Pipeline: &PipelineServiceOp{
			pipelinePath: fmt.Sprintf("%s/%s", crmPath, pipelineBasePath),
			client:       c,
		},
		Imports: &CrmImportsServiceOp{
			crmImportsPath: fmt.Sprintf("%s/%s", crmPath, crmImportsBasePath),
			client:         c,
		},
		Schemas: &CrmSchemasServiceOp{
			crmSchemasPath: fmt.Sprintf("%s/%s", crmPath, crmSchemasPath),
			client:         c,
		},
		Properties: &CrmPropertiesServiceOp{
			crmPropertiesPath: fmt.Sprintf("%s/%s", crmPath, crmPropertiesPath),
			client:            c,
		},
		Tickets: &CrmTicketsServivceOp{
			crmTicketsPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, crmTicketsBasePath),
			client:         c,
		},
	}
}
