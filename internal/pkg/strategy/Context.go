package strategy

type ContextData struct {
	Value int
	From  string
	To    string
}

type Context struct {
	str  Strategy
	data ContextData
}

func InitContext(str Strategy, data ContextData) *Context {
	return &Context{
		str:  str,
		data: data,
	}
}

func (ctx *Context) SetStrategy(str Strategy) {
	ctx.str = str
}

func (ctx *Context) Execute() float64 {
	return float64(ctx.str.execute(ctx))
}
