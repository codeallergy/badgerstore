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
	"github.com/dgraph-io/badger/v3"
	"reflect"
	"strings"
	"time"
)

func OpenDatabase(dataDir string, options ...Option) (*badger.DB, *StoreOptions, error) {

	storeOpts := DefaultStoreOptions()

	opts := badger.DefaultOptions(dataDir)
	opts.ValueLogMaxEntries = DefaultValueLogMaxEntries

	for _, opt := range options {
		opt.apply(&opts, storeOpts)
	}

	deadline := time.Now().Add(storeOpts.OpenTimeout)
	for {

		db, err := badger.Open(opts)
		if err != nil {
			if strings.Contains(err.Error(), "Cannot acquire directory lock") && time.Now().Before(deadline) {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			return nil, storeOpts, err
		}

		return db, storeOpts, nil
	}

}

func ObjectType() reflect.Type {
	return BadgerStoreClass
}


