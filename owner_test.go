package hubspot_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/theovassiliou/go-hubspot"
)

func TestOwnerServiceOp_Get(t *testing.T) {
	type CustomFields struct {
		hubspot.Owner
		CustomName string          `json:"custom_name,omitempty"`
		CustomDate *hubspot.HsTime `json:"custom_date,omitempty"`
	}

	type fields struct {
		ownerPath string
		client    *hubspot.Client
	}
	type args struct {
		ownerID string
		owner   interface{}
		option  *hubspot.RequestQueryOption
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *hubspot.ResponseResource
		wantV3  *hubspot.ResponseResourceAll
		wantErr error
	}{
		{
			name: "Successfully get an owner",
			fields: fields{
				ownerPath: hubspot.ExportOwnerBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusOK,
					Header: http.Header{},
					Body:   []byte(`{"results":{"id":"512","email":"owner@hubspot.com","firstName":"John","lastName":"Doe","userId":"1234","createdAt":"2019-12-07T16:50:06.678Z","updatedAt":"2019-10-30T03:30:17.883Z","archived":false}`),
				}),
			},
			args: args{
				ownerID: "512",
				owner:   &hubspot.Owner{},
			},
			wantV3: &hubspot.ResponseResourceAll{
				Results: []interface{}{&hubspot.Owner{
					ID:        hubspot.NewString("512"),
					Email:     hubspot.NewString("owner@hubspot.com"),
					FirstName: hubspot.NewString("John"),
					LastName:  hubspot.NewString("Doe"),
					//UserID:    hubspot.NewString("1234"),
					UserID:    1234,
					CreatedAt: &createdAt,
					UpdatedAt: &updatedAt,
					Archived:  true,
				},
				},
			},
			wantErr: nil,
		},
		{
			name: "Received invalid request",
			fields: fields{
				ownerPath: hubspot.ExportOwnerBasePath,
				client: hubspot.NewMockClient(&hubspot.MockConfig{
					Status: http.StatusBadRequest,
					Header: http.Header{},
					Body:   []byte(`{"message": "Invalid input (details will vary based on the error)","correlationId": "aeb5f871-7f07-4993-9211-075dc63e7cbf","category": "VALIDATION_ERROR","links": {"knowledge-base": "https://www.hubspot.com/products/service/knowledge-base"}}`),
				}),
			},
			args: args{
				ownerID: "owner001",
				owner:   &hubspot.Owner{},
			},
			want: nil,
			wantErr: &hubspot.APIError{
				HTTPStatusCode: http.StatusBadRequest,
				Message:        "Invalid input (details will vary based on the error)",
				CorrelationID:  "aeb5f871-7f07-4993-9211-075dc63e7cbf",
				Category:       "VALIDATION_ERROR",
				Links: hubspot.ErrLinks{
					KnowledgeBase: "https://www.hubspot.com/products/service/knowledge-base",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.CRM.Owner.Get(tt.args.ownerID, tt.args.owner, tt.args.option)
			fmt.Printf("Result: %+v\n", got)
			fmt.Printf("Error: %+v\n", err)
			if !reflect.DeepEqual(tt.wantErr, err) {
				t.Errorf("Get() error mismatch: want %s got %s", tt.wantErr, err)
				return
			}
			if tt.want != nil {
				if diff := cmp.Diff(tt.want, got, cmpTimeOption); diff != "" {
					t.Errorf("Get() response mismatch (-want +got):%s", diff)
				}
			} else {
				if diff := cmp.Diff(tt.wantV3, got, cmpTimeOption); diff != "" {
					t.Errorf("Get() response all mismatch (-want +got):%s", diff)
				}
			}
		})
	}
}
