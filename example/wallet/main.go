package main

//
//import (
//	"context"
//	"encoding/base64"
//	"fmt"
//	"log"
//	"strings"
//
//	"github.com/xssnick/tonutils-go/address"
//	"github.com/xssnick/tonutils-go/liteclient"
//	"github.com/xssnick/tonutils-go/tlb"
//	"github.com/xssnick/tonutils-go/ton"
//	"github.com/xssnick/tonutils-go/ton/wallet"
//)
//
//func main() {
//	client := liteclient.NewConnectionPool()
//
//	// get config
//	cfg, err := liteclient.GetConfigFromUrl(context.Background(), "https://ton.org/testnet-global.config.json")
//	if err != nil {
//		log.Fatalln("get config err: ", err.Error())
//		return
//	}
//
//	// connect to mainnet lite servers
//	err = client.AddConnectionsFromConfig(context.Background(), cfg)
//	if err != nil {
//		log.Fatalln("connection err: ", err.Error())
//		return
//	}
//
//	// api client with full proof checks
//	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry()
//	api.SetTrustedBlockFromConfig(cfg)
//
//	// bound all requests to single ton node
//	ctx := client.StickyContext(context.Background())
//
//	// seed words of account, you can generate them with any wallet or using wallet.NewSeed() method
//	words := strings.Split("wonder final foil crowd clip struggle tonight morning today fabric group measure lottery again chunk buzz unfair spin film angry trumpet broken another warrior", " ")
//
//	w, err := wallet.FromSeed(api, words, wallet.V4R2)
//	if err != nil {
//		log.Fatalln("FromSeed err:", err.Error())
//		return
//	}
//	//取出私钥
//	// 打印私钥
//	fmt.Printf("Wallet Address: %s\n", w.WalletAddress())
//
//	log.Println("wallet address:", w.WalletAddress())
//
//	privateKey := w.GetPrivateKey()
//	fmt.Printf("Private Key: %x\n", privateKey)
//
//	log.Println("fetching and checking proofs since config init block, it may take near a minute...")
//	block, err := api.CurrentMasterchainInfo(context.Background())
//	if err != nil {
//		log.Fatalln("get masterchain info err: ", err.Error())
//		return
//	}
//	log.Println("master proof checks are completed successfully, now communication is 100% safe!")
//
//	balance, err := w.GetBalance(ctx, block)
//	if err != nil {
//		log.Fatalln("GetBalance err:", err.Error())
//		return
//	}
//
//	if balance.Nano().Uint64() >= 3000000 {
//		addr := address.MustParseAddr("EQCD39VS5jcptHL8vMjEXrzGaRcCVYto7HUn4bpAOg8xqB2N")
//
//		log.Println("sending transaction and waiting for confirmation...")
//
//		// if destination wallet is not initialized (or you don't care)
//		// you should set bounce to false to not get money back.
//		// If bounce is true, money will be returned in case of not initialized destination wallet or smart-contract error
//		bounce := false
//
//		transfer, err := w.BuildTransfer(addr, tlb.MustFromTON("0.003"), bounce, "Hello from tonutils-go!")
//		if err != nil {
//			log.Fatalln("Transfer err:", err.Error())
//			return
//		}
//
//		tx, block, err := w.SendWaitTransaction(ctx, transfer)
//		if err != nil {
//			log.Fatalln("SendWaitTransaction err:", err.Error())
//			return
//		}
//
//		balance, err = w.GetBalance(ctx, block)
//		if err != nil {
//			log.Fatalln("GetBalance err:", err.Error())
//			return
//		}
//
//		log.Printf("transaction confirmed at block %d, hash: %s balance left: %s", block.SeqNo,
//			base64.StdEncoding.EncodeToString(tx.Hash), balance.String())
//
//		return
//	}
//
//	log.Println("not enough balance:", balance.String())
//}
