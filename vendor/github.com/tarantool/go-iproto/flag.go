// Code generated by generate.sh; DO NOT EDIT.

package iproto

// IPROTO flag constants, generated from
// tarantool/src/box/iproto_constants.h
type Flag int

const (
	// Set for the last xrow in a transaction.
	IPROTO_FLAG_COMMIT Flag = 1 << 0
	// Set for the last row of a tx residing in limbo.
	IPROTO_FLAG_WAIT_SYNC Flag = 1 << 1
	// Set for the last row of a synchronous tx.
	IPROTO_FLAG_WAIT_ACK Flag = 1 << 2
)