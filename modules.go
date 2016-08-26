package main

import (
	_ "github.com/actuosus/registrator/consul"
	_ "github.com/actuosus/registrator/consulkv"
	_ "github.com/actuosus/registrator/etcd"
	_ "github.com/actuosus/registrator/skydns2"
	_ "github.com/actuosus/registrator/zookeeper"
)
