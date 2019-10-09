const { weiToEther, etherToWei, currentBalance } = require('./utils');
const { BN, expectEvent, expectRevert } = require('openzeppelin-test-helpers');
const { expect } = require('chai');

const Betting = artifacts.require('Betting');

// TODO: The tests (it) in one describe are not independent.
// If tests fail, look at the first one failing.

contract('Betting', accounts => {
    const [owner, oracle , winner1, winner2, loser1, loser2] = accounts;
    let betting = null;
    const team1 = web3.utils.soliditySha3("team1");
    const team2 = web3.utils.soliditySha3("team2");
    const team3 = web3.utils.soliditySha3("team3");

    describe('constructor', () => {
        it('should successfully create the contract with nonempty outcomes', async () => {
            betting = await Betting.new([team1, team2]); 
            (await betting.owner()).should.equal(owner);
        });
        
        it('should revert when creating contract with empty outcomes', async ()=> {
            expectRevert(Betting.new([]), 'must register at least 2 bets');
        })

        it('should revert when creating contract with only one outcome', async ()=> {
            await expectRevert(Betting.new([team1]), 'must register at least 2 bets');
        })
    });

    describe('constructor and checkOutcome', () => {
        before(async () => {
            betting = await Betting.new([team1, team2]);
        });
        it('should be able to check first registered outcomes after creation', async () => {
            expect(await betting.checkOutcome(team1)).to.be.bignumber.equal(new BN(0));
        });
        it('should be able to check second registered outcomes after creation', async () => {
            expect(await betting.checkOutcome(team2)).to.be.bignumber.equal(new BN(0));
        });

        it('checkOutcome for unregistered outcome should revert', async () => {
            await expectRevert(betting.checkOutcome(team3), 'outcome not registered');
        });
    });

    describe('choose Oracle and isOracle', () => {
        before(async () => {
            betting = await Betting.new([team1, team2]);
        });
        it("owner can asign oracle and isOracle checks correctly", async () => {
            await betting.chooseOracle(winner1);
            expect(await betting.isOracle(winner1)).to.equal(true);
        });
        it("owner can reasign oracle and isOracle checks correctly", async () => {
            await betting.chooseOracle(oracle);
            expect(await betting.isOracle(oracle)).to.equal(true);
        });
        it("isOracle correctly checks oracle", async () => {
            expect(await betting.isOracle(oracle)).to.equal(true);
            expect(await betting.isOracle(owner)).to.equal(false);
            expect(await betting.isOracle(winner1)).to.equal(false);
        })
        it("non-owner cannot reasign oracle", async () => {
            await expectRevert(betting.chooseOracle(winner1, {"from": winner1}),"sender isn't the owner");
        })

        it('should emit an event when reasigning oracle', async () => {
            await betting.chooseOracle(oracle);
            const { logs } = await betting.chooseOracle(winner1);

            expectEvent.inLogs(logs, 'OracleChanged', {
                previousOracle: oracle,
                newOracle: winner1
            });
        });
    });

    describe('makeBet', () => {
        before(async () => {
            betting = await Betting.new([team1, team2]);
        });
        it('makeBet revert if oracle is not assigned', async ()=> {
            await expectRevert(betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')}),"no oracle found");
        });
        it('makeBet revert if called by owner', async ()=> {
            await betting.chooseOracle(oracle);
            await expectRevert(betting.makeBet(team1, {"from": owner, "value": etherToWei('1')}),"the contract owner cannot bet");
        });
        it('makeBet revert if called by oracle', async ()=> {
            await expectRevert(betting.makeBet(team1, {"from": oracle, "value": etherToWei('1')}),"the oracle of the betting cannot bet");
        });
        it('cannot bet on unregistered outcome', async ()=> {
            await expectRevert(betting.makeBet(team3, {"from": winner1, "value": etherToWei('1')}),"outcome don't exists");
        });
        it('better can makeBet', async ()=> {
            await betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')});
        });
        it('cannot bet twice', async ()=> {
            await expectRevert(betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')}),"each gambler can only bet once");
            await expectRevert(betting.makeBet(team2, {"from": winner1, "value": etherToWei('1')}),"each gambler can only bet once");
        });
        it('should emit an event when making bet', async () => {
            const { logs } = await betting.makeBet(team1, {"from": winner2, "value": etherToWei('1')});

            expectEvent.inLogs(logs, 'BetMade', {
                gambler: winner2,
                outcome: team1,
                amount: etherToWei('1')
            });
        });
    });
    describe('makeBet and checkOutcome', () => {
        before(async () => {
            betting = await Betting.new([team1, team2]);
            await betting.chooseOracle(oracle);
        });
        it('checkOutcome shows sum betted on outcome', async ()=> {
            await betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')});
            expect(await betting.checkOutcome(team1)).to.be.bignumber.equal(new BN(etherToWei('1')));
            await betting.makeBet(team1, {"from": winner2, "value": etherToWei('2')});
            expect(await betting.checkOutcome(team1)).to.be.bignumber.equal(new BN(etherToWei('3')));
            await betting.makeBet(team2, {"from": loser1, "value": etherToWei('2')});
            expect(await betting.checkOutcome(team2)).to.be.bignumber.equal(new BN(etherToWei('2')));
        });
    });
    describe('makeDecision', () => {
        before(async () => {
            betting = await Betting.new([team1, team2]);
            await betting.chooseOracle(oracle);
        });
        it('revert if makeDecision for unregistered outcome', async ()=>{
            await expectRevert(betting.makeDecision(team3, {"from": oracle}), "outcome don't exists");
        });
        it('revert if makeDecision called by non oracle', async ()=>{
            await expectRevert(betting.makeDecision(team1, {"from": owner}), "sender isn't the oracle");
        });
    });
    describe('makeBet and makeDecision and checkWinnings winner takes it all', () => {
        before(async () => {
            betting = await Betting.new([team1, team2]);
            await betting.chooseOracle(oracle);
            await betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')});
            await betting.makeBet(team2, {"from": loser1, "value": etherToWei('2')});
        });
        it('can makeDecision and check winnings', async () => {
            await betting.makeDecision(team1, {"from":oracle});
            expect(await betting.checkWinnings({"from":winner2})).to.be.bignumber.equal(new BN(etherToWei('0')));
            expect(await betting.checkWinnings({"from":loser1})).to.be.bignumber.equal(new BN('0'));
        });
        it('winner get his reward and bid back', async () => {
            expect(await betting.checkWinnings({"from":winner1})).to.be.bignumber.equal(new BN(etherToWei('3')));
        });
        it('cannot makeDecision again', async () => {
            await expectRevert(betting.makeDecision(team2, {"from":oracle}), 'can make decision only once');
        });
        it('cannot make new bet after decision', async () => {
            await expectRevert(betting.makeBet(team1, {"from":loser2, "value": etherToWei('1')}), "cannot bet after decision was made");
        });
    });
    describe('makeBet and makeDecision and checkWinnings relative winnings', () => {
        before(async () => {
            betting = await Betting.new([team1, team2, team3]);
            await betting.chooseOracle(oracle);
            await betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')});
            await betting.makeBet(team1, {"from": winner2, "value": etherToWei('2')});
            await betting.makeBet(team2, {"from": loser1, "value": etherToWei('2')});
            await betting.makeBet(team3, {"from": loser2, "value": etherToWei('1')});
            await betting.makeDecision(team1, {"from":oracle});
        });
        it('winners get proportional reward', async () => {
            expect(await betting.checkWinnings({"from":winner1})).to.be.bignumber.equal(new BN(etherToWei('2')));
            expect(await betting.checkWinnings({"from":winner2})).to.be.bignumber.equal(new BN(etherToWei('4')));
            expect(await betting.checkWinnings({"from":loser1})).to.be.bignumber.equal(new BN('0'));
            expect(await betting.checkWinnings({"from":loser2})).to.be.bignumber.equal(new BN('0'));
        });
    });
    describe('makeBet and makeDecision and checkWinnings, oracle wins', () => {
        before(async () => {
            betting = await Betting.new([team1, team2, team3]);
            await betting.chooseOracle(oracle);
            await betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')});
            await betting.makeBet(team2, {"from": winner2, "value": etherToWei('2')});
            await betting.makeDecision(team3, {"from":oracle});
        });
        it('winners get proportional reward', async () => {
            expect(await betting.checkWinnings({"from":winner1})).to.be.bignumber.equal(new BN(etherToWei('0')));
            expect(await betting.checkWinnings({"from":winner2})).to.be.bignumber.equal(new BN(etherToWei('0')));
            expect(await betting.checkWinnings({"from":oracle})).to.be.bignumber.equal(new BN(etherToWei('3')));
        });
    });
    describe('reset and checkWinnings', () => {
        before(async () => {
            betting = await Betting.new([team1, team2, team3]);
            await betting.chooseOracle(oracle);
            await betting.makeBet(team1, {"from": winner1, "value": etherToWei('1')});
            await betting.makeBet(team1, {"from": winner2, "value": etherToWei('2')});
            await betting.makeBet(team2, {"from": loser1, "value": etherToWei('2')});
            await betting.makeBet(team3, {"from": loser2, "value": etherToWei('1')});
        });
        it('reset before decision reverts', async ()=> {
            await expectRevert(betting.contractReset(), "cannot reset before decision");
        });
        it('reset works after decision', async ()=> {
            await betting.makeDecision(team1, {"from":oracle});
            await betting.contractReset();
        });
        it('winners have reward even after decision', async () => {
            expect(await betting.checkWinnings({"from":winner1})).to.be.bignumber.equal(new BN(etherToWei('2')));
            expect(await betting.checkWinnings({"from":winner2})).to.be.bignumber.equal(new BN(etherToWei('4')));
            expect(await betting.checkWinnings({"from":loser1})).to.be.bignumber.equal(new BN('0'));
            expect(await betting.checkWinnings({"from":loser2})).to.be.bignumber.equal(new BN('0'));
        });
    });
});
