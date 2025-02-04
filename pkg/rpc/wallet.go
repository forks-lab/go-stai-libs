package rpc

import (
	"net/http"

	"github.com/forks-lab/go-stai-libs/pkg/rpcinterface"
	"github.com/forks-lab/go-stai-libs/pkg/types"
)

// WalletService encapsulates wallet RPC methods
type WalletService struct {
	client *Client
}

// NewRequest returns a new request specific to the wallet service
func (s *WalletService) NewRequest(rpcEndpoint rpcinterface.Endpoint, opt interface{}) (*rpcinterface.Request, error) {
	return s.client.NewRequest(rpcinterface.ServiceWallet, rpcEndpoint, opt)
}

// Do is just a shortcut to the client's Do method
func (s *WalletService) Do(req *rpcinterface.Request, v interface{}) (*http.Response, error) {
	return s.client.Do(req, v)
}

// GetWalletSyncStatusResponse Response for get_sync_status on wallet
type GetWalletSyncStatusResponse struct {
	Success            bool `json:"success"`
	GenesisInitialized bool `json:"genesis_initialized"`
	Synced             bool `json:"synced"`
	Syncing            bool `json:"syncing"`
}

// GetSyncStatus wallet rpc -> get_sync_status
func (s *WalletService) GetSyncStatus() (*GetWalletSyncStatusResponse, *http.Response, error) {
	request, err := s.NewRequest("get_sync_status", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletSyncStatusResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletHeightInfoResponse response for get_height_info on wallet
type GetWalletHeightInfoResponse struct {
	Success bool   `json:"success"`
	Height  uint32 `json:"height"`
}

// GetHeightInfo wallet rpc -> get_height_info
func (s *WalletService) GetHeightInfo() (*GetWalletHeightInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_height_info", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletHeightInfoResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletNetworkInfoResponse response for get_height_info on wallet
type GetWalletNetworkInfoResponse struct {
	Success       bool   `json:"success"`
	NetworkName   string `json:"network_name"`
	NetworkPrefix string `json:"network_prefix"`
}

// GetNetworkInfo wallet rpc -> get_network_info
func (s *WalletService) GetNetworkInfo() (*GetWalletNetworkInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("get_network_info", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletNetworkInfoResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletsResponse wallet rpc -> get_wallets
type GetWalletsResponse struct {
	Success     bool                `json:"success"`
	Fingerprint int                 `json:"fingerprint"`
	Wallets     []*types.WalletInfo `json:"wallets"`
}

// GetWallets wallet rpc -> get_wallets
func (s *WalletService) GetWallets() (*GetWalletsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_wallets", nil)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletBalanceOptions request options for get_wallet_balance
type GetWalletBalanceOptions struct {
	WalletID uint32 `json:"wallet_id"`
}

// GetWalletBalanceResponse is the wallet balance RPC response
type GetWalletBalanceResponse struct {
	Success bool                 `json:"success"`
	Balance *types.WalletBalance `json:"wallet_balance"`
}

// GetWalletBalance returns wallet balance
func (s *WalletService) GetWalletBalance(opts *GetWalletBalanceOptions) (*GetWalletBalanceResponse, *http.Response, error) {
	request, err := s.NewRequest("get_wallet_balance", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletBalanceResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletTransactionCountOptions options for get transaction count
type GetWalletTransactionCountOptions struct {
	WalletID uint32 `json:"wallet_id"`
}

// GetWalletTransactionCountResponse response for get_transaction_count
type GetWalletTransactionCountResponse struct {
	Success  bool   `json:"success"`
	WalletID uint32 `json:"wallet_id"`
	Count    int    `json:"count"`
}

// GetTransactionCount returns the total count of transactions for the specific wallet ID
func (s *WalletService) GetTransactionCount(opts *GetWalletTransactionCountOptions) (*GetWalletTransactionCountResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transaction_count", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionCountResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletTransactionsOptions options for get wallet transactions
type GetWalletTransactionsOptions struct {
	WalletID  uint32 `json:"wallet_id"`
	Start     *int   `json:"start,omitempty"`
	End       *int   `json:"end,omitempty"`
	ToAddress string `json:"to_address,omitempty"`
}

// GetWalletTransactionsResponse response for get_wallet_transactions
type GetWalletTransactionsResponse struct {
	Success      bool                       `json:"success"`
	WalletID     uint32                     `json:"wallet_id"`
	Transactions []*types.TransactionRecord `json:"transactions"`
}

// GetTransactions wallet rpc -> get_transactions
func (s *WalletService) GetTransactions(opts *GetWalletTransactionsOptions) (*GetWalletTransactionsResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transactions", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetWalletTransactionOptions options for getting a single wallet transaction
type GetWalletTransactionOptions struct {
	WalletID      uint32 `json:"wallet_id"`
	TransactionID string `json:"transaction_id"`
}

// GetWalletTransactionResponse response for get_wallet_transactions
type GetWalletTransactionResponse struct {
	Success       bool                     `json:"success"`
	Transaction   *types.TransactionRecord `json:"transaction"`
	TransactionID string                   `json:"transaction_id"`
}

// GetTransaction returns a single transaction record
func (s *WalletService) GetTransaction(opts *GetWalletTransactionOptions) (*GetWalletTransactionResponse, *http.Response, error) {
	request, err := s.NewRequest("get_transaction", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetWalletTransactionResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// SendTransactionOptions represents the options for send_transaction
type SendTransactionOptions struct {
	WalletID uint32 `json:"wallet_id"`
	Amount   uint64 `json:"amount"`
	Address  string `json:"address"`
	Fee      uint64 `json:"fee"`
}

// SendTransactionResponse represents the response from send_transaction
type SendTransactionResponse struct {
	Success       bool                    `json:"success"`
	TransactionID string                  `json:"transaction_id"`
	Transaction   types.TransactionRecord `json:"transaction"`
}

// SendTransaction sends a transaction
func (s *WalletService) SendTransaction(opts *SendTransactionOptions) (*SendTransactionResponse, *http.Response, error) {
	request, err := s.NewRequest("send_transaction", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &SendTransactionResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// CatSpendOptions represents the options for cat_spend
type CatSpendOptions struct {
	WalletID uint32 `json:"wallet_id"`
	Amount   uint64 `json:"amount"`
	Address  string `json:"inner_address"`
	Fee      uint64 `json:"fee"`
}

// CatSpendResponse represents the response from cat_spend
type CatSpendResponse struct {
	Success       bool                    `json:"success"`
	TransactionID string                  `json:"transaction_id"`
	Transaction   types.TransactionRecord `json:"transaction"`
}

// CatSpend sends a transaction
func (s *WalletService) CatSpend(opts *CatSpendOptions) (*CatSpendResponse, *http.Response, error) {
	request, err := s.NewRequest("cat_spend", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &CatSpendResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// MintNFTOptions represents the options for nft_get_info
type MintNFTOptions struct {
	DidID             string   `json:"did_id"`             // not required
	EditionNumber     uint32   `json:"edition_number"`     // not required
	EditionCount      uint32   `json:"edition_count"`      // not required
	Fee               uint64   `json:"fee"`                // not required
	LicenseHash       string   `json:"license_hash"`       //not required
	LicenseURIs       []string `json:"license_uris"`       // not required
	MetaHash          string   `json:"meta_hash"`          // not required
	MetaURIs          []string `json:"meta_uris"`          // not required
	RoyaltyAddress    string   `json:"royalty_address"`    // not required
	RoyaltyPercentage uint32   `json:"royalty_percentage"` // not required
	TargetAddress     string   `json:"target_address"`     // not required
	Hash              string   `json:"hash"`
	URIs              []string `json:"uris"`
	WalletID          uint32   `json:"wallet_id"`
}

// MintNFTResponse represents the response from nft_get_info
type MintNFTResponse struct {
	SpendBundle types.SpendBundle `json:"spend_bundle"`
	Success     bool              `json:"success"`
	WalletID    uint32            `json:"wallet_id"`
}

// MintNFT Mint a new NFT
func (s *WalletService) MintNFT(opts *MintNFTOptions) (*MintNFTResponse, *http.Response, error) {
	request, err := s.NewRequest("nft_mint_nft", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &MintNFTResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetNFTsOptions represents the options for nft_get_nfts
type GetNFTsOptions struct {
	WalletID uint32 `json:"wallet_id"`
}

// GetNFTsResponse represents the response from nft_get_nfts
type GetNFTsResponse struct {
	Success  bool        `json:"success"`
	WalletID uint32      `json:"wallet_id"`
	NFTList  []types.NFT `json:"nft_list"`
}

// GetNFTs Show all NFTs in a given wallet
func (s *WalletService) GetNFTs(opts *GetNFTsOptions) (*GetNFTsResponse, *http.Response, error) {
	request, err := s.NewRequest("nft_get_nfts", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetNFTsResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// TransferNFTOptions represents the options for nft_get_info
type TransferNFTOptions struct {
	Fee           uint64 `json:"fee"` // not required
	NFTCoinID     string `json:"nft_coin_id"`
	TargetAddress string `json:"target_address"`
	WalletID      uint32 `json:"wallet_id"`
}

// TransferNFTResponse represents the response from nft_get_info
type TransferNFTResponse struct {
	SpendBundle types.SpendBundle `json:"spend_bundle"`
	Success     bool              `json:"success"`
	WalletID    uint32            `json:"wallet_id"`
}

// TransferNFT Get info about an NFT
func (s *WalletService) TransferNFT(opts *TransferNFTOptions) (*TransferNFTResponse, *http.Response, error) {
	request, err := s.NewRequest("nft_transfer_nft", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &TransferNFTResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetNFTInfoOptions represents the options for nft_get_info
type GetNFTInfoOptions struct {
	CoinID   string `json:"coin_id"`
	WalletID uint32 `json:"wallet_id"`
}

// GetNFTInfoResponse represents the response from nft_get_info
type GetNFTInfoResponse struct {
	NFTInfo types.NFT `json:"nft_info"`
	Success bool      `json:"success"`
}

// GetNFTInfo Get info about an NFT
func (s *WalletService) GetNFTInfo(opts *GetNFTInfoOptions) (*GetNFTInfoResponse, *http.Response, error) {
	request, err := s.NewRequest("nft_get_info", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &GetNFTInfoResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// NFTAddURIOptions represents the options for nft_add_uri
type NFTAddURIOptions struct {
	Fee       uint64 `json:"fee"` // not required
	Key       string `json:"key"`
	NFTCoinID string `json:"nft_coin_id"`
	URI       string `json:"uri"`
	WalletID  uint32 `json:"wallet_id"`
}

// NFTAddURIResponse represents the response from nft_add_uri
type NFTAddURIResponse struct {
	SpendBundle types.SpendBundle `json:"spend_bundle"`
	Success     bool              `json:"success"`
	WalletID    uint32            `json:"wallet_id"`
}

// NFTAddURI Get info about an NFT
func (s *WalletService) NFTAddURI(opts *NFTAddURIOptions) (*NFTAddURIResponse, *http.Response, error) {
	request, err := s.NewRequest("nft_add_uri", opts)
	if err != nil {
		return nil, nil, err
	}

	r := &NFTAddURIResponse{}
	resp, err := s.Do(request, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
