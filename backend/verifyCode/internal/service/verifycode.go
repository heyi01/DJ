package service

import (
	"context"
	"math/rand"

	pb "verifyCode/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}
func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:

		return randCode("0123456789", l)
	case pb.TYPE_LETTER:
		return randCode("qwertyuiopkjhffd", l)
	case pb.TYPE_MIXED:

		return randCode("0123456789qwertyuiopkjhffd", l)

	}
	return ""
}
func randCode(chars string, l int) string {
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		randIndex := rand.Intn(len(chars))
		result[i] = chars[randIndex]
	}
	return string(result)
}
