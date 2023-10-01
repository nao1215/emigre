package api

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

var _ = Describe("User api test", func() {
	var (
		t   GinkgoTInterface
		api *API
	)

	BeforeEach(func() {
		t = GinkgoT()
		api = NewAPI()
	})

	Context("success to create user record in DB", func() {
		It("return 201(Created) and created user data", func() {
			apitest.New("success to POST /users").
				Report(apitest.SequenceDiagram(documentDirPath())).
				Handler(api).
				Post("/v1/users").
				ContentType("application/json").
				JSON(`{"name": "nao1215", "biography": "I'm emigre author", "email": "nao1215@example.com"}`).
				Expect(t).
				// id is ULID, so it is not fixed value. not check it.
				Assert(jsonpath.Equal(`$.name`, "nao1215")).
				Assert(jsonpath.Equal(`$.biography`, "I'm emigre author")).
				Assert(jsonpath.Equal(`$.email`, "nao1215@example.com")).
				Status(http.StatusCreated).
				End()
		})
	})
})
