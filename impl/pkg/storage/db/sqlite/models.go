// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlite

type DhtRecord struct {
	ID    int64
	Key   []byte
	Value []byte
	Sig   []byte
	Seq   int64
}

type FailedRecord struct {
	ID           []byte
	FailureCount int64
}