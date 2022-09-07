package lingxing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdService_Groups(t *testing.T) {
	params := AdGroupsQueryParams{
		SID:       2345,
		StartDate: "2022-09-01",
		EndDate:   "2022-09-01",
		Type:      1,
	}
	params.Limit = 1
	_, _, _, err := lingXingClient.Services.Ad.Groups(params)
	assert.Equal(t, nil, err, "error")
}

func Test_adService_ProductTargets(t *testing.T) {
	tests := []struct {
		name           string
		args           AdProductTargetsQueryParams
		wantItems      int
		wantNextOffset int
		wantIsLastPage bool
		wantErr        bool
	}{
		{"t1", AdProductTargetsQueryParams{}, 0, 0, false, true},
		{"t2", AdProductTargetsQueryParams{
			Paging:    Paging{Limit: 1},
			SID:       2345,
			StartDate: "2022-09-01",
			EndDate:   "2022-09-01",
		}, 1, 1, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItems, gotNextOffset, gotIsLastPage, err := lingXingClient.Services.Ad.ProductTargets(tt.args)
			assert.Equal(t, tt.wantErr, err != nil, "err ProductTargets(%#v)", tt.args)
			assert.Equalf(t, tt.wantItems, len(gotItems), "items ProductTargets(%#v)", tt.args)
			assert.Equalf(t, tt.wantNextOffset, gotNextOffset, "nextOffset ProductTargets(%#v)", tt.args)
			assert.Equalf(t, tt.wantIsLastPage, gotIsLastPage, "isLastPage ProductTargets(%#v)", tt.args)
		})
	}
}
