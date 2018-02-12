// Copyright 2016 The go-daylight Authors
// This file is part of the go-daylight library.
//
// The go-daylight library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-daylight library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-daylight library. If not, see <http://www.gnu.org/licenses/>.

package tcpserver

import (
	"errors"

	"github.com/GenesisKernel/go-genesis/packages/consts"
	"github.com/GenesisKernel/go-genesis/packages/model"
	"github.com/GenesisKernel/go-genesis/packages/utils"

	log "github.com/sirupsen/logrus"
)

// Type7 writes the body of the specified block
// blocksCollection and queue_parser_blocks daemons send the request through p.GetBlocks()
func Type7(request *GetBodyRequest) (*GetBodyResponse, error) {
	block := &model.Block{}
	found, err := block.Get(int64(request.BlockID))
	if err != nil {
		log.WithFields(log.Fields{"type": consts.DBError, "error": err, "block_id": request.BlockID}).Error("Error getting block with id")
		return nil, utils.ErrInfo(err)
	}
	if !found {
		log.WithFields(log.Fields{"type": consts.NotFound, "block_id": request.BlockID}).Error("block with id not found")
		return nil, errors.New("Block not found. ID: " + string(request.BlockID))
	}
	return &GetBodyResponse{Data: block.Data}, nil
}
