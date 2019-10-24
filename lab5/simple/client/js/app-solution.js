import Web3 from "web3";
import Simple from "../../build/contracts/Simple.json";

const App = {
    web3: null,
    contract: null,
    defaultAccount: null,

    start: async function () {
        const { web3 } = this;

        let deployedNetwork = Simple.networks["5777"];
        this.contract = new web3.eth.Contract(
            Simple.abi,
            deployedNetwork.address,
        );

        let accounts = await web3.eth.getAccounts();
        this.defaultAccount = accounts[0];

        this.get();
    },

    get: async function () {
        let value = await this.contract.methods.get().call();
        let contractValue = document.getElementById("contractValue");
        contractValue.innerHTML = value;

        let cbalance = await this.contract.methods.balance().call();
        let contractBalance = document.getElementById("contractBalance");
        contractBalance.innerHTML = cbalance;

        let ubalance = await this.web3.eth.getBalance(this.defaultAccount);
        let userBalance = document.getElementById("userBalance");
        userBalance.innerHTML = ubalance;
    },

    set: async function () {
        const { set } = this.contract.methods;

        let text = document.getElementById("textValue").value;
        let amount = document.getElementById("amount").value;
        await set(text).send({ from: this.defaultAccount, value: amount });
        this.get();
        // NOTE: you can still able to send money to the contract address after the the self destruct invocation
    },

    stop: async function () {
        const { stop } = this.contract.methods;
        await stop().send({ from: this.defaultAccount });
        this.get();
    }
}

window.App = App;

window.addEventListener("load", function () {
    App.web3 = new Web3(new Web3.providers.HttpProvider('http://127.0.0.1:7545'));
    App.start();
});