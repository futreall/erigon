// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

//go:build integration

package tests

import (
	"fmt"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/erigontech/erigon-lib/log/v3"
)

func TestExecutionSpec(t *testing.T) {
	defer log.Root().SetHandler(log.Root().GetHandler())
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlError, log.StderrHandler))

	bt := new(testMatcher)

	dir := filepath.Join(".", "execution-spec-tests")
	checkStateRoot := true

	fmt.Println("Running TestExecutionSpec tests")
	testTimes := make(map[string]time.Duration)
	startTime := time.Now()
	bt.walk(t, dir, func(t *testing.T, name string, test *BlockTest) {
		// import pre accounts & construct test genesis block & state root
		testStart := time.Now()
		if err := bt.checkFailure(t, test.Run(t, checkStateRoot)); err != nil {
			t.Error(err)
		}
		testTimes[name] = time.Since(testStart)
	})

	fmt.Println("TestExecutionSpec test times:")
	for _, name := range sortMapByValue(testTimes) {
		fmt.Println(name, testTimes[name])
	}

	averageTime := time.Duration(0)
	for _, time := range testTimes {
		averageTime += time
	}
	averageTime /= time.Duration(len(testTimes))

	fmt.Println("Average blockchain test time:", averageTime)
	fmt.Println("Test count:", len(testTimes))
	fmt.Println("TestExecutionSpec tests took", time.Since(startTime))
}

func sortMapByValue(m map[string]time.Duration) []string {
	type kv struct {
		Key   string
		Value time.Duration
	}
	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})
	var keys []string
	for _, kv := range ss {
		keys = append(keys, kv.Key)
	}
	return keys
}
