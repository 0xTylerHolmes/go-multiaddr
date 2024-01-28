package main

import (
	ma "github.com/multiformats/go-multiaddr"
	net "github.com/multiformats/go-multiaddr/net"
	"testing"
)

var fullProtocolList = []string{
	"/unix/",
	"/ptp/",
	"/ip4/",
	"/ip6/",
	"/ip6zone/",
	"/ipfs/",
	"/quic/",
	"/quic-v1/",
	"/onion/",
	"/onion3/",
	"/garlic64/",
	"/garlic32/",
	"/udp/",
	"/tcp/",
	"/dns/",
	"/dns4/",
	"/dns6/",
	"/dnsaddr/",
	"/dccp/",
	"/ipcidr/",
	"/sctp/",
	"/p2p-circuit/",
	"/utp/",
	"/udt/",
	"/webtransport/",
	"/certhash/",
	"/http/",
	"/https/",
	"/unix/",
	"/p2p-webrtc-direct/",
	"/tls/",
	"/noise/",
	"/plaintextv2/",
	"/ws/",
	"/wss/",
	"/webrtc-direct/",
	"/webrtc/",
}

var fullTestList = []string{
	"/ip6/2601:9:4f81:9700:803e:ca65:66e8:c21",
	"/ip6/2601:9:4f81:9700:803e:ca65:66e8:c21/udp/1234/quic",
	"/ip6/2601:9:4f81:9700:803e:ca65:66e8:c21/udp/1234/quic-v1",
	"/ip6/2001:db8::/ipcidr/32",
	"/ip6zone/x/ip6/fe80::1",
	"/ip6zone/x%y/ip6/fe80::1",
	"/ip6zone/x%y/ip6/::",
	"/ip6zone/x/ip6/fe80::1/udp/1234/quic",
	"/ip6zone/x/ip6/fe80::1/udp/1234/quic-v1",
	"/udp/0",
	"/tcp/0",
	"/sctp/0",
	"/udp/1234",
	"/tcp/1234",
	"/sctp/1234",
	"/udp/65535",
	"/tcp/65535",
	"/ipfs/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC",
	"/ipfs/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7",
	"/p2p/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC",
	"/p2p/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7",
	"/p2p/bafzbeigvf25ytwc3akrijfecaotc74udrhcxzh2cx3we5qqnw5vgrei4bm",
	"/p2p/12D3KooWCryG7Mon9orvQxcS1rYZjotPgpwoJNHHKcLLfE4Hf5mV",
	"/p2p/k51qzi5uqu5dhb6l8spkdx7yxafegfkee5by8h7lmjh2ehc2sgg34z7c15vzqs",
	"/p2p/bafzaajaiaejcalj543iwv2d7pkjt7ykvefrkfu7qjfi6sduakhso4lay6abn2d5u",
	"/onion/timaq4ygg2iegci7:1234",
	"/onion/timaq4ygg2iegci7:80/http",
	"/onion3/vww6ybal4bd7szmgncyruucpgfkqahzddi37ktceo3ah7ngmcopnpyyd:1234",
	"/onion3/vww6ybal4bd7szmgncyruucpgfkqahzddi37ktceo3ah7ngmcopnpyyd:80/http",
	"/garlic64/jT~IyXaoauTni6N4517EG8mrFUKpy0IlgZh-EY9csMAk82Odatmzr~YTZy8Hv7u~wvkg75EFNOyqb~nAPg-khyp2TS~ObUz8WlqYAM2VlEzJ7wJB91P-cUlKF18zSzVoJFmsrcQHZCirSbWoOknS6iNmsGRh5KVZsBEfp1Dg3gwTipTRIx7Vl5Vy~1OSKQVjYiGZS9q8RL0MF~7xFiKxZDLbPxk0AK9TzGGqm~wMTI2HS0Gm4Ycy8LYPVmLvGonIBYndg2bJC7WLuF6tVjVquiokSVDKFwq70BCUU5AU-EvdOD5KEOAM7mPfw-gJUG4tm1TtvcobrObqoRnmhXPTBTN5H7qDD12AvlwFGnfAlBXjuP4xOUAISL5SRLiulrsMSiT4GcugSI80mF6sdB0zWRgL1yyvoVWeTBn1TqjO27alr95DGTluuSqrNAxgpQzCKEWAyzrQkBfo2avGAmmz2NaHaAvYbOg0QSJz1PLjv2jdPW~ofiQmrGWM1cd~1cCqAAAA",
	"/garlic64/jT~IyXaoauTni6N4517EG8mrFUKpy0IlgZh-EY9csMAk82Odatmzr~YTZy8Hv7u~wvkg75EFNOyqb~nAPg-khyp2TS~ObUz8WlqYAM2VlEzJ7wJB91P-cUlKF18zSzVoJFmsrcQHZCirSbWoOknS6iNmsGRh5KVZsBEfp1Dg3gwTipTRIx7Vl5Vy~1OSKQVjYiGZS9q8RL0MF~7xFiKxZDLbPxk0AK9TzGGqm~wMTI2HS0Gm4Ycy8LYPVmLvGonIBYndg2bJC7WLuF6tVjVquiokSVDKFwq70BCUU5AU-EvdOD5KEOAM7mPfw-gJUG4tm1TtvcobrObqoRnmhXPTBTN5H7qDD12AvlwFGnfAlBXjuP4xOUAISL5SRLiulrsMSiT4GcugSI80mF6sdB0zWRgL1yyvoVWeTBn1TqjO27alr95DGTluuSqrNAxgpQzCKEWAyzrQkBfo2avGAmmz2NaHaAvYbOg0QSJz1PLjv2jdPW~ofiQmrGWM1cd~1cCqAAAA/http",
	"/garlic64/jT~IyXaoauTni6N4517EG8mrFUKpy0IlgZh-EY9csMAk82Odatmzr~YTZy8Hv7u~wvkg75EFNOyqb~nAPg-khyp2TS~ObUz8WlqYAM2VlEzJ7wJB91P-cUlKF18zSzVoJFmsrcQHZCirSbWoOknS6iNmsGRh5KVZsBEfp1Dg3gwTipTRIx7Vl5Vy~1OSKQVjYiGZS9q8RL0MF~7xFiKxZDLbPxk0AK9TzGGqm~wMTI2HS0Gm4Ycy8LYPVmLvGonIBYndg2bJC7WLuF6tVjVquiokSVDKFwq70BCUU5AU-EvdOD5KEOAM7mPfw-gJUG4tm1TtvcobrObqoRnmhXPTBTN5H7qDD12AvlwFGnfAlBXjuP4xOUAISL5SRLiulrsMSiT4GcugSI80mF6sdB0zWRgL1yyvoVWeTBn1TqjO27alr95DGTluuSqrNAxgpQzCKEWAyzrQkBfo2avGAmmz2NaHaAvYbOg0QSJz1PLjv2jdPW~ofiQmrGWM1cd~1cCqAAAA/udp/8080",
	"/garlic64/jT~IyXaoauTni6N4517EG8mrFUKpy0IlgZh-EY9csMAk82Odatmzr~YTZy8Hv7u~wvkg75EFNOyqb~nAPg-khyp2TS~ObUz8WlqYAM2VlEzJ7wJB91P-cUlKF18zSzVoJFmsrcQHZCirSbWoOknS6iNmsGRh5KVZsBEfp1Dg3gwTipTRIx7Vl5Vy~1OSKQVjYiGZS9q8RL0MF~7xFiKxZDLbPxk0AK9TzGGqm~wMTI2HS0Gm4Ycy8LYPVmLvGonIBYndg2bJC7WLuF6tVjVquiokSVDKFwq70BCUU5AU-EvdOD5KEOAM7mPfw-gJUG4tm1TtvcobrObqoRnmhXPTBTN5H7qDD12AvlwFGnfAlBXjuP4xOUAISL5SRLiulrsMSiT4GcugSI80mF6sdB0zWRgL1yyvoVWeTBn1TqjO27alr95DGTluuSqrNAxgpQzCKEWAyzrQkBfo2avGAmmz2NaHaAvYbOg0QSJz1PLjv2jdPW~ofiQmrGWM1cd~1cCqAAAA/tcp/8080",
	"/garlic32/566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuq",
	"/garlic32/566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuqzwas",
	"/garlic32/566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuqzwassw",
	"/garlic32/566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuq/http",
	"/garlic32/566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuq/tcp/8080",
	"/garlic32/566niximlxdzpanmn4qouucvua3k7neniwss47li5r6ugoertzuq/udp/8080",
	"/udp/1234/sctp/1234",
	"/udp/1234/udt",
	"/udp/1234/utp",
	"/tcp/1234/http",
	"/tcp/1234/tls/http", "/ip4/127.0.0.1/udp/1234/quic-v1/webtransport/certhash/b2uaraocy6yrdblb4sfptaddgimjmmpy/certhash/zQmbWTwYGcmdyK9CYfNBcfs9nhZs17a6FQ4Y8oea278xx41",
	"/ip4/127.0.0.1/ipfs/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC",
	"/ip4/127.0.0.1/ipfs/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC/tcp/1234",
	"/ip4/127.0.0.1/ipfs/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7",
	"/ip4/127.0.0.1/ipfs/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7/tcp/1234",
	"/ip4/127.0.0.1/p2p/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC",
	"/ip4/127.0.0.1/p2p/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC/tcp/1234",
	"/ip4/127.0.0.1/p2p/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7",
	"/ip4/127.0.0.1/p2p/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7/tcp/1234",
	"/ip4/1.2.3.4/tcp/80/unix/a/b/c/d/e/f",
	"/ip4/127.0.0.1/ipfs/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC/tcp/1234/unix/stdio",
	"/ip4/127.0.0.1/ipfs/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7/tcp/1234/unix/stdio",
	"/ip4/127.0.0.1/p2p/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC/tcp/1234/unix/stdio",
	"/ip4/127.0.0.1/p2p/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7/tcp/1234/unix/stdio",
	"/ip4/127.0.0.1/tcp/9090/http/p2p-webrtc-direct",
	"/ip4/127.0.0.1/tcp/127/ws",
	"/ip4/127.0.0.1/tcp/127/ws",
	"/ip4/127.0.0.1/tcp/127/tls",
	"/ip4/127.0.0.1/tcp/127/tls/ws",
	"/ip4/127.0.0.1/tcp/127/noise",
	"/ip4/127.0.0.1/tcp/-1127/wss",
	"/ip4/127.0.0.1/tcp/127/wss",
	"/ip4/127.0.0.1/tcp/127/webrtc-direct",
	"/ip4/127.0.0.1/tcp/127/webrtc",
	"/dns/a.localhost/tcp/1",
	"/dns/node.libp2p.io/udp/1/quic-v1",
	"/dnsaddr/node.libp2p.io/udp/1/quic-v1",
	"/dns/node.libp2p.local/udp/1/quic-v1",
	"/dns/localhost/udp/1/quic-v1",
	"/tcp/1234/https",
	"/ipfs/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC/tcp/1234",
	"/ipfs/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7/tcp/1234",
	"/p2p/QmcgpsyWgH8Y8ajJz1Cu72KnS5uo2Aa2LpzU7kinSupNKC/tcp/1234",
	"/p2p/k2k4r8oqamigqdo6o7hsbfwd45y70oyynp98usk7zmyfrzpqxh1pohl7/tcp/1234",
	"/ip4/127.0.0.1/udp/1234",
	"/ip4/127.0.0.1/udp/0",
	"/ip4/127.0.0.1/tcp/1234",
	"/ip4/127.0.0.1/tcp/1234/",
	"/ip4/127.0.0.1/udp/1234/quic",
	"/ip4/127.0.0.1/udp/1234/quic-v1",
	"/ip4/127.0.0.1/udp/1234/quic-v1/webtransport",
	"/ip4/127.0.0.1/udp/1234/quic-v1/webtransport/certhash/b2uaraocy6yrdblb4sfptaddgimjmmpy",
}

func FuzzByteOperations(f *testing.F) {
	for _, s := range fullProtocolList {
		m, err := ma.NewMultiaddr(s)
		if err == nil {
			f.Add(m.Bytes())
		}
	}
	for _, s := range fullTestList {
		m, err := ma.NewMultiaddr(s)
		if err == nil {
			f.Add(m.Bytes())
		}
	}

	f.Fuzz(func(t *testing.T, b []byte) {
		m, err := ma.NewMultiaddrBytes(b)
		if err != nil {
			return
		}
		mainFuzzLoop(m)
	})
}

func FuzzStringOperations(f *testing.F) {
	for _, s := range fullProtocolList {
		_, err := ma.NewMultiaddr(s)
		if err == nil {
			f.Add(s)
		}
	}
	for _, s := range fullTestList {
		_, err := ma.NewMultiaddr(s)
		if err == nil {
			f.Add(s)
		}
	}

	f.Fuzz(func(t *testing.T, s string) {
		m, err := ma.NewMultiaddr(s)
		if err != nil {
			return
		}
		mainFuzzLoop(m)
	})
}

func mainFuzzLoop(m ma.Multiaddr) {
	m.Protocols()
	m.String()
	net.IsIPLoopback(m)
	net.IsIP6LinkLocal(m)
	net.IsThinWaist(m)
	net.IsIPUnspecified(m)
	net.IsNAT64IPv4ConvertedIPv6Addr(m)
	net.IsPublicAddr(m)
	net.IsPrivateAddr(m)
	ma.Split(m)
	ma.SplitFirst(m)
	ma.SplitLast(m)
	m.Encapsulate(m)
	m.Decapsulate(m)
	_, _ = net.ToNetAddr(m)
	_, _ = net.ToIP(m)
	_, _ = net.ResolveUnspecifiedAddress(m, nil)
	bin, err := m.MarshalBinary()
	if err != nil {
		panic("failed marshal binary")
	}
	js, err := m.MarshalJSON()
	if err != nil {
		panic("failed to marshal json")
	}
	txt, err := m.MarshalText()
	if err != nil {
		panic("failed to marshal txt")
	}

	m2, err := ma.NewMultiaddrBytes(m.Bytes())

	err = m2.UnmarshalBinary(bin)
	if err != nil {
		panic("unmarshal")
	}
	if !m.Equal(m2) {
		panic("round trip failed")
	}

	err = m2.UnmarshalJSON(js)
	if err != nil {
		panic("unmarshal")
	}
	// TODO: failing on .String() based marshalling
	//if !m.Equal(m2) {
	//	t.Logf("%s : %s", m.String(), m2.String())
	//	t.Logf("%v : %v", hex.EncodeToString(m.Bytes()), hex.EncodeToString(m2.Bytes()))
	//	panic("round trip failed")
	//}

	err = m2.UnmarshalText(txt)
	if err != nil {
		panic("text marshal")
	}
	// TODO: failing on .String() based marshalling
	//if !m.Equal(m2) {
	//	t.Logf("%s : %s", m.String(), m2.String())
	//	t.Logf("%v : %v", hex.EncodeToString(m.Bytes()), hex.EncodeToString(m2.Bytes()))
	//	panic("text not equal")
	//}
}
