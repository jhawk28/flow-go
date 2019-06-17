package collection_builder

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/dapperlabs/bamboo-emulator/data"
	"github.com/dapperlabs/bamboo-emulator/tests"
)

func TestCollectionBuilder(t *testing.T) {
	gomega := NewWithT(t)
	log := logrus.New()

	state := data.NewWorldState(log)

	transactionsIn := make(chan *data.Transaction)
	collectionsOut := make(chan *data.Collection)

	collectionBuilder := NewCollectionBuilder(state, transactionsIn, collectionsOut, log)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go collectionBuilder.Start(ctx, time.Millisecond)

	txA := tests.MockTransaction(1)
	txB := tests.MockTransaction(2)
	txC := tests.MockTransaction(3)

	transactionsIn <- txA
	transactionsIn <- txB
	transactionsIn <- txC

	collection := <-collectionsOut

	gomega.Expect(collection.TransactionHashes).To(HaveLen(3))
	gomega.Expect(collection.TransactionHashes).To(ContainElement(txA.Hash()))
	gomega.Expect(collection.TransactionHashes).To(ContainElement(txB.Hash()))
	gomega.Expect(collection.TransactionHashes).To(ContainElement(txC.Hash()))
}
