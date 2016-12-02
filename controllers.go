package main

import (
	"encoding/json"
//	"errors"
	"fmt"
//	"sort"
//	"strconv"

//	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const PREFIX = "[VEHICLE0000]"

type UserInfo struct {
	RoleId string
	UserId string
}

//func (this *HDLS) clearTableBondTermSheet(args []string) error {
//func (this *HDLS) clear_data(args []string) error {
//	this.clearTableBondTermSheet(args)
//	return errors.New("OK|Function clear_data completed.")
//}

func (this *VEHICLE) add_EcuTestHistory(args []string) error {
	this.logger.Infof(PREFIX+"Call: add_EcuTestHistory")
	this.logger.Debugf(PREFIX+"%s",args[0])

//	var userInfo UserInfo
//	err0 := json.Unmarshal([]byte(args[0]), &userInfo)
//	if err0 != nil {
//		fmt.Printf(ERR_PREFIX+"Failed to unmarshal args[0]\n")
//		return err0
//	}
	// Authentication
	//	if userInfo.RoleId != "admin" {

	var ETH EcuTestHistory

	err1 := json.Unmarshal([]byte(args[0]), &ETH)
	if err1 != nil {
		this.logger.Errorf(ERR_PREFIX+"Failed to unmarshal args[0]\n")
		return err1
	}
	// Data Verification

	// register EcuTestHistory
	err2 := this.addEcuTestHistory(args[0])
	if err2 != nil {
		this.logger.Errorf(ERR_PREFIX+"Failed to execute addEcuTestHistory\n")
		return err2
	}
	
	
	var ELTK EcuLatestTestKey
	ELTK.ProtoVehicleId	= ETH.ProtoVehicleId	
	ELTK.EcuId			= ETH.EcuId			
	ELTK.SupplierId		= ETH.SupplierId		
	ELTK.TestVehicleId	= ETH.TestVehicleId	
	
	
	
	
//	return errors.New("OK|Function add_EcuTestHistory completed.")
	this.logger.Debugf("OK|Function add_EcuTestHistory completed.")
	return nil
}


///////////////////////////////////// query ////////////////////////////////////////////

func (this *VEHICLE) get_EcuTestHistory_list(args []string) (*[]EcuTestHistory, error) {
//	var userInfo UserInfo
//	err0 := json.Unmarshal([]byte(args[0]), &userInfo)
//	if err0 != nil {
//		fmt.Printf(ERR_PREFIX+"Failed to unmarshal args[0]\n")
//		return err0
//	}
	// Authentication
//	if userInfo.RoleId != "admin" {

	this.logger.Infof("Call: get_EcuTestHistory_list")
	this.logger.Debugf("%s",args[0])

	EcuTestHistoryList, err := this.listEcuTestHistoryBySearchKey(args[0])
	if err != nil {
		return nil, err
	}

//	var ret EbBondArray
//	for _, lbts := range (*latestBondTermSheetList).Data {
//		bts, err4 := this.getBondTermSheetFromLatestBts(&lbts)
//		if err4 != nil {
//			return nil, err4
//		}
//
//		permitted := this.check_read_permission_by_state(bts.Status,userInfo.RoleId)

//		sort.Sort(ebBond.Indications)
//		ret = append(ret,ebBond)
//	}

	return EcuTestHistoryList, nil
}

