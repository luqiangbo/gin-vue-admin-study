package tables

type InitDbFunc interface {
	Init() (err error)
}
