// Copyright 2018 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/maichain/eth-indexer/client/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cache Test", func() {
	var (
		mockClient  *mocks.EthClient
		cacheClient *cacheMiddleware
	)

	ctx := context.Background()
	block := types.NewBlockWithHeader(
		&types.Header{
			Number:     big.NewInt(100),
			Root:       common.StringToHash("12345678900"),
			Difficulty: big.NewInt(100),
		},
	)
	tx := types.NewTransaction(1, common.HexToAddress("1234567890"), big.NewInt(1000), 100, big.NewInt(10000), nil)
	unknownErr := errors.New("unknown error")

	BeforeEach(func() {
		mockClient = new(mocks.EthClient)
		cacheClient = newCacheMiddleware(mockClient).(*cacheMiddleware)
	})

	AfterEach(func() {
		mockClient.AssertExpectations(GinkgoT())
	})

	Context("BlockByNumber()", func() {
		It("find block successfully", func() {
			mockClient.On("BlockByNumber", ctx, block.Number()).Return(block, nil).Once()
			resBlock, err := cacheClient.BlockByNumber(ctx, block.Number())
			Expect(err).Should(BeNil())
			Expect(resBlock).Should(Equal(block))

			value, ok := cacheClient.blockCache.Get(block.Hash().Hex())
			Expect(ok).Should(BeTrue())
			block1 := value.(*types.Block)
			Expect(block1).Should(Equal(block))
		})
		It("failed to find block", func() {
			mockClient.On("BlockByNumber", ctx, block.Number()).Return(nil, unknownErr).Once()
			resBlock, err := cacheClient.BlockByNumber(ctx, block.Number())
			Expect(err).Should(Equal(unknownErr))
			Expect(resBlock).Should(BeNil())

			_, ok := cacheClient.blockCache.Get(block.Hash().Hex())
			Expect(ok).Should(BeFalse())
		})
	})

	Context("BlockByHash()", func() {
		It("in cache", func() {
			By("wrong in cache")
			cacheClient.blockCache.Add(block.Hash().Hex(), "wrong data")
			mockClient.On("BlockByHash", ctx, block.Hash()).Return(nil, unknownErr).Once()
			resBlock, err := cacheClient.BlockByHash(ctx, block.Hash())
			Expect(err).Should(Equal(unknownErr))
			Expect(resBlock).Should(BeNil())

			By("add in cache")
			mockClient.On("BlockByHash", ctx, block.Hash()).Return(block, nil).Once()
			resBlock, err = cacheClient.BlockByHash(ctx, block.Hash())
			Expect(err).Should(BeNil())
			Expect(resBlock).Should(Equal(block))

			By("already in cache")
			resBlock, err = cacheClient.BlockByHash(ctx, block.Hash())
			Expect(err).Should(BeNil())
			Expect(resBlock).Should(Equal(block))
		})
		Context("not in cache", func() {
			It("find block successfully", func() {
				mockClient.On("BlockByHash", ctx, block.Hash()).Return(block, nil).Once()
				resBlock, err := cacheClient.BlockByHash(ctx, block.Hash())
				Expect(err).Should(BeNil())
				Expect(resBlock).Should(Equal(block))

				resBlock, err = cacheClient.BlockByHash(ctx, block.Hash())
				Expect(err).Should(BeNil())
				Expect(resBlock).Should(Equal(block))
			})
			It("failed to find block", func() {
				mockClient.On("BlockByHash", ctx, block.Hash()).Return(nil, unknownErr).Once()
				resBlock, err := cacheClient.BlockByHash(ctx, block.Hash())
				Expect(err).Should(Equal(unknownErr))
				Expect(resBlock).Should(BeNil())

				_, ok := cacheClient.blockCache.Get(block.Hash().Hex())
				Expect(ok).Should(BeFalse())
			})
		})
	})

	Context("TransactionByHash()", func() {
		It("in cache", func() {
			By("wrong in cache")
			cacheClient.txCache.Add(tx.Hash().Hex(), "wrong data")
			mockClient.On("TransactionByHash", ctx, tx.Hash()).Return(nil, false, unknownErr).Once()
			resTX, _, err := cacheClient.TransactionByHash(ctx, tx.Hash())
			Expect(err).Should(Equal(unknownErr))
			Expect(resTX).Should(BeNil())

			By("add in cache")
			mockClient.On("TransactionByHash", ctx, tx.Hash()).Return(tx, true, nil).Once()
			resTX, pending, err := cacheClient.TransactionByHash(ctx, tx.Hash())
			Expect(err).Should(BeNil())
			Expect(pending).Should(BeTrue())
			Expect(resTX).Should(Equal(tx))

			By("already in cache")
			resTX, pending, err = cacheClient.TransactionByHash(ctx, tx.Hash())
			Expect(err).Should(BeNil())
			Expect(pending).Should(BeFalse())
			Expect(resTX).Should(Equal(tx))
		})
		Context("not in cache", func() {
			It("find tx successfully", func() {
				mockClient.On("TransactionByHash", ctx, tx.Hash()).Return(tx, true, nil).Once()
				resTX, pending, err := cacheClient.TransactionByHash(ctx, tx.Hash())
				Expect(err).Should(BeNil())
				Expect(pending).Should(BeTrue())
				Expect(resTX).Should(Equal(tx))

				resTX, pending, err = cacheClient.TransactionByHash(ctx, tx.Hash())
				Expect(err).Should(BeNil())
				Expect(pending).Should(BeFalse())
				Expect(resTX).Should(Equal(tx))
			})
			It("failed to find tx", func() {
				mockClient.On("TransactionByHash", ctx, tx.Hash()).Return(nil, false, unknownErr).Once()
				resTX, pending, err := cacheClient.TransactionByHash(ctx, tx.Hash())
				Expect(err).Should(Equal(unknownErr))
				Expect(pending).Should(BeFalse())
				Expect(resTX).Should(BeNil())

				_, ok := cacheClient.txCache.Get(tx.Hash().Hex())
				Expect(ok).Should(BeFalse())
			})
		})
	})
})

func TestClientServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Test")
}