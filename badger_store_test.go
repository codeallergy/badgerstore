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

package badgerstore_test

import (
	"context"
	"github.com/codeallergy/badgerstore"
	"github.com/codeallergy/store"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestBadgerStore(t *testing.T) {

	dir, err := os.MkdirTemp(os.TempDir(), "badgerstoretest")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	s, err := badgerstore.New("test", dir)
	require.NoError(t, err)
	defer s.Destroy()

	err = populateData(s)
	require.NoError(t, err)

	var visited bool
	var cnt int
	err = s.Enumerate(context.Background()).ByPrefix("test").Do(func(entry *store.RawEntry) bool {
		if string(entry.Key) == "test" && string(entry.Value) == "abc" {
			visited = true
		}
		cnt++
		return true
	})
	require.NoError(t, err)
	require.True(t, visited)
	require.Equal(t, 1, cnt)

	err = cleanData(s)
	require.NoError(t, err)
}

func populateData(store store.TransactionalDataStore) (err error) {

	ctx := store.BeginTransaction(context.Background(), false)
	defer func() {
		err = store.EndTransaction(ctx, err)
	}()

	err = store.Set(ctx).ByKey("test").String("abc")

	value, err := store.Get(ctx).ByKey("test").ToString()

	if value != "abc" {
		err = errors.Errorf("value not found")
	}

	return
}

func cleanData(store store.TransactionalDataStore) (err error) {

	ctx := store.BeginTransaction(context.Background(), false)
	defer func() {
		err = store.EndTransaction(ctx, err)
	}()

	err = store.Remove(ctx).ByKey("test").Do()

	return
}