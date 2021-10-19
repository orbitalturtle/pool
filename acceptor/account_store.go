package acceptor

import (
	"github.com/lightninglabs/pool/account"
	"github.com/lightninglabs/pool/clientdb"
)

// AccountStore is a clientdb.DB wrapper to implement the account.Store
// interface.
type AccountStore struct {
	*clientdb.DB
}

var _ account.Store = (*AccountStore)(nil)

func (s *AccountStore) PendingBatch() error {
	_, err := s.DB.PendingBatchSnapshot()
	return err
}
