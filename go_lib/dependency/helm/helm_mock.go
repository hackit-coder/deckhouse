// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package helm

//go:generate minimock -i github.com/deckhouse/deckhouse/go_lib/dependency/helm.Client -o helm_mock.go -n ClientMock -p helm

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"helm.sh/helm/v3/pkg/postrender"
)

// ClientMock implements Client
type ClientMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcDelete          func(releaseName string) (err error)
	inspectFuncDelete   func(releaseName string)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mClientMockDelete

	funcUpgrade          func(releaseName string, releaseNamespace string, templates map[string]interface{}, values map[string]interface{}, debug bool, pr ...postrender.PostRenderer) (err error)
	inspectFuncUpgrade   func(releaseName string, releaseNamespace string, templates map[string]interface{}, values map[string]interface{}, debug bool, pr ...postrender.PostRenderer)
	afterUpgradeCounter  uint64
	beforeUpgradeCounter uint64
	UpgradeMock          mClientMockUpgrade
}

// NewClientMock returns a mock for Client
func NewClientMock(t minimock.Tester) *ClientMock {
	m := &ClientMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteMock = mClientMockDelete{mock: m}
	m.DeleteMock.callArgs = []*ClientMockDeleteParams{}

	m.UpgradeMock = mClientMockUpgrade{mock: m}
	m.UpgradeMock.callArgs = []*ClientMockUpgradeParams{}

	return m
}

type mClientMockDelete struct {
	mock               *ClientMock
	defaultExpectation *ClientMockDeleteExpectation
	expectations       []*ClientMockDeleteExpectation

	callArgs []*ClientMockDeleteParams
	mutex    sync.RWMutex
}

// ClientMockDeleteExpectation specifies expectation struct of the Client.Delete
type ClientMockDeleteExpectation struct {
	mock    *ClientMock
	params  *ClientMockDeleteParams
	results *ClientMockDeleteResults
	Counter uint64
}

// ClientMockDeleteParams contains parameters of the Client.Delete
type ClientMockDeleteParams struct {
	releaseName string
}

// ClientMockDeleteResults contains results of the Client.Delete
type ClientMockDeleteResults struct {
	err error
}

// Expect sets up expected params for Client.Delete
func (mmDelete *mClientMockDelete) Expect(releaseName string) *mClientMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ClientMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ClientMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &ClientMockDeleteParams{releaseName}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the Client.Delete
func (mmDelete *mClientMockDelete) Inspect(f func(releaseName string)) *mClientMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for ClientMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by Client.Delete
func (mmDelete *mClientMockDelete) Return(err error) *ClientMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ClientMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ClientMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &ClientMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the Client.Delete method
func (mmDelete *mClientMockDelete) Set(f func(releaseName string) (err error)) *ClientMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the Client.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the Client.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the Client.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mClientMockDelete) When(releaseName string) *ClientMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ClientMock.Delete mock is already set by Set")
	}

	expectation := &ClientMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &ClientMockDeleteParams{releaseName},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up Client.Delete return parameters for the expectation previously defined by the When method
func (e *ClientMockDeleteExpectation) Then(err error) *ClientMock {
	e.results = &ClientMockDeleteResults{err}
	return e.mock
}

// Delete implements Client
func (mmDelete *ClientMock) Delete(releaseName string) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(releaseName)
	}

	mm_params := ClientMockDeleteParams{releaseName}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, &mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := ClientMockDeleteParams{releaseName}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("ClientMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the ClientMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(releaseName)
	}
	mmDelete.t.Fatalf("Unexpected call to ClientMock.Delete. %v", releaseName)
	return
}

// DeleteAfterCounter returns a count of finished ClientMock.Delete invocations
func (mmDelete *ClientMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of ClientMock.Delete invocations
func (mmDelete *ClientMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to ClientMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mClientMockDelete) Calls() []*ClientMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*ClientMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *ClientMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMock.Delete")
		} else {
			m.t.Errorf("Expected call to ClientMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to ClientMock.Delete")
	}
}

type mClientMockUpgrade struct {
	mock               *ClientMock
	defaultExpectation *ClientMockUpgradeExpectation
	expectations       []*ClientMockUpgradeExpectation

	callArgs []*ClientMockUpgradeParams
	mutex    sync.RWMutex
}

// ClientMockUpgradeExpectation specifies expectation struct of the Client.Upgrade
type ClientMockUpgradeExpectation struct {
	mock    *ClientMock
	params  *ClientMockUpgradeParams
	results *ClientMockUpgradeResults
	Counter uint64
}

// ClientMockUpgradeParams contains parameters of the Client.Upgrade
type ClientMockUpgradeParams struct {
	releaseName      string
	releaseNamespace string
	templates        map[string]interface{}
	values           map[string]interface{}
	debug            bool
	pr               []postrender.PostRenderer
}

// ClientMockUpgradeResults contains results of the Client.Upgrade
type ClientMockUpgradeResults struct {
	err error
}

// Expect sets up expected params for Client.Upgrade
func (mmUpgrade *mClientMockUpgrade) Expect(releaseName string, releaseNamespace string, templates map[string]interface{}, values map[string]interface{}, debug bool, pr ...postrender.PostRenderer) *mClientMockUpgrade {
	if mmUpgrade.mock.funcUpgrade != nil {
		mmUpgrade.mock.t.Fatalf("ClientMock.Upgrade mock is already set by Set")
	}

	if mmUpgrade.defaultExpectation == nil {
		mmUpgrade.defaultExpectation = &ClientMockUpgradeExpectation{}
	}

	mmUpgrade.defaultExpectation.params = &ClientMockUpgradeParams{releaseName, releaseNamespace, templates, values, debug, pr}
	for _, e := range mmUpgrade.expectations {
		if minimock.Equal(e.params, mmUpgrade.defaultExpectation.params) {
			mmUpgrade.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpgrade.defaultExpectation.params)
		}
	}

	return mmUpgrade
}

// Inspect accepts an inspector function that has same arguments as the Client.Upgrade
func (mmUpgrade *mClientMockUpgrade) Inspect(f func(releaseName string, releaseNamespace string, templates map[string]interface{}, values map[string]interface{}, debug bool, pr ...postrender.PostRenderer)) *mClientMockUpgrade {
	if mmUpgrade.mock.inspectFuncUpgrade != nil {
		mmUpgrade.mock.t.Fatalf("Inspect function is already set for ClientMock.Upgrade")
	}

	mmUpgrade.mock.inspectFuncUpgrade = f

	return mmUpgrade
}

// Return sets up results that will be returned by Client.Upgrade
func (mmUpgrade *mClientMockUpgrade) Return(err error) *ClientMock {
	if mmUpgrade.mock.funcUpgrade != nil {
		mmUpgrade.mock.t.Fatalf("ClientMock.Upgrade mock is already set by Set")
	}

	if mmUpgrade.defaultExpectation == nil {
		mmUpgrade.defaultExpectation = &ClientMockUpgradeExpectation{mock: mmUpgrade.mock}
	}
	mmUpgrade.defaultExpectation.results = &ClientMockUpgradeResults{err}
	return mmUpgrade.mock
}

// Set uses given function f to mock the Client.Upgrade method
func (mmUpgrade *mClientMockUpgrade) Set(f func(releaseName string, releaseNamespace string, templates map[string]interface{}, values map[string]interface{}, debug bool, pr ...postrender.PostRenderer) (err error)) *ClientMock {
	if mmUpgrade.defaultExpectation != nil {
		mmUpgrade.mock.t.Fatalf("Default expectation is already set for the Client.Upgrade method")
	}

	if len(mmUpgrade.expectations) > 0 {
		mmUpgrade.mock.t.Fatalf("Some expectations are already set for the Client.Upgrade method")
	}

	mmUpgrade.mock.funcUpgrade = f
	return mmUpgrade.mock
}

// When sets expectation for the Client.Upgrade which will trigger the result defined by the following
// Then helper
func (mmUpgrade *mClientMockUpgrade) When(releaseName string, releaseNamespace string, templates map[string]interface{}, values map[string]interface{}, debug bool, pr ...postrender.PostRenderer) *ClientMockUpgradeExpectation {
	if mmUpgrade.mock.funcUpgrade != nil {
		mmUpgrade.mock.t.Fatalf("ClientMock.Upgrade mock is already set by Set")
	}

	expectation := &ClientMockUpgradeExpectation{
		mock:   mmUpgrade.mock,
		params: &ClientMockUpgradeParams{releaseName, releaseNamespace, templates, values, debug, pr},
	}
	mmUpgrade.expectations = append(mmUpgrade.expectations, expectation)
	return expectation
}

// Then sets up Client.Upgrade return parameters for the expectation previously defined by the When method
func (e *ClientMockUpgradeExpectation) Then(err error) *ClientMock {
	e.results = &ClientMockUpgradeResults{err}
	return e.mock
}

// Upgrade implements Client
func (mmUpgrade *ClientMock) Upgrade(releaseName string, releaseNamespace string, templates map[string]interface{}, values map[string]interface{}, debug bool, pr ...postrender.PostRenderer) (err error) {
	mm_atomic.AddUint64(&mmUpgrade.beforeUpgradeCounter, 1)
	defer mm_atomic.AddUint64(&mmUpgrade.afterUpgradeCounter, 1)

	if mmUpgrade.inspectFuncUpgrade != nil {
		mmUpgrade.inspectFuncUpgrade(releaseName, releaseNamespace, templates, values, debug, pr...)
	}

	mm_params := ClientMockUpgradeParams{releaseName, releaseNamespace, templates, values, debug, pr}

	// Record call args
	mmUpgrade.UpgradeMock.mutex.Lock()
	mmUpgrade.UpgradeMock.callArgs = append(mmUpgrade.UpgradeMock.callArgs, &mm_params)
	mmUpgrade.UpgradeMock.mutex.Unlock()

	for _, e := range mmUpgrade.UpgradeMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpgrade.UpgradeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpgrade.UpgradeMock.defaultExpectation.Counter, 1)
		mm_want := mmUpgrade.UpgradeMock.defaultExpectation.params
		mm_got := ClientMockUpgradeParams{releaseName, releaseNamespace, templates, values, debug, pr}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpgrade.t.Errorf("ClientMock.Upgrade got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpgrade.UpgradeMock.defaultExpectation.results
		if mm_results == nil {
			mmUpgrade.t.Fatal("No results are set for the ClientMock.Upgrade")
		}
		return (*mm_results).err
	}
	if mmUpgrade.funcUpgrade != nil {
		return mmUpgrade.funcUpgrade(releaseName, releaseNamespace, templates, values, debug, pr...)
	}
	mmUpgrade.t.Fatalf("Unexpected call to ClientMock.Upgrade. %v %v %v %v %v %v", releaseName, releaseNamespace, templates, values, debug, pr)
	return
}

// UpgradeAfterCounter returns a count of finished ClientMock.Upgrade invocations
func (mmUpgrade *ClientMock) UpgradeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpgrade.afterUpgradeCounter)
}

// UpgradeBeforeCounter returns a count of ClientMock.Upgrade invocations
func (mmUpgrade *ClientMock) UpgradeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpgrade.beforeUpgradeCounter)
}

// Calls returns a list of arguments used in each call to ClientMock.Upgrade.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpgrade *mClientMockUpgrade) Calls() []*ClientMockUpgradeParams {
	mmUpgrade.mutex.RLock()

	argCopy := make([]*ClientMockUpgradeParams, len(mmUpgrade.callArgs))
	copy(argCopy, mmUpgrade.callArgs)

	mmUpgrade.mutex.RUnlock()

	return argCopy
}

// MinimockUpgradeDone returns true if the count of the Upgrade invocations corresponds
// the number of defined expectations
func (m *ClientMock) MinimockUpgradeDone() bool {
	for _, e := range m.UpgradeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpgradeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpgradeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpgrade != nil && mm_atomic.LoadUint64(&m.afterUpgradeCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpgradeInspect logs each unmet expectation
func (m *ClientMock) MinimockUpgradeInspect() {
	for _, e := range m.UpgradeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ClientMock.Upgrade with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpgradeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpgradeCounter) < 1 {
		if m.UpgradeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ClientMock.Upgrade")
		} else {
			m.t.Errorf("Expected call to ClientMock.Upgrade with params: %#v", *m.UpgradeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpgrade != nil && mm_atomic.LoadUint64(&m.afterUpgradeCounter) < 1 {
		m.t.Error("Expected call to ClientMock.Upgrade")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ClientMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockDeleteInspect()

			m.MinimockUpgradeInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ClientMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ClientMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteDone() &&
		m.MinimockUpgradeDone()
}
