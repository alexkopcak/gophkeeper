package services

import (
	"context"
	"log"

	querypb "github.com/alexkopcak/gophkeeper/api-gateway/pkg/query/pb"
	servicespb "github.com/alexkopcak/gophkeeper/api-gateway/pkg/services/pb"
	"github.com/alexkopcak/gophkeeper/query-service/internal/db"
	"github.com/alexkopcak/gophkeeper/query-service/internal/models"
)

type QueryServer struct {
	querypb.UnimplementedQueryServiceServer
	Handler *db.Handler
}

var _ querypb.QueryServiceServer = (*QueryServer)(nil)

func NewQueryServer() *QueryServer {
	return &QueryServer{}
}

func (s *QueryServer) Query(ctx context.Context, in *querypb.QueryRequest) (*servicespb.QueryResponseArray, error) {
	var record models.Record
	if in.Type == querypb.MessageType_ANY {
		record = models.Record{UserId: in.UserID}
	} else {
		record = models.Record{UserId: in.UserID, MessageType: byte(in.Type)}
	}

	rows, err := s.Handler.DB.Where(&record).Rows()

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	outArray := &servicespb.QueryResponseArray{}

	for rows.Next() {
		err = s.Handler.DB.ScanRows(rows, &record)
		if err != nil {
			outArray.Error = err.Error()
			return outArray, err
		}

		out := &servicespb.QueryResponseArray_QueryResponse{
			Id:   record.Id,
			Type: servicespb.MessageType(record.MessageType),
			Data: record.Data,
			Meta: record.Meta,
		}

		outArray.Count += 1
		outArray.Items = append(outArray.Items, out)
	}

	return outArray, err
}
