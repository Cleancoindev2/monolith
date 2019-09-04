pragma solidity ^0.5.10;

import "../externals/SafeMath.sol";

interface TokenHolder {
    function burn(address, uint ) external returns (bool);
}

contract BurnerToken {

    using SafeMath for uint256;

    event Transfer(address indexed from, address indexed to, uint value);
    event Approval(address indexed owner, address indexed spender, uint value);

    uint public currentSupply;

    address public owner;
    string public name;
    uint8 public decimals;
    string public symbol;

    mapping(address => uint) public balanceOf;
    mapping(address => mapping (address => uint)) public allowance;

    //Holds accumulated dividend tokens other than TKN
    TokenHolder public tokenholder;

    constructor() public {
        owner = msg.sender;
        name = "Monolith TKN";
        decimals = 8;
        symbol = "TKN";
    }

    function totalSupply() external view returns(uint) {
        return currentSupply;
    }

    function mint(address addr, uint amount) external {
        balanceOf[addr] = balanceOf[addr].add(amount);
        currentSupply = currentSupply.add(amount);
        emit Transfer(address(0), addr, amount);
    }

    function transfer(address _to, uint _value) external returns (bool success) {
        if (balanceOf[msg.sender] < _value) return false;
        if (_to == address(0)) return false;

        balanceOf[msg.sender] = balanceOf[msg.sender].sub(_value);
        balanceOf[_to] = balanceOf[_to].add(_value);
        emit Transfer(msg.sender, _to, _value);
        return true;
    }

    function transferFrom(address _from, address _to, uint _value) public returns (bool success) {
        if (_to == address(0)) return false;
        if (balanceOf[_from] < _value) return false;

        uint allowed = allowance[_from][msg.sender];
        if (allowed < _value) return false; //PROBLEM!
        /* require(_value <= allowed, "amount exceeds allowance"); */

        balanceOf[_to] = balanceOf[_to].add(_value);
        balanceOf[_from] = balanceOf[_from].sub(_value);
        allowance[_from][msg.sender] = allowed.sub(_value);
        emit Transfer(_from, _to, _value);
        return true;
    }

    function approve(address _spender, uint _value) external returns (bool success) {
        //require user to set to zero before resetting to nonzero
        if ((_value != 0) && (allowance[msg.sender][_spender] != 0)) {
            return false;
        }

        allowance[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }

    function increaseApproval (address _spender, uint _addedValue) external returns (bool success) {
        uint oldValue = allowance[msg.sender][_spender];
        allowance[msg.sender][_spender] = oldValue.add(_addedValue);
        emit Approval(msg.sender, _spender, allowance[msg.sender][_spender]);
        return true;
    }

    function decreaseApproval (address _spender, uint _subtractedValue) external returns (bool success) {
        uint oldValue = allowance[msg.sender][_spender];
        if (_subtractedValue > oldValue) {
            allowance[msg.sender][_spender] = 0;
        } else {
            allowance[msg.sender][_spender] = oldValue.sub(_subtractedValue);
        }
        emit Approval(msg.sender, _spender, allowance[msg.sender][_spender]);
        return true;
    }

    function setTokenHolder(address _th) external {
        tokenholder = TokenHolder(_th);
    }

    function burn(uint _amount) public returns (bool result)  {
        if (_amount > balanceOf[msg.sender]) return false;

        balanceOf[msg.sender] = balanceOf[msg.sender].sub(_amount);
        currentSupply  = currentSupply.sub(_amount);
        result = tokenholder.burn(msg.sender, _amount);
        if (!result) revert();
        emit Transfer(msg.sender, address(0), _amount);
    }

    // save some gas by making only one contract call
    //lift from KNC
    function burnFrom(address _from, uint256 _value) public returns (bool) {
        assert( transferFrom( _from, msg.sender, _value ) );
        return burn(_value);
    }
}
