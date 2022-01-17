package kugo

type BaseRequest struct {
	KubeContext string
	Namespace string
	ExtraArgs []string
}

type GetRequest struct {
	BaseRequest

	Kind string
	FieldSelectors []string
	Label string
	Output string
}

type ExecRequest struct {
	BaseRequest

	Command string
	Target string
	Container string
}
