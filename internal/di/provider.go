package di

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-bff/internal/config"
	rssiclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/rssi"
	rssistatclient "github.com/ZecretBone/ips-bff/internal/repository/RSSIClient/stat"
	"github.com/ZecretBone/ips-bff/internal/repository/cache"
	"github.com/ZecretBone/ips-bff/internal/repository/minio"
	"github.com/ZecretBone/ips-bff/internal/repository/mongodb"
	"github.com/ZecretBone/ips-bff/internal/services/rssi"
	"github.com/google/wire"
)

var DatabaseSet = wire.NewSet(
	minio.ProvideMinioService,
	mongodb.ProvideMongoDBService,
	cache.ProvideCacheService,
)

var ProviderSet = wire.NewSet(
	rssi.ProvideRSSIService,
	rssiclient.ProvideRSSIService,
	rssistatclient.ProvideRSSIService,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
	config.ProvideRSSIConfig,
	config.ProvideGRPCServiceConfig,
)

type Locator struct {
	MongoDBConn     wiremongo.Connection
	MinioXConn      wireminio.Connection
	RSSIService     rssi.Service
	RSSIGRPCService rssiclient.Service
	CacheService    cache.Service
}
