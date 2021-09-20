package flip

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

var mod *mold.Transformer
var val *validator.Validate

func init() {
	val = validator.New()
	val.RegisterValidationCtx("bank_code", validateBankCode)

	mod = modifiers.New()
	mod.Register("no_space", modNoSpace)
}

func validate(data interface{}) error {
	if err := mod.Struct(context.Background(), data); err != nil {
		return err
	}
	if err := val.Struct(data); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				return errRequiredField(e.Field())
			case "gt":
				return errGTField(e.Field(), e.Param())
			case "gte":
				return errGTEField(e.Field(), e.Param())
			case "max":
				return errMaxField(e.Field(), e.Param())
			case "numeric":
				return errNumericField(e.Field())
			case "bank_code":
				return errBankCodeField(e.Field())
			default:
				return errInvalidFormatField(e.Field())
			}
		}
		return err
	}
	return nil
}

func modNoSpace(ctx context.Context, fl mold.FieldLevel) error {
	switch fl.Field().Kind() {
	case reflect.String:
		fl.Field().SetString(strings.Replace(fl.Field().String(), " ", "", -1))
	}
	return nil
}

func validateBankCode(ctx context.Context, fl validator.FieldLevel) bool {
	return map[BankCode]bool{
		BankMandiri:           true,
		BankBRI:               true,
		BankBNI:               true,
		BankBCA:               true,
		BankBSM:               true,
		BankCIMB:              true,
		BankMuamalat:          true,
		BankDanamon:           true,
		BankPermata:           true,
		BankBII:               true,
		BankPanin:             true,
		BankUOB:               true,
		BankOCBC:              true,
		BankCitibank:          true,
		BankArtha:             true,
		BankTokyo:             true,
		BankDBS:               true,
		BankStandardChartered: true,
		BankCapital:           true,
		BankANZ:               true,
		BankBOC:               true,
		BankBumiArta:          true,
		BankHSBC:              true,
		BankRabobank:          true,
		BankMayapada:          true,
		BankBJB:               true,
		BankDKI:               true,
		BankDaerahIstimewa:    true,
		BankJateng:            true,
		BankJatim:             true,
		BankJambi:             true,
		BankSumut:             true,
		BankSumbar:            true,
		BankRiau:              true,
		BankSumsel:            true,
		BankLampung:           true,
		BankKalse:             true,
		BankKalbar:            true,
		BankKaltim:            true,
		BankKalteng:           true,
		BankSulselbar:         true,
		BankSulut:             true,
		BankNTB:               true,
		BankBali:              true,
		BankNTT:               true,
		BankMaluku:            true,
		BankPapua:             true,
		BankBengkulu:          true,
		BankSulawesi:          true,
		BankSultra:            true,
		BankNusantara:         true,
		BankIndia:             true,
		BankMestika:           true,
		BankSinarmas:          true,
		BankMaspion:           true,
		BankGanesha:           true,
		BankICBC:              true,
		BankQNB:               true,
		BankBTN:               true,
		BankWoori:             true,
		BankBTPN:              true,
		BankBTPNSyariah:       true,
		BankBJBSyariah:        true,
		BankMega:              true,
		BankBukopin:           true,
		BankBukopinSyariah:    true,
		BankJasaJakarta:       true,
		BankHana:              true,
		BankMNC:               true,
		BankAgroniaga:         true,
		BankSBI:               true,
		BankBCADigital:        true,
		BankNobu:              true,
		BankMegaSyariah:       true,
		BankInaPerdana:        true,
		BankSahabatSampoerna:  true,
		BankBKE:               true,
		BankBCASyariah:        true,
		BankJago:              true,
		BankMayora:            true,
		BankIndexSelindo:      true,
		BankVictoria:          true,
		BankIBK:               true,
		BankCTBC:              true,
		BankCommonwealth:      true,
		BankVictoriaSyariah:   true,
		BankBanten:            true,
		BankMutiara:           true,
		BankPaninSyariah:      true,
		BankAceh:              true,
		BankAntardaerah:       true,
		BankCCB:               true,
		BankCNB:               true,
		BankDinar:             true,
		BankEKA:               true,
		BankHarda:             true,
		BankMantap:            true,
		BankMAS:               true,
		BankPrima:             true,
		BankShinhan:           true,
		BankYudha:             true,
		BankGoPay:             true,
		BankOVO:               true,
		BankShopeePay:         true,
		BankDana:              true,
		BankLinkAja:           true,
	}[BankCode(fl.Field().String())]
}
