package context

var (
	Context *context
)

type context struct {
	ContextMap map[string]string
}

func init() {
	Context = &context{ContextMap: make(map[string]string)}
}
