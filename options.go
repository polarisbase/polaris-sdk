package pbsdk

import "github.com/polarisbase/polaris-sdk/pboptions"

var Options option = option{
	ApiService: pboptions.ApiServiceOptions,
	PostgresService:  pboptions.PostgresServiceOptions,
}

type option struct {
	ApiService pboptions.ApiServiceOption
	PostgresService  pboptions.PostgresServiceOption
}
