package common

type OptionService interface {
	Option
}

type OptionServiceApi interface {
	OptionService
}

type OptionServicePostgres interface {
	OptionService
}

type OptionServiceAuthentication interface {
	OptionService
}

type OptionServicePointmass interface {
	OptionService
}

type OptionServiceWebHtml interface {
	OptionService
}

type OptionServiceAuthWebHtml interface {
	OptionService
}
