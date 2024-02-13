package di

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-bff/internal/config"
	mapgrpcclient "github.com/ZecretBone/ips-bff/internal/repository/MapGRPCClient"
	rssiclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/stat"
	rssistatclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/stat"
	"github.com/ZecretBone/ips-bff/internal/repository/cache"
	"github.com/ZecretBone/ips-bff/internal/repository/minio"
	"github.com/ZecretBone/ips-bff/internal/repository/mongodb"
	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(
	minio.ProvideMinioService,
	mongodb.ProvideMongoDBService,
	cache.ProvideCacheService,
)

var ProviderSet = wire.NewSet(
	mapgrpcclient.ProvideMapService,
	rssistatclient.ProvideRSSIService,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
	config.ProvideGRPCServiceConfig,
)

type Locator struct {
	MongoDBConn     wiremongo.Connection
	MinioXConn      wireminio.Connection
	RSSIGRPCService rssiclient.Service
	CacheService    cache.Service
}
