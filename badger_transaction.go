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
	"github.com/codeallergy/store"
	"github.com/dgraph-io/badger/v3"
)

type implBadgerTransaction struct {
	tx *badger.Txn
	readOnly bool
}

func NewTransaction(tx *badger.Txn, readOnly bool) store.Transaction {
	return &implBadgerTransaction{tx: tx, readOnly: readOnly}
}

func (t *implBadgerTransaction) ReadOnly() bool {
	return t.readOnly
}

func (t *implBadgerTransaction) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return wrapError(err)
	}
	return nil
}

func (t *implBadgerTransaction) Rollback() {
	t.tx.Discard()
}

func (t *implBadgerTransaction) Instance() interface{} {
	return t.tx
}
