// Code generated by liverpcgen, DO NOT EDIT.
// source: *.proto files under this directory
// If you want to change this file, Please see README in go-common/app/tool/liverpc/protoc-gen-liverpc/
package liverpc

import (
	"go-common/app/service/live/userext/api/liverpc/v0"
	"go-common/app/service/live/userext/api/liverpc/v1"
	"go-common/library/net/rpc/liverpc"
)

// Client that represents a liverpc userext service api
type Client struct {
	cli *liverpc.Client
	// V0Conf presents the controller in liverpc
	V0Conf v0.ConfRPCClient
	// V1Bubble presents the controller in liverpc
	V1Bubble v1.BubbleRPCClient
	// V1Color presents the controller in liverpc
	V1Color v1.ColorRPCClient
	// V1DanmuConf presents the controller in liverpc
	V1DanmuConf v1.DanmuConfRPCClient
	// V1GrayRule presents the controller in liverpc
	V1GrayRule v1.GrayRuleRPCClient
	// V1Remind presents the controller in liverpc
	V1Remind v1.RemindRPCClient
	// V1Tag presents the controller in liverpc
	V1Tag v1.TagRPCClient
}

// DiscoveryAppId the discovery id is not the tree name
var DiscoveryAppId = "live.userext"

// New a Client that represents a liverpc live.userext service api
// conf can be empty, and it will use discovery to find service by default
// conf.AppID will be overwrite by a fixed value DiscoveryAppId
// therefore is no need to set
func New(conf *liverpc.ClientConfig) *Client {
	if conf == nil {
		conf = &liverpc.ClientConfig{}
	}
	conf.AppID = DiscoveryAppId
	var realCli = liverpc.NewClient(conf)
	cli := &Client{cli: realCli}
	cli.clientInit(realCli)
	return cli
}

func (cli *Client) GetRawCli() *liverpc.Client {
	return cli.cli
}

func (cli *Client) clientInit(realCli *liverpc.Client) {
	cli.V0Conf = v0.NewConfRPCClient(realCli)
	cli.V1Bubble = v1.NewBubbleRPCClient(realCli)
	cli.V1Color = v1.NewColorRPCClient(realCli)
	cli.V1DanmuConf = v1.NewDanmuConfRPCClient(realCli)
	cli.V1GrayRule = v1.NewGrayRuleRPCClient(realCli)
	cli.V1Remind = v1.NewRemindRPCClient(realCli)
	cli.V1Tag = v1.NewTagRPCClient(realCli)
}