// Copyright (C) 2022 Jonathan Giannuzzi <jonathan@giannuzzi.me>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sqlite3

/*
#cgo CFLAGS: -DUSE_SQLCIPHER
#cgo CFLAGS: -DSQLITE_HAS_CODEC
#cgo CFLAGS: -DSQLITE_TEMP_STORE=2
#cgo !darwin LDFLAGS: -lcrypto
#cgo !darwin CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL
#cgo darwin LDFLAGS: -framework CoreFoundation -framework Security
#cgo darwin CFLAGS: -DSQLCIPHER_CRYPTO_CC
*/
import "C"
import (
	"bytes"
	"errors"
	"os"
)

// sqlite3Header defines the header string used by SQLite 3.
var sqlite3Header = []byte("SQLite format 3\000")

// IsEncrypted returns true, if the database with the given filename is
// encrypted, and false otherwise.
// If the database header cannot be read properly an error is returned.
func IsEncrypted(filename string) (bool, error) {
	// open file
	db, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer db.Close()
	// read header
	var header [16]byte
	n, err := db.Read(header[:])
	if err != nil {
		return false, err
	}
	if n != len(header) {
		return false, errors.New("go-sqlcipher: could not read full header")
	}
	// SQLCipher encrypts also the header, the file is encrypted if the read
	// header does not equal the header string used by SQLite 3.
	encrypted := !bytes.Equal(header[:], sqlite3Header)
	return encrypted, nil
}
