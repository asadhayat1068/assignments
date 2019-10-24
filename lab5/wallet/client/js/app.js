import Web3 from "web3";
import { fromWei } from "web3-utils";
import MyWalletContract from "../../build/contracts/MyWallet.json";

const App = {
    web3: null,
    accounts: null,
    defaultAccount: null,
    contract: null,

    start: async function () {
        const { web3 } = this;

        try {
            // get contract instance
            const networkId = await web3.eth.net.getId();
            const deployedNetwork = MyWalletContract.networks[networkId];
            this.contract = new web3.eth.Contract(
                MyWalletContract.abi,
                deployedNetwork.address,
            );

            // get accounts
            this.accounts = await this.getAccounts();
            this.setDefaultAccount();
            this.buildAccountList();
            this.refreshPageInfo();
        } catch (error) {
            console.error(`Could not connect to contract or chain. Due the following error: ${error}`);
        }
    },

    deposit: async function () {
        const amount = document.getElementById("depositAmount").value;

        this.setStatus("Initiating transaction... (please wait)");

        const { deposit } = this.contract.methods;
        await deposit().send({ from: this.defaultAccount, value: amount });

        this.setStatus(`Transaction complete! Successfully deposit ${amount} wei.`);
        this.refreshPageInfo();
    },

    withdraw: async function () {
        const amount = document.getElementById("withdrawAmount").value;

        this.setStatus("Initiating transaction... (please wait)");

        const { withdraw } = this.contract.methods;
        await withdraw(amount).send({ from: this.defaultAccount });

        this.setStatus(`Transaction complete! Successfully withdrawn ${amount} wei.`);
        this.refreshPageInfo();
    },

    transfer: async function () {
        const amount = document.getElementById("transferAmount").value;
        const recipient = document.getElementById("transferRecipient").value;

        this.setStatus("Initiating transaction... (please wait)");

        const { transfer } = this.contract.methods;
        await transfer(recipient, amount).send({ from: this.defaultAccount });

        this.setStatus(`Transaction complete! Successfully transferred ${amount} ether to ${recipient}.`);
        this.refreshPageInfo();
    },

    // Helper functions
    refreshWalletBalance: async function () {
        const { getBalance } = this.contract.methods;
        const balance = await getBalance().call();
        const balanceElement = document.getElementById("walletBalance");
        balanceElement.innerHTML = balance;
    },

    refreshAccountsInfo: async function () {
        const accountBalanceElements = document.getElementsByClassName("accountBalance");
        var i = 0;
        for (var account of this.accounts.keys()) {
            let accountBalance = await this.getAccountBalance(account);
            accountBalanceElements[i].innerHTML = accountBalance;
            i++;
        };
    },

    buildAccountList: async function () {
        const accountListElement = document.getElementById('accountList');
        var i = 0;
        this.accounts.forEach(function (balance, account, index) {
            let li = document.createElement('li');
            li.className = "list-group-item d-flex justify-content-between align-items-center";
            if (i == 0) {
                li.className += " active";
            }
            let spanAccount = document.createElement('span');
            spanAccount.className = "account";
            spanAccount.innerHTML = account;
            let spanBalance = document.createElement('span');
            spanBalance.className = "accountBalance";
            spanBalance.innerHTML = balance;
            li.appendChild(spanAccount);
            li.appendChild(spanBalance);
            accountListElement.appendChild(li);
            i++;
        });
    },

    refreshPageInfo: async function () {
        this.refreshWalletBalance();
        this.refreshAccountsInfo();
    },

    setStatus: function (message) {
        const status = document.getElementById("status");
        // TODO: display transaction errors
        status.innerHTML = message;
    },

    getAccounts: async function () {
        let accountsBalances = new Map();
        let accounts = await this.web3.eth.getAccounts();
        for (var i = 0; i < accounts.length; i++) {
            let accountBalance = await this.getAccountBalance(accounts[i]);
            accountsBalances.set(accounts[i], accountBalance)
        }
        return accountsBalances;
    },

    // TODO: allow user to select default account
    setDefaultAccount: async function () {
        let accounts = await this.web3.eth.getAccounts();
        this.defaultAccount = accounts[0]; // set default account
    },

    getAccountBalance: async function (account) {
        const balance = await this.web3.eth.getBalance(account);
        return fromWei(balance, 'ether');
    },
};

window.App = App;

const loadWeb3 = function () {
    let web3 = null;
    if (window.ethereum) {
        console.log("Using Metamask provider");
        web3 = new Web3(window.ethereum);
        window.ethereum.enable(); // get permission to access accounts
    } else {
        console.log("No web3 detected. Falling back to http://127.0.0.1:7545");
        // fallback to local node (ganache)
        web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:7545"));
    }
    return web3;
};

window.addEventListener("load", function () {
    App.web3 = loadWeb3();
    App.start();
});