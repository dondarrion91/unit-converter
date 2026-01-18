package strategy

type Strategy interface {
	execute(ctx *Context) float64
}
