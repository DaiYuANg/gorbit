package cache

// NewCacheModule 创建 fx.Module
//func NewCacheModule(opts ...CacheOption) fx.Option {
//	options := defaultCacheOptions()
//	for _, o := range opts {
//		o(options)
//	}
//
//	return fx.Module("cache_module",
//		fx.Provide(func() gocache.Cache {
//			cacheInstance := gocache.New(options.Store)
//			cacheInstance.SetTTL(options.DefaultTTL)
//			if options.Serializer != nil {
//				cacheInstance.SetSerializer(options.Serializer)
//			}
//			return cacheInstance
//		}),
//	)
//}
