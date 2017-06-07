package main

import (
	"encoding/json"
	"errors"
	"fmt"
//	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// FFP is a high level smart contract that FFPs together business artifact based smart contracts
type FFP struct {

}

type BusinessEntity struct {	
	 CompanyId string `json:"companyId"`
	 DoingBusinessAsName string `json:"dbaName"`
	 Country string `json:"country"`
	 Address string `json:"address"`
	 BusinessType string `json:"businessType"` //Payee or Payor
	 CompanyStatus string `json:"companyStatus"`
}

// UserDetails is for storing User Details

type UserDetails struct{	
	UId string `json:"uId"`
	Title string `json:"title"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Country string `json:"country"`
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
	CreatedBy string `json:"createdBy"`
	UserType string `json:"userType"`
	UserStatus string `json:"userStatus"`
	CompanyId string `json:"companyId"`
}




// PayorAccount is for storing account

type PayorAccount struct{	
	 PayorAcctId string `json:"payorAcctId"`
	 CompanyId string `json:"companyId"`
	 RoutingNbr string `json:"routingNbr"`
	 AccountNbr string `json:"accountNbr"`
	 NameOnAcct string `json:"nameOnAcct"`
	 Country string `json:"country"`
	 Address string `json:"address"`
	 AccType string `json:"accType"`
	 AccStatus string `json:"accStatus"`
}

// PayeeAccount is for storing account

type PayeeAccount struct{
	 PayeeAccountId string `json:"payeeAccountId"`	
	 CompanyId string `json:"companyId"`
	 NameOnCheck string `json:"nameOnCheck"`
	 AccountNbr string `json:"accountNbr"`
	 NameOnAcct string `json:"nameOnAcct"`
	 Country string `json:"country"`
	 MailingAddress string `json:"mailingAddress"`
	 MailingZip string `json:"mailingZip"`
	 MailingCountry string `json:"mailingCountry"`
}

// Payment is for storing account
type Payment struct{	
	 SysPayId string `json:"sysPayId"`
	 CompanyId string `json:"companyId"`
	 PayorAcctId string `json:"payorAcctId"`
	 PayeeAccountId string `json:"payeeAccountId"`	
	 CreatTimeStamp string `json:"creatTimeStamp"`
	 LstUpdTimeStamp string `json:"lstUpdTimeStamp"`
	 PayAmt string `json:"payAmt"`
	 PayStatus string `json:"payStatus"`
	
}

// Relationship is for storing account
type Relationship struct{	
	 RelId string `json:"relId"`
	 PayorAcctId string `json:"payorAcctId"`
	 PayeeAccountId string `json:"payeeAccountId"`		
	 RelStatus string `json:"relStatus"`
	 
}

// Screen is for storing account
type ScreenCustomer struct{	
	 CustomerId string `json:"customerId"`
	 ScreenStatus string `json:"screenStatus"`	 
}

/* // GetMile is for storing retreived Get the total Points

type GetMile struct{	
	TotalPoint string `json:"totalPoint"`
}

// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
} */
	



// Init initializes the smart contracts
func (t *FFP) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

// Check if table already exists
	_, err := stub.GetTable("BusinessEntity")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("BusinessEntity", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "companyId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "dbaName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "businessType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "companyStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}

	// Check if table already exists
	_, err = stub.GetTable("UserDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("UserDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "uId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "userType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "userStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "companyId", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}
	
	

	

	
	
	// Check if table already exists
	_, err = stub.GetTable("PayorAccount")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("PayorAccount", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "payorAcctId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "companyId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "routingNbr", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "accountNbr", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nameOnAcct", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "accType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "accStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}

	// Check if table already exists
	_, err = stub.GetTable("PayeeAccount")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("PayeeAccount", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "payeeAccountId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "companyId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nameOnCheck", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "accountNbr", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nameOnAcct", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "mailingAddress", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "mailingZip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "mailingCountry", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
		// Check if table already exists
	_, err = stub.GetTable("Payment")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	
	
	// Create application Table
	err = stub.CreateTable("Payment", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "sysPayId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "companyId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "payorAccount", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "payeeAccount", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "creatTimeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lstUpdTimeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "payAmt", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "payStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}

		// Check if table already exists
	_, err = stub.GetTable("Relationship")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("Relationship", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "relId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "payorAccount", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "payeeAccount", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "relStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
	
	// Check if table already exists
	_, err = stub.GetTable("ScreenCustomer")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ScreenCustomer", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "customerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "screenStatus", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
		
	// setting up the users role
	stub.PutState("user_type1_1", []byte("etihad"))
	stub.PutState("user_type1_2", []byte("hertz"))
	stub.PutState("user_type1_3", []byte("marriot"))
	stub.PutState("user_type1_4", []byte("amazon"))	
	
	return nil, nil
}
	

	
//registerUser to register a user
func (t *FFP) registerUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 12 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 12. Got: %d.", len(args))
		}
		
		uId:=args[0]
		title:=args[1]
		firstName:=args[2]
		lastName:=args[3]
		email:=args[4]
		country:=args[5]
		address:=args[6]
		city:=args[7]
		zip:=args[8]
		createdBy:=args[9]
		userType:=args[10]
		userStatus:=args[11]
		companyId:=args[12]
		

		
		// Insert a row
		ok, err := stub.InsertRow("UserDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: uId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: userType}},
				&shim.Column{Value: &shim.Column_String_{String_: userStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: companyId}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}



func (t *FFP) addPayment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 5 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 5. Got: %d.", len(args))
		}
		
		fmt.Errorf("Correct Arg:", len(args))
		 
		sysPayId:=args[0]
		companyId:=args[1]
		payorAcctId:=args[2]
		payeeAccountId:=args[3]
		creatTimeStamp:=args[4]
		lstUpdTimeStamp:=args[5]
		payAmt:=args[6]
		payStatus:=args[7]
			
        fmt.Println("1========>",string(args[0]))
		fmt.Println("1========>",string(args[7]))
//		
//		assignerOrg1, err := stub.GetState(args[11])
//		assignerOrg := string(assignerOrg1)
//		
//		createdBy:=assignerOrg
//		totalPoint:="0"
//		
	//	err = stub.CreateTable("Payment", []*shim.ColumnDefinition{
	//	&shim.ColumnDefinition{Name: "payId", Type: shim.ColumnDefinition_STRING, Key: true},
	//	&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
	//	&shim.ColumnDefinition{Name: "payAmt", Type: shim.ColumnDefinition_STRING, Key: false},
	//	&shim.ColumnDefinition{Name: "payor", Type: shim.ColumnDefinition_STRING, Key: false},
	//	&shim.ColumnDefinition{Name: "payee", Type: shim.ColumnDefinition_STRING, Key: false},


		// Insert a row
		ok, err := stub.InsertRow("Payment", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: sysPayId}},
				&shim.Column{Value: &shim.Column_String_{String_: companyId}},
				&shim.Column{Value: &shim.Column_String_{String_: payorAcctId}},
				&shim.Column{Value: &shim.Column_String_{String_: payeeAccountId}},
				&shim.Column{Value: &shim.Column_String_{String_: creatTimeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: lstUpdTimeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: payAmt}},
				&shim.Column{Value: &shim.Column_String_{String_: payStatus}},
		 }})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}



/* // add or delete points and insert the transaction(irrespective of org)
func (t *FFP) addDeleteMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	trxId := args[0]
	timeStamp:=args[1]
	ffId := args[2]
	
	assignerOrg1, err := stub.GetState(args[3])
	assignerOrg := string(assignerOrg1)
	
	source := assignerOrg
	points := args[4]
	trxntype := args[5]
	trxnSubType := args[6]
	remarks := args[7]
	
	newPoints, _ := strconv.ParseInt(points, 10, 0)
	
	//whether ADD_PENDING, DELETE_PENDING 
	if trxnSubType == "ADD_PENDING" || trxnSubType == "DELETE_PENDING"{
		newPoints = 0
	}
	

	// Get the row pertaining to this ffid
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving user with ffid %s. Error %s", ffId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}

	newRoyaltyPoint := row.Columns[12].GetString_()
	
	if trxntype=="add"{
		earlierMile:=row.Columns[12].GetString_()
		earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
		newRoyaltyPoint = strconv.Itoa(int(earlierRoyalty) + int(newPoints))
	}else if trxntype=="delete"{
	
		earlierMile:=row.Columns[12].GetString_()
		earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
		newRoyaltiPointtoTest := int(earlierRoyalty) - int(newPoints)
		
		if newRoyaltiPointtoTest < 0 {
			return nil, errors.New("can't deduct as the resulting royalty becoming less than zero.")
		}
		newRoyaltyPoint = strconv.Itoa(int(earlierRoyalty) - int(newPoints))
	}else{
		return nil, fmt.Errorf("Error: Failed retrieving user with ffid %s. Error %s", ffId, err.Error())
	}
	
	
	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this ffid
	err = stub.DeleteRow(
		"UserDetails",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	
	//ffId := row.Columns[0].GetString_()
	
	title := row.Columns[1].GetString_()
	gender := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	dob := row.Columns[5].GetString_()
	email := row.Columns[6].GetString_()
	country := row.Columns[7].GetString_()
	address := row.Columns[8].GetString_()
	city := row.Columns[9].GetString_()
	zip := row.Columns[10].GetString_()
	createdBy := row.Columns[11].GetString_()
	totalPoint := newRoyaltyPoint


		// Insert a row
		ok, err := stub.InsertRow("UserDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: ffId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: totalPoint}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		
		//inserting the transaction
		
		// Insert a row
		ok, err = stub.InsertRow("Transaction", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: ffId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: points}},
				&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
				&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}		
	return nil, nil

} */


/* //get the miles against the ffid (irrespective of org)
func (t *FFP) getMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the ffId " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the ffId " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := GetMile{}
	
	res2E.TotalPoint = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

} */



/* //get all transaction against the ffid (depends on org)
func (t *FFP) getTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	assignerOrg1, err := stub.GetState(assignerRole)
	assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}	
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.FfId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Points = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.Remarks = row.Columns[7].GetString_()
		
		if newApp.FfId == ffId && newApp.Source == assignerOrg{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

} */




/* //get All transaction against ffid (irrespective of org)
func (t *FFP) getAllTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}	
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.FfId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Points = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.Remarks = row.Columns[7].GetString_()
		
		if newApp.FfId == ffId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

} */

/* //get All transaction against ffid (irrespective of org)
func (t *FFP) getAllPayments(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	payId := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Payment", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Payment{}	
	
	//PayID string `json:"payId"`
	//TimeStamp string `json:"timeStamp"`
	//PayAmt string `json:"payAmt"`
	//Payor string `json:"payor"`
	//Payee string `json:"payee"`
	
	for row := range rows {		
		newApp:= new(Payment)
		newApp.PayID = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.PayAmt = row.Columns[2].GetString_()
		newApp.Payor = row.Columns[3].GetString_()
		newApp.Payee = row.Columns[4].GetString_()
		//newApp.Trxntype = row.Columns[5].GetString_()
		//newApp.TrxnSubType = row.Columns[6].GetString_()
		//newApp.Remarks = row.Columns[7].GetString_()
		
		if newApp.PayID == payId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

} */


 // to get the deatils of a user against ffid (for internal testing, irrespective of org)
func (t *FFP) getUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	uId := args[0]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: uId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + uId + "\"}"
		return nil, errors.New(jsonResp)
	}



	
	
	res2E := UserDetails{}
	
	res2E.UId = row.Columns[0].GetString_()
	res2E.Title = row.Columns[1].GetString_()
	res2E.FirstName = row.Columns[2].GetString_()
	res2E.LastName = row.Columns[3].GetString_()
	res2E.Email = row.Columns[4].GetString_()
	res2E.Country = row.Columns[5].GetString_()
	res2E.Address = row.Columns[6].GetString_()
	res2E.City = row.Columns[7].GetString_()
	res2E.Zip = row.Columns[8].GetString_()
	res2E.CreatedBy = row.Columns[9].GetString_()
	res2E.UserType = row.Columns[10].GetString_()
	res2E.UserStatus = row.Columns[11].GetString_()
	res2E.CompanyId = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

} 

 // to get the deatils of a user against ffid (for internal testing, irrespective of org)
func (t *FFP) getPayment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	sysPayId := args[0]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: sysPayId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("Payment", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + sysPayId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + sysPayId + "\"}"
		return nil, errors.New(jsonResp)
	}
//payId:=args[0]
//		timeStamp:=args[1]
//		payAmt:=args[2]
//		payor:=args[3]
//		payee:=args[4]
	
	
	
	res2E := Payment{}
	
	res2E.SysPayId = row.Columns[0].GetString_()
	res2E.CompanyId = row.Columns[1].GetString_()
	res2E.PayorAcctId = row.Columns[2].GetString_()
	res2E.PayeeAccountId = row.Columns[3].GetString_()
	res2E.CreatTimeStamp = row.Columns[4].GetString_()
	res2E.LstUpdTimeStamp = row.Columns[5].GetString_()
	res2E.PayAmt = row.Columns[6].GetString_()
	res2E.PayStatus = row.Columns[7].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


/*
// verify the user is present or not (for internal testing, irrespective of org)
func (t *FFP) verifyUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	dob := args[1]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	userDob := row.Columns[5].GetString_()
	
	res2E := VerifyU{}
	
	if dob == userDob{
		res2E.Result="success"
	}else{
		res2E.Result="failed"
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

} */



// Invoke invokes the chaincode
func (t *FFP) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerUser" {
		t := FFP{}
		return t.registerUser(stub, args)	
	}else if function == "addPayment" { 
		t := FFP{}
		return t.addPayment(stub, args)
	}  
	/* else if function == "addDeleteMile" { 
		t := FFP{}
		return t.addDeleteMile(stub, args)
	} else if function == "addPayment" { 
		t := FFP{}
		return t.addPayment(stub, args)
	} */
	

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *FFP) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getUser" {
		t := FFP{}
		return t.getUser(stub, args)		
	}else if function == "getPayment" { 
		t := FFP{}
		return t.getPayment(stub, args)
	}
	/* else if function == "getTransaction" { 
		t := FFP{}
		return t.getTransaction(stub, args)
	}else if function == "getAllTransaction" { 
		t := FFP{}
		return t.getAllTransaction(stub, args)
	} else if function == "getUser" { 
		t := FFP{}
		return t.getUser(stub, args)
	}else if function == "verifyUser" { 
		t := FFP{}
		return t.verifyUser(stub, args) 
	}*/
	
	/* else if function == "getAllPayments" { 
		t := FFP{}
		return t.getAllPayments(stub, args)
	} */
	
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(FFP))
	if err != nil {
		fmt.Printf("Error starting FFP: %s", err)
	}
} 
