package db

import "time"

type testContext struct {
	TableA *MemSet
}

type testTable struct {
	Record
	Name string
	Age  int
}

type testTableB struct {
	Record
	Details     string
	Relation    *testTable
	Collections []*testTable
}

var testCtx *testContext

func newTestTable() testTable {
	return testTable{
		Record: Record{
			ID:         0,
			Deleted:    false,
			CreateDate: time.Now(),
		},
	}
}

func init() {
	testCtx = &testContext{
		TableA: NewMemSet(testTable{}),
	}
}

func (t testTable) Validate() (bool, error) {
	return true, nil
}
