package boltdb

import (
	"errors"
	"io"
	"path"
	"sync"

	bolt "go.etcd.io/bbolt"

	"github.com/drand/drand/chain"
	"github.com/drand/drand/log"
)

// BoltStore implements the Store interface using the kv storage boltdb (native
// golang implementation). Internally, Beacons are stored as JSON-encoded in the
// db file.
type BoltStore struct {
	sync.Mutex
	db  *bolt.DB
	log log.Logger
}

var beaconBucket = []byte("beacons")

// BoltFileName is the name of the file boltdb writes to
const BoltFileName = "drand.db"

// BoltStoreOpenPerm is the permission we will use to read bolt store file from disk
const BoltStoreOpenPerm = 0660

// NewBoltStore returns a Store implementation using the boltdb storage engine.
func NewBoltStore(l log.Logger, folder string, opts *bolt.Options) (*BoltStore, error) {
	dbPath := path.Join(folder, BoltFileName)
	db, err := bolt.Open(dbPath, BoltStoreOpenPerm, opts)
	if err != nil {
		return nil, err
	}
	// create the bucket already
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(beaconBucket)
		if err != nil {
			return err
		}
		return nil
	})

	return &BoltStore{
		log: l,
		db:  db,
	}, err
}

func (b *BoltStore) Len() (int, error) {
	var length = 0
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(beaconBucket)
		length = bucket.Stats().KeyN
		return nil
	})
	if err != nil {
		b.log.Warnw("", "boltdb", "error getting length", "err", err)
	}
	return length, err
}

func (b *BoltStore) Close() error {
	err := b.db.Close()
	if err != nil {
		b.log.Errorw("", "boltdb", "close", "err", err)
	}

	return err
}

// Put implements the Store interface. WARNING: It does NOT verify that this
// beacon is not already saved in the database or not and will overwrite it.
func (b *BoltStore) Put(beacon *chain.Beacon) error {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(beaconBucket)
		key := chain.RoundToBytes(beacon.Round)
		buff, err := beacon.Marshal()
		if err != nil {
			return err
		}
		return bucket.Put(key, buff)
	})
	if err != nil {
		return err
	}
	return nil
}

// ErrNoBeaconSaved is the error returned when no beacon have been saved in the
// database yet.
var ErrNoBeaconSaved = errors.New("beacon not found in database")

// Last returns the last beacon signature saved into the db
func (b *BoltStore) Last() (*chain.Beacon, error) {
	var beacon *chain.Beacon
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(beaconBucket)
		cursor := bucket.Cursor()
		_, v := cursor.Last()
		if v == nil {
			return ErrNoBeaconSaved
		}
		b := &chain.Beacon{}
		if err := b.Unmarshal(v); err != nil {
			return err
		}
		beacon = b
		return nil
	})
	return beacon, err
}

// Get returns the beacon saved at this round
func (b *BoltStore) Get(round uint64) (*chain.Beacon, error) {
	var beacon *chain.Beacon
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(beaconBucket)
		v := bucket.Get(chain.RoundToBytes(round))
		if v == nil {
			return ErrNoBeaconSaved
		}
		b := &chain.Beacon{}
		if err := b.Unmarshal(v); err != nil {
			return err
		}
		beacon = b
		return nil
	})
	if err != nil {
		return nil, err
	}
	return beacon, err
}

func (b *BoltStore) Del(round uint64) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(beaconBucket)
		return bucket.Delete(chain.RoundToBytes(round))
	})
}

func (b *BoltStore) Cursor(fn func(chain.Cursor) error) error {
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(beaconBucket)
		c := bucket.Cursor()
		return fn(&boltCursor{Cursor: c})
	})
	if err != nil {
		b.log.Warnw("", "boltdb", "error getting cursor", "err", err)
	}

	return err
}

// SaveTo saves the bolt database to an alternate file.
func (b *BoltStore) SaveTo(w io.Writer) error {
	return b.db.View(func(tx *bolt.Tx) error {
		_, err := tx.WriteTo(w)
		return err
	})
}

type boltCursor struct {
	*bolt.Cursor
}

func (c *boltCursor) First() (*chain.Beacon, error) {
	k, v := c.Cursor.First()
	if k == nil {
		return nil, ErrNoBeaconSaved
	}
	b := new(chain.Beacon)
	if err := b.Unmarshal(v); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *boltCursor) Next() (*chain.Beacon, error) {
	k, v := c.Cursor.Next()
	if k == nil {
		return nil, ErrNoBeaconSaved
	}
	b := new(chain.Beacon)
	if err := b.Unmarshal(v); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *boltCursor) Seek(round uint64) (*chain.Beacon, error) {
	k, v := c.Cursor.Seek(chain.RoundToBytes(round))
	if k == nil {
		return nil, ErrNoBeaconSaved
	}
	b := new(chain.Beacon)
	if err := b.Unmarshal(v); err != nil {
		return nil, err
	}
	return b, nil
}

func (c *boltCursor) Last() (*chain.Beacon, error) {
	k, v := c.Cursor.Last()
	if k == nil {
		return nil, ErrNoBeaconSaved
	}
	b := new(chain.Beacon)
	if err := b.Unmarshal(v); err != nil {
		return nil, err
	}
	return b, nil
}
