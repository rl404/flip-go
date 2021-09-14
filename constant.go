package flip

// EnvironmentType is type for environment.
type EnvironmentType int8

// Available options for EnvironmentType.
const (
	Sandbox EnvironmentType = iota
	Production
)

var envURL = map[EnvironmentType]string{
	Sandbox:    "https://bigflip.id/big_sandbox_api/v2",
	Production: "https://bigflip.id/api/v2",
}

var envLog = map[EnvironmentType]LogLevel{
	Sandbox:    LogDebug,
	Production: LogError,
}

// BankStatus is type for bank status.
type BankStatus string

// Available options for BankStatus.
const (
	BankStatusOperational      BankStatus = "OPERATIONAL"
	BankStatusDisturbed        BankStatus = "DISTURBED"
	BankStatusHeavilyDisturbed BankStatus = "HEAVILY_DISTURBED"
)

// AccountStatus is type for bank account status.
type AccountStatus string

// Available options for AccountStatus.
const (
	AccountStatusPending              AccountStatus = "PENDING"
	AccountStatusSuccess              AccountStatus = "SUCCESS"
	AccountStatusInvalidAccountNumber AccountStatus = "INVALID_ACCOUNT_NUMBER"
	AccountStatusSuspectedAccount     AccountStatus = "SUSPECTED_ACCOUNT"
	AccountStatusBlackListed          AccountStatus = "BLACK_LISTED"
)

// TransactionStatus is type for disbursement transaction status.
type TransactionStatus string

// Available options for TransactionStatus.
const (
	TransactionStatusPending   TransactionStatus = "PENDING"
	TransactionStatusCancelled TransactionStatus = "CANCELLED"
	TransactionStatusDone      TransactionStatus = "DONE"
)

// Source is type for disbursement created from.
type Source string

// Available options for Source.
const (
	SourceAPI       Source = "API"
	SourceDashboard Source = "DASHBOARD"
)

// Direction is type for transaction direction.
type Direction string

// Available options for Direction.
const (
	DirectionDomestic        Direction = "DOMESTIC_TRANSFER"
	DirectionDomesticSpecial Direction = "DOMESTIC_SPECIAL_TRANSFER"
	DirectionForeign         Direction = "FOREIGN_INBOUND_SPECIAL_TRANSFER"
)

// IdentityType is type for sender identity type.
type IdentityType string

// Available options for IdentityType.
const (
	IdentityNationalID   IdentityType = "nat_id"
	IdentityDriveLicense IdentityType = "drv_lic"
	IdentityPassport     IdentityType = "passport"
	IdentityBankAccount  IdentityType = "bank_acc"
)

// JobType is type for sender job.
type JobType string

// Available options for JobType.
const (
	JobHouseWife               JobType = "housewife"
	JobEntrepreneur            JobType = "entrepreneur"
	JobPrivateEmployee         JobType = "private_employee"
	JobGovernmentEmployee      JobType = "government_employee"
	JobFoundationBoard         JobType = "foundation_board"
	JobIndonesianMigrantWorker JobType = "indonesian_migrant_worker"
	JobCompany                 JobType = "company"
	JobOthers                  JobType = "others"
)

// Sort is type for get disbursement list sorting type.
type Sort string

// Available options for Sort.
const (
	SortIDAsc             Sort = "id"
	SortIDDesc            Sort = "-id"
	SortAmountAsc         Sort = "amount"
	SortAmountDesc        Sort = "-amount"
	SortStatusAsc         Sort = "status"
	SortStatusDesc        Sort = "-status"
	SortTimestampAsc      Sort = "timestamp"
	SortTimestampDesc     Sort = "-timestamp"
	SortBankCodeAsc       Sort = "bank_code"
	SortBankCodeDesc      Sort = "-bank_code"
	SortAccountNumberAsc  Sort = "account_number"
	SortAccountNumberDesc Sort = "-account_number"
	SortRecipientNameAsc  Sort = "recipient_name"
	SortRecipientNameDesc Sort = "-recipient_name"
	SortRemarkAsc         Sort = "remark"
	SortRemarkDesc        Sort = "-remark"
	SortTimeServedAsc     Sort = "time_served"
	SortTimeServedDesc    Sort = "-time_served"
	SortCreatedFromAsc    Sort = "created_from"
	SortCreatedFromDesc   Sort = "-created_from"
	SortDirectionAsc      Sort = "direction"
	SortDirectionDesc     Sort = "-direction"
)

// BankCode is code for banks.
type BankCode string

// Available options for BankCode.
const (
	BankMandiri           BankCode = "mandiri"                     //	Bank Mandiri
	BankBRI               BankCode = "bri"                         //	Bank Rakyat Indonesia
	BankBNI               BankCode = "bni"                         //	BNI (Bank Negara Indonesia) & BNI Syariah
	BankBCA               BankCode = "bca"                         //	Bank Central Asia
	BankBSM               BankCode = "bsm"                         //	Bank Syariah Mandiri/BSI
	BankCIMB              BankCode = "cimb"                        //	CIMB Niaga & CIMB Niaga Syariah
	BankMuamalat          BankCode = "muamalat"                    //	Muamalat
	BankDanamon           BankCode = "danamon"                     //	Bank Danamon & Danamon Syariah
	BankPermata           BankCode = "permata"                     //	Bank Permata & Permata Syariah
	BankBII               BankCode = "bii"                         //	Maybank Indonesia
	BankPanin             BankCode = "panin"                       //	Panin Bank
	BankUOB               BankCode = "uob"                         //	TMRW/UOB
	BankOCBC              BankCode = "ocbc"                        //	OCBC NISP
	BankCitibank          BankCode = "citibank"                    //	Citibank
	BankArtha             BankCode = "artha"                       //	Bank Artha Graha Internasional
	BankTokyo             BankCode = "tokyo"                       //	Bank of Tokyo Mitsubishi UFJ
	BankDBS               BankCode = "dbs"                         //	DBS Indonesia
	BankStandardChartered BankCode = "standard_chartered"          //	Standard Chartered Bank
	BankCapital           BankCode = "capital"                     //	Bank Capital Indonesia
	BankANZ               BankCode = "anz"                         //	ANZ Indonesia
	BankBOC               BankCode = "boc"                         //	Bank of China (Hong Kong) Limited
	BankBumiArta          BankCode = "bumi_arta"                   //	Bank Bumi Arta
	BankHSBC              BankCode = "hsbc"                        //	HSBC Indonesia
	BankRabobank          BankCode = "rabobank"                    //	Rabobank International Indonesia
	BankMayapada          BankCode = "mayapada"                    //	Bank Mayapada
	BankBJB               BankCode = "bjb"                         //	BJB
	BankDKI               BankCode = "dki"                         //	Bank DKI Jakarta
	BankDaerahIstimewa    BankCode = "daerah_istimewa"             //	BPD DIY
	BankJateng            BankCode = "jawa_tengah"                 //	Bank Jateng
	BankJatim             BankCode = "jawa_timur"                  //	Bank Jatim
	BankJambi             BankCode = "jambi"                       //	Bank Jambi
	BankSumut             BankCode = "sumut"                       //	Bank Sumut
	BankSumbar            BankCode = "sumatera_barat"              //	Bank Sumbar (Bank Nagari)
	BankRiau              BankCode = "riau_dan_kepri"              //	Bank Riau Kepri
	BankSumsel            BankCode = "sumsel_dan_babel"            //	Bank Sumsel Babel
	BankLampung           BankCode = "lampung"                     //	Bank Lampung
	BankKalse             BankCode = "kalimantan_selatan"          //	Bank Kalsel
	BankKalbar            BankCode = "kalimantan_barat"            //	Bank Kalbar
	BankKaltim            BankCode = "kalimantan_timur"            //	Bank Kaltim
	BankKalteng           BankCode = "kalimantan_tengah"           //	Bank Kalteng
	BankSulselbar         BankCode = "sulselbar"                   //	Bank Sulselbar
	BankSulut             BankCode = "sulut"                       //	Bank SulutGo
	BankNTB               BankCode = "nusa_tenggara_barat"         //	Bank NTB
	BankBali              BankCode = "bali"                        //	BPD Bali
	BankNTT               BankCode = "nusa_tenggara_timur"         //	Bank NTT
	BankMaluku            BankCode = "maluku"                      //	Bank Maluku
	BankPapua             BankCode = "papua"                       //	Bank Papua
	BankBengkulu          BankCode = "bengkulu"                    //	Bank Bengkulu
	BankSulawesi          BankCode = "sulawesi"                    //	Bank Sulteng
	BankSultra            BankCode = "sulawesi_tenggara"           //	Bank Sultra
	BankNusantara         BankCode = "nusantara_parahyangan"       //	Bank Nusantara Parahyangan
	BankIndia             BankCode = "india"                       //	Bank of India Indonesia
	BankMestika           BankCode = "mestika_dharma"              //	Bank Mestika Dharma
	BankSinarmas          BankCode = "sinarmas"                    //	Bank Sinarmas
	BankMaspion           BankCode = "maspion"                     //	Bank Maspion Indonesia
	BankGanesha           BankCode = "ganesha"                     //	Bank Ganesha
	BankICBC              BankCode = "icbc"                        //	ICBC Indonesia
	BankQNB               BankCode = "qnb_kesawan"                 //	QNB Indonesia
	BankBTN               BankCode = "btn"                         //	BTN (Bank Tabungan Negara)
	BankWoori             BankCode = "woori"                       //	Bank Woori Saudara
	BankBTPN              BankCode = "tabungan_pensiunan_nasional" //	BTPN
	BankBTPNSyariah       BankCode = "btpn_syr"                    //	Bank BTPN Syariah
	BankBJBSyariah        BankCode = "bjb_syr"                     //	BJB Syariah
	BankMega              BankCode = "mega"                        //	Bank Mega
	BankBukopin           BankCode = "bukopin"                     //	Wokee/Bukopin
	BankBukopinSyariah    BankCode = "bukopin_syr"                 //	Bank Bukopin Syariah
	BankJasaJakarta       BankCode = "jasa_jakarta"                //	Bank Jasa Jakarta
	BankHana              BankCode = "hana"                        //	LINE Bank/KEB Hana
	BankMNC               BankCode = "mnc_internasional"           //	Motion/MNC Bank
	BankAgroniaga         BankCode = "agroniaga"                   //	BRI Agroniaga
	BankSBI               BankCode = "sbi_indonesia"               //	SBI Indonesia
	BankBCADigital        BankCode = "royal"                       //	Blu/BCA Digital
	BankNobu              BankCode = "nationalnobu"                //	Nobu (Nationalnobu) Bank
	BankMegaSyariah       BankCode = "mega_syr"                    //	Bank Mega Syariah
	BankInaPerdana        BankCode = "ina_perdana"                 //	Bank Ina Perdana
	BankSahabatSampoerna  BankCode = "sahabat_sampoerna"           //	Bank Sahabat Sampoerna
	BankBKE               BankCode = "kesejahteraan_ekonomi"       //	Seabank/Bank BKE
	BankBCASyariah        BankCode = "bca_syr"                     //	BCA (Bank Central Asia) Syariah
	BankJago              BankCode = "artos"                       //	Jago/Artos
	BankMayora            BankCode = "mayora"                      //	Bank Mayora Indonesia
	BankIndexSelindo      BankCode = "index_selindo"               //	Bank Index Selindo
	BankVictoria          BankCode = "victoria_internasional"      //	Bank Victoria International
	BankIBK               BankCode = "agris"                       //	Bank IBK Indonesia
	BankCTBC              BankCode = "chinatrust"                  //	CTBC (Chinatrust) Indonesia
	BankCommonwealth      BankCode = "commonwealth"                //	Commonwealth Bank
	BankVictoriaSyariah   BankCode = "victoria_syr"                //	Bank Victoria Syariah
	BankBanten            BankCode = "banten"                      //	BPD Banten
	BankMutiara           BankCode = "mutiara"                     //	Bank Mutiara
	BankPaninSyariah      BankCode = "panin_syr"                   //	Panin Dubai Syariah
	BankAceh              BankCode = "aceh"                        //	Bank Aceh
	BankAntardaerah       BankCode = "antardaerah"                 //	Bank Antardaerah
	BankCCB               BankCode = "ccb"                         //	Bank China Construction Bank Indonesia
	BankCNB               BankCode = "cnb"                         //	Bank CNB (Centratama Nasional Bank)
	BankDinar             BankCode = "dinar"                       //	Bank Dinar Indonesia
	BankEKA               BankCode = "eka"                         //	BPR EKA (Bank Eka)
	BankHarda             BankCode = "harda"                       //	Allo Bank/Bank Harda Internasional
	BankMantap            BankCode = "mantap"                      //	BANK MANTAP (Mandiri Taspen)
	BankMAS               BankCode = "mas"                         //	Bank Multi Arta Sentosa (Bank MAS)
	BankPrima             BankCode = "prima"                       //	Bank Prima Master
	BankShinhan           BankCode = "shinhan"                     //	Bank Shinhan Indonesia
	BankYudha             BankCode = "yudha_bakti"                 //	Neo Commerce/Yudha Bhakti
	BankGoPay             BankCode = "gopay"                       //	GoPay
	BankOVO               BankCode = "ovo"                         //	OVO
	BankShopeePay         BankCode = "shopeepay"                   //	ShopeePay
	BankDana              BankCode = "dana"                        //	Dana
	BankLinkAja           BankCode = "linkaja"                     //	LinkAja
)
