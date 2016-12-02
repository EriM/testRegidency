package main

import (
	"encoding/json"
//	"fmt"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//type Dump struct {

//func (this *HDLS) dump() (*Dump, error) {

//func (this *HDLS) imprt(dump *Dump) (error) {

//func (this *HDLS) imprtJson(jsonStr string) (error) {

//------------------------
// createSchema 
//------------------------
func (this *VEHICLE) createSchema() error {
	this.logger.Infof("Call: createSchema")
	this.createEcuTestHistory()

	return nil
}

//------------------------
// 1. ProtoVehicle 
//------------------------
func (this *VEHICLE) createProtoVehicle() error {
	this.logger.Infof("Call: createProtoVehicle")

	err := this.stub.DeleteTable("ProtoVehicle")
	if err != nil {
		this.logger.Errorf("delete table ProtoVehicle error")
	}
	else {
    	this.logger.Debugf("delete table ProtoVehicle ok")
	}

	err = this.stub.CreateTable("ProtoVehicle", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"ProtoVehicleId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"ProtoVehicleName", shim.ColumnDefinition_STRING, false},
	})
	if err != nil {
		this.logger.Errorf("create table ProtoVehicle error")
		return err
	}
    this.logger.Debugf("create table ProtoVehicle ok")

	return nil
}

func (this *VEHICLE) putProtoVehicle(x *ProtoVehicle) error {
	this.logger.Infof("Call: putProtoVehicle")
	this.logger.Debugf("ProtoVehicleId           %s\n",x.ProtoVehicleId)
	this.logger.Debugf("ProtoVehicleName         %s\n",x.ProtoVehicleName)

	// InsertRow to ProtoVehicle table
	ok, err := this.stub.InsertRow("ProtoVehicle", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleName}},
		},
	})

	if err != nil {
		this.logger.Errorf("Failed inserting row.")
		return errors.New("Failed inserting row.")
	}
	if !ok {
		this.logger.Errorf("insertRow operation failed. Row with given key already exists")
		return errors.New("insertRow operation failed. Row with given key already exists")
	}
	this.logger.Debugf("InsertRow - ProtoVehicle OK")
	return nil
}

func (this *VEHICLE) addProtoVehicle(jsonStr string) error {
	this.logger.Infof("Call: addProtoVehicle")
	this.logger.Debugf("%s",jsonStr)

	var x ProtoVehicle 

	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		this.logger.Errorf("Failed Unmarshal JSON.")
		return err
	}
	return this.putProtoVehicle(&x)
}

func (this *VEHICLE) listProtoVehicle() (*[]ProtoVehicle, error) {
	this.logger.Infof("Call: listProtoVehicle")
	this.logger.Debugf("%s",jsonStr)

	var x ProtoVehicle
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		this.logger.Errorf("Failed Unmarshal JSON.")
		return nil,err
	}

	rowChannel, err := this.stub.GetRows("ProtoVehicle", nil)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed retrieveing . Error " + err.Error() + ". \"}"
		this.logger.Errorf("%s",jsonResp)
		return nil, errors.New(jsonResp)
	}

	var myResponces []ProtoVehicle
	
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				//
				var myRes ProtoVehicle
				myRes.ProtoVehicleId  = row.Columns[0].GetString_()
				myRes.ProtoVehicleName= row.Columns[1].GetString_()

				myResponces = append(myResponces, myRes)

				this.logger.Debugf("==================================================")
				this.logger.Debugf("ProtoVehicleId           %s\n",myRes.ProtoVehicleId)
				this.logger.Debugf("ProtoVehicleName         %s\n",myRes.ProtoVehicleName)

			}
		}
		if rowChannel == nil {
			break
		}
	}
	return &myResponces, nil
}

//------------------------
// 2. EcuTestHistory 
//------------------------
func (this *VEHICLE) createEcuTestHistory() error {
	this.logger.Infof("Call: createEcuTestHistory")

	err := this.stub.DeleteTable("EcuTestHistory")
	if err != nil {
		this.logger.Errorf("delete table EcuTestHistory error")
	}
	else {
    	this.logger.Debugf("delete table EcuTestHistory ok")
	}

	err = this.stub.CreateTable("EcuTestHistory", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"ProtoVehicleId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"ProtoVehicleName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"EcuId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"EcuName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"EcuVer", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"TestDate", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"TestLocation", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"SupplierId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"SupplierName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"TestType", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Remarks", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"TestVehicleId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"TestVehicleName", shim.ColumnDefinition_STRING, false},
	})
	if err != nil {
		this.logger.Errorf("create table EcuTestHistory error")
		return err
	}
    this.logger.Debugf("create table EcuTestHistory ok")

	return nil
}

func (this *VEHICLE) putEcuTestHistory(x *EcuTestHistory) error {
	this.logger.Infof("Call: putEcuTestHistory")
	this.logger.Debugf("ProtoVehicleId           %s\n",x.ProtoVehicleId)
	this.logger.Debugf("ProtoVehicleName         %s\n",x.ProtoVehicleName)
	this.logger.Debugf("EcuId                    %s\n",x.EcuId)
	this.logger.Debugf("EcuName                  %s\n",x.EcuName)
	this.logger.Debugf("EcuVer                   %s\n",x.EcuVer)
	this.logger.Debugf("TestDate                 %s\n",x.TestDate)
	this.logger.Debugf("TestLocation             %s\n",x.TestLocation)
	this.logger.Debugf("SupplierId               %s\n",x.SupplierId)
	this.logger.Debugf("SupplierName             %s\n",x.SupplierName)
	this.logger.Debugf("TestType                 %s\n",x.TestType)
	this.logger.Debugf("Remarks                  %s\n",x.Remarks)
	this.logger.Debugf("TestVehicleId            %s\n",x.TestVehicleId)
	this.logger.Debugf("TestVehicleName          %s\n",x.TestVehicleName)

	// InsertRow to EcuTestHistory table
	ok, err := this.stub.InsertRow("EcuTestHistory", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleName}},
			&shim.Column{Value: &shim.Column_String_{String_: x.EcuId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.EcuName}},
			&shim.Column{Value: &shim.Column_String_{String_: x.EcuVer}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestDate}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestLocation}},
			&shim.Column{Value: &shim.Column_String_{String_: x.SupplierId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.SupplierName}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestType}},
			&shim.Column{Value: &shim.Column_String_{String_: x.Remarks}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestVehicleId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestVehicleName}},
		},
	})

	if err != nil {
		this.logger.Errorf("Failed inserting row.")
		return errors.New("Failed inserting row.")
	}
	if !ok {
		this.logger.Errorf("insertRow operation failed. Row with given key already exists")
		return errors.New("insertRow operation failed. Row with given key already exists")
	}
	this.logger.Debugf("InsertRow - EcuTestHistory OK")
	return nil
}

func (this *VEHICLE) addEcuTestHistory(jsonStr string) error {
	this.logger.Infof("Call: addEcuTestHistory")
	this.logger.Debugf("%s",jsonStr)

	var x EcuTestHistory 

	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		this.logger.Errorf("Failed Unmarshal JSON.")
		return err
	}
	return this.putEcuTestHistory(&x)
}

//Query for EcuTestHistory
//	SearchKey
//		ProtoVehicleId			string
//		EcuId					string
//		SupplierName			string

func (this *VEHICLE) listEcuTestHistoryBySearchKey(jsonStr string) (*[]EcuTestHistory, error) {
	this.logger.Infof("Call: listEcuTestHistoryBySearchKey")
	this.logger.Debugf("%s",jsonStr)

	var x EcuTestHistory
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		this.logger.Errorf("Failed Unmarshal JSON.")
		return nil,err
	}

	// AccesControl
	// user := readAttributeName()
	// this.logger.Debugf("user = %s",user)
	// if user == "MK_A" {
	// 		x.SupplierId = nil
	// }
	// else {
	//		x.SupplierId = user
	// }
	this.logger.Debugf("ProtoVehicleId           %s\n",x.ProtoVehicleId)
	this.logger.Debugf("EcuId                    %s\n",x.EcuId)
	this.logger.Debugf("TestDate                 %s\n",x.TestDate)
	this.logger.Debugf("TestVehicleId            %s\n",x.TestVehicleId)
	this.logger.Debugf("SupplierId               %s\n",x.SupplierId)

	var columns []shim.Column
    col01 := shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleId  }}
    col02 := shim.Column{Value: &shim.Column_String_{String_: x.EcuId           }}
    col03 := shim.Column{Value: &shim.Column_String_{String_: x.TestDate        }}
    col04 := shim.Column{Value: &shim.Column_String_{String_: x.TestVehicleId   }}
    col05 := shim.Column{Value: &shim.Column_String_{String_: x.SupplierId      }}
    if len(x.ProtoVehicleId  )>0 { columns = append(columns, col01)}
    if len(x.EcuId           )>0 { columns = append(columns, col02)}
    if len(x.TestDate        )>0 { columns = append(columns, col03)}
    if len(x.TestVehicleId   )>0 { columns = append(columns, col04)}
    if len(x.SupplierName    )>0 { columns = append(columns, col05)}

    rowChannel, err := this.stub.GetRows("EcuTestHistory", columns)
//	rowChannel, err := this.stub.GetRows("EcuTestHistory", nil)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed retrieveing . Error " + err.Error() + ". \"}"
		this.logger.Errorf("%s",jsonResp)
		return nil, errors.New(jsonResp)
	}

	var myResponces []EcuTestHistory
	
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				//
				var myRes EcuTestHistory
				myRes.ProtoVehicleId  = row.Columns[0].GetString_()
				myRes.ProtoVehicleName= row.Columns[1].GetString_()
				myRes.EcuId           = row.Columns[2].GetString_()
				myRes.EcuName         = row.Columns[3].GetString_()
				myRes.EcuVer          = row.Columns[4].GetString_()
				myRes.TestDate        = row.Columns[5].GetString_()
				myRes.TestLocation    = row.Columns[6].GetString_()
				myRes.SupplierId      = row.Columns[7].GetString_()
				myRes.SupplierName    = row.Columns[8].GetString_()
				myRes.TestType        = row.Columns[9].GetString_()
				myRes.Remarks         = row.Columns[10].GetString_()
				myRes.TestVehicleId   = row.Columns[11].GetString_()
				myRes.TestVehicleName = row.Columns[12].GetString_()

				myResponces = append(myResponces, myRes)

				this.logger.Debugf("==================================================")
				this.logger.Debugf("ProtoVehicleId           %s\n",myRes.ProtoVehicleId)
				this.logger.Debugf("ProtoVehicleName         %s\n",myRes.ProtoVehicleName)
				this.logger.Debugf("EcuId                    %s\n",myRes.EcuId)
				this.logger.Debugf("EcuName                  %s\n",myRes.EcuName)
				this.logger.Debugf("EcuVer                   %s\n",myRes.EcuVer)
				this.logger.Debugf("TestDate                 %s\n",myRes.TestDate)
				this.logger.Debugf("TestLocation             %s\n",myRes.TestLocation)
				this.logger.Debugf("SupplierId               %s\n",myRes.SupplierId)
				this.logger.Debugf("SupplierName             %s\n",myRes.SupplierName)
				this.logger.Debugf("TestType                 %s\n",myRes.TestType)
				this.logger.Debugf("Remarks                  %s\n",myRes.Remarks)
				this.logger.Debugf("TestVehicleId            %s\n",myRes.TestVehicleId)
				this.logger.Debugf("TestVehicleName          %s\n",myRes.TestVehicleName)

			}
		}
		if rowChannel == nil {
			break
		}
	}
	return &myResponces, nil
}

//------------------------
// 3. EcuLatestTest 
//------------------------
func (this *VEHICLE) createEcuLatestTest() error {
	this.logger.Infof("Call: createEcuLatestTest")

	err := this.stub.DeleteTable("EcuLatestTest")
	if err != nil {
		this.logger.Errorf("delete table EcuLatestTest error")
	}
	else {
    	this.logger.Debugf("delete table EcuLatestTest ok")
	}

	err = this.stub.CreateTable("EcuLatestTest", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{"ProtoVehicleId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"ProtoVehicleName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"EcuId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"EcuName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"EcuVer", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"TestDate", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"TestLocation", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"SupplierId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"SupplierName", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"TestType", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"Remarks", shim.ColumnDefinition_STRING, false},
		&shim.ColumnDefinition{"TestVehicleId", shim.ColumnDefinition_STRING, true},
		&shim.ColumnDefinition{"TestVehicleName", shim.ColumnDefinition_STRING, false},
	})
	if err != nil {
		this.logger.Errorf("create table EcuLatestTest error")
		return err
	}
    this.logger.Debugf("create table EcuLatestTest ok")

	return nil
}

//
//func (this *VEHICLE) delEcuLatestTestByKey(jsonStr string) error {
//
//


func (this *VEHICLE) putEcuLatestTest(x *EcuLatestTest) error {
	this.logger.Infof("Call: EcuLatestTest")
	this.logger.Debugf("ProtoVehicleId           %s\n",x.ProtoVehicleId)
	this.logger.Debugf("ProtoVehicleName         %s\n",x.ProtoVehicleName)
	this.logger.Debugf("EcuId                    %s\n",x.EcuId)
	this.logger.Debugf("EcuName                  %s\n",x.EcuName)
	this.logger.Debugf("EcuVer                   %s\n",x.EcuVer)
	this.logger.Debugf("TestDate                 %s\n",x.TestDate)
	this.logger.Debugf("TestLocation             %s\n",x.TestLocation)
	this.logger.Debugf("SupplierId               %s\n",x.SupplierId)
	this.logger.Debugf("SupplierName             %s\n",x.SupplierName)
	this.logger.Debugf("TestType                 %s\n",x.TestType)
	this.logger.Debugf("Remarks                  %s\n",x.Remarks)
	this.logger.Debugf("TestVehicleId            %s\n",x.TestVehicleId)
	this.logger.Debugf("TestVehicleName          %s\n",x.TestVehicleName)

	// InsertRow to EcuLatestTest table
	ok, err := this.stub.InsertRow("EcuLatestTest", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleName}},
			&shim.Column{Value: &shim.Column_String_{String_: x.EcuId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.EcuName}},
			&shim.Column{Value: &shim.Column_String_{String_: x.EcuVer}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestDate}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestLocation}},
			&shim.Column{Value: &shim.Column_String_{String_: x.SupplierId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.SupplierName}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestType}},
			&shim.Column{Value: &shim.Column_String_{String_: x.Remarks}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestVehicleId}},
			&shim.Column{Value: &shim.Column_String_{String_: x.TestVehicleName}},
		},
	})

	if err != nil {
		this.logger.Errorf("Failed inserting row.")
		return errors.New("Failed inserting row.")
	}
	if !ok {
		this.logger.Errorf("insertRow operation failed. Row with given key already exists")
		return errors.New("insertRow operation failed. Row with given key already exists")
	}
	this.logger.Debugf("InsertRow - EcuLatestTest OK")
	return nil
}

func (this *VEHICLE) addEcuLatestTest(jsonStr string) error {
	this.logger.Infof("Call: addEcuLatestTest")
	this.logger.Debugf("%s",jsonStr)

	var x EcuLatestTest 

	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		this.logger.Errorf("Failed Unmarshal JSON.")
		return err
	}
	return this.putEcuLatestTest(&x)
}

//Query for EcuLatestTest
//	SearchKey
//		ProtoVehicleId			string
//		SupplierName			string

func (this *VEHICLE) listEcuLatestTestBySearchKey(jsonStr string) (*[]EcuLatestTest, error) {
	this.logger.Infof("Call: listEcuLatestTestBySearchKey")
	this.logger.Debugf("%s",jsonStr)

	var x EcuLatestTest
	err := json.Unmarshal([]byte(jsonStr), &x)
	if err != nil {
		this.logger.Errorf("Failed Unmarshal JSON.")
		return nil,err
	}

	// AccesControl
	// user := readAttributeName()
	// this.logger.Debugf("user = %s",user)
	// if user == "MK_A" {
	// 		x.SupplierId = nil
	// }
	// else {
	//		x.SupplierId = user
	// }
	this.logger.Debugf("ProtoVehicleId           %s\n",x.ProtoVehicleId)
	this.logger.Debugf("EcuId                    %s\n",x.EcuId)
	this.logger.Debugf("TestVehicleId            %s\n",x.TestVehicleId)
	this.logger.Debugf("SupplierId               %s\n",x.SupplierId)

	var columns []shim.Column
    col01 := shim.Column{Value: &shim.Column_String_{String_: x.ProtoVehicleId  }}
    col02 := shim.Column{Value: &shim.Column_String_{String_: x.EcuId           }}
    col03 := shim.Column{Value: &shim.Column_String_{String_: x.TestVehicleId   }}
    col04 := shim.Column{Value: &shim.Column_String_{String_: x.SupplierId      }}
    if len(x.ProtoVehicleId  )>0 { columns = append(columns, col01)}
    if len(x.EcuId           )>0 { columns = append(columns, col02)}
    if len(x.TestVehicleId   )>0 { columns = append(columns, col03)}
    if len(x.SupplierName    )>0 { columns = append(columns, col04)}

    rowChannel, err := this.stub.GetRows("EcuLatestTest", columns)
//	rowChannel, err := this.stub.GetRows("EcuLatestTest", nil)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed retrieveing . Error " + err.Error() + ". \"}"
		this.logger.Errorf("%s",jsonResp)
		return nil, errors.New(jsonResp)
	}

	var myResponces []EcuLatestTest
	
	for {
		select {
		case row, ok := <-rowChannel:
			if !ok {
				rowChannel = nil
			} else {
				//
				var myRes EcuLatestTest
				myRes.ProtoVehicleId  = row.Columns[0].GetString_()
				myRes.ProtoVehicleName= row.Columns[1].GetString_()
				myRes.EcuId           = row.Columns[2].GetString_()
				myRes.EcuName         = row.Columns[3].GetString_()
				myRes.EcuVer          = row.Columns[4].GetString_()
				myRes.TestDate        = row.Columns[5].GetString_()
				myRes.TestLocation    = row.Columns[6].GetString_()
				myRes.SupplierId      = row.Columns[7].GetString_()
				myRes.SupplierName    = row.Columns[8].GetString_()
				myRes.TestType        = row.Columns[9].GetString_()
				myRes.Remarks         = row.Columns[10].GetString_()
				myRes.TestVehicleId   = row.Columns[11].GetString_()
				myRes.TestVehicleName = row.Columns[12].GetString_()

				myResponces = append(myResponces, myRes)

				this.logger.Debugf("==================================================")
				this.logger.Debugf("ProtoVehicleId           %s\n",myRes.ProtoVehicleId)
				this.logger.Debugf("ProtoVehicleName         %s\n",myRes.ProtoVehicleName)
				this.logger.Debugf("EcuId                    %s\n",myRes.EcuId)
				this.logger.Debugf("EcuName                  %s\n",myRes.EcuName)
				this.logger.Debugf("EcuVer                   %s\n",myRes.EcuVer)
				this.logger.Debugf("TestDate                 %s\n",myRes.TestDate)
				this.logger.Debugf("TestLocation             %s\n",myRes.TestLocation)
				this.logger.Debugf("SupplierId               %s\n",myRes.SupplierId)
				this.logger.Debugf("SupplierName             %s\n",myRes.SupplierName)
				this.logger.Debugf("TestType                 %s\n",myRes.TestType)
				this.logger.Debugf("Remarks                  %s\n",myRes.Remarks)
				this.logger.Debugf("TestVehicleId            %s\n",myRes.TestVehicleId)
				this.logger.Debugf("TestVehicleName          %s\n",myRes.TestVehicleName)

			}
		}
		if rowChannel == nil {
			break
		}
	}
	return &myResponces, nil
}

