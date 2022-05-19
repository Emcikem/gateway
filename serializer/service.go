package serializer

import (
	"gateway/dao"
	"gateway/service/model"
)

// BuildServiceVO 序列化服务
func BuildServiceVO(service *dao.GatewayService) model.ServiceDetailVO {
	return model.ServiceDetailVO{
		ID:           service.ID,
		ServiceName:  service.ServiceName,
		ServiceDesc:  service.ServiceDesc,
		LoadType:     service.LoadType,
		ServiceAddr:  service.ServiceAddr,
		OpenAuth:     service.OpenAuth,
		ClientLimit:  service.ClientLimit,
		ServerLimit:  service.ServerLimit,
		RoundType:    service.RoundType,
		IpList:       service.IpList,
		WeightList:   service.WeightList,
		WhiteIpList:  service.WhiteIpList,
		BlackIpList:  service.BlackIpList,
		RemoteParams: service.RemoteParams,
	}
}

func BuildServiceEntity(service *model.ServiceDetailVO) dao.GatewayService {
	return dao.GatewayService{
		ID:           service.ID,
		ServiceName:  service.ServiceName,
		ServiceDesc:  service.ServiceDesc,
		LoadType:     service.LoadType,
		ServiceAddr:  service.ServiceAddr,
		OpenAuth:     service.OpenAuth,
		ClientLimit:  service.ClientLimit,
		ServerLimit:  service.ServerLimit,
		RoundType:    service.RoundType,
		IpList:       service.IpList,
		WeightList:   service.WeightList,
		WhiteIpList:  service.WhiteIpList,
		BlackIpList:  service.BlackIpList,
		RemoteParams: service.RemoteParams,
	}
}
