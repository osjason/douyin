package store

// Factory Store层的调用金字塔最顶端，工厂模式
type Factory interface {
	BearerAuthorization() BearerAuthStoreInterface
	BasicAuthorization() BasicAuthStoreInterface
}
