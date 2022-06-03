package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// testnet
		"enode://62b41585d48b3ef7fb417afafb701870fc6de5f171e343d135307c89d73d19952de42d9c108feb1cefdfe414e39634dd437a39a579a25e70d63ed6fa5f2d65da@65.108.43.238:18887",
		"enode://80ac277343411bcdc37d03d10329398fccb798cf791d869a8d4d69ede31a27ce417b6560e11a1f1e24121cae96f2fd76acc753df42a3ad82f2536379514961d2@65.108.124.81:18887",
		"enode://c12d3af7947b99654e5b9ec44eb24e908955fd55570083f0c66fc3dd5312983a97e20d4d178edbf31fe14b3d4489b5110a85576f3ac1b99b37530379e40c30eb@65.108.2.89:18887",
		// mainnet
		"enode://f424c35794f0d73ea8a6cfb740865d910568ef5df18efe5258510711efe1c8dd107100e95bbe9d4788bad3434fd52c863740bdce7858a2b1e80919fe1af07acb@65.108.141.55:18887",
		"enode://7d45a6987abf274aeae326fa3a99bf18d90e9a695d9504150aae44290ee6830e2366d4f0b84754c0921be782297dd8c2eb848690cfde8ec28966a2abfa443ba0@65.21.89.25:18887",
		"enode://38a5d52445c70cf99fed60e79149faa0cd539fee9c105e6196ebbdfd3895f8b284899f439c2412ad4dc4b4276044d22d855ef5d9abcccd0797a19f4c23816411@95.217.205.113:18887",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
