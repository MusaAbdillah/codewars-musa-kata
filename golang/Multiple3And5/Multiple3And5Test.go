package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)
var _ = Describe("Multiples of 3 and 5", func() {
   
   It("should handle basic cases", func() {
     Expect(Multiple3And5(10)).To(Equal(23))
   })
})