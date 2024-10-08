package entstorage

import (
	"fmt"

	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage/ent"
)

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
