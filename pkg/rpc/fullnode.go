package rpc

import (
	"net/http"

	"github.com/forks-lab/go-stai-libs/pkg/rpcinterface"
	"github.com/forks-lab/go-stai-libs/pkg/types"
)

// FullNodeService encapsulates full node RPC methods
type FullNodeService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *FullNodeService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceFullNode, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *FullNodeService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetConnectionsOptions options to filter get_connections
type GetConnectionsOptions struct {
	NodeType types.NodeType `json:"node_type,omitempty"`
}

// GetConnectionsResponse get_connections response format
type GetConnectionsResponse struct {
	Success     bool                `json:"success"`
	Connections []*types.Connection `json:"connections"`
}

// GetConnections returns connections
func (s *FullNodeService) GetConnections(opts *GetConnectionsOptions) (*GetConnectionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_connections", opts)
	if err != nil {
		return nil, nil, err
	}

	c := &GetConnectionsResponse{}
	resp, err := s.Do(request, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// GetBlockchainStateResponse is the blockchain state RPC response
type GetBlockchainStateResponse struct {
	Success         bool                   `json:"success"`
	BlockchainState *types.BlockchainState `json:"blockchain_state"`
}

// GetBlockchainState returns blockchain state
func (s *FullNodeService) GetBlockchainState() (*GetBlockchainStateResponse, *http.Response, error) {
	request, err := s.NewRequest("get_blockchain_state", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlockchainStateResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockOptions options for get_block rpc call
type GetBlockOptions struct {
	HeaderHash string `json:"header_hash"`
}

// GetBlockResponse response for get_block rpc call
type GetBlockResponse struct {
	Success bool             `json:"success"`
	Block   *types.FullBlock `json:"block"`
}

// GetBlock full_node->get_block RPC method
func (s *FullNodeService) GetBlock(opts *GetBlockOptions) (*GetBlockResponse, *http.Response, error) {
	request, err := s.NewRequest("get_block", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlockResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlocksOptions options for get_blocks rpc call
type GetBlocksOptions struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// GetBlocksResponse response for get_blocks rpc call
type GetBlocksResponse struct {
	Success bool               `json:"success"`
	Blocks  []*types.FullBlock `json:"blocks"`
}

// GetBlocks full_node->get_blocks RPC method
func (s *FullNodeService) GetBlocks(opts *GetBlocksOptions) (*GetBlocksResponse, *http.Response, error) {
	request, err := s.NewRequest("get_blocks", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlocksResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockCountMetricsResponse response for get_block_count_metrics rpc call
type GetBlockCountMetricsResponse struct {
	Success bool                     `json:"success"`
	Metrics *types.BlockCountMetrics `json:"metrics"`
}

// GetBlockCountMetrics gets metrics about blocks
func (s *FullNodeService) GetBlockCountMetrics() (*GetBlockCountMetricsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_block_count_metrics", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetBlockCountMetricsResponse{}

	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetBlockByHeightOptions options for get_block_record_by_height and get_block rpc call
type GetBlockByHeightOptions struct {
	BlockHeight int `json:"height"`
}

// GetBlockRecordResponse response from get_block_record_by_height
type GetBlockRecordResponse struct {
	Success     bool               `json:"success"`
	BlockRecord *types.BlockRecord `json:"block_record"`
}

// GetBlockRecordByHeight full_node->get_block_record_by_height RPC method
func (s *FullNodeService) GetBlockRecordByHeight(opts *GetBlockByHeightOptions) (*GetBlockRecordResponse, *http.Response, error) {
	// Get Block Record
	request, err := s.NewRequest("get_block_record_by_height", opts)
	if err != nil {
		return nil, nil, err
	}

	record := &GetBlockRecordResponse{}
	resp, err := s.Do(request, record)
	if err != nil {
		return nil, resp, err
	}

	// @TODO handle this correctly
	// I believe this happens when the node is not yet synced to this height
	if record == nil || record.BlockRecord == nil {
		return nil, nil, nil
	}

	return record, resp, nil
}

// GetBlockByHeight helper function to get a full block by height, calls full_node->get_block_record_by_height RPC method then full_node->get_block RPC method
func (s *FullNodeService) GetBlockByHeight(opts *GetBlockByHeightOptions) (*GetBlockResponse, *http.Response, error) {
	// Get Block Record
	record, resp, err := s.GetBlockRecordByHeight(opts)
	if err != nil {
		return nil, resp, err
	}

	request, err := s.NewRequest("get_block", GetBlockOptions{
		HeaderHash: record.BlockRecord.HeaderHash,
	})
	if err != nil {
		return nil, nil, err
	}

	// Get Full Block
	block := &GetBlockResponse{}
	resp, err = s.Do(request, block)
	if err != nil {
		return nil, resp, err
	}

	return block, resp, nil
}
