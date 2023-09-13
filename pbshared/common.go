package pbshared

import "github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"

type Option interface{
	~*common.BaseOption
}
