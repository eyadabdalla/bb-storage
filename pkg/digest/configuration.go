package digest

import (
	"github.com/buildbarn/bb-storage/pkg/clock"
	"github.com/buildbarn/bb-storage/pkg/eviction"
	pb "github.com/buildbarn/bb-storage/pkg/proto/configuration/digest"
	"github.com/buildbarn/bb-storage/pkg/util"
	"github.com/golang/protobuf/ptypes"
)

// NewExistenceCacheFromConfiguration is identical to
// NewExistenceCache(), except that it takes a specification for the
// object to be created from a configuration file message.
func NewExistenceCacheFromConfiguration(configuration *pb.ExistenceCacheConfiguration, keyFormat KeyFormat, name string) (*ExistenceCache, error) {
	cacheDuration, err := ptypes.Duration(configuration.CacheDuration)
	if err != nil {
		return nil, util.StatusWrap(err, "Cache duration")
	}
	evictionSet, err := eviction.NewSetFromConfiguration(configuration.CacheReplacementPolicy)
	if err != nil {
		return nil, util.StatusWrap(err, "Cache replacement policy")
	}
	return NewExistenceCache(
		clock.SystemClock,
		keyFormat,
		int(configuration.CacheSize),
		cacheDuration,
		eviction.NewMetricsSet(evictionSet, name)), nil
}
