package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"

	"cosmossdk.io/math"
	rpc "github.com/DonggyuLim/grc20/protos/RPC"
	u "github.com/DonggyuLim/grc20/utils"
	"google.golang.org/grpc"
)

const (
	RPCPort = "9001"
)

type RPCServer struct {
	rpc.RPCServer
}

func Grpc(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", ":"+RPCPort)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	rpc.RegisterRPCServer(grpcServer, &RPCServer{})

	log.Printf("start gRPC server on %s port", RPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server:%s", err)
	}
	defer wg.Done()
}

func (r *RPCServer) Transfer(ctx context.Context, req *rpc.TransferRequest) (*rpc.TransferResponse, error) {
	fmt.Println("Transfer!!!")
	tokenName := req.TokenName
	to := req.To
	from := req.From

	t, err := u.GetToken(tokenName)
	if err != nil {
		return &rpc.TransferResponse{
			ToBalance:   "0",
			FromBalance: "0",
		}, err
	}
	amount, _ := math.NewIntFromString(req.Amount)
	t.Transfer(from, to, amount)
	u.SaveToken(tokenName, t)
	return &rpc.TransferResponse{

		ToBalance:   t.BalanceOf(to).String(),
		FromBalance: t.BalanceOf(from).String(),
	}, nil
}

func (r *RPCServer) Approve(ctx context.Context, req *rpc.ApproveRequest) (*rpc.ApproveResponse, error) {
	fmt.Println("Approve!!!")
	tokenName := req.GetTokenName()
	owner := req.GetOwner()
	spender := req.GetSpender()
	amount, _ := math.NewIntFromString(req.Amount)

	t, err := u.GetToken(tokenName)
	if err != nil {
		return &rpc.ApproveResponse{
			Allowance: "0",
		}, err
	}

	t.Approve(owner, spender, amount)

	u.SaveToken(tokenName, t)
	return &rpc.ApproveResponse{

		Allowance: t.Allowance(owner, spender).String(),
	}, nil
}

func (r *RPCServer) TransferFrom(ctx context.Context, req *rpc.TransferFromRequest) (*rpc.TransferFromResponse, error) {
	fmt.Println("TransferFrom!!")
	tokenName := req.GetTokenName()
	owner := req.GetOwner()
	spender := req.GetSpender()
	to := req.GetTo()
	amount, _ := math.NewIntFromString(req.GetAmount())

	t, err := u.GetToken(tokenName)
	if err != nil {
		return &rpc.TransferFromResponse{

			ToBalance: "0",
		}, err
	}
	err = t.TransferFrom(owner, spender, to, amount)
	if err != nil {
		return &rpc.TransferFromResponse{

			ToBalance: "0",
		}, err
	}
	u.SaveToken(tokenName, t)
	return &rpc.TransferFromResponse{

		ToBalance: t.BalanceOf(to).String(),
	}, nil
}

func (r *RPCServer) GetBalance(ctx context.Context, req *rpc.GetBalanceRequest) (*rpc.GetBalanceResponse, error) {
	fmt.Println("GetBalance!!")
	t, err := u.GetToken(req.GetTokenName())
	if err != nil {
		return &rpc.GetBalanceResponse{

			Balance: "0",
		}, err
	}
	balance := t.BalanceOf(req.GetAccount())
	return &rpc.GetBalanceResponse{
		Balance: balance.String(),
	}, nil
}

func (r *RPCServer) GetTokenInfo(ctx context.Context, req *rpc.TokenInfoRequest) (*rpc.TokenInfoResponse, error) {
	fmt.Println("GetTokenInfo!!")
	t, err := u.GetToken(req.GetTokenName())
	if err != nil {
		return &rpc.TokenInfoResponse{

			TokenName:   "",
			Symbol:      "",
			Decimal:     "0",
			TotalSupply: "0",
		}, err
	}
	return &rpc.TokenInfoResponse{
		TokenName:   t.GetName(),
		Symbol:      t.GetSymbol(),
		Decimal:     strconv.Itoa(int(t.GetDecimal())),
		TotalSupply: t.GetTotalSupply().String(),
	}, nil
}

func (r *RPCServer) GetAllowance(ctx context.Context, req *rpc.AllowanceRequest) (*rpc.AllowanceResponse, error) {
	fmt.Println("GetAllowance!!!")
	t, err := u.GetToken(req.GetTokenName())
	if err != nil {
		return &rpc.AllowanceResponse{

			Allowance: "0",
		}, err
	}
	allowance := t.Allowance(req.GetOwner(), req.GetSpender())
	return &rpc.AllowanceResponse{
		Allowance: allowance.String(),
	}, nil
}
