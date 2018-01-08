package db

import "testing"

func TestNewContext(t *testing.T) {
	if testCtx.TableA == nil {
		t.Error("Context didn't initialize.")
	}
}

func TestSet_Create(t *testing.T) {
	row := newTestTable()
	row.Name = "ABC"
	row.ID = 99

	testCtx.TableA.Create(row)

	var records []*testTable
	err := testCtx.TableA.Read(testTable{}, &records)

	if err != nil {
		t.Error(err)
	}

	if len(records) < 1 {
		t.Error("Record wasn't added to Context")
	}
}

func TestSet_Read(t *testing.T) {
	row := newTestTable()
	row.Name = "ABC"
	row.ID = 99
	testCtx.TableA.Create(row)

	rowb := newTestTable()
	rowb.Name = "DEF"
	rowb.ID = 98
	rowb.Deleted = true
	testCtx.TableA.Create(rowb)

	var records []*testTable
	err := testCtx.TableA.Read(testTable{}, &records)

	if err != nil {
		t.Error(err)
	}

	for _, v := range records {
		if v.IsDeleted() {
			t.Error("records shouldn't contain deleted records.")
			break
		}
	}
}

func TestSet_ReadOne_Nil(t *testing.T) {
	row := newTestTable()
	row.Name = "ABC"
	row.ID = 55
	row.Deleted = true
	testCtx.TableA.Create(row)

	record, err := testCtx.TableA.ReadOne(testTable{Record: Record{ID: 55}})

	if err != nil {
		t.Error(err)
	}

	if record != nil {
		t.Error("Deleted records shouldn't be returned.")
	}
}

func TestSet_ReadOne(t *testing.T) {
	row := newTestTable()
	row.Name = "ABC"
	row.ID = 56
	testCtx.TableA.Create(row)

	record, err := testCtx.TableA.ReadOne(testTable{Record: Record{ID: 56}})

	if err != nil {
		t.Error(err)
	}

	if record == nil {
		t.Error("Added record not found.")
	}
}

func TestSet_Update(t *testing.T) {
	row := newTestTable()
	row.Name = "ABC"
	row.ID = 56
	testCtx.TableA.Create(row)

	row.Age = 18

	testCtx.TableA.Update(row)

	record, err := testCtx.TableA.ReadOne(testTable{Record: Record{ID: 56}})

	if err != nil {
		t.Error(err)
	}

	tblRec := record.(testTable)
	if tblRec.Age != 18 {
		t.Error("Record didn't update.")
	}
}

func TestSet_Delete(t *testing.T) {
	row := newTestTable()
	row.Name = "ABC"
	row.ID = 99
	row.Deleted = false

	testCtx.TableA.Create(row)

	record, err := testCtx.TableA.ReadOne(testTable{Record: Record{ID: 99}})

	if err != nil {
		t.Error(err)
	}

	if record == nil {
		t.Error("Record is nil")
	}

	record.Disable()

	if !record.IsDeleted() {
		t.Error("Record hasn't been set as Deleted")
	}
}
