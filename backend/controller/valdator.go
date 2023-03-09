package controller

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var Val *validator.Validate

func ValidatorInit() error {
	Val = validator.New()
	if err := Val.RegisterValidation("valStuNum", valStuNum); err != nil {
		return err
	}
	if err := Val.RegisterValidation("valStuName", valStuName); err != nil {
		return err

	}
	if err := Val.RegisterValidation("valStuGender", valStuGender); err != nil {
		return err

	}
	if err := Val.RegisterValidation("valStuMajor", valStuMajor); err != nil {
		return err

	}
	if err := Val.RegisterValidation("valStuQq", valStuQq); err != nil {
		return err

	}
	if err := Val.RegisterValidation("valStuMobile", valStuMobile); err != nil {
		return err

	}
	if err := Val.RegisterValidation("valStuProvince", valStuProvince); err != nil {
		return err

	}
	if err := Val.RegisterValidation("valStuIntroduce", valStuIntroduce); err != nil {
		return err

	}
	if err := Val.RegisterValidation("valStuPassword", valStuPassword); err != nil {
		return err

	}
	return nil
}

func valStuNum(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^22(999|006|180|008)([0-9][0-9][0-9])$", fl.Field().String())
	return ok
}

func valStuName(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]{2,10}$", fl.Field().String())
	return ok
}
func valStuGender(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^女|男$", fl.Field().String())
	return ok
}
func valStuMajor(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^计算机科学与技术|计科师范|软件工程|网络工程$", fl.Field().String())
	return ok
}
func valStuQq(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^[1-9][0-9]{4,10}$", fl.Field().String())
	return ok
}
func valStuMobile(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^(?:(?:\\+|00)86)?1\\d{10}$", fl.Field().String())
	return ok
}
func valStuProvince(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^浙江|上海|北京|天津|重庆|黑龙江|吉林|辽宁|内蒙古|河北|新疆|甘肃|青海|陕西|宁夏|河南|山东|山西|安徽|湖北|湖南|江苏|四川|贵州|云南|广西|西藏|江西|广东|福建|台湾|海南|香港|澳门$", fl.Field().String())
	return ok
}

func valStuIntroduce(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^.{10,1000}$", fl.Field().String())
	return ok
}
func valStuPassword(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^[A-Za-z0-9]{8,10}$", fl.Field().String())
	return ok
}
