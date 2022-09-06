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
