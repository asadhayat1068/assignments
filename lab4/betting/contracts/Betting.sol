pragma solidity 0.5.11;

contract Betting {
    /* Constructor function, where owner and outcomes are set */
    constructor(bytes32[] memory initOutcomes) public {
        // should register at least 2 possible outcomes.
        
    }

    /* Standard state variables */
    address public owner;
    address public oracle;
    bool public decisionMade;
    
    
    /* Structs are custom data structures with self-defined parameters */
    struct Bet {
        bytes32 outcome;
        uint amount;
    }

    /* Keep track of who the gambler's are */
    address[] public gamblers; // list of all gamblers for iteration
    mapping(address => bool) public isGambler; // map for checking for a gambler

    /* Keep track of every gambler's bet */
    mapping (address => Bet) bets;

    /* Keep track of every player's wins and their amount (if any) 
     zero the amount after withdraw to avoid replay attacks*/
    mapping (address => uint) winners;
    
    /* Keep track of all possible outcomes */
    bytes32[] public outcomes;
    mapping (bytes32 => bool) public validOutcomes;

    /* Keep track of the total bet amount in all outcomes */
    mapping (bytes32 => uint) public outcomeBets;

    /* Add any events you think are necessary */
    event BetMade(address indexed gambler, bytes32 indexed outcome, uint amount);
    event Winners(address[] indexed winners, uint totalPrize);
    event OracleChanged(address indexed previousOracle, address indexed newOracle);
    event Withdrawn(address indexed gambler, uint amount);

    /* Uh Oh, what are these? */
    modifier ownerOnly() {
        require(msg.sender == owner, "sender isn't the owner");
        _;
    }

    modifier oracleOnly() {
        require(isOracle(msg.sender), "sender isn't the oracle");
        _;
    }

    modifier requireOracle() {
        require(oracle != address(0),"no oracle found");
        _;
    }

    modifier outcomeExists(bytes32 outcome) {
        require(validOutcomes[outcome], "outcome don't exists");
        _;
    }

    function isOracle(address sender) public view returns (bool) {
        return sender == oracle;
    }

    /* Owner chooses their trusted Oracle, returns new oracle */
    function chooseOracle(address newOracle) public ownerOnly {
        // the oracle cannot be a gambler
    }

    /* Gamblers place their bets, preferably after calling checkOutcomes */
    function makeBet(bytes32 outcome) public payable  {
        // owner and oracle cannot make a bet
        // a gambler cannot bet twice
        // a gambler can only bet on a registered outcome
        // an oracle should be assigned before starting bets
        // impossible to bet after decision was made and before contract was reset
        // Betters are registered by placing a bet
        // emit BetMade event
    }

    /* The oracle chooses which outcome wins */
    function makeDecision(bytes32 decidedOutcome) public  {
        // only called by the oracle
        // should be called only once before reset
        // winning outcome must exist
        // The winners receive a proportional share of the total funds at stake if they all bet on the correct outcome
        // All gamblers betting on the correct outcome, get reimbursed their funds
        // If no gamblers bet on the correct outcome, the oracle wins the sum of the funds
    }

    /* Allow anyone to withdraw their winnings safely (if they have enough) */
    function withdraw(uint amount) public returns (uint) {
    }
    
    /* Allow anyone to check the outcomes they can bet on */
    function checkOutcome(bytes32 outcome) public view returns (uint) {
        // return amount bet un current outcome. 
        // revert if outcome does not exist
    }

    /* The same as checkOutcome, but first hashes the outcome string with keccak to bytes32 */
    function checkOutcomeString(string memory outcomestring) public view returns (uint) {
        // hash outcomestring using keccak256. Then checkOutcome
    }
    
    /* Allow anyone to check their winnings */
    function checkWinnings() public view returns(uint) {
    }

    /* Call delete() to reset certain state variables. Which ones? That's upto you to decide */
    function contractReset() public ownerOnly() {
        // revert if decision has not been made
    }
}