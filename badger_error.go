/*
 * Copyright (c) 2022-2023 Zander Schwid & Co. LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

package badgerstore

import (
	"context"
	"github.com/codeallergy/store"
	"github.com/dgraph-io/badger/v3"
)

func wrapError(err error) error {
	switch err {
	case context.DeadlineExceeded:
		return err
	case context.Canceled:
		return err
	case badger.ErrConflict:
		return store.ErrConcurrentTxn
	case badger.ErrReadOnlyTxn:
		return store.ErrReadOnlyTxn
	case badger.ErrInvalidRequest:
		return store.ErrInvalidRequest
	case badger.ErrKeyNotFound:
		return store.ErrNotFound
	case badger.ErrEmptyKey:
		return store.ErrEmptyKey
	case badger.ErrInvalidKey:
		return store.ErrInvalidKey
	case badger.ErrDiscardedTxn:
		return store.ErrDiscardedTxn
	case badger.ErrTxnTooBig:
		return store.ErrTooBigTxn
	case badger.ErrDBClosed:
		return store.ErrAlreadyClosed
	case ErrTransactionCanceled:
		return store.ErrCanceledTxn
	default:
		return err
	}
}
