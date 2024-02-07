package di

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-bff/internal/config"
	"github.com/ZecretBone/ips-bff/internal/repository/cache"
	mapgrpcclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/mapclient"
	"github.com/ZecretBone/ips-bff/internal/repository/grpc/presenceclient"
	rssiclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/rssiclient"
	"github.com/ZecretBone/ips-bff/internal/repository/grpc/statclient"
	"github.com/ZecretBone/ips-bff/internal/repository/minio"
	"github.com/ZecretBone/ips-bff/internal/repository/mongodb"
	"github.com/ZecretBone/ips-bff/internal/services/realtime"
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
	statclient.ProvideStatService,
	rssiclient.ProvideRSSIService,
	presenceclient.ProvidePresenceServiceClient,
	mapgrpcclient.ProvideMapService,
	realtime.ProvideRealtimeService,
)

var ConfigSet = wire.NewSet(
	config.ProvideMongoxConfig,
	config.ProvideMinioXConfig,
	config.ProvideCacheConfig,
	config.ProvideRSSIConfig,
	config.ProvideGRPCServiceConfig,
	config.ProvideRealtimeServiceConfig,
)

type Locator struct {
	MongoDBConn     wiremongo.Connection
	MinioXConn      wireminio.Connection
	RSSIService     rssi.Service
	RSSIGRPCService rssiclient.Service
	CacheService    cache.Service
}
