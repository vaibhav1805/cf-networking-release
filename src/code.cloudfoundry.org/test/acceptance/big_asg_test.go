package acceptance_test

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"code.cloudfoundry.org/lib/testsupport"
	"github.com/cloudfoundry/cf-test-helpers/v2/workflowhelpers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container startup time with a big ASG", func() {
	var (
		ASGFilepath string
	)

	AfterEach(func() {
		workflowhelpers.AsUser(TestSetup.AdminUserContext(), Timeout_Push, func() {
			Expect(cfCLI.DeleteSecurityGroup("big-asg")).To(Succeed())
		})
		os.Remove(ASGFilepath)
	})

	Describe("Pushing app with and without a large application security group (ASG)", func() {
		It("should not give large time difference", func() {
			var (
				durationWithoutASG time.Duration
				durationWithASG    time.Duration
			)

			By("pushing an app", func() {
				start := time.Now()
				appB := fmt.Sprintf("appB-%d", rand.Int31())
				pushProxy(appB)
				durationWithoutASG = time.Since(start)
			})

			By("creating a large ASG", func() {
				asg := testsupport.BuildASG(1000)
				var err error
				ASGFilepath, err = testsupport.CreateTempFile(asg)
				Expect(err).NotTo(HaveOccurred())
				workflowhelpers.AsUser(TestSetup.AdminUserContext(), Timeout_Push, func() {
					Expect(cfCLI.CreateSecurityGroup("big-asg", ASGFilepath)).To(Succeed())
					Expect(cfCLI.BindSecurityGroup("big-asg", TestSetup.TestSpace.OrganizationName(), TestSetup.TestSpace.SpaceName())).To(Succeed())
				})
			})

			By("pushing another app", func() {
				start := time.Now()
				appA := fmt.Sprintf("appA-%d", rand.Int31())
				pushProxy(appA)
				durationWithASG = time.Since(start)
			})

			fmt.Println("##############################################")
			fmt.Println("push app without ASG took:", durationWithoutASG)
			fmt.Println("push app with big ASG took:", durationWithASG)
			fmt.Println("##############################################")
		})
	})
})
