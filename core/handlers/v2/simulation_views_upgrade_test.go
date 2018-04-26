package v2

import (
	"testing"

	"github.com/SpectoLabs/hoverfly/core/util"
	. "github.com/onsi/gomega"
)

var v1Meta = MetaView{
	SchemaVersion:   "v1",
	HoverflyVersion: "test",
	TimeExported:    "today",
}
var v2Meta = MetaView{
	SchemaVersion:   "v2",
	HoverflyVersion: "test",
	TimeExported:    "today",
}

func Test_upgradeV1_ReturnsAnUpgradedSimulation(t *testing.T) {
	RegisterTestingT(t)

	v1Simulation := SimulationViewV1{
		DataViewV1{
			RequestResponsePairViewV1: []RequestResponsePairViewV1{
				{
					Request: RequestDetailsView{
						RequestType: util.StringToPointer("recording"),
						Scheme:      util.StringToPointer("http"),
						Body:        util.StringToPointer("body"),
						Destination: util.StringToPointer("destination"),
						Method:      util.StringToPointer("GET"),
						Path:        util.StringToPointer("/path"),
						Query:       util.StringToPointer("query=query"),
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "body",
						EncodedBody: false,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v1Meta,
	}

	upgradedSimulation := upgradeV1(v1Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme[0].Value).To(Equal("http"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body[0].Value).To(Equal("body"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination[0].Value).To(Equal("destination"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Value).To(Equal("GET"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path[0].Value).To(Equal("/path"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Value).To(Equal("query=query"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Headers).To(BeEmpty())

	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Status).To(Equal(200))
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Templated).To(BeFalse())
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Body).To(Equal("body"))
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.EncodedBody).To(BeFalse())
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Headers).To(HaveKeyWithValue("Test", []string{"headers"}))

	Expect(upgradedSimulation.SchemaVersion).To(Equal("v5"))
	Expect(upgradedSimulation.HoverflyVersion).To(Equal("test"))
	Expect(upgradedSimulation.TimeExported).To(Equal("today"))
}
func Test_upgradeV1_HandlesTemplates(t *testing.T) {
	RegisterTestingT(t)

	v1Simulation := SimulationViewV1{
		DataViewV1{
			RequestResponsePairViewV1: []RequestResponsePairViewV1{
				{
					Request: RequestDetailsView{
						RequestType: util.StringToPointer("template"),
						Scheme:      util.StringToPointer("http"),
						Body:        util.StringToPointer("body"),
						Destination: util.StringToPointer("destination"),
						Method:      util.StringToPointer("GET"),
						Path:        util.StringToPointer("/path"),
						Query:       util.StringToPointer("query=query"),
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "body",
						EncodedBody: false,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v1Meta,
	}

	upgradedSimulation := upgradeV1(v1Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme[0].Value).To(Equal("http"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body[0].Value).To(Equal("body"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination[0].Value).To(Equal("destination"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Value).To(Equal("GET"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path[0].Value).To(Equal("/path"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Value).To(Equal("query=query"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Headers).To(BeEmpty())
}
func Test_upgradeV1_HandlesIncompleteRequest(t *testing.T) {
	RegisterTestingT(t)

	v1Simulation := SimulationViewV1{
		DataViewV1{
			RequestResponsePairViewV1: []RequestResponsePairViewV1{
				{
					Request: RequestDetailsView{
						Method: util.StringToPointer("POST"),
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "body",
						EncodedBody: false,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v1Meta,
	}

	upgradedSimulation := upgradeV1(v1Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme).To(HaveLen(0))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body).To(HaveLen(0))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination).To(HaveLen(0))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path).To(HaveLen(0))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query).To(HaveLen(0))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Value).To(Equal("POST"))

	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Status).To(Equal(200))
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Body).To(Equal("body"))
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.EncodedBody).To(BeFalse())
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Headers).To(HaveKeyWithValue("Test", []string{"headers"}))
}

func Test_upgradeV1_Upgrade_UnescapesRequestQueryParameters(t *testing.T) {
	RegisterTestingT(t)

	v1Simulation := SimulationViewV1{
		DataViewV1{
			RequestResponsePairViewV1: []RequestResponsePairViewV1{
				{
					Request: RequestDetailsView{
						Query: util.StringToPointer("q=10%20Downing%20Street%20London"),
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "body",
						EncodedBody: false,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v1Meta,
	}

	upgradedSimulation := upgradeV1(v1Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Value).To(Equal("q=10 Downing Street London"))
}

func Test_upgradeV2_ReturnsAnUpgradedSimulation(t *testing.T) {
	RegisterTestingT(t)

	v2Simulation := SimulationViewV2{
		DataViewV2{
			RequestResponsePairs: []RequestMatcherResponsePairViewV2{
				{
					RequestMatcher: RequestMatcherViewV2{
						Scheme: &RequestFieldMatchersView{
							RegexMatch: util.StringToPointer("http"),
						},
						Method: &RequestFieldMatchersView{
							XpathMatch: util.StringToPointer("*"),
						},
						Query: &RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("query=query"),
						},
						Destination: &RequestFieldMatchersView{
							GlobMatch: util.StringToPointer("*"),
						},
						Path: &RequestFieldMatchersView{
							JsonMatch: util.StringToPointer("*"),
						},
						Body: &RequestFieldMatchersView{
							XmlMatch: util.StringToPointer("*"),
						},
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "body",
						EncodedBody: false,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v2Meta,
	}

	upgradedSimulation := upgradeV2(v2Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme[0].Matcher).To(Equal("regex"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Scheme[0].Value).To(Equal("http"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body[0].Matcher).To(Equal("xml"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Body[0].Value).To(Equal("*"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Destination[0].Value).To(Equal("*"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Matcher).To(Equal("xpath"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Method[0].Value).To(Equal("*"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path[0].Matcher).To(Equal("json"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Path[0].Value).To(Equal("*"))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Value).To(Equal("query=query"))

	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Status).To(Equal(200))
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Templated).To(BeFalse())
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Body).To(Equal("body"))
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.EncodedBody).To(BeFalse())
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Headers).To(HaveKeyWithValue("Test", []string{"headers"}))

	Expect(upgradedSimulation.SchemaVersion).To(Equal("v5"))
	Expect(upgradedSimulation.HoverflyVersion).To(Equal("test"))
	Expect(upgradedSimulation.TimeExported).To(Equal("today"))
}

func Test_upgradeV2_UnescapesExactMatchRequestQueryParameters(t *testing.T) {
	RegisterTestingT(t)

	v2Simulation := SimulationViewV2{
		DataViewV2{
			RequestResponsePairs: []RequestMatcherResponsePairViewV2{
				{
					RequestMatcher: RequestMatcherViewV2{
						Query: &RequestFieldMatchersView{
							ExactMatch: util.StringToPointer("q=10%20Downing%20Street%20London"),
						},
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "body",
						EncodedBody: false,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v2Meta,
	}

	upgradedSimulation := upgradeV2(v2Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Matcher).To(Equal("exact"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Value).To(Equal("q=10 Downing Street London"))
}

func Test_upgradeV2_UnescapesGlobMatchRequestQueryParameters(t *testing.T) {
	RegisterTestingT(t)

	v2Simulation := SimulationViewV2{
		DataViewV2{
			RequestResponsePairs: []RequestMatcherResponsePairViewV2{
				{
					RequestMatcher: RequestMatcherViewV2{
						Query: &RequestFieldMatchersView{
							GlobMatch: util.StringToPointer("q=*%20London"),
						},
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "body",
						EncodedBody: false,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v2Meta,
	}

	upgradedSimulation := upgradeV2(v2Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))

	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Matcher).To(Equal("glob"))
	Expect(upgradedSimulation.RequestResponsePairs[0].RequestMatcher.Query[0].Value).To(Equal("q=* London"))
}

func Test_upgradeV2_Upgrade_KeepsEncodedResponsesEncoded(t *testing.T) {
	RegisterTestingT(t)

	v2Simulation := SimulationViewV2{
		DataViewV2{
			RequestResponsePairs: []RequestMatcherResponsePairViewV2{
				{
					RequestMatcher: RequestMatcherViewV2{
						Query: &RequestFieldMatchersView{
							GlobMatch: util.StringToPointer("q=*%20London"),
						},
					},
					Response: ResponseDetailsView{
						Status:      200,
						Body:        "YmFzZTY0IGVuY29kZWQ=",
						EncodedBody: true,
						Headers: map[string][]string{
							"Test": []string{"headers"},
						},
					},
				},
			},
		},
		v2Meta,
	}

	upgradedSimulation := upgradeV2(v2Simulation)

	Expect(upgradedSimulation.RequestResponsePairs).To(HaveLen(1))
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.EncodedBody).To(BeTrue())
	Expect(upgradedSimulation.RequestResponsePairs[0].Response.Body).To(Equal("YmFzZTY0IGVuY29kZWQ="))
}
