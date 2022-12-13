package keeper

import (
	"context"
	"github.com/SaoNetwork/sao-did/parser"
	"strings"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidateDid(goCtx context.Context, req *types.QueryValidateDidRequest) (*types.QueryValidateDidResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	did := req.Did
	parsedDid, err := parser.Parse(did)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	switch parsedDid.Method {
	case "key":
		// check payment address
		if _, found := k.GetPaymentAddress(ctx, did); !found {
			return nil, status.Error(codes.NotFound, "payment address not found")
		}
	case "sid":
		// check sid document
		if _, found := k.GetSidDocument(ctx, parsedDid.ID); !found {
			return nil, status.Error(codes.NotFound, "sid document not found")
		}

		// check sid document version
		versionList, found := k.GetSidDocumentVersion(ctx, parsedDid.ID)
		if !found {
			return nil, status.Error(codes.InvalidArgument, "sidId should be a rootDocId")
		}

		// check version
		version := getVersionInfo(parsedDid.Query)
		if version != "" {
			if !inVersionList(version, versionList.VersionList) {
				return nil, status.Error(codes.NotFound, "sid version not found")
			}

			if _, found := k.GetSidDocument(ctx, version); !found {
				return nil, status.Error(codes.NotFound, "versioned sid document not found")
			}
		}

		// check payment address
		if _, found := k.GetPaymentAddress(ctx, did); !found {
			return nil, status.Error(codes.NotFound, "payment address not found")
		}

		// check account auth
		if _, found := k.GetAccountList(ctx, did); !found {
			return nil, status.Error(codes.NotFound, "account list not found")
		}

		// TODO: check pastSeeds, check binding accounts
	}

	return &types.QueryValidateDidResponse{}, nil
}

func getVersionInfo(query string) string {
	// version-id was changed to versionId in the latest did-core spec
	// https://github.com/w3c/did-core/pull/553
	var versionId string
	for _, q := range strings.Split(query, "&") {
		if strings.Contains(q, "versionId") || strings.Contains(q, "version-id") {
			versionId = strings.Split(q, "=")[1]
			break
		}
	}

	return versionId
}

func inVersionList(version string, list []string) bool {
	for _, v := range list {
		if v == version {
			return true
		}
	}
	return false
}
