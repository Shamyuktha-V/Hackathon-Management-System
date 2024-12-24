package graph

type Resolver struct {
	QueryResolver    queryResolver
	MutationResolver mutationResolver
}
