package core

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/immesys/wave/eapi"
	"github.com/immesys/wave/iapi"
	"github.com/immesys/wave/localdb/lls"
	"github.com/immesys/wave/localdb/poc"
	"github.com/immesys/wave/storage/overlay"
	"github.com/immesys/wave/waved"
	"github.com/immesys/wave/wve"
	pb "github.com/immesys/wavemq/mqpb"
)

type AuthModule struct {
	cfg     *waved.Configuration
	wave    *eapi.EAPI
	cachemu sync.Mutex
	cache   map[cacheKey]time.Time
}

type cacheKey struct {
	Namespace  [32]byte
	URI        string
	Permission string
	ProofHash  [32]byte
}

func NewAuthModule(cfg *waved.Configuration) (*AuthModule, error) {
	llsdb, err := lls.NewLowLevelStorage(cfg.Database)
	if err != nil {
		return nil, err
	}
	si, err := overlay.NewOverlay(cfg.Storage)
	if err != nil {
		fmt.Printf("storage overlay error: %v\n", err)
		os.Exit(1)
	}

	iapi.InjectStorageInterface(si)
	ws := poc.NewPOC(llsdb)
	eapi := eapi.NewEAPI(ws)
	return &AuthModule{
		cfg:   cfg,
		wave:  eapi,
		cache: make(map[cacheKey]time.Time),
	}, nil
}

//This checks that a publish message is authorized for the given URI
func (am *AuthModule) CheckMessage(m *pb.Message) wve.WVE {
	//TODO
	return nil
}

//Check that the given proof is valid for subscription on the given URI
func (am *AuthModule) CheckSubscription(s *pb.PeerSubscribeParams) wve.WVE {
	//TODO
	return nil
}

func (am *AuthModule) FormMessage(p *pb.PublishParams, routerID string) (*pb.Message, wve.WVE) {
	//TODO
	return &pb.Message{
		Tbs: &pb.MessageTBS{
			Namespace:    p.Namespace,
			Uri:          p.Uri,
			Payload:      p.Content,
			OriginRouter: routerID,
		},
		Persist: p.Persist,
	}, nil
}

func (am *AuthModule) FormSubRequest(p *pb.SubscribeParams, routerID string) (*pb.PeerSubscribeParams, wve.WVE) {
	//TODO
	return &pb.PeerSubscribeParams{
		Tbs: &pb.PeerSubscriptionTBS{
			Namespace: p.Namespace,
			Uri:       p.Uri,
			Id:        p.Identifier,
			RouterID:  routerID,
		},
	}, nil
}