package main

type ProtoVehicle struct {
		ProtoVehicleId			string	//@PK
		ProtoVehicleName		string
}

type EcuTestHistory struct {
		ProtoVehicleId			string	//@PK
		ProtoVehicleName		string
		EcuId					string	//@PK
		EcuName					string
		EcuVer					string
		TestDate				string	//@PK
		TestLocation			string
		SupplierId				string	//@PK
		SupplierName			string
		TestType				string
		Remarks					string
		TestVehicleId			string	//@PK
		TestVehicleName			string
}

type EcuLatestTest struct {
		ProtoVehicleId			string	//@PK
		ProtoVehicleName		string
		EcuId					string	//@PK
		EcuName					string
		EcuVer					string
		TestDate				string
		TestLocation			string
		SupplierId				string	//@PK
		SupplierName			string
		TestType				string
		Remarks					string
		TestVehicleId			string	//@PK
		TestVehicleName			string
}

type EcuLatestTestKey struct {
		ProtoVehicleId			string	//@PK
		EcuId					string	//@PK
		SupplierId				string	//@PK
		TestVehicleId			string	//@PK
}
