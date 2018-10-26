// Copyright 2018 Wolk Inc.
// This file is part of the Wolk go-plasma library.
package smt

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/wolkdb/go-plasma/deep"
)

func TestCheckProof(t *testing.T) {
	var proof Proof
	proof.key = common.Hex2Bytes("79e4453dcbc77b29")
	fmt.Printf("Index: %x\n", proof.key)
	proof.proofBits = uint64(0xc800000000000000)
	proof.proof = make([][]byte, 3)

	proof.proof[0] = common.Hex2Bytes("a5d59db538d26bd26e86b7fab2d688f8c03ab9d0dbf1adf2ef9bfa82de04b82b")
	proof.proof[1] = common.Hex2Bytes("49b4e065d6289c39dd4bb46545fd87a65edc5b9f9c8cc2fc6dfe9dc23b43d5a4")
	proof.proof[2] = common.Hex2Bytes("e8512edfdb95ea0eba5bdf718b981b3e845526b5d3ce2c463bc927cd5ad79a67")
	v := common.Hex2Bytes("7f2867b83f19a1443f67910d3f999a0385bbe50bf61c0df3795fbf23c081dd44")
	root := common.Hex2Bytes("ab06ee97217a525d229fe2f0ba129834b8a83742ae176b4987c5fdb95dc58797")
	//var defaultHashes [TreeDepth][]byte
	defaultHashes := ComputeDefaultHashes()
	fmt.Printf("Proof Bytes: %x\n", proof.Bytes())
	fmt.Printf("Proof: %s\n", proof.String())
	if !proof.Check(v, root, defaultHashes, true) {
		t.Fatalf("CheckProof Error\n")
	}
	recoveredProof := ToProof(deep.BytesToUint64(proof.key), proof.Bytes())
	if !recoveredProof.Check(v, root, defaultHashes, false) {
		t.Fatalf("recoveredProof Error\n")
	}
	invalidProof := recoveredProof
	invalidProof.proof[2] = common.Hex2Bytes("e8512edfdb95ea0eba5bdf718b981b3e845526b5d3ce2c463bc927cd5ad79a61")
	if invalidProof.Check(v, root, defaultHashes, false) {
		t.Fatalf("Uncaught Error: invalidProof pass\n")
	}
}
