package main

//go:generate mockgen -package=main -destination=iredis_resolver_mock.go -source=iredis_resolver.go
type RedisResolverInterface interface {
	Clauses(action string) IRedisClientInterface
}
