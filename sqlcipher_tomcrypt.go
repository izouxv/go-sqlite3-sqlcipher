//go:build tomcrypt
// +build tomcrypt

package sqlite3

/*
// make go-sqlite3 use embedded library without code changes
#cgo CFLAGS: -DUSE_SQLCIPHER

// enable encryption codec in sqlite
#cgo CFLAGS: -DSQLITE_HAS_CODEC

// use memory for temporay storage in sqlite
#cgo CFLAGS: -DSQLITE_TEMP_STORE=2

// use libtomcrypt implementation in sqlcipher
#cgo CFLAGS: -DSQLCIPHER_CRYPTO_LIBTOMCRYPT

// disable anything "not portable" in libtomcrypt
#cgo CFLAGS: -DLTC_NO_ASM

// disable assertions
#cgo CFLAGS: -DNDEBUG

// set operating specific sqlite flags
#cgo linux CFLAGS: -DSQLITE_OS_UNIX=1
#cgo windows CFLAGS: -DSQLITE_OS_WIN=1

#include "./tomcrypt/hmac_memory.c"
#include "./tomcrypt/aes.c"
#include "./tomcrypt/compare_testvector.c"
#include "./tomcrypt/crypt_prng_descriptor.c"
#include "./tomcrypt/hmac_init.c"
#include "./tomcrypt/crypt_cipher_is_valid.c"
#include "./tomcrypt/cbc_decrypt.c"
#include "./tomcrypt/crypt_hash_descriptor.c"
#include "./tomcrypt/fortuna.c"
#include "./tomcrypt/cbc_done.c"
#include "./tomcrypt/crypt_register_prng.c"
#include "./tomcrypt/crypt_register_cipher.c"
#include "./tomcrypt/crypt_find_hash.c"
#include "./tomcrypt/sha1.c"
#include "./tomcrypt/crypt_cipher_descriptor.c"
#include "./tomcrypt/burn_stack.c"
#include "./tomcrypt/hmac_done.c"
#include "./tomcrypt/crypt_find_cipher.c"
#include "./tomcrypt/crypt_register_hash.c"
#include "./tomcrypt/crypt_hash_is_valid.c"
#include "./tomcrypt/hash_memory.c"
#include "./tomcrypt/sha256.c"
#include "./tomcrypt/crypt_argchk.c"
#include "./tomcrypt/cbc_start.c"
#include "./tomcrypt/pkcs_5_2.c"
#include "./tomcrypt/cbc_encrypt.c"
#include "./tomcrypt/zeromem.c"
#include "./tomcrypt/sha512.c"
#include "./tomcrypt/hmac_process.c"
*/
import "C"

// go test  --tags tomcrypt
