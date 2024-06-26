// Copyright 2024 The Solaris Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chunkfs

import (
	"context"
	"github.com/solarisdb/solaris/golibs/errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestChunkAccessor_OpenChunk(t *testing.T) {
	ca := NewChunkAccessor()

	assert.Nil(t, ca.openChunk(context.Background(), "ll"))
	assert.NotNil(t, ca.openChunk(context.Background(), "ll"))
	assert.Nil(t, ca.closeChunk("ll"))
	assert.Nil(t, ca.openChunk(context.Background(), "ll"))

	assert.Nil(t, ca.closeChunk("ll"))
	assert.True(t, ca.setDeleting("ll"))
	assert.NotNil(t, ca.openChunk(context.Background(), "ll"))
	ca.SetIdle("ll")
	assert.Nil(t, ca.openChunk(context.Background(), "ll"))

	assert.Nil(t, ca.closeChunk("ll"))
	ca.SetWriting(context.Background(), "ll")
	go func() {
		time.Sleep(50 * time.Millisecond)
		ca.SetIdle("ll")
	}()
	assert.Nil(t, ca.openChunk(context.Background(), "ll"))
	assert.NotNil(t, ca.openChunk(context.Background(), "ll"))

	assert.Nil(t, ca.closeChunk("ll"))
	ca.SetWriting(context.Background(), "ll")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	err := ca.openChunk(ctx, "ll")
	assert.Equal(t, ctx.Err(), err)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ca.Shutdown()
	}()
	assert.True(t, errors.ErrClosed == ca.openChunk(context.Background(), "ll"))
	assert.True(t, errors.ErrClosed == ca.openChunk(context.Background(), "ll"))
}

func TestChunkAccessor_SetWriting(t *testing.T) {
	ca := NewChunkAccessor()

	assert.Nil(t, ca.SetWriting(context.Background(), "ll"))
	ca.SetIdle("ll")
	assert.Nil(t, ca.openChunk(context.Background(), "ll"))
	assert.Nil(t, ca.SetWriting(context.Background(), "ll"))

	go func() {
		time.Sleep(50 * time.Millisecond)
		ca.SetIdle("ll")
	}()
	assert.Nil(t, ca.SetWriting(context.Background(), "ll"))

	ca.SetIdle("ll")
	assert.Nil(t, ca.closeChunk("ll"))
	assert.Nil(t, ca.SetWriting(context.Background(), "ll"))
	go func() {
		time.Sleep(50 * time.Millisecond)
		ca.SetIdle("ll")
	}()
	assert.Nil(t, ca.SetWriting(context.Background(), "ll"))
	ca.SetIdle("ll")
	assert.Equal(t, 0, len(ca.chunks))

	ca.SetWriting(context.Background(), "ll")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	err := ca.SetWriting(ctx, "ll")
	assert.Equal(t, ctx.Err(), err)
	ca.SetIdle("ll")

	// check the waiting channel
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			assert.Nil(t, ca.SetWriting(context.Background(), "l1"))
			ca.SetIdle("l1")
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Nil(t, ca.SetWriting(context.Background(), "ll"))
	go func() {
		time.Sleep(50 * time.Millisecond)
		ca.Shutdown()
	}()
	assert.True(t, errors.ErrClosed == ca.SetWriting(context.Background(), "ll"))
	assert.True(t, errors.ErrClosed == ca.SetWriting(context.Background(), "ll"))
}

func TestChunkAccessor_SetDeleting(t *testing.T) {
	ca := NewChunkAccessor()
	assert.True(t, ca.setDeleting("ll"))
	assert.False(t, ca.setDeleting("ll"))
	ca.SetIdle("ll")
	assert.True(t, ca.setDeleting("ll"))
	ca.SetIdle("ll")

	ca.Shutdown()
	assert.False(t, ca.setDeleting("ll"))
}

func TestChunkAccessor_SetIdle(t *testing.T) {
	ca := NewChunkAccessor()

	ca.SetWriting(context.Background(), "l1")
	ca.openChunk(context.Background(), "l2")
	ca.Shutdown()
	assert.Equal(t, 2, len(ca.chunks))
	ca.SetIdle("l1")
	ca.SetIdle("l2")
	assert.Equal(t, 1, len(ca.chunks))
	assert.NotNil(t, ca.closeChunk("l1"))
	assert.Nil(t, ca.closeChunk("l2"))
	assert.Equal(t, 0, len(ca.chunks))
}
