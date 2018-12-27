package asb2018

import (
	"github.com/orbs-network/asb2018/driver"
	"testing"
)

// EDIT THIS CONFIGURATION TO CONTROL THE TEST SCENARIO
// DON'T FORGET TO UPDATE VALUES ACCORDING TO INSTRUCTIONS AFTER DEPLOY
var configGanache = &driver.Config{
	DebugLogs:                        true,              // shows detailed responses for every command
	EthereumErc20Address:             "",                // update after deploy with the resulting value
	OrbsErc20ContractName:            "ERC20TokenProxy", // choose different names to redeploy
	OrbsAsbContractName:              "ASBEthereum",     // choose different names to redeploy
	UserAccountOnEthereum:            "",                // one of your ganache accounts
	UserAccountOnOrbs:                "user1",           // one of the IDs in orbs-test-keys.json
	UserInitialBalanceOnEthereum:     20000,
	UserTransferAmountToOrbs:         130,
	UserTransferAmountBackToEthereum: 30,
}

// before starting:
// 1. make sure ganache is running locally on port 7545
// 2. make sure gamma server is running with `gamma-cli start-local`

func TestDeployOnGanache(t *testing.T) {
	orbs := driver.AdapterForGammaCliLocal(configGanache)
	ethereum := driver.AdapterForTruffleGanache(configGanache)
	driver.RunDeployFlow(t, configGanache, orbs, ethereum)
}

func TestEthToOrbsOnGanache(t *testing.T) {
	orbs := driver.AdapterForGammaCliLocal(configGanache)
	ethereum := driver.AdapterForTruffleGanache(configGanache)
	driver.RunEthToOrbsFlow(t, configGanache, orbs, ethereum)
}

func TestOrbsToEthOnGanache(t *testing.T) {
	orbs := driver.AdapterForGammaCliLocal(configGanache)
	ethereum := driver.AdapterForTruffleGanache(configGanache)
	driver.RunOrbsToEthFlow(t, configGanache, orbs, ethereum)
}
