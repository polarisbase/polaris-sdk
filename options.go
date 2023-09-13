package pbsdk

import "github.com/polarisbase/polaris-sdk/pboptions"

var Options option = option{
	ApiService: pboptions.ApiServiceOptions,
}

type option struct {
	ApiService pboptions.ApiServiceOption
}
