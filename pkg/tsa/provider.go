package tsa

import "github.com/nurmanhabib/go-tsa-client/domain/entity"

type Provider interface {
	TSARequest(tsq []byte) (*entity.TSReply, error)
}
