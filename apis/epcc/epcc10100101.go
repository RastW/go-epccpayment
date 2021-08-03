package epcc

type epcc10100101 struct {
	msgHander msgHeader `xml:"MsgHeader"`
	msgBody   msgBody   `xml:"MsgBody"`
}

type msgHeader struct {
	sndDt     string `xml:"SndDt"`
	msgTp     string `xml:"MsgTp"`
	issrId    string `xml:"IssrId"`
	drctn     string `xml:"Drctn"`
	signSN    string `xml:"SignSN"`
	ncrptnSN  string `xml:"NcrptnSN"`
	dgtlEnvlp string `xml:"DgtlEnvlp"`
}

type msgBody struct {
	sigint sigInf `xml:"SigInf"`
}

type sigInf struct {
	sgnAcctIssrId string `xml:"SgnAcctIssrId"`
	sgnAcctTp     string `xml:"SgnAcctTp"`
	sgnAcctId     string `xml:"SgnAcctId"`
	sgnAcctNm     string `xml:"SgnAcctNm"`
	iDTp          string `xml:"IDTp"`
	iDNo          string `xml:"IDNo"`
	mobNo         string `xml:"MobNo"`
}

type trxInf struct {
	trxCtgy  string `xml:"TrxCtgy"`
	trxId    string `xml:"Trxid"`
	trxDtTm  string `xml:"TrxDtTm"`
	authMsg  string `xml:"AuthMsg"`
	clbckUrl string `xml:"ClbckUrl"`
}
