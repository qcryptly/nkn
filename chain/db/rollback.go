package db

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/nknorg/nkn/block"
	"github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/common/serialization"
	"github.com/nknorg/nkn/pb"
	"github.com/nknorg/nkn/transaction"
	"github.com/nknorg/nkn/util/config"
	"github.com/nknorg/nkn/util/log"
)

func (cs *ChainStore) Rollback(b *block.Block) error {
	log.Warning("start rollback.")

	if err := cs.st.NewBatch(); err != nil {
		return err
	}

	if b.Header.UnsignedHeader.Height == 0 {
		return errors.New("the genesis block need not be rolled back.")
	}

	if err := cs.rollbackHeader(b); err != nil {
		return err
	}

	if err := cs.rollbackTransaction(b); err != nil {
		return err
	}

	if err := cs.rollbackBlockHash(b); err != nil {
		return err
	}

	if err := cs.rollbackCurrentBlockHash(b); err != nil {
		return err
	}

	if err := cs.rollbackDonation(b); err != nil {
		return err
	}

	if err := cs.rollbackNames(b); err != nil {
		return err
	}

	if err := cs.rollbackPubSub(b); err != nil {
		return err
	}

	if err := cs.rollbackStates(b); err != nil {
		return err
	}

	if err := cs.rollbackHeaderCache(b); err != nil {
		return err
	}

	if err := cs.st.BatchCommit(); err != nil {
		return err
	}

	return nil
}

func (cs *ChainStore) rollbackHeader(b *block.Block) error {
	blockHash := b.Hash()
	return cs.st.BatchDelete(append([]byte{byte(DATA_Header)}, blockHash[:]...))
}

func (cs *ChainStore) rollbackTransaction(b *block.Block) error {
	for _, txn := range b.Transactions {
		txHash := txn.Hash()
		if err := cs.st.BatchDelete(append([]byte{byte(DATA_Transaction)}, txHash[:]...)); err != nil {
			return err
		}
	}

	return nil
}

func (cs *ChainStore) rollbackBlockHash(b *block.Block) error {
	height := make([]byte, 4)
	binary.LittleEndian.PutUint32(height[:], b.Header.UnsignedHeader.Height)
	return cs.st.BatchDelete(append([]byte{byte(DATA_BlockHash)}, height...))
}

func (cs *ChainStore) rollbackCurrentBlockHash(b *block.Block) error {
	value := new(bytes.Buffer)
	prevHash, _ := common.Uint256ParseFromBytes(b.Header.UnsignedHeader.PrevBlockHash)
	if _, err := prevHash.Serialize(value); err != nil {
		return err
	}
	if err := serialization.WriteUint32(value, b.Header.UnsignedHeader.Height-1); err != nil {
		return err
	}

	return cs.st.BatchPut([]byte{byte(SYS_CurrentBlock)}, value.Bytes())
}

func (cs *ChainStore) rollbackNames(b *block.Block) error {
	for _, txn := range b.Transactions {
		if txn.UnsignedTx.Payload.Type == pb.RegisterNameType {
			pl, err := transaction.Unpack(txn.UnsignedTx.Payload)
			if err != nil {
				return err
			}

			registerNamePayload := pl.(*pb.RegisterName)
			err = cs.DeleteName(registerNamePayload.Registrant)
			if err != nil {
				return err
			}
		}
	}

	for _, txn := range b.Transactions {
		if txn.UnsignedTx.Payload.Type == pb.DeleteNameType {
			pl, err := transaction.Unpack(txn.UnsignedTx.Payload)
			if err != nil {
				return err
			}

			deleteNamePayload := pl.(*pb.DeleteName)
			err = cs.SaveName(deleteNamePayload.Registrant, deleteNamePayload.Name)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (cs *ChainStore) rollbackPubSub(b *block.Block) error {
	height := b.Header.UnsignedHeader.Height

	for _, txn := range b.Transactions {
		if txn.UnsignedTx.Payload.Type == pb.SubscribeType {
			pl, err := transaction.Unpack(txn.UnsignedTx.Payload)
			if err != nil {
				return err
			}

			subscribePayload := pl.(*pb.Subscribe)
			err = cs.Unsubscribe(subscribePayload.Subscriber, subscribePayload.Identifier, subscribePayload.Topic, subscribePayload.Bucket, subscribePayload.Duration, height)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (cs *ChainStore) rollbackStates(b *block.Block) error {
	//TODO add err statements
	prevHash, _ := common.Uint256ParseFromBytes(b.Header.UnsignedHeader.PrevBlockHash)
	prevHead, _ := cs.GetHeader(prevHash)
	root, _ := common.Uint256ParseFromBytes(prevHead.UnsignedHeader.StateRoot)
	cs.States, _ = NewStateDB(root, NewTrieStore(cs.GetDatabase()))

	err := cs.st.BatchPut(currentStateTrie(), root.ToArray())
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChainStore) rollbackHeaderCache(b *block.Block) error {
	cs.headerCache.RollbackHeader(b.Header)
	return nil
}

func (cs *ChainStore) rollbackDonation(b *block.Block) error {
	if b.Header.UnsignedHeader.Height%uint32(config.RewardAdjustInterval) != 0 {
		return nil
	}

	cs.st.BatchDelete(donationKey(b.Header.UnsignedHeader.Height))
	return nil
}