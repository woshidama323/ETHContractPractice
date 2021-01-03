// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OnesplitauditABI is the input ABI used to generate the binding from.
const OnesplitauditABI = "[{\"inputs\":[{\"internalType\":\"contractIOneSplit\",\"name\":\"impl\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImpl\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_AAVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_BANCOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_BDAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CHAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_COMPOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_BINANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_COMPOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_SYNTHETIX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_USDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_CURVE_Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_FULCRUM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_IEARN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_KYBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_OASIS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_SMART_TOKEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_UNISWAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_DISABLE_WETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_KYBER_BANCOR_RESERVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_KYBER_OASIS_RESERVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_KYBER_UNISWAP_RESERVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_MULTI_PATH_DAI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_MULTI_PATH_ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_MULTI_PATH_USDC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FLAG_ENABLE_UNISWAP_COMPOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"featureFlags\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oneSplitImpl\",\"outputs\":[{\"internalType\":\"contractIOneSplit\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIOneSplit\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setNewImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"featureFlags\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OnesplitauditBin is the compiled bytecode used for deploying new contracts.
var OnesplitauditBin = "0x60806040523480156200001157600080fd5b506040516200269c3803806200269c833981810160405260208110156200003757600080fd5b810190808051906020019092919050505060006200005a6200011060201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35062000109816200011860201b60201c565b5062000288565b600033905090565b620001286200022260201b60201c565b6200019b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0360405160405180910390a250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166200026c6200011060201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b61240480620002986000396000f3fe6080604052600436106101ee5760003560e01c806375a8b0121161010d578063ba4917b3116100a0578063d393c3e91161006f578063d393c3e914610887578063dc1536b2146108b2578063e2a7515e146108dd578063f2fde38b146109f3578063f56e281f14610a44576101ee565b8063ba4917b3146107ab578063c762a46c14610806578063c77b9de614610831578063c9b42c671461085c576101ee565b80638f32d59b116100dc5780638f32d59b146106d5578063b0a7ef2914610704578063b26413f81461072f578063b3bc784414610780576101ee565b806375a8b012146105d15780637a88bdbd146105fc578063867807ca146106275780638da5cb5b1461067e576101ee565b80634a7101d51161018557806364ec4e5c1161015457806364ec4e5c1461053957806368e2a014146105645780636cbc4a6e1461058f578063715018a6146105ba576101ee565b80634a7101d51461048d5780635aa8fb48146104b85780635ae51b82146104e35780635c0cb4791461050e576101ee565b80632d3b5207116101c15780632d3b5207146103e15780632e707bd21461040c57806334b4dabb1461043757806344211d6214610462576101ee565b8063085e2c5b1461027557806313989140146103605780632113240d1461038b57806321a360f5146103b6575b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610273576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806122bf6022913960400191505060405180910390fd5b005b34801561028157600080fd5b50610302600480360360a081101561029857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190505050610a6f565b6040518083815260200180602001828103825283818151815260200191508051906020019060200280838360005b8381101561034b578082015181840152602081019050610330565b50505050905001935050505060405180910390f35b34801561036c57600080fd5b50610375610c4b565b6040518082815260200191505060405180910390f35b34801561039757600080fd5b506103a0610c51565b6040518082815260200191505060405180910390f35b3480156103c257600080fd5b506103cb610c57565b6040518082815260200191505060405180910390f35b3480156103ed57600080fd5b506103f6610c60565b6040518082815260200191505060405180910390f35b34801561041857600080fd5b50610421610c69565b6040518082815260200191505060405180910390f35b34801561044357600080fd5b5061044c610c6e565b6040518082815260200191505060405180910390f35b34801561046e57600080fd5b50610477610c73565b6040518082815260200191505060405180910390f35b34801561049957600080fd5b506104a2610c78565b6040518082815260200191505060405180910390f35b3480156104c457600080fd5b506104cd610c7d565b6040518082815260200191505060405180910390f35b3480156104ef57600080fd5b506104f8610c83565b6040518082815260200191505060405180910390f35b34801561051a57600080fd5b50610523610c89565b6040518082815260200191505060405180910390f35b34801561054557600080fd5b5061054e610c8e565b6040518082815260200191505060405180910390f35b34801561057057600080fd5b50610579610c95565b6040518082815260200191505060405180910390f35b34801561059b57600080fd5b506105a4610c9c565b6040518082815260200191505060405180910390f35b3480156105c657600080fd5b506105cf610ca3565b005b3480156105dd57600080fd5b506105e6610ddc565b6040518082815260200191505060405180910390f35b34801561060857600080fd5b50610611610de2565b6040518082815260200191505060405180910390f35b34801561063357600080fd5b5061063c610de7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561068a57600080fd5b50610693610e0d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106e157600080fd5b506106ea610e36565b604051808215151515815260200191505060405180910390f35b34801561071057600080fd5b50610719610e94565b6040518082815260200191505060405180910390f35b34801561073b57600080fd5b5061077e6004803603602081101561075257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e9a565b005b34801561078c57600080fd5b50610795610f9b565b6040518082815260200191505060405180910390f35b3480156107b757600080fd5b50610804600480360360408110156107ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610fa4565b005b34801561081257600080fd5b5061081b61104e565b6040518082815260200191505060405180910390f35b34801561083d57600080fd5b50610846611053565b6040518082815260200191505060405180910390f35b34801561086857600080fd5b50610871611059565b6040518082815260200191505060405180910390f35b34801561089357600080fd5b5061089c611060565b6040518082815260200191505060405180910390f35b3480156108be57600080fd5b506108c7611067565b6040518082815260200191505060405180910390f35b6109f1600480360360c08110156108f357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001909291908035906020019064010000000081111561096457600080fd5b82018360208201111561097657600080fd5b8035906020019184602083028401116401000000008311171561099857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919050505061106d565b005b3480156109ff57600080fd5b50610a4260048036036020811015610a1657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611514565b005b348015610a5057600080fd5b50610a5961159a565b6040518082815260200191505060405180910390f35b60006060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663085e2c5b88888888886040518663ffffffff1660e01b8152600401808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020018281526020019550505050505060006040518083038186803b158015610b5e57600080fd5b505afa158015610b72573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506040811015610b9c57600080fd5b810190808051906020019092919080516040519392919084640100000000821115610bc657600080fd5b83820191506020820185811115610bdc57600080fd5b8251866020820283011164010000000082111715610bf957600080fd5b8083526020830192505050908051906020019060200280838360005b83811015610c30578082015181840152602081019050610c15565b50505050905001604052505050915091509550959350505050565b61200081565b61800081565b64020000000081565b64010000000081565b608081565b604081565b601081565b602081565b61400081565b61080081565b600881565b6202000081565b6210000081565b6208000081565b610cab610e36565b610d1d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b61040081565b600281565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610e7861159f565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b61100081565b610ea2610e36565b610f14576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0360405160405180910390a250565b64040000000081565b610fac610e36565b61101e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61104933828473ffffffffffffffffffffffffffffffffffffffff166115a79092919063ffffffff16565b505050565b600181565b61020081565b6204000081565b6201000081565b61010081565b8473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141580156110a95750600084115b61111b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f4f6e6553706c69743a2073776170206d616b6573206e6f2073656e736500000081525060200191505060405180910390fd5b61113a8673ffffffffffffffffffffffffffffffffffffffff16611650565b15156000341415151514611199576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001806123666034913960400191505060405180910390fd5b60006111d6346111c8308a73ffffffffffffffffffffffffffffffffffffffff166116d290919063ffffffff16565b6117c290919063ffffffff16565b90506000611203308873ffffffffffffffffffffffffffffffffffffffff166116d290919063ffffffff16565b905061122e868973ffffffffffffffffffffffffffffffffffffffff1661180c90919063ffffffff16565b61127b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16878a73ffffffffffffffffffffffffffffffffffffffff166118c29092919063ffffffff16565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e2a7515e348a8a8a8a8a8a6040518863ffffffff1660e01b8152600401808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200180602001838152602001828103825284818151815260200191508051906020019060200280838360005b8381101561138857808201518184015260208101905061136d565b505050509050019750505050505050506000604051808303818588803b1580156113b157600080fd5b505af11580156113c5573d6000803e3d6000fd5b505050505060006113f5308a73ffffffffffffffffffffffffffffffffffffffff166116d290919063ffffffff16565b90506000611422308a73ffffffffffffffffffffffffffffffffffffffff166116d290919063ffffffff16565b9050600061143984836117c290919063ffffffff16565b905087811015611494576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001806123076035913960400191505060405180910390fd5b6114bf33828c73ffffffffffffffffffffffffffffffffffffffff166115a79092919063ffffffff16565b508483111561150757611505336114df87866117c290919063ffffffff16565b8d73ffffffffffffffffffffffffffffffffffffffff166115a79092919063ffffffff16565b505b5050505050505050505050565b61151c610e36565b61158e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61159781611a2d565b50565b600481565b600033905090565b6000808214156115ba5760019050611649565b6115c384611650565b15611614578273ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505015801561160e573d6000803e3d6000fd5b50611648565b61163f83838673ffffffffffffffffffffffffffffffffffffffff16611b719092919063ffffffff16565b60019050611649565b5b9392505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806116cb575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b9050919050565b60006116dd83611650565b15611701578173ffffffffffffffffffffffffffffffffffffffff163190506117bc565b8273ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561177e57600080fd5b505afa158015611792573d6000803e3d6000fd5b505050506040513d60208110156117a857600080fd5b810190808051906020019092919050505090505b92915050565b600061180483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611c42565b905092915050565b600081141561181a576118be565b61182382611650565b1561188f578034111561188a573373ffffffffffffffffffffffffffffffffffffffff166108fc61185d83346117c290919063ffffffff16565b9081150290604051600060405180830381858888f19350505050158015611888573d6000803e3d6000fd5b505b6118bd565b6118bc3330838573ffffffffffffffffffffffffffffffffffffffff16611d02909392919063ffffffff16565b5b5b5050565b6118cb83611650565b611a28576000811180156119ca575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b15801561198d57600080fd5b505afa1580156119a1573d6000803e3d6000fd5b505050506040513d60208110156119b757600080fd5b8101908080519060200190929190505050115b156119fc576119fb8260008573ffffffffffffffffffffffffffffffffffffffff16611e089092919063ffffffff16565b5b611a2782828573ffffffffffffffffffffffffffffffffffffffff16611e089092919063ffffffff16565b5b505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611ab3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806122e16026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611c3d838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb905060e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612028565b505050565b6000838311158290611cef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611cb4578082015181840152602081019050611c99565b50505050905090810190601f168015611ce15780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b611e02848573ffffffffffffffffffffffffffffffffffffffff166323b872dd905060e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612028565b50505050565b6000811480611f02575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b158015611ec557600080fd5b505afa158015611ed9573d6000803e3d6000fd5b505050506040513d6020811015611eef57600080fd5b8101908080519060200190929190505050145b611f57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061239a6036913960400191505060405180910390fd5b612023838473ffffffffffffffffffffffffffffffffffffffff1663095ea7b3905060e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612028565b505050565b6120478273ffffffffffffffffffffffffffffffffffffffff16612273565b6120b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740081525060200191505060405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b6020831061210857805182526020820191506020810190506020830392506120e5565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461216a576040519150601f19603f3d011682016040523d82523d6000602084013e61216f565b606091505b5091509150816121e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656481525060200191505060405180910390fd5b60008151111561226d5780806020019051602081101561220657600080fd5b810190808051906020019092919050505061226c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061233c602a913960400191505060405180910390fd5b5b50505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f91508082141580156122b557506000801b8214155b9250505091905056fe4f6e6553706c69743a20646f206e6f742073656e6420455448206469726563746c794f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f6e6553706c69743a2061637475616c2072657475726e20616d6f756e74206973206c657373207468616e206d696e52657475726e5361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565644f6e6553706c69743a206d73672e76616c75652073686f756c652062652075736564206f6e6c7920666f722045544820737761705361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a265627a7a723158200faba43f4cc875b37fce9291739f04d2ce7071bdfa970c2dad82f02e0604da4964736f6c63430005110032"

// DeployOnesplitaudit deploys a new Ethereum contract, binding an instance of Onesplitaudit to it.
func DeployOnesplitaudit(auth *bind.TransactOpts, backend bind.ContractBackend, impl common.Address) (common.Address, *types.Transaction, *Onesplitaudit, error) {
	parsed, err := abi.JSON(strings.NewReader(OnesplitauditABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OnesplitauditBin), backend, impl)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Onesplitaudit{OnesplitauditCaller: OnesplitauditCaller{contract: contract}, OnesplitauditTransactor: OnesplitauditTransactor{contract: contract}, OnesplitauditFilterer: OnesplitauditFilterer{contract: contract}}, nil
}

// Onesplitaudit is an auto generated Go binding around an Ethereum contract.
type Onesplitaudit struct {
	OnesplitauditCaller     // Read-only binding to the contract
	OnesplitauditTransactor // Write-only binding to the contract
	OnesplitauditFilterer   // Log filterer for contract events
}

// OnesplitauditCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnesplitauditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnesplitauditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnesplitauditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnesplitauditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnesplitauditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnesplitauditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnesplitauditSession struct {
	Contract     *Onesplitaudit    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnesplitauditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnesplitauditCallerSession struct {
	Contract *OnesplitauditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OnesplitauditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnesplitauditTransactorSession struct {
	Contract     *OnesplitauditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OnesplitauditRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnesplitauditRaw struct {
	Contract *Onesplitaudit // Generic contract binding to access the raw methods on
}

// OnesplitauditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnesplitauditCallerRaw struct {
	Contract *OnesplitauditCaller // Generic read-only contract binding to access the raw methods on
}

// OnesplitauditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnesplitauditTransactorRaw struct {
	Contract *OnesplitauditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnesplitaudit creates a new instance of Onesplitaudit, bound to a specific deployed contract.
func NewOnesplitaudit(address common.Address, backend bind.ContractBackend) (*Onesplitaudit, error) {
	contract, err := bindOnesplitaudit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Onesplitaudit{OnesplitauditCaller: OnesplitauditCaller{contract: contract}, OnesplitauditTransactor: OnesplitauditTransactor{contract: contract}, OnesplitauditFilterer: OnesplitauditFilterer{contract: contract}}, nil
}

// NewOnesplitauditCaller creates a new read-only instance of Onesplitaudit, bound to a specific deployed contract.
func NewOnesplitauditCaller(address common.Address, caller bind.ContractCaller) (*OnesplitauditCaller, error) {
	contract, err := bindOnesplitaudit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnesplitauditCaller{contract: contract}, nil
}

// NewOnesplitauditTransactor creates a new write-only instance of Onesplitaudit, bound to a specific deployed contract.
func NewOnesplitauditTransactor(address common.Address, transactor bind.ContractTransactor) (*OnesplitauditTransactor, error) {
	contract, err := bindOnesplitaudit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnesplitauditTransactor{contract: contract}, nil
}

// NewOnesplitauditFilterer creates a new log filterer instance of Onesplitaudit, bound to a specific deployed contract.
func NewOnesplitauditFilterer(address common.Address, filterer bind.ContractFilterer) (*OnesplitauditFilterer, error) {
	contract, err := bindOnesplitaudit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnesplitauditFilterer{contract: contract}, nil
}

// bindOnesplitaudit binds a generic wrapper to an already deployed contract.
func bindOnesplitaudit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnesplitauditABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Onesplitaudit *OnesplitauditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Onesplitaudit.Contract.OnesplitauditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Onesplitaudit *OnesplitauditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.OnesplitauditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Onesplitaudit *OnesplitauditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.OnesplitauditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Onesplitaudit *OnesplitauditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Onesplitaudit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Onesplitaudit *OnesplitauditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Onesplitaudit *OnesplitauditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.contract.Transact(opts, method, params...)
}

// FLAGDISABLEAAVE is a free data retrieval call binding the contract method 0x2e707bd2.
//
// Solidity: function FLAG_DISABLE_AAVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEAAVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_AAVE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEAAVE is a free data retrieval call binding the contract method 0x2e707bd2.
//
// Solidity: function FLAG_DISABLE_AAVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEAAVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEAAVE(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEAAVE is a free data retrieval call binding the contract method 0x2e707bd2.
//
// Solidity: function FLAG_DISABLE_AAVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEAAVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEAAVE(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEBANCOR is a free data retrieval call binding the contract method 0xf56e281f.
//
// Solidity: function FLAG_DISABLE_BANCOR() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEBANCOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_BANCOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEBANCOR is a free data retrieval call binding the contract method 0xf56e281f.
//
// Solidity: function FLAG_DISABLE_BANCOR() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEBANCOR() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEBANCOR(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEBANCOR is a free data retrieval call binding the contract method 0xf56e281f.
//
// Solidity: function FLAG_DISABLE_BANCOR() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEBANCOR() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEBANCOR(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEBDAI is a free data retrieval call binding the contract method 0x75a8b012.
//
// Solidity: function FLAG_DISABLE_BDAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEBDAI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_BDAI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEBDAI is a free data retrieval call binding the contract method 0x75a8b012.
//
// Solidity: function FLAG_DISABLE_BDAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEBDAI() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEBDAI(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEBDAI is a free data retrieval call binding the contract method 0x75a8b012.
//
// Solidity: function FLAG_DISABLE_BDAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEBDAI() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEBDAI(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECHAI is a free data retrieval call binding the contract method 0x34b4dabb.
//
// Solidity: function FLAG_DISABLE_CHAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLECHAI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_CHAI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLECHAI is a free data retrieval call binding the contract method 0x34b4dabb.
//
// Solidity: function FLAG_DISABLE_CHAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLECHAI() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECHAI(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECHAI is a free data retrieval call binding the contract method 0x34b4dabb.
//
// Solidity: function FLAG_DISABLE_CHAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLECHAI() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECHAI(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECOMPOUND is a free data retrieval call binding the contract method 0x44211d62.
//
// Solidity: function FLAG_DISABLE_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLECOMPOUND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_COMPOUND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLECOMPOUND is a free data retrieval call binding the contract method 0x44211d62.
//
// Solidity: function FLAG_DISABLE_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLECOMPOUND() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECOMPOUND(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECOMPOUND is a free data retrieval call binding the contract method 0x44211d62.
//
// Solidity: function FLAG_DISABLE_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLECOMPOUND() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECOMPOUND(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVEBINANCE is a free data retrieval call binding the contract method 0x2113240d.
//
// Solidity: function FLAG_DISABLE_CURVE_BINANCE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLECURVEBINANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_CURVE_BINANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLECURVEBINANCE is a free data retrieval call binding the contract method 0x2113240d.
//
// Solidity: function FLAG_DISABLE_CURVE_BINANCE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLECURVEBINANCE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVEBINANCE(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVEBINANCE is a free data retrieval call binding the contract method 0x2113240d.
//
// Solidity: function FLAG_DISABLE_CURVE_BINANCE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLECURVEBINANCE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVEBINANCE(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVECOMPOUND is a free data retrieval call binding the contract method 0xb0a7ef29.
//
// Solidity: function FLAG_DISABLE_CURVE_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLECURVECOMPOUND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_CURVE_COMPOUND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLECURVECOMPOUND is a free data retrieval call binding the contract method 0xb0a7ef29.
//
// Solidity: function FLAG_DISABLE_CURVE_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLECURVECOMPOUND() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVECOMPOUND(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVECOMPOUND is a free data retrieval call binding the contract method 0xb0a7ef29.
//
// Solidity: function FLAG_DISABLE_CURVE_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLECURVECOMPOUND() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVECOMPOUND(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVESYNTHETIX is a free data retrieval call binding the contract method 0xc9b42c67.
//
// Solidity: function FLAG_DISABLE_CURVE_SYNTHETIX() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLECURVESYNTHETIX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_CURVE_SYNTHETIX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLECURVESYNTHETIX is a free data retrieval call binding the contract method 0xc9b42c67.
//
// Solidity: function FLAG_DISABLE_CURVE_SYNTHETIX() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLECURVESYNTHETIX() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVESYNTHETIX(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVESYNTHETIX is a free data retrieval call binding the contract method 0xc9b42c67.
//
// Solidity: function FLAG_DISABLE_CURVE_SYNTHETIX() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLECURVESYNTHETIX() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVESYNTHETIX(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVEUSDT is a free data retrieval call binding the contract method 0x13989140.
//
// Solidity: function FLAG_DISABLE_CURVE_USDT() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLECURVEUSDT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_CURVE_USDT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLECURVEUSDT is a free data retrieval call binding the contract method 0x13989140.
//
// Solidity: function FLAG_DISABLE_CURVE_USDT() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLECURVEUSDT() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVEUSDT(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVEUSDT is a free data retrieval call binding the contract method 0x13989140.
//
// Solidity: function FLAG_DISABLE_CURVE_USDT() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLECURVEUSDT() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVEUSDT(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVEY is a free data retrieval call binding the contract method 0x5aa8fb48.
//
// Solidity: function FLAG_DISABLE_CURVE_Y() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLECURVEY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_CURVE_Y")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLECURVEY is a free data retrieval call binding the contract method 0x5aa8fb48.
//
// Solidity: function FLAG_DISABLE_CURVE_Y() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLECURVEY() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVEY(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLECURVEY is a free data retrieval call binding the contract method 0x5aa8fb48.
//
// Solidity: function FLAG_DISABLE_CURVE_Y() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLECURVEY() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLECURVEY(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEFULCRUM is a free data retrieval call binding the contract method 0x4a7101d5.
//
// Solidity: function FLAG_DISABLE_FULCRUM() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEFULCRUM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_FULCRUM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEFULCRUM is a free data retrieval call binding the contract method 0x4a7101d5.
//
// Solidity: function FLAG_DISABLE_FULCRUM() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEFULCRUM() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEFULCRUM(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEFULCRUM is a free data retrieval call binding the contract method 0x4a7101d5.
//
// Solidity: function FLAG_DISABLE_FULCRUM() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEFULCRUM() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEFULCRUM(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEIEARN is a free data retrieval call binding the contract method 0x5ae51b82.
//
// Solidity: function FLAG_DISABLE_IEARN() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEIEARN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_IEARN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEIEARN is a free data retrieval call binding the contract method 0x5ae51b82.
//
// Solidity: function FLAG_DISABLE_IEARN() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEIEARN() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEIEARN(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEIEARN is a free data retrieval call binding the contract method 0x5ae51b82.
//
// Solidity: function FLAG_DISABLE_IEARN() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEIEARN() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEIEARN(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEKYBER is a free data retrieval call binding the contract method 0x7a88bdbd.
//
// Solidity: function FLAG_DISABLE_KYBER() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEKYBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_KYBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEKYBER is a free data retrieval call binding the contract method 0x7a88bdbd.
//
// Solidity: function FLAG_DISABLE_KYBER() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEKYBER() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEKYBER(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEKYBER is a free data retrieval call binding the contract method 0x7a88bdbd.
//
// Solidity: function FLAG_DISABLE_KYBER() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEKYBER() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEKYBER(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEOASIS is a free data retrieval call binding the contract method 0x5c0cb479.
//
// Solidity: function FLAG_DISABLE_OASIS() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEOASIS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_OASIS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEOASIS is a free data retrieval call binding the contract method 0x5c0cb479.
//
// Solidity: function FLAG_DISABLE_OASIS() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEOASIS() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEOASIS(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEOASIS is a free data retrieval call binding the contract method 0x5c0cb479.
//
// Solidity: function FLAG_DISABLE_OASIS() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEOASIS() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEOASIS(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLESMARTTOKEN is a free data retrieval call binding the contract method 0xdc1536b2.
//
// Solidity: function FLAG_DISABLE_SMART_TOKEN() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLESMARTTOKEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_SMART_TOKEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLESMARTTOKEN is a free data retrieval call binding the contract method 0xdc1536b2.
//
// Solidity: function FLAG_DISABLE_SMART_TOKEN() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLESMARTTOKEN() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLESMARTTOKEN(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLESMARTTOKEN is a free data retrieval call binding the contract method 0xdc1536b2.
//
// Solidity: function FLAG_DISABLE_SMART_TOKEN() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLESMARTTOKEN() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLESMARTTOKEN(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEUNISWAP is a free data retrieval call binding the contract method 0xc762a46c.
//
// Solidity: function FLAG_DISABLE_UNISWAP() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEUNISWAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_UNISWAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEUNISWAP is a free data retrieval call binding the contract method 0xc762a46c.
//
// Solidity: function FLAG_DISABLE_UNISWAP() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEUNISWAP() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEUNISWAP(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEUNISWAP is a free data retrieval call binding the contract method 0xc762a46c.
//
// Solidity: function FLAG_DISABLE_UNISWAP() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEUNISWAP() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEUNISWAP(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEWETH is a free data retrieval call binding the contract method 0x6cbc4a6e.
//
// Solidity: function FLAG_DISABLE_WETH() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGDISABLEWETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_DISABLE_WETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGDISABLEWETH is a free data retrieval call binding the contract method 0x6cbc4a6e.
//
// Solidity: function FLAG_DISABLE_WETH() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGDISABLEWETH() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEWETH(&_Onesplitaudit.CallOpts)
}

// FLAGDISABLEWETH is a free data retrieval call binding the contract method 0x6cbc4a6e.
//
// Solidity: function FLAG_DISABLE_WETH() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGDISABLEWETH() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGDISABLEWETH(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEKYBERBANCORRESERVE is a free data retrieval call binding the contract method 0xb3bc7844.
//
// Solidity: function FLAG_ENABLE_KYBER_BANCOR_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGENABLEKYBERBANCORRESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_ENABLE_KYBER_BANCOR_RESERVE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGENABLEKYBERBANCORRESERVE is a free data retrieval call binding the contract method 0xb3bc7844.
//
// Solidity: function FLAG_ENABLE_KYBER_BANCOR_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGENABLEKYBERBANCORRESERVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEKYBERBANCORRESERVE(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEKYBERBANCORRESERVE is a free data retrieval call binding the contract method 0xb3bc7844.
//
// Solidity: function FLAG_ENABLE_KYBER_BANCOR_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGENABLEKYBERBANCORRESERVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEKYBERBANCORRESERVE(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEKYBEROASISRESERVE is a free data retrieval call binding the contract method 0x21a360f5.
//
// Solidity: function FLAG_ENABLE_KYBER_OASIS_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGENABLEKYBEROASISRESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_ENABLE_KYBER_OASIS_RESERVE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGENABLEKYBEROASISRESERVE is a free data retrieval call binding the contract method 0x21a360f5.
//
// Solidity: function FLAG_ENABLE_KYBER_OASIS_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGENABLEKYBEROASISRESERVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEKYBEROASISRESERVE(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEKYBEROASISRESERVE is a free data retrieval call binding the contract method 0x21a360f5.
//
// Solidity: function FLAG_ENABLE_KYBER_OASIS_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGENABLEKYBEROASISRESERVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEKYBEROASISRESERVE(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEKYBERUNISWAPRESERVE is a free data retrieval call binding the contract method 0x2d3b5207.
//
// Solidity: function FLAG_ENABLE_KYBER_UNISWAP_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGENABLEKYBERUNISWAPRESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_ENABLE_KYBER_UNISWAP_RESERVE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGENABLEKYBERUNISWAPRESERVE is a free data retrieval call binding the contract method 0x2d3b5207.
//
// Solidity: function FLAG_ENABLE_KYBER_UNISWAP_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGENABLEKYBERUNISWAPRESERVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEKYBERUNISWAPRESERVE(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEKYBERUNISWAPRESERVE is a free data retrieval call binding the contract method 0x2d3b5207.
//
// Solidity: function FLAG_ENABLE_KYBER_UNISWAP_RESERVE() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGENABLEKYBERUNISWAPRESERVE() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEKYBERUNISWAPRESERVE(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEMULTIPATHDAI is a free data retrieval call binding the contract method 0xd393c3e9.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_DAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGENABLEMULTIPATHDAI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_ENABLE_MULTI_PATH_DAI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGENABLEMULTIPATHDAI is a free data retrieval call binding the contract method 0xd393c3e9.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_DAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGENABLEMULTIPATHDAI() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEMULTIPATHDAI(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEMULTIPATHDAI is a free data retrieval call binding the contract method 0xd393c3e9.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_DAI() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGENABLEMULTIPATHDAI() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEMULTIPATHDAI(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEMULTIPATHETH is a free data retrieval call binding the contract method 0xc77b9de6.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_ETH() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGENABLEMULTIPATHETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_ENABLE_MULTI_PATH_ETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGENABLEMULTIPATHETH is a free data retrieval call binding the contract method 0xc77b9de6.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_ETH() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGENABLEMULTIPATHETH() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEMULTIPATHETH(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEMULTIPATHETH is a free data retrieval call binding the contract method 0xc77b9de6.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_ETH() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGENABLEMULTIPATHETH() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEMULTIPATHETH(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEMULTIPATHUSDC is a free data retrieval call binding the contract method 0x64ec4e5c.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_USDC() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGENABLEMULTIPATHUSDC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_ENABLE_MULTI_PATH_USDC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGENABLEMULTIPATHUSDC is a free data retrieval call binding the contract method 0x64ec4e5c.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_USDC() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGENABLEMULTIPATHUSDC() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEMULTIPATHUSDC(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEMULTIPATHUSDC is a free data retrieval call binding the contract method 0x64ec4e5c.
//
// Solidity: function FLAG_ENABLE_MULTI_PATH_USDC() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGENABLEMULTIPATHUSDC() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEMULTIPATHUSDC(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEUNISWAPCOMPOUND is a free data retrieval call binding the contract method 0x68e2a014.
//
// Solidity: function FLAG_ENABLE_UNISWAP_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCaller) FLAGENABLEUNISWAPCOMPOUND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "FLAG_ENABLE_UNISWAP_COMPOUND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLAGENABLEUNISWAPCOMPOUND is a free data retrieval call binding the contract method 0x68e2a014.
//
// Solidity: function FLAG_ENABLE_UNISWAP_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditSession) FLAGENABLEUNISWAPCOMPOUND() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEUNISWAPCOMPOUND(&_Onesplitaudit.CallOpts)
}

// FLAGENABLEUNISWAPCOMPOUND is a free data retrieval call binding the contract method 0x68e2a014.
//
// Solidity: function FLAG_ENABLE_UNISWAP_COMPOUND() view returns(uint256)
func (_Onesplitaudit *OnesplitauditCallerSession) FLAGENABLEUNISWAPCOMPOUND() (*big.Int, error) {
	return _Onesplitaudit.Contract.FLAGENABLEUNISWAPCOMPOUND(&_Onesplitaudit.CallOpts)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address toToken, uint256 amount, uint256 parts, uint256 featureFlags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Onesplitaudit *OnesplitauditCaller) GetExpectedReturn(opts *bind.CallOpts, fromToken common.Address, toToken common.Address, amount *big.Int, parts *big.Int, featureFlags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "getExpectedReturn", fromToken, toToken, amount, parts, featureFlags)

	outstruct := new(struct {
		ReturnAmount *big.Int
		Distribution []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReturnAmount = out[0].(*big.Int)
	outstruct.Distribution = out[1].([]*big.Int)

	return *outstruct, err

}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address toToken, uint256 amount, uint256 parts, uint256 featureFlags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Onesplitaudit *OnesplitauditSession) GetExpectedReturn(fromToken common.Address, toToken common.Address, amount *big.Int, parts *big.Int, featureFlags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _Onesplitaudit.Contract.GetExpectedReturn(&_Onesplitaudit.CallOpts, fromToken, toToken, amount, parts, featureFlags)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address toToken, uint256 amount, uint256 parts, uint256 featureFlags) view returns(uint256 returnAmount, uint256[] distribution)
func (_Onesplitaudit *OnesplitauditCallerSession) GetExpectedReturn(fromToken common.Address, toToken common.Address, amount *big.Int, parts *big.Int, featureFlags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _Onesplitaudit.Contract.GetExpectedReturn(&_Onesplitaudit.CallOpts, fromToken, toToken, amount, parts, featureFlags)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Onesplitaudit *OnesplitauditCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Onesplitaudit *OnesplitauditSession) IsOwner() (bool, error) {
	return _Onesplitaudit.Contract.IsOwner(&_Onesplitaudit.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Onesplitaudit *OnesplitauditCallerSession) IsOwner() (bool, error) {
	return _Onesplitaudit.Contract.IsOwner(&_Onesplitaudit.CallOpts)
}

// OneSplitImpl is a free data retrieval call binding the contract method 0x867807ca.
//
// Solidity: function oneSplitImpl() view returns(address)
func (_Onesplitaudit *OnesplitauditCaller) OneSplitImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "oneSplitImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OneSplitImpl is a free data retrieval call binding the contract method 0x867807ca.
//
// Solidity: function oneSplitImpl() view returns(address)
func (_Onesplitaudit *OnesplitauditSession) OneSplitImpl() (common.Address, error) {
	return _Onesplitaudit.Contract.OneSplitImpl(&_Onesplitaudit.CallOpts)
}

// OneSplitImpl is a free data retrieval call binding the contract method 0x867807ca.
//
// Solidity: function oneSplitImpl() view returns(address)
func (_Onesplitaudit *OnesplitauditCallerSession) OneSplitImpl() (common.Address, error) {
	return _Onesplitaudit.Contract.OneSplitImpl(&_Onesplitaudit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Onesplitaudit *OnesplitauditCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Onesplitaudit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Onesplitaudit *OnesplitauditSession) Owner() (common.Address, error) {
	return _Onesplitaudit.Contract.Owner(&_Onesplitaudit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Onesplitaudit *OnesplitauditCallerSession) Owner() (common.Address, error) {
	return _Onesplitaudit.Contract.Owner(&_Onesplitaudit.CallOpts)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xba4917b3.
//
// Solidity: function claimAsset(address asset, uint256 amount) returns()
func (_Onesplitaudit *OnesplitauditTransactor) ClaimAsset(opts *bind.TransactOpts, asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Onesplitaudit.contract.Transact(opts, "claimAsset", asset, amount)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xba4917b3.
//
// Solidity: function claimAsset(address asset, uint256 amount) returns()
func (_Onesplitaudit *OnesplitauditSession) ClaimAsset(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.ClaimAsset(&_Onesplitaudit.TransactOpts, asset, amount)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xba4917b3.
//
// Solidity: function claimAsset(address asset, uint256 amount) returns()
func (_Onesplitaudit *OnesplitauditTransactorSession) ClaimAsset(asset common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.ClaimAsset(&_Onesplitaudit.TransactOpts, asset, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Onesplitaudit *OnesplitauditTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onesplitaudit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Onesplitaudit *OnesplitauditSession) RenounceOwnership() (*types.Transaction, error) {
	return _Onesplitaudit.Contract.RenounceOwnership(&_Onesplitaudit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Onesplitaudit *OnesplitauditTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Onesplitaudit.Contract.RenounceOwnership(&_Onesplitaudit.TransactOpts)
}

// SetNewImpl is a paid mutator transaction binding the contract method 0xb26413f8.
//
// Solidity: function setNewImpl(address impl) returns()
func (_Onesplitaudit *OnesplitauditTransactor) SetNewImpl(opts *bind.TransactOpts, impl common.Address) (*types.Transaction, error) {
	return _Onesplitaudit.contract.Transact(opts, "setNewImpl", impl)
}

// SetNewImpl is a paid mutator transaction binding the contract method 0xb26413f8.
//
// Solidity: function setNewImpl(address impl) returns()
func (_Onesplitaudit *OnesplitauditSession) SetNewImpl(impl common.Address) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.SetNewImpl(&_Onesplitaudit.TransactOpts, impl)
}

// SetNewImpl is a paid mutator transaction binding the contract method 0xb26413f8.
//
// Solidity: function setNewImpl(address impl) returns()
func (_Onesplitaudit *OnesplitauditTransactorSession) SetNewImpl(impl common.Address) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.SetNewImpl(&_Onesplitaudit.TransactOpts, impl)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Onesplitaudit *OnesplitauditTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Onesplitaudit.contract.Transact(opts, "swap", fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Onesplitaudit *OnesplitauditSession) Swap(fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.Swap(&_Onesplitaudit.TransactOpts, fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Onesplitaudit *OnesplitauditTransactorSession) Swap(fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.Swap(&_Onesplitaudit.TransactOpts, fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Onesplitaudit *OnesplitauditTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Onesplitaudit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Onesplitaudit *OnesplitauditSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.TransferOwnership(&_Onesplitaudit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Onesplitaudit *OnesplitauditTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.TransferOwnership(&_Onesplitaudit.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Onesplitaudit *OnesplitauditTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Onesplitaudit.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Onesplitaudit *OnesplitauditSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.Fallback(&_Onesplitaudit.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Onesplitaudit *OnesplitauditTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Onesplitaudit.Contract.Fallback(&_Onesplitaudit.TransactOpts, calldata)
}

// OnesplitauditImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the Onesplitaudit contract.
type OnesplitauditImplementationUpdatedIterator struct {
	Event *OnesplitauditImplementationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnesplitauditImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnesplitauditImplementationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnesplitauditImplementationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnesplitauditImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnesplitauditImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnesplitauditImplementationUpdated represents a ImplementationUpdated event raised by the Onesplitaudit contract.
type OnesplitauditImplementationUpdated struct {
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address indexed newImpl)
func (_Onesplitaudit *OnesplitauditFilterer) FilterImplementationUpdated(opts *bind.FilterOpts, newImpl []common.Address) (*OnesplitauditImplementationUpdatedIterator, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _Onesplitaudit.contract.FilterLogs(opts, "ImplementationUpdated", newImplRule)
	if err != nil {
		return nil, err
	}
	return &OnesplitauditImplementationUpdatedIterator{contract: _Onesplitaudit.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address indexed newImpl)
func (_Onesplitaudit *OnesplitauditFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *OnesplitauditImplementationUpdated, newImpl []common.Address) (event.Subscription, error) {

	var newImplRule []interface{}
	for _, newImplItem := range newImpl {
		newImplRule = append(newImplRule, newImplItem)
	}

	logs, sub, err := _Onesplitaudit.contract.WatchLogs(opts, "ImplementationUpdated", newImplRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnesplitauditImplementationUpdated)
				if err := _Onesplitaudit.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseImplementationUpdated is a log parse operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address indexed newImpl)
func (_Onesplitaudit *OnesplitauditFilterer) ParseImplementationUpdated(log types.Log) (*OnesplitauditImplementationUpdated, error) {
	event := new(OnesplitauditImplementationUpdated)
	if err := _Onesplitaudit.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OnesplitauditOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Onesplitaudit contract.
type OnesplitauditOwnershipTransferredIterator struct {
	Event *OnesplitauditOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OnesplitauditOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnesplitauditOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OnesplitauditOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OnesplitauditOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnesplitauditOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnesplitauditOwnershipTransferred represents a OwnershipTransferred event raised by the Onesplitaudit contract.
type OnesplitauditOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Onesplitaudit *OnesplitauditFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OnesplitauditOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Onesplitaudit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OnesplitauditOwnershipTransferredIterator{contract: _Onesplitaudit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Onesplitaudit *OnesplitauditFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OnesplitauditOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Onesplitaudit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnesplitauditOwnershipTransferred)
				if err := _Onesplitaudit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Onesplitaudit *OnesplitauditFilterer) ParseOwnershipTransferred(log types.Log) (*OnesplitauditOwnershipTransferred, error) {
	event := new(OnesplitauditOwnershipTransferred)
	if err := _Onesplitaudit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
