package pbsdk

import "github.com/polarisbase/polaris-sdk/pboptions"

var Options Option = Option{
	ApiService: pboptions.ApiServiceOptions,
}

type Option struct {
	ApiService pboptions.ApiServiceOption
}
