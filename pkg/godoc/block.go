// Package block contains common functionality for interacting with TSDB blocks
// in the context of Thanos.
package block 

import(
	"context"
	"fmt"
	"github.com/oklog/ulid"
)

const(
	// MetaFilename is the known JSON filename for meta information
	MetaFilename = "meta.json"
)

func Downloaded(ctx context.Context, id ulid.ULID, dst string) error {
	fmt.Println("Downloaded")
	return nil 
}

func cleanUp(ctx context.Context, id ulid.ULID) error {
	fmt.Println("Cleaned")
	return nil 
}

