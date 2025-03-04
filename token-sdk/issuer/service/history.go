/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"fmt"
	"github.com/hyperledger-labs/fabric-token-sdk/token"
	"github.com/pkg/errors"
	"strconv"
)

// SERVICE

// GetHistory returns the history of issuance for an issuer.
func (s TokenService) GetHistory() (tokens []string, err error) {
	w := token.GetManagementService(s.FSC).WalletManager().IssuerWallet("")
	if w == nil {
		return nil, errors.New("failed getting default issuer wallet")
	}

	res, err := w.ListIssuedTokens()
	if err != nil {
		return nil, err
	}

	for _, issuedToken := range res.Tokens {
		id := issuedToken.Id.String()
		tokenType := issuedToken.Type
		//owner := issuedToken.Owner TODO - Q. how to get the owner details? I get bytes array
		quantity, err := strconv.ParseInt(issuedToken.Quantity, 0, 64)
		if err != nil {
			return nil, errors.New("failed to parse quantity")
		}
		tokens = append(tokens, fmt.Sprintf("%s Issued %d of %s", id, quantity, tokenType))
	}

	return tokens, nil
}
