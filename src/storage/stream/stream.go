/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package stream TODO
package stream

import (
	"context"
	"fmt"
	"time"

	"configcenter/src/apimachinery/discovery"
	"configcenter/src/storage/dal/mongo/local"
	"configcenter/src/storage/stream/event"
	"configcenter/src/storage/stream/loop"
	"configcenter/src/storage/stream/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// Interface TODO
// Stream Interface defines all the functionality it have.
type Interface interface {
	List(ctx context.Context, opts *types.ListOptions) (ch chan *types.Event, err error)
	Watch(ctx context.Context, opts *types.WatchOptions) (*types.Watcher, error)
	ListWatch(ctx context.Context, opts *types.ListWatchOptions) (*types.Watcher, error)
}

// NewStream create a list watch event stream
func NewStream(conf local.MongoConf) (Interface, error) {
	connStr, err := connstring.Parse(conf.URI)
	if nil != err {
		return nil, err
	}
	if conf.RsName == "" {
		return nil, fmt.Errorf("rsName not set")
	}

	timeout := 15 * time.Second
	conOpt := options.ClientOptions{
		MaxPoolSize:    &conf.MaxOpenConns,
		MinPoolSize:    &conf.MaxIdleConns,
		ConnectTimeout: &timeout,
	}

	// 区分建连方式
	if conf.ClusterMode == "shard" { // 分片集群模式
		// 读关注
		conOpt.SetReadConcern(readconcern.Majority()) // 指定查询应返回实例的最新数据确认为 已写入集群中的大多数成员
		// 写关注
		wc := writeconcern.New(writeconcern.WMajority()) // 请求确认写操作传播到大多数mongod实例
		wc = wc.WithOptions(writeconcern.WTimeout(30 * time.Second))
		conOpt.SetWriteConcern(wc)
	} else { // 副本集模式
		conOpt.ReplicaSet = &conf.RsName
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.URI), &conOpt)
	if nil != err {
		return nil, err
	}
	if err := client.Connect(context.TODO()); nil != err {
		return nil, err
	}

	event, err := event.NewEvent(client, connStr.Database)
	if err != nil {
		return nil, fmt.Errorf("new event failed, err: %v", err)
	}
	return event, nil
}

// LoopInterface TODO
type LoopInterface interface {
	WithOne(opts *types.LoopOneOptions) error
	WithBatch(opts *types.LoopBatchOptions) error
}

// NewLoopStream create a new event loop stream.
func NewLoopStream(conf local.MongoConf, isMaster discovery.ServiceManageInterface) (LoopInterface, error) {
	connStr, err := connstring.Parse(conf.URI)
	if nil != err {
		return nil, err
	}
	if conf.RsName == "" {
		return nil, fmt.Errorf("rsName not set")
	}

	timeout := 15 * time.Second
	conOpt := options.ClientOptions{
		MaxPoolSize:    &conf.MaxOpenConns,
		MinPoolSize:    &conf.MaxIdleConns,
		ConnectTimeout: &timeout,
	}

	// 区分建连方式
	if conf.ClusterMode == "shard" { // 分片集群模式
		// 读关注
		conOpt.SetReadConcern(readconcern.Majority()) // 指定查询应返回实例的最新数据确认为 已写入集群中的大多数成员
		// 写关注
		wc := writeconcern.New(writeconcern.WMajority()) // 请求确认写操作传播到大多数mongod实例
		wc = wc.WithOptions(writeconcern.WTimeout(30 * time.Second))
		conOpt.SetWriteConcern(wc)
	} else { // 副本集模式
		conOpt.ReplicaSet = &conf.RsName
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.URI), &conOpt)
	if nil != err {
		return nil, err
	}
	if err := client.Connect(context.TODO()); nil != err {
		return nil, err
	}

	event, err := event.NewEvent(client, connStr.Database)
	if err != nil {
		return nil, fmt.Errorf("new event failed, err: %v", err)
	}

	loop, err := loop.NewLoopWatch(event, isMaster)
	if err != nil {
		return nil, err
	}

	return loop, nil
}
