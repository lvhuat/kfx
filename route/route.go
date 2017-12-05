package route

import (
	"io"
	"keto/kfx/pb"

	"github.com/gogo/protobuf/proto"
)

func HandleStream(stream pb.KfxService_DoWorksServer) error {
	return nil
}

type Router interface {
	Route(stream pb.KfxService_DoWorksServer) error
}

type Mt4Route struct {
}

func (mt4Route *Mt4Route) configService(req *pb.Request, stream pb.KfxService_DoWorksServer) {

}

func (mt4Route *Mt4Route) getCommonSetting(req *pb.Request, stream pb.KfxService_DoWorksServer) {

}

func (mt4Route *Mt4Route) Route(stream pb.KfxService_DoWorksServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		newRes := func(code pb.Codes, msg proto.Message) *pb.Response {
			return &pb.Response{
				Rid:  req.Rid,
				Code: int32(code),
			}
		}
		switch pb.Codes(req.Code) {
		case pb.Codes_Ping:
			stream.Send(newRes(pb.Codes(req.Code), nil))
		case pb.Codes_ConfigService:
			mt4Route.configService(req, stream)
		case pb.Codes_GetCoreStatus:
		case pb.Codes_GetCommonSetting:

		case pb.Codes_GetManagerSetting:
		case pb.Codes_GetHolidaySetting:
		case pb.Codes_GetTimeSetting:
		case pb.Codes_Apply:
		case pb.Codes_GetAccountInfo:
		case pb.Codes_GetAccountMargin:
		case pb.Codes_CheckPassword:
		case pb.Codes_UpdatePassword:
		case pb.Codes_GetGroups:
		case pb.Codes_GetGroup:
		case pb.Codes_GetSymbols:
		case pb.Codes_GetCharts:
		case pb.Codes_GetPrices:
		case pb.Codes_GetOrders:
		case pb.Codes_GetPositions:
		case pb.Codes_GetClosed:
		case pb.Codes_OpenMarket:
		case pb.Codes_OpenLimit:
		case pb.Codes_OpenStop:
		case pb.Codes_UpdateOrder:
		case pb.Codes_UpdatePosition:
		case pb.Codes_ClosePosition:
		case pb.Codes_Deposit:
		case pb.Codes_Withdraw:
		case pb.Codes_CreditIn:
		case pb.Codes_CreditOut:
		case pb.Codes_GetHistoryCashflows:
		case pb.Codes_GetHistoryCredits:
		case pb.Codes_CalPositionProfit:
		case pb.Codes_TradeEvent:
		case pb.Codes_CashflowEvent:
		case pb.Codes_PriceEvent:
		case pb.Codes_GroupEvent:
		case pb.Codes_AccountEvent:
		case pb.Codes_SymbolEvent:
		}
	}
	return nil
}
