# sap-api-integrations-ehs-incident-reads  
sap-api-integrations-ehs-incident-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で EHSに関する データを取得するマイクロサービスです。  
sap-api-integrations-ehs-incident-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-ehs-incident-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_EHS_REPORT_INCIDENT_SRV_0001/overview

## 動作環境
sap-api-integrations-ehs-incident-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-ehs-incident-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-ehs-incident-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_EHS_REPORT_INCIDENT_SRV_0001/overview
* APIサービス名(=baseURL): API_EHS_REPORT_INCIDENT_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-ehs-incident-reads には、次の API をコールするためのリソースが含まれています。  

* Incident（EHS - 出来事）

## API への 値入力条件 の 初期値
sap-api-integrations-ehs-incident-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.Incident.IncidentTitle（出来事）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Incident" が指定されています。    
  
```
	"api_schema": "sap.s4.beh.functionallocation.v1.FunctionalLocation.Created.v1",
	"accepter": ["Incident"],
	"functional_location_code": "1010-CWS",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.functionallocation.v1.FunctionalLocation.Created.v1",
	"accepter": ["All"],
	"functional_location_code": "1010-CWS",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetFunctionalLocation(functionalLocation, functionalLocationName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(functionalLocation)
				wg.Done()
			}()
		case "FunctionalLocationName":
			func() {
				c.FunctionalLocationName(functionalLocationName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 機能場所 の ヘッダ が取得された結果の JSON の例です。  
以下の項目のうち、"FunctionalLocation" ～ "CreationDate" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-ehs-incident-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-ehs-incident-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"FunctionalLocation": "1010-CWS",
			"FunctionalLocationLabelName": "1010-CWS",
			"FunctionalLocationName": "Cooling Water System",
			"FuncLocationStructure": "YBPM",
			"FunctionalLocationCategory": "M",
			"SuperiorFunctionalLocation": "1010",
			"SuperiorFuncnlLocLabelName": "1010",
			"TechnicalObjectType": "",
			"GrossWeight": "0.000",
			"GrossWeightUnit": "",
			"OperationStartDate": "",
			"InventoryNumber": "",
			"AcquisitionValue": "0.00",
			"Currency": "",
			"AcquisitionDate": "",
			"AssetManufacturerName": "",
			"ManufacturerPartNmbr": "",
			"ManufacturerCountry": "",
			"ManufacturerPartTypeName": "",
			"ConstructionMonth": "",
			"ConstructionYear": "",
			"ManufacturerSerialNumber": "",
			"MaintenancePlant": "1010",
			"AssetLocation": "YB_1006",
			"AssetRoom": "",
			"PlantSection": "YIS",
			"WorkCenter": "",
			"WorkCenterInternalID": "0",
			"WorkCenterPlant": "",
			"ABCIndicator": "",
			"MaintObjectFreeDefinedAttrib": "",
			"FormOfAddress": "",
			"BusinessPartnerName1": "",
			"BusinessPartnerName2": "",
			"CityName": "",
			"PostalCode": "",
			"StreetName": "",
			"Region": "",
			"Country": "",
			"PhoneNumber": "",
			"FaxNumber": "",
			"CompanyCode": "1010",
			"BusinessArea": "",
			"MasterFixedAsset": "",
			"FixedAsset": "",
			"CostCenter": "10101301",
			"ControllingArea": "A000",
			"WBSElementExternalID": "",
			"SettlementOrder": "",
			"ConstructionMaterial": "",
			"MaintenancePlannerGroup": "920",
			"MaintenancePlanningPlant": "1010",
			"MainWorkCenterPlant": "1010",
			"MainWorkCenter": "RES-0100",
			"MainWorkCenterInternalID": "10000026",
			"CatalogProfile": "YB-PM0003",
			"EquipmentInstallationIsAllowed": true,
			"OnePieceOfEquipmentIsAllowed": false,
			"SalesOrganization": "",
			"DistributionChannel": "",
			"SalesOffice": "",
			"OrganizationDivision": "",
			"SalesGroup": "",
			"FunctionalLocationHasEquipment": "",
			"FuncnlLocHasSubOrdinateFuncLoc": "X",
			"LastChangeDateTime": "",
			"FuncnlLocIsMarkedForDeletion": false,
			"FuncnlLocIsDeleted": false,
			"FunctionalLocationIsActive": true,
			"FuncnlLocIsDeactivated": false,
			"CreationDate": "2016-06-24T09:00:00+09:00"
		}
	],
	"time": "2022-01-28T16:52:14+09:00"
}
```
