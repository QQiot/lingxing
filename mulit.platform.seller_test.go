package lingxing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_multiPlatformSellerService_All(t *testing.T) {
	type args struct {
		params MultiPlatformSellersQueryParams
	}
	tests := []struct {
		name           string
		args           args
		wantItems      []MultiPlatformSeller
		wantNextOffset int
		wantIsLastPage bool
		// wantErr        assert.ErrorAssertionFunc
	}{
		{"test1", args{MultiPlatformSellersQueryParams{}}, []MultiPlatformSeller{}, 10, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItems, gotNextOffset, gotIsLastPage, err := lingXingClient.Services.MultiPlatform.Seller.All(tt.args.params)
			_= err
			// if !tt.wantErr(t, err, fmt.Sprintf("All(%v)", tt.args.params)) {
			// 	return
			// }
			assert.Equalf(t, tt.wantItems, gotItems, "All(%v)", tt.args.params)
			assert.Equalf(t, tt.wantNextOffset, gotNextOffset, "All(%v)", tt.args.params)
			assert.Equalf(t, tt.wantIsLastPage, gotIsLastPage, "All(%v)", tt.args.params)
		})
	}
}
