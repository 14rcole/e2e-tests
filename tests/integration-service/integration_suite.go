package integration

var _ = BeforeSuite(func() {
	Expect(true).To(BeFalse())
})

var _ = AfterSuite(func() {
	Expect(true).To(BeFalse())
})
