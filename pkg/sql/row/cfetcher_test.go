// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package row

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCFetcherUninitialized(t *testing.T) {
	// Regression test for #36570: make sure it's okay to call GetRangesInfo even
	// before the fetcher was fully initialized.
	var fetcher CFetcher

	assert.Nil(t, fetcher.GetRangesInfo())
}
