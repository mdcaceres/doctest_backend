package execution

type IExecutable interface {
	Execute() IExecutable
}
