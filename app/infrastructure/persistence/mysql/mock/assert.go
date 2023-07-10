package mock

import "testing"

// AssertAllTxCommitted asserts that all transactions are committed.
func (db *DB) AssertAllTxCommitted(t *testing.T) {
	t.Helper()

	db.mu.Lock()
	defer db.mu.Unlock()

	for _, tx := range db.txs {
		if tx.state != txStateCommitted {
			t.Errorf("tx %p is not committed", tx)
		}
	}
}

// AssertAllTxRolledBack asserts that all transactions are rolled back.
func (db *DB) AssertAllTxRolledBack(t *testing.T) {
	t.Helper()

	db.mu.Lock()
	defer db.mu.Unlock()

	for _, tx := range db.txs {
		if tx.state != txStateRolledBack {
			t.Errorf("tx %p is not rolled back", tx)
		}
	}
}

// AssertAllTxCommittedOrRolledBack asserts that all transactions are committed or rolled back.
func (db *DB) AssertAllTxCommittedOrRolledBack(t *testing.T) {
	t.Helper()

	db.mu.Lock()
	defer db.mu.Unlock()

	for _, tx := range db.txs {
		if tx.state != txStateCommitted && tx.state != txStateRolledBack {
			t.Errorf("tx %p is not committed or rolled back", tx)
		}
	}
}

// AssertAnyTxNotBegin asserts that any transaction is not begin.
func (db *DB) AssertAnyTxNotBegin(t *testing.T) {
	t.Helper()

	db.mu.Lock()
	defer db.mu.Unlock()

	for _, tx := range db.txs {
		if tx.state == txStateActive {
			t.Errorf("tx %p is begin", tx)
		}
	}
}

// AssertAnyTxNotActive asserts that any transaction is not active.
func (db *DB) AssertAnyTxNotActive(t *testing.T) {
	t.Helper()

	db.mu.Lock()
	defer db.mu.Unlock()

	for _, tx := range db.txs {
		if tx.state == txStateActive {
			t.Errorf("tx %p is active", tx)
		}
	}
}
