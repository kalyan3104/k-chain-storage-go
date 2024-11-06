package testscommon

import (
	"github.com/kalyan3104/k-chain-storage-go/common"
	"github.com/kalyan3104/k-chain-storage-go/leveldb"
	"github.com/kalyan3104/k-chain-storage-go/memorydb"
	"github.com/kalyan3104/k-chain-storage-go/storageUnit"
	"github.com/kalyan3104/k-chain-storage-go/types"
)

type persisterFactoryHandlerMock struct {
	dbType            storageUnit.DBType
	batchDelaySeconds int
	maxBatchSize      int
	maxOpenFiles      int
}

// NewPersisterFactoryHandlerMock -
func NewPersisterFactoryHandlerMock(dbType storageUnit.DBType, batchDelaySeconds int, maxBatchSize int, maxOpenFiles int) *persisterFactoryHandlerMock {
	return &persisterFactoryHandlerMock{
		dbType:            dbType,
		batchDelaySeconds: batchDelaySeconds,
		maxBatchSize:      maxBatchSize,
		maxOpenFiles:      maxOpenFiles,
	}
}

// Create -
func (mock *persisterFactoryHandlerMock) Create(path string) (types.Persister, error) {
	switch mock.dbType {
	case storageUnit.LvlDB:
		return leveldb.NewDB(path, mock.batchDelaySeconds, mock.maxBatchSize, mock.maxOpenFiles)
	case storageUnit.LvlDBSerial:
		return leveldb.NewSerialDB(path, mock.batchDelaySeconds, mock.maxBatchSize, mock.maxOpenFiles)
	case storageUnit.MemoryDB:
		return memorydb.New(), nil
	default:
		return nil, common.ErrNotSupportedDBType
	}
}

// IsInterfaceNil -
func (mock *persisterFactoryHandlerMock) IsInterfaceNil() bool {
	return mock == nil
}
