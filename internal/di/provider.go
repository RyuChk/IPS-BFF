package di

import (
	wireminio "git.cie.com/ips/wire-provider/minio"
	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/ZecretBone/ips-bff/internal/config"
	"github.com/ZecretBone/ips-bff/internal/repository/cache"
	datacollectionclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/dataCollectionClient"
	mapgrpcclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/mapClient"
	usertrackingclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/userTrackingClient"
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
	datacollectionclient.ProvideDataCollectionGRPCClient,
	usertrackingclient.ProvideUserTrackingService,
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
	RSSIGRPCService datacollectionclient.Service
	CacheService    cache.Service
}
