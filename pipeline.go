package hubspot

const (
	pipelineBasePath = "pipelines"
)

// Pipeline is an interface of pipeline endpoints of the HubSpot API.
// HubSpot uses owners to assign specific users to contacts, companies, deals, tickets, or engagements.
// Reference: https://developers.hubspot.com/docs/api/crm/pipeline
type PipelineService interface {
	Get(pipelineID string, pipeline interface{}, option *RequestQueryOption) (ResponseResourceNonObject, error)
	GetAll(pipeline interface{}, option *RequestQueryOption) (*ResponseResourceAll, error)
}

// PipelineServiceOp handles communication with the product related methods of the HubSpot API.
type PipelineServiceOp struct {
	pipelinePath string
	client       *Client
}

type Pipeline struct {
	Label        *HsStr   `json:"label,omitempty"`
	DisplayOrder *HsInt   `json:"displayOrder,omitempty"`
	CreatedAt    *HsTime  `json:"createdAt,omitempty"`
	UpdatedAt    *HsTime  `json:"updatedAt,omitempty"`
	Archived     *HsBool  `json:"archived,omitempty"`
	Stages       *[]Stage `json:"stages,omitempty"`
	PipelineID   *HsStr   `json:"id,omitempty"`
}

type Stage struct {
	Label        *HsStr    `json:"label,omitempty"`
	DisplayOrder *HsInt    `json:"displayOrder,omitempty"`
	CreatedAt    *HsTime   `json:"createdAt,omitempty"`
	UpdatedAt    *HsTime   `json:"updatedAt,omitempty"`
	Archived     *HsBool   `json:"archived,omitempty"`
	Metadata     *Metadata `json:"metadata,omitempty"`
	StageID      *HsStr    `json:"id,omitempty"`
}

type Metadata struct {
	IsClosed    *HsBool `json:"isClosed,omitempty"`
	Probability *HsStr  `json:"probability,omitempty"`
}

var defaultPipelineFields = []string{
	"label",
	"displayOrder",
	"createdAt",
	"updatedAt",
	"archived",
	"stages",
}

// Get gets a pipeline.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *PipelineServiceOp) Get(pipelineID string, pipeline interface{}, option *RequestQueryOption) (ResponseResourceNonObject, error) {
	//resource := &ResponseResource{Properties: pipeline}
	//if err := s.client.Get(s.pipelinePath+"/deals/"+pipelineID, resource, option.setupProperties(defaultPipelineFields)); err != nil {
	if err := s.client.Get(s.pipelinePath+"/deals/"+pipelineID, pipeline, option.setupProperties(defaultPipelineFields)); err != nil {
		return nil, err
	}
	//return resource, nil
	return pipeline, nil
}

// GetAll gets all pipelines.
// In order to bind the get content, a structure must be specified as an argument.
// Also, if you want to gets a custom field, you need to specify the field name.
// If you specify a non-existent field, it will be ignored.
// e.g. &hubspot.RequestQueryOption{ Properties: []string{"custom_a", "custom_b"}}
func (s *PipelineServiceOp) GetAll(pipeline interface{}, option *RequestQueryOption) (*ResponseResourceAll, error) {
	result := []interface{}{}
	result = append(result, pipeline)
	resource := &ResponseResourceAll{Results: result}
	if err := s.client.Get(s.pipelinePath+"/deals", resource, option.setupProperties(defaultPipelineFields)); err != nil {
		return nil, err
	}
	return resource, nil
}
