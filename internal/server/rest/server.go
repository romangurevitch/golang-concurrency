package rest

type Server interface {
	Run(addr ...string) error
}
