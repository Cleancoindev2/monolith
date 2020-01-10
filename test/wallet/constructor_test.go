package wallet_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tokencard/contracts/v2/pkg/bindings"
	. "github.com/tokencard/contracts/v2/test/shared"
)

var _ = Describe("wallet constructor", func() {


	It("Should deploy a new wallet", func() {
		_, tx, _, err := bindings.DeployWallet(
			BankAccount.TransactOpts(),
			Backend,
			Owner.Address(),
			true,
			ENSRegistryAddress,
			TokenWhitelistName,
			ControllerName,
			LicenceName,
			EthToWei(100),
		)
		Expect(err).ToNot(HaveOccurred())
		Backend.Commit()
		Expect(isSuccessful(tx)).To(BeTrue())
	})

})
