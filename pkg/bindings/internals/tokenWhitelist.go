// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internals

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenWhitelistABI is the input ABI used to generate the binding from.
const TokenWhitelistABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"getTokenInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"},{\"name\":\"_symbols\",\"type\":\"bytes32[]\"},{\"name\":\"_magnitude\",\"type\":\"uint256[]\"},{\"name\":\"_loadable\",\"type\":\"bool[]\"},{\"name\":\"_burnable\",\"type\":\"bool[]\"},{\"name\":\"_lastUpdate\",\"type\":\"uint256\"}],\"name\":\"addTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnableCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStablecoinInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddressArray\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ensRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracleNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_transferable\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnableTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_updateDate\",\"type\":\"uint256\"}],\"name\":\"updateTokenRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controllerNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stablecoin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens_\",\"type\":\"address\"},{\"name\":\"_oracleNameHash_\",\"type\":\"bytes32\"},{\"name\":\"_controllerNameHash_\",\"type\":\"bytes32\"},{\"name\":\"_owner_\",\"type\":\"address\"},{\"name\":\"_transferable_\",\"type\":\"bool\"},{\"name\":\"_stabelcoinAddress_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"UpdatedTokenRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_magnitude\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_loadable\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_burnable\",\"type\":\"bool\"}],\"name\":\"AddedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RemovedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"TransferredOwnership\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_locked\",\"type\":\"address\"}],\"name\":\"LockedOwnership\",\"type\":\"event\"}]"

// TokenWhitelistBin is the compiled bytecode used for deploying new contracts.
const TokenWhitelistBin = `608060405234801561001057600080fd5b5060405160c080611e84833981018060405260c081101561003057600080fd5b508051602082015160408301516060840151608085015160a090950151600180546001600160a01b038088166001600160a01b031992831617928390556000805483169382169390931790925560028590556003805489151574010000000000000000000000000000000000000000908102600160a01b60ff021995881692909416919091179390931691909117908190559596949593949293849184910460ff1661011357604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b60408051600081526001600160a01b038416602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a150506008949094555050600780546001600160a01b0319166001600160a01b039093169290921790915550611cf19050806101936000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80637d73b231116100a2578063b242e53411610071578063b242e53414610506578063bf25811f14610534578063d545782e1461053c578063e2b4ce971461056e578063e9cbd822146105765761010b565b80637d73b2311461049c57806380cc0dda146104c05780638da5cb5b146104c8578063996cba68146104d05761010b565b80633efec5e9116100de5780633efec5e9146103c6578063443dd2a4146103ce5780636c3824ef14610426578063715018a6146104945761010b565b80631f69565f146101105780632121dc75146101e057806334c73edc146101fc5780633da9a430146103ac575b600080fd5b6101366004803603602081101561012657600080fd5b50356001600160a01b031661057e565b6040805160208082018990529181018790528515156060820152841515608082015283151560a082015260c0810183905260e08082528951908201528851909182916101008301918b019080838360005b8381101561019f578181015183820152602001610187565b50505050905090810190601f1680156101cc5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b6101e861067a565b604080519115158252519081900360200190f35b6103aa600480360360c081101561021257600080fd5b810190602081018135600160201b81111561022c57600080fd5b82018360208201111561023e57600080fd5b803590602001918460208302840111600160201b8311171561025f57600080fd5b919390929091602081019035600160201b81111561027c57600080fd5b82018360208201111561028e57600080fd5b803590602001918460208302840111600160201b831117156102af57600080fd5b919390929091602081019035600160201b8111156102cc57600080fd5b8201836020820111156102de57600080fd5b803590602001918460208302840111600160201b831117156102ff57600080fd5b919390929091602081019035600160201b81111561031c57600080fd5b82018360208201111561032e57600080fd5b803590602001918460208302840111600160201b8311171561034f57600080fd5b919390929091602081019035600160201b81111561036c57600080fd5b82018360208201111561037e57600080fd5b803590602001918460208302840111600160201b8311171561039f57600080fd5b91935091503561068b565b005b6103b4610b2d565b60408051918252519081900360200190f35b610136610b33565b6103d6610c2f565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156104125781810151838201526020016103fa565b505050509050019250505060405180910390f35b6103aa6004803603602081101561043c57600080fd5b810190602081018135600160201b81111561045657600080fd5b82018360208201111561046857600080fd5b803590602001918460208302840111600160201b8311171561048957600080fd5b509092509050610c91565b6103aa610f48565b6104a4611053565b604080516001600160a01b039092168252519081900360200190f35b6103b4611062565b6104a4611068565b6103aa600480360360608110156104e657600080fd5b506001600160a01b03813581169160208101359091169060400135611077565b6103aa6004803603604081101561051c57600080fd5b506001600160a01b03813516906020013515156110df565b6103d66112ba565b6103aa6004803603606081101561055257600080fd5b506001600160a01b038135169060208101359060400135611388565b6103b46114e6565b6104a46114ec565b6001600160a01b0381166000908152600460208181526040808420600180820154600280840154600385015497850154855487516101009682161587026000190190911693909304601f81018990048902840189019097528683526060999889988998899889988998919788979296929560ff80821696938204811695620100009092041693928991908301828280156106595780601f1061062e57610100808354040283529160200191610659565b820191906000526020600020905b81548152906001019060200180831161063c57829003601f168201915b50505050509650975097509750975097509750975050919395979092949650565b600354600160a01b900460ff165b90565b610694336114fb565b6106e85760408051600160e51b62461bcd02815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b89881480156106f657508986145b801561070157508984145b801561070c57508984145b6107605760408051600160e51b62461bcd02815260206004820152601e60248201527f706172616d65746572206c656e6774687320646f206e6f74206d617463680000604482015290519081900360640190fd5b60005b8a811015610b1f57600460008d8d8481811061077b57fe5b602090810292909201356001600160a01b03168352508101919091526040016000206003015460ff16156107f95760408051600160e51b62461bcd02815260206004820152601760248201527f746f6b656e20616c726561647920617661696c61626c65000000000000000000604482015290519081900360640190fd5b606061081e6108198c8c8581811061080d57fe5b90506020020135611591565b6115bd565b90506040518060e001604052808281526020018a8a8581811061083d57fe5b9050602002013581526020016000815260200160011515815260200188888581811061086557fe5b9050602002013515151515815260200186868581811061088157fe5b9050602002013515151515815260200184815250600460008f8f868181106108a557fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060008201518160000190805190602001906108f1929190611b5b565b5060208201516001820155604082015160028201556060820151600382018054608085015160a08601511515620100000262ff0000199115156101000261ff001995151560ff199094169390931794909416919091171691909117905560c09091015160049091015560058d8d8481811061096857fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055508484838181106109b057fe5b90506020020135156109d4576006546109d090600163ffffffff61161416565b6006555b7f1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b338e8e85818110610a0257fe5b905060200201356001600160a01b0316838c8c87818110610a1f57fe5b905060200201358b8b88818110610a3257fe5b9050602002013515158a8a89818110610a4757fe5b90506020020135151560405180876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b03168152602001806020018581526020018415151515815260200183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015610ad7578181015183820152602001610abf565b50505050905090810190601f168015610b045780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150600101610763565b505050505050505050505050565b60065490565b6007546001600160a01b03166000908152600460208181526040808420600180820154600280840154600385015497850154855487516101009682161587026000190190911693909304601f81018990048902840189019097528683526060999889988998899889988998919788979296929560ff8082169693820481169562010000909204169392899190830182828015610c105780601f10610be557610100808354040283529160200191610c10565b820191906000526020600020905b815481529060010190602001808311610bf357829003601f168201915b5050505050965097509750975097509750975097505090919293949596565b60606005805480602002602001604051908101604052809291908181526020018280548015610c8757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610c69575b5050505050905090565b610c9a336114fb565b610cee5760408051600160e51b62461bcd02815260206004820152601a60248201527f73656e646572206973206e6f74206120636f6e74726f6c6c6572000000000000604482015290519081900360640190fd5b60005b81811015610f43576000838383818110610d0757fe5b602090810292909201356001600160a01b0316600081815260049093526040909220600301549192505060ff16610d885760408051600160e51b62461bcd02815260206004820152601660248201527f746f6b656e206973206e6f7420617661696c61626c6500000000000000000000604482015290519081900360640190fd5b6001600160a01b03811660009081526004602052604090206003015462010000900460ff1615610dca57600654610dc690600163ffffffff61167816565b6006555b6001600160a01b038116600090815260046020526040812090610ded8282611bd9565b506000600182018190556002820181905560038201805462ffffff1916905560049091018190555b600554610e2990600163ffffffff61167816565b811015610ee257816001600160a01b031660058281548110610e4757fe5b6000918252602090912001546001600160a01b03161415610eda5760058054610e7790600163ffffffff61167816565b81548110610e8157fe5b600091825260209091200154600580546001600160a01b039092169183908110610ea757fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550610ee2565b600101610e15565b506005805490610ef6906000198301611c20565b50604080513381526001600160a01b038316602082015281517f703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0929181900390910190a150600101610cf1565b505050565b610f51336116d8565b610fa55760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b600354600160a01b900460ff166110065760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b600380546001600160a01b0319169055604080516000808252602082015281517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5929181900390910190a1565b6001546001600160a01b031690565b60085490565b6003546001600160a01b031690565b611080336116d8565b6110d45760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b610f438383836116ec565b6110e8336116d8565b61113c5760408051600160e51b62461bcd02815260206004820152601660248201527f73656e646572206973206e6f7420616e206f776e657200000000000000000000604482015290519081900360640190fd5b600354600160a01b900460ff1661119d5760408051600160e51b62461bcd02815260206004820152601d60248201527f6f776e657273686970206973206e6f74207472616e7366657261626c65000000604482015290519081900360640190fd5b6001600160a01b0382166111e557604051600160e51b62461bcd028152600401808060200182810382526023815260200180611c796023913960400191505060405180910390fd5b6003805474ff00000000000000000000000000000000000000001916600160a01b831515021790558061124f57604080516001600160a01b038416815290517f808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec1229181900360200190a15b600354604080516001600160a01b039283168152918416602083015280517f850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea59281900390910190a150600380546001600160a01b0319166001600160a01b0392909216919091179055565b6060806006546040519080825280602002602001820160405280156112e9578160200160208202803883390190505b5090506000805b6005548110156113805760006005828154811061130957fe5b60009182526020808320909101546001600160a01b0316808352600490915260409091206003015490915060ff620100009091041615611377578084848151811061135057fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506001830192505b506001016112f0565b509091505090565b600061139560085461179f565b90506113a0336114fb565b806113b35750336001600160a01b038216145b6114075760408051600160e51b62461bcd02815260206004820152601b60248201527f656974686572206f7261636c65206f7220636f6e74726f6c6c65720000000000604482015290519081900360640190fd5b6001600160a01b03841660009081526004602052604090206003015460ff1661147a5760408051600160e51b62461bcd02815260206004820152601660248201527f746f6b656e206973206e6f7420617661696c61626c6500000000000000000000604482015290519081900360640190fd5b6001600160a01b03841660008181526004602081815260409283902060028101889055909101859055815133815290810192909252818101859052517fdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d9181900360600190a150505050565b60025490565b6007546001600160a01b031690565b600061150860025461179f565b6001600160a01b031663b429afeb836040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561155d57600080fd5b505afa158015611571573d6000803e3d6000fd5b505050506040513d602081101561158757600080fd5b505190505b919050565b611599611c44565b60405160208101604052828152806020830152506115b682611867565b8152919050565b60608082600001516040519080825280601f01601f1916602001820160405280156115ef576020820181803883390190505b509050600060208201905061160d8185602001518660000151611900565b5092915050565b6000828201838110156116715760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6000828211156116d25760408051600160e51b62461bcd02815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6003546001600160a01b0390811691161490565b6001600160a01b038216611736576040516001600160a01b0384169082156108fc029083906000818181858888f19350505050158015611730573d6000803e3d6000fd5b50611750565b6117506001600160a01b038316848363ffffffff61193e16565b604080516001600160a01b0380861682528416602082015280820183905290517ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd39926839181900360600190a1505050565b6000805460408051600160e01b630178b8bf0281526004810185905290516001600160a01b0390921691630178b8bf91602480820192602092909190829003018186803b1580156117ef57600080fd5b505afa158015611803573d6000803e3d6000fd5b505050506040513d602081101561181957600080fd5b505160408051600160e11b631d9dabef0281526004810185905290516001600160a01b0390921691633b3b57de91602480820192602092909190829003018186803b15801561155d57600080fd5b6000808261187957600091505061158c565b6fffffffffffffffffffffffffffffffff831661189d57601001600160801b830492505b67ffffffffffffffff83166118be5760080168010000000000000000830492505b63ffffffff83166118d657600401600160201b830492505b61ffff83166118eb5760020162010000830492505b60ff83166118f7576001015b60200392915050565b5b60208110611920578151835260209283019290910190601f1901611901565b905182516020929092036101000a6000190180199091169116179052565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b0316600160e01b63a9059cbb02179052610f439084906119a0826001600160a01b0316611b55565b6119f45760408051600160e51b62461bcd02815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310611a325780518252601f199092019160209182019101611a13565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611a94576040519150601f19603f3d011682016040523d82523d6000602084013e611a99565b606091505b509150915081611af35760408051600160e51b62461bcd02815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115611b4f57808060200190516020811015611b0f57600080fd5b5051611b4f57604051600160e51b62461bcd02815260040180806020018281038252602a815260200180611c9c602a913960400191505060405180910390fd5b50505050565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611b9c57805160ff1916838001178555611bc9565b82800160010185558215611bc9579182015b82811115611bc9578251825591602001919060010190611bae565b50611bd5929150611c5e565b5090565b50805460018160011615610100020316600290046000825580601f10611bff5750611c1d565b601f016020900490600052602060002090810190611c1d9190611c5e565b50565b815481835581811115610f4357600083815260209020610f43918101908301611c5e565b604051806040016040528060008152602001600081525090565b61068891905b80821115611bd55760008155600101611c6456fe6f776e65722063616e6e6f742062652073657420746f207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a165627a7a723058208f7e3393418f54904356484aa42c1e1de972e9670877ca7c872d98136b72ce600029`

// DeployTokenWhitelist deploys a new Ethereum contract, binding an instance of TokenWhitelist to it.
func DeployTokenWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend, _ens_ common.Address, _oracleNameHash_ [32]byte, _controllerNameHash_ [32]byte, _owner_ common.Address, _transferable_ bool, _stabelcoinAddress_ common.Address) (common.Address, *types.Transaction, *TokenWhitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenWhitelistBin), backend, _ens_, _oracleNameHash_, _controllerNameHash_, _owner_, _transferable_, _stabelcoinAddress_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenWhitelist{TokenWhitelistCaller: TokenWhitelistCaller{contract: contract}, TokenWhitelistTransactor: TokenWhitelistTransactor{contract: contract}, TokenWhitelistFilterer: TokenWhitelistFilterer{contract: contract}}, nil
}

// TokenWhitelist is an auto generated Go binding around an Ethereum contract.
type TokenWhitelist struct {
	TokenWhitelistCaller     // Read-only binding to the contract
	TokenWhitelistTransactor // Write-only binding to the contract
	TokenWhitelistFilterer   // Log filterer for contract events
}

// TokenWhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenWhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenWhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenWhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenWhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenWhitelistSession struct {
	Contract     *TokenWhitelist   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenWhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenWhitelistCallerSession struct {
	Contract *TokenWhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenWhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenWhitelistTransactorSession struct {
	Contract     *TokenWhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenWhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenWhitelistRaw struct {
	Contract *TokenWhitelist // Generic contract binding to access the raw methods on
}

// TokenWhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenWhitelistCallerRaw struct {
	Contract *TokenWhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// TokenWhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenWhitelistTransactorRaw struct {
	Contract *TokenWhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenWhitelist creates a new instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelist(address common.Address, backend bind.ContractBackend) (*TokenWhitelist, error) {
	contract, err := bindTokenWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelist{TokenWhitelistCaller: TokenWhitelistCaller{contract: contract}, TokenWhitelistTransactor: TokenWhitelistTransactor{contract: contract}, TokenWhitelistFilterer: TokenWhitelistFilterer{contract: contract}}, nil
}

// NewTokenWhitelistCaller creates a new read-only instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistCaller(address common.Address, caller bind.ContractCaller) (*TokenWhitelistCaller, error) {
	contract, err := bindTokenWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistCaller{contract: contract}, nil
}

// NewTokenWhitelistTransactor creates a new write-only instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenWhitelistTransactor, error) {
	contract, err := bindTokenWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistTransactor{contract: contract}, nil
}

// NewTokenWhitelistFilterer creates a new log filterer instance of TokenWhitelist, bound to a specific deployed contract.
func NewTokenWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenWhitelistFilterer, error) {
	contract, err := bindTokenWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistFilterer{contract: contract}, nil
}

// bindTokenWhitelist binds a generic wrapper to an already deployed contract.
func bindTokenWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenWhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelist *TokenWhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelist.Contract.TokenWhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelist *TokenWhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenWhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelist *TokenWhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TokenWhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenWhitelist *TokenWhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenWhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenWhitelist *TokenWhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenWhitelist *TokenWhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.contract.Transact(opts, method, params...)
}

// BurnableCounter is a free data retrieval call binding the contract method 0x3da9a430.
//
// Solidity: function burnableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCaller) BurnableCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "burnableCounter")
	return *ret0, err
}

// BurnableCounter is a free data retrieval call binding the contract method 0x3da9a430.
//
// Solidity: function burnableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistSession) BurnableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.BurnableCounter(&_TokenWhitelist.CallOpts)
}

// BurnableCounter is a free data retrieval call binding the contract method 0x3da9a430.
//
// Solidity: function burnableCounter() constant returns(uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) BurnableCounter() (*big.Int, error) {
	return _TokenWhitelist.Contract.BurnableCounter(&_TokenWhitelist.CallOpts)
}

// BurnableTokens is a free data retrieval call binding the contract method 0xbf25811f.
//
// Solidity: function burnableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCaller) BurnableTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "burnableTokens")
	return *ret0, err
}

// BurnableTokens is a free data retrieval call binding the contract method 0xbf25811f.
//
// Solidity: function burnableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) BurnableTokens() ([]common.Address, error) {
	return _TokenWhitelist.Contract.BurnableTokens(&_TokenWhitelist.CallOpts)
}

// BurnableTokens is a free data retrieval call binding the contract method 0xbf25811f.
//
// Solidity: function burnableTokens() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCallerSession) BurnableTokens() ([]common.Address, error) {
	return _TokenWhitelist.Contract.BurnableTokens(&_TokenWhitelist.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCaller) ControllerNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "controllerNode")
	return *ret0, err
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) ControllerNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.CallOpts)
}

// ControllerNode is a free data retrieval call binding the contract method 0xe2b4ce97.
//
// Solidity: function controllerNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCallerSession) ControllerNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.ControllerNode(&_TokenWhitelist.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) EnsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "ensRegistry")
	return *ret0, err
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.CallOpts)
}

// EnsRegistry is a free data retrieval call binding the contract method 0x7d73b231.
//
// Solidity: function ensRegistry() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) EnsRegistry() (common.Address, error) {
	return _TokenWhitelist.Contract.EnsRegistry(&_TokenWhitelist.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetStablecoinInfo(opts *bind.CallOpts) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getStablecoinInfo")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.CallOpts)
}

// GetStablecoinInfo is a free data retrieval call binding the contract method 0x3efec5e9.
//
// Solidity: function getStablecoinInfo() constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetStablecoinInfo() (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetStablecoinInfo(&_TokenWhitelist.CallOpts)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCaller) GetTokenInfo(opts *bind.CallOpts, _a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(bool)
		ret5 = new(bool)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _TokenWhitelist.contract.Call(opts, out, "getTokenInfo", _a)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.CallOpts, _a)
}

// GetTokenInfo is a free data retrieval call binding the contract method 0x1f69565f.
//
// Solidity: function getTokenInfo(address _a) constant returns(string, uint256, uint256, bool, bool, bool, uint256)
func (_TokenWhitelist *TokenWhitelistCallerSession) GetTokenInfo(_a common.Address) (string, *big.Int, *big.Int, bool, bool, bool, *big.Int, error) {
	return _TokenWhitelist.Contract.GetTokenInfo(&_TokenWhitelist.CallOpts, _a)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCaller) IsTransferable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "isTransferable")
	return *ret0, err
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_TokenWhitelist *TokenWhitelistSession) IsTransferable() (bool, error) {
	return _TokenWhitelist.Contract.IsTransferable(&_TokenWhitelist.CallOpts)
}

// IsTransferable is a free data retrieval call binding the contract method 0x2121dc75.
//
// Solidity: function isTransferable() constant returns(bool)
func (_TokenWhitelist *TokenWhitelistCallerSession) IsTransferable() (bool, error) {
	return _TokenWhitelist.Contract.IsTransferable(&_TokenWhitelist.CallOpts)
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCaller) OracleNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "oracleNode")
	return *ret0, err
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistSession) OracleNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.CallOpts)
}

// OracleNode is a free data retrieval call binding the contract method 0x80cc0dda.
//
// Solidity: function oracleNode() constant returns(bytes32)
func (_TokenWhitelist *TokenWhitelistCallerSession) OracleNode() ([32]byte, error) {
	return _TokenWhitelist.Contract.OracleNode(&_TokenWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) Owner() (common.Address, error) {
	return _TokenWhitelist.Contract.Owner(&_TokenWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) Owner() (common.Address, error) {
	return _TokenWhitelist.Contract.Owner(&_TokenWhitelist.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCaller) Stablecoin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "stablecoin")
	return *ret0, err
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistSession) Stablecoin() (common.Address, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.CallOpts)
}

// Stablecoin is a free data retrieval call binding the contract method 0xe9cbd822.
//
// Solidity: function stablecoin() constant returns(address)
func (_TokenWhitelist *TokenWhitelistCallerSession) Stablecoin() (common.Address, error) {
	return _TokenWhitelist.Contract.Stablecoin(&_TokenWhitelist.CallOpts)
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCaller) TokenAddressArray(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenWhitelist.contract.Call(opts, out, "tokenAddressArray")
	return *ret0, err
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.CallOpts)
}

// TokenAddressArray is a free data retrieval call binding the contract method 0x443dd2a4.
//
// Solidity: function tokenAddressArray() constant returns(address[])
func (_TokenWhitelist *TokenWhitelistCallerSession) TokenAddressArray() ([]common.Address, error) {
	return _TokenWhitelist.Contract.TokenAddressArray(&_TokenWhitelist.CallOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _burnable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) AddTokens(opts *bind.TransactOpts, _tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _burnable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "addTokens", _tokens, _symbols, _magnitude, _loadable, _burnable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _burnable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _burnable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _burnable, _lastUpdate)
}

// AddTokens is a paid mutator transaction binding the contract method 0x34c73edc.
//
// Solidity: function addTokens(address[] _tokens, bytes32[] _symbols, uint256[] _magnitude, bool[] _loadable, bool[] _burnable, uint256 _lastUpdate) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) AddTokens(_tokens []common.Address, _symbols [][32]byte, _magnitude []*big.Int, _loadable []bool, _burnable []bool, _lastUpdate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.AddTokens(&_TokenWhitelist.TransactOpts, _tokens, _symbols, _magnitude, _loadable, _burnable, _lastUpdate)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) Claim(opts *bind.TransactOpts, _to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "claim", _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Claim(&_TokenWhitelist.TransactOpts, _to, _asset, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x996cba68.
//
// Solidity: function claim(address _to, address _asset, uint256 _amount) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) Claim(_to common.Address, _asset common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.Claim(&_TokenWhitelist.TransactOpts, _to, _asset, _amount)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) RemoveTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "removeTokens", _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RemoveTokens(&_TokenWhitelist.TransactOpts, _tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] _tokens) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) RemoveTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RemoveTokens(&_TokenWhitelist.TransactOpts, _tokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWhitelist *TokenWhitelistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWhitelist *TokenWhitelistSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RenounceOwnership(&_TokenWhitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenWhitelist.Contract.RenounceOwnership(&_TokenWhitelist.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, _account common.Address, _transferable bool) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "transferOwnership", _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_TokenWhitelist *TokenWhitelistSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TransferOwnership(&_TokenWhitelist.TransactOpts, _account, _transferable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address _account, bool _transferable) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) TransferOwnership(_account common.Address, _transferable bool) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.TransferOwnership(&_TokenWhitelist.TransactOpts, _account, _transferable)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistTransactor) UpdateTokenRate(opts *bind.TransactOpts, _token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.contract.Transact(opts, "updateTokenRate", _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.UpdateTokenRate(&_TokenWhitelist.TransactOpts, _token, _rate, _updateDate)
}

// UpdateTokenRate is a paid mutator transaction binding the contract method 0xd545782e.
//
// Solidity: function updateTokenRate(address _token, uint256 _rate, uint256 _updateDate) returns()
func (_TokenWhitelist *TokenWhitelistTransactorSession) UpdateTokenRate(_token common.Address, _rate *big.Int, _updateDate *big.Int) (*types.Transaction, error) {
	return _TokenWhitelist.Contract.UpdateTokenRate(&_TokenWhitelist.TransactOpts, _token, _rate, _updateDate)
}

// TokenWhitelistAddedTokenIterator is returned from FilterAddedToken and is used to iterate over the raw logs and unpacked data for AddedToken events raised by the TokenWhitelist contract.
type TokenWhitelistAddedTokenIterator struct {
	Event *TokenWhitelistAddedToken // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistAddedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistAddedToken)
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
		it.Event = new(TokenWhitelistAddedToken)
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
func (it *TokenWhitelistAddedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistAddedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistAddedToken represents a AddedToken event raised by the TokenWhitelist contract.
type TokenWhitelistAddedToken struct {
	Sender    common.Address
	Token     common.Address
	Symbol    string
	Magnitude *big.Int
	Loadable  bool
	Burnable  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedToken is a free log retrieval operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _burnable)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterAddedToken(opts *bind.FilterOpts) (*TokenWhitelistAddedTokenIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistAddedTokenIterator{contract: _TokenWhitelist.contract, event: "AddedToken", logs: logs, sub: sub}, nil
}

// WatchAddedToken is a free log subscription operation binding the contract event 0x1802e89da3f6ef84e024e37454c226b1e13bf846ce71cd2a1d24faef9cbf779b.
//
// Solidity: event AddedToken(address _sender, address _token, string _symbol, uint256 _magnitude, bool _loadable, bool _burnable)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchAddedToken(opts *bind.WatchOpts, sink chan<- *TokenWhitelistAddedToken) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "AddedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistAddedToken)
				if err := _TokenWhitelist.contract.UnpackLog(event, "AddedToken", log); err != nil {
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

// TokenWhitelistClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the TokenWhitelist contract.
type TokenWhitelistClaimedIterator struct {
	Event *TokenWhitelistClaimed // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistClaimed)
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
		it.Event = new(TokenWhitelistClaimed)
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
func (it *TokenWhitelistClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistClaimed represents a Claimed event raised by the TokenWhitelist contract.
type TokenWhitelistClaimed struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterClaimed(opts *bind.FilterOpts) (*TokenWhitelistClaimedIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistClaimedIterator{contract: _TokenWhitelist.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address _to, address _asset, uint256 _amount)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *TokenWhitelistClaimed) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistClaimed)
				if err := _TokenWhitelist.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// TokenWhitelistLockedOwnershipIterator is returned from FilterLockedOwnership and is used to iterate over the raw logs and unpacked data for LockedOwnership events raised by the TokenWhitelist contract.
type TokenWhitelistLockedOwnershipIterator struct {
	Event *TokenWhitelistLockedOwnership // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistLockedOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistLockedOwnership)
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
		it.Event = new(TokenWhitelistLockedOwnership)
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
func (it *TokenWhitelistLockedOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistLockedOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistLockedOwnership represents a LockedOwnership event raised by the TokenWhitelist contract.
type TokenWhitelistLockedOwnership struct {
	Locked common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedOwnership is a free log retrieval operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterLockedOwnership(opts *bind.FilterOpts) (*TokenWhitelistLockedOwnershipIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistLockedOwnershipIterator{contract: _TokenWhitelist.contract, event: "LockedOwnership", logs: logs, sub: sub}, nil
}

// WatchLockedOwnership is a free log subscription operation binding the contract event 0x808639ff9c8e4732d60b6c2330de498035416d229f27a77d259680895efec122.
//
// Solidity: event LockedOwnership(address _locked)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchLockedOwnership(opts *bind.WatchOpts, sink chan<- *TokenWhitelistLockedOwnership) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "LockedOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistLockedOwnership)
				if err := _TokenWhitelist.contract.UnpackLog(event, "LockedOwnership", log); err != nil {
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

// TokenWhitelistRemovedTokenIterator is returned from FilterRemovedToken and is used to iterate over the raw logs and unpacked data for RemovedToken events raised by the TokenWhitelist contract.
type TokenWhitelistRemovedTokenIterator struct {
	Event *TokenWhitelistRemovedToken // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistRemovedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistRemovedToken)
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
		it.Event = new(TokenWhitelistRemovedToken)
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
func (it *TokenWhitelistRemovedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistRemovedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistRemovedToken represents a RemovedToken event raised by the TokenWhitelist contract.
type TokenWhitelistRemovedToken struct {
	Sender common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemovedToken is a free log retrieval operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterRemovedToken(opts *bind.FilterOpts) (*TokenWhitelistRemovedTokenIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistRemovedTokenIterator{contract: _TokenWhitelist.contract, event: "RemovedToken", logs: logs, sub: sub}, nil
}

// WatchRemovedToken is a free log subscription operation binding the contract event 0x703f7e3f084d5b8dcc12fddcfd9a70d65b6b21ec7659e4608dbaf4419ede3ad0.
//
// Solidity: event RemovedToken(address _sender, address _token)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchRemovedToken(opts *bind.WatchOpts, sink chan<- *TokenWhitelistRemovedToken) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "RemovedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistRemovedToken)
				if err := _TokenWhitelist.contract.UnpackLog(event, "RemovedToken", log); err != nil {
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

// TokenWhitelistTransferredOwnershipIterator is returned from FilterTransferredOwnership and is used to iterate over the raw logs and unpacked data for TransferredOwnership events raised by the TokenWhitelist contract.
type TokenWhitelistTransferredOwnershipIterator struct {
	Event *TokenWhitelistTransferredOwnership // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistTransferredOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistTransferredOwnership)
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
		it.Event = new(TokenWhitelistTransferredOwnership)
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
func (it *TokenWhitelistTransferredOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistTransferredOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistTransferredOwnership represents a TransferredOwnership event raised by the TokenWhitelist contract.
type TokenWhitelistTransferredOwnership struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferredOwnership is a free log retrieval operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterTransferredOwnership(opts *bind.FilterOpts) (*TokenWhitelistTransferredOwnershipIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistTransferredOwnershipIterator{contract: _TokenWhitelist.contract, event: "TransferredOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferredOwnership is a free log subscription operation binding the contract event 0x850b3df64837d7d518b45f5aa64d104652c3b80eb5b34a8e3d9eb666cb7cdea5.
//
// Solidity: event TransferredOwnership(address _from, address _to)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchTransferredOwnership(opts *bind.WatchOpts, sink chan<- *TokenWhitelistTransferredOwnership) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "TransferredOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistTransferredOwnership)
				if err := _TokenWhitelist.contract.UnpackLog(event, "TransferredOwnership", log); err != nil {
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

// TokenWhitelistUpdatedTokenRateIterator is returned from FilterUpdatedTokenRate and is used to iterate over the raw logs and unpacked data for UpdatedTokenRate events raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRateIterator struct {
	Event *TokenWhitelistUpdatedTokenRate // Event containing the contract specifics and raw log

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
func (it *TokenWhitelistUpdatedTokenRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenWhitelistUpdatedTokenRate)
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
		it.Event = new(TokenWhitelistUpdatedTokenRate)
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
func (it *TokenWhitelistUpdatedTokenRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenWhitelistUpdatedTokenRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenWhitelistUpdatedTokenRate represents a UpdatedTokenRate event raised by the TokenWhitelist contract.
type TokenWhitelistUpdatedTokenRate struct {
	Sender common.Address
	Token  common.Address
	Rate   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedTokenRate is a free log retrieval operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) FilterUpdatedTokenRate(opts *bind.FilterOpts) (*TokenWhitelistUpdatedTokenRateIterator, error) {

	logs, sub, err := _TokenWhitelist.contract.FilterLogs(opts, "UpdatedTokenRate")
	if err != nil {
		return nil, err
	}
	return &TokenWhitelistUpdatedTokenRateIterator{contract: _TokenWhitelist.contract, event: "UpdatedTokenRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedTokenRate is a free log subscription operation binding the contract event 0xdb3a4cfb4cd8ac94343ff7440cee8d05ade309056203f0e53ca49b6db8197c7d.
//
// Solidity: event UpdatedTokenRate(address _sender, address _token, uint256 _rate)
func (_TokenWhitelist *TokenWhitelistFilterer) WatchUpdatedTokenRate(opts *bind.WatchOpts, sink chan<- *TokenWhitelistUpdatedTokenRate) (event.Subscription, error) {

	logs, sub, err := _TokenWhitelist.contract.WatchLogs(opts, "UpdatedTokenRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenWhitelistUpdatedTokenRate)
				if err := _TokenWhitelist.contract.UnpackLog(event, "UpdatedTokenRate", log); err != nil {
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
