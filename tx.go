package bigquery

import logrus "github.com/bonitoo-io/go-sql-bigquery/nolog"

type tx struct {
	c *Conn
}

func newTx(c *Conn) (*tx, error) {
	return &tx{c: c}, nil
}

// Commit currently just  passes through
func (t *tx) Commit() (err error) {
	logrus.Debug("Got tx.Commit")
	return
}

// Rollback currently just  passes through
func (t *tx) Rollback() (err error) {
	logrus.Debug("Got tx.Rollback")
	return
}
