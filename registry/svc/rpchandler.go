package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
	"strings"
	"github.com/pkg/errors"
)

type rpchandler struct {

}

func (h* rpchandler) Registry(ctx context.Context, req *proto.ConnectMessageRequest, rsp *proto.ConnectMessageResponse) error {
	logrus.Infof("registry connect message, username : %s, client id : %s", req.Username, req.ClientId)
	// 1. update device info to mysql database and get guid from mysql database
	// Password hold "product_key:sign", sign is generated by device_name device_secret product_key
	info := strings.Split(string(req.Password), ":")
	if len(info) != 2 {
		rsp.Code = 400
		rsp.Message = "password must be contain product key and sign."
		return errors.New(rsp.Message)
	}
	guid, err := Global.DeviceSvc.Register(info[0], req.ClientId, info[1])
	if err != nil {
		rsp.Code = 600
		rsp.Message = err.Error()
		return err
	}
	rsp.Guid = guid
	rsp.Code = 200
	rsp.Message = "registry success."
	return nil
}
