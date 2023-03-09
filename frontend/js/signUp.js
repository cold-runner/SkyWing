// 学号校验
var stuNum = document.getElementById('stuNum')
var numV = document.getElementById('stuNumV')
stuNum.addEventListener("input", function (event) {
	var numPattern = /^22(999|006|180|008)([0-9][0-9][0-9])$/
	if (!numPattern.test(stuNum.value)) {
		numV.style.display = "block"
	} else {
		numV.style.display = "none"
	}
})
// 姓名校验
var stuName = document.getElementById('stuName')
var nameV = document.getElementById('stuNameV')
stuName.addEventListener("input", function (event) {
	var namePattern = /^[\u4E00-\u9FA5]{2,10}$/
	if (!namePattern.test(stuName.value)) {
		nameV.style.display = "block"
	} else {
		nameV.style.display = "none"
	}
})
// 性别校验
var stuGender = document.getElementById('stuGender')
var genderV = document.getElementById('stuGenderV')
stuGender.addEventListener("input", function (event) {
	var genderPattern = /^女|男$/
	if (!genderPattern.test(stuGender.value)) {
		genderV.style.display = "block"
	} else {
		genderV.style.display = "none"
	}
})
// 班级校验
var stuClass = document.getElementById('stuClass')
var classV = document.getElementById('stuClassV')
stuClass.addEventListener("input", function (event) {
	var classPattern = /^(1|2|3|4|5|6|7|11|12|13|14|15|16)$/
	if (!classPattern.test(stuClass.value)) {
		classV.style.display = "block"
	} else {
		classV.style.display = "none"
	}
})
// 省份校验
var stuProvince = document.getElementById('stuProvince')
var provinceV = document.getElementById('stuProvinceV')
stuProvince.addEventListener("change", function (event) {
	var provincePattern = /^浙江|上海|北京|天津|重庆|黑龙江|吉林|辽宁|内蒙古|河北|新疆|甘肃|青海|陕西|宁夏|河南|山东|山西|安徽|湖北|湖南|江苏|四川|贵州|云南|广西|西藏|江西|广东|福建|台湾|海南|香港|澳门$/
	if (!provincePattern.test(stuProvince.value)) {
		provinceV.style.display = "block"
	} else {
		provinceV.style.display = "none"
	}
})
// qq校验
var stuQq = document.getElementById('stuQq')
var qqV = document.getElementById('stuQqV')
stuQq.addEventListener("input", function (event) {
	var qqPattern = /^[1-9][0-9]{4,10}$/
	if (!qqPattern.test(stuQq.value)) {
		qqV.style.display = "block"
	} else {
		qqV.style.display = "none"
	}
})
// 专业校验
var stuMajor = document.getElementById('stuMajor')
var majorV = document.getElementById('stuMajorV')
stuMajor.addEventListener("input", function (event) {
	var majorPattern = /^计算机科学与技术|计科师范|软件工程|网络工程$/
	if (!majorPattern.test(stuMajor.value)) {
		majorV.style.display = "block"
	} else {
		majorV.style.display = "none"

	}
})
// 照片校验
var preview = document.getElementById('preview')
var stuPhoto = document.getElementById('stuPhoto')
var photoV = document.getElementById('photoV')
// 获取 FileReader 的实例
const reader = new FileReader()

stuPhoto.addEventListener("change", function (event) {
	// 获取所选文件
	const file = event.target.files[0]
	// 创建引用该文件的新 URL
	const url = URL.createObjectURL(file)
	preview.src = url

	// 上传后获取文件对象，以 URL 二进制字符串的形式读取数据
	reader.readAsDataURL(file)
	// 加载后，对字符串进行处理
	reader.addEventListener('load', (e) => {
		// 设置预览元素的源
		preview.src = reader.result
	})
	if (file.size > 1024 * 1024 * 2) {
		photoV.style.display = "block"
	} else {
		photoV.style.display = "none"
	}
})
function checkFile() {
	fileSize = stuPhoto.files[0].size
	if (preview.src.length == 0 || fileSize > 1024 * 1024 * 2) {
		return false;
	}
	else {
		return true;
	}
}
// 校验日期
var stuBirth = document.getElementById('stuBirth')
var birthV = document.getElementById('stuBirthV')
stuBirth.addEventListener("input", function (event) {
	birthPattern = /^[0-9]{4}-(((0[13578]|(10|12))-(0[1-9]|[1-2][0-9]|3[0-1]))|(02-(0[1-9]|[1-2][0-9]))|((0[469]|11)-(0[1-9]|[1-2][0-9]|30)))$/
	if (!birthPattern.test(stuBirth.value)) {
		birthV.style.display = "block"
	} else {
		birthV.style.display = "none"
	}
})
// 校验手机号
var stuPhone = document.getElementById('stuPhone')
var phoneV = document.getElementById('stuPhoneV')
stuPhone.addEventListener("input", function (event) {
	var phonePattern = /^(?:(?:\+|00)86)?1\d{10}$/
	if (!phonePattern.test(stuPhone.value)) {
		phoneV.style.display = "block"
	} else {
		phoneV.style.display = "none"
	}
})
// 校验长文本输入
var smscode = document.getElementById('smsCode')
var smsV = document.getElementById('smsV')
smscode.addEventListener("input", function (event) {
	var smsPattern = /^[0-9]{4}$/
	if (!smsPattern.test(smscode.value)) {
		smsV.style.display = "block"
	} else {
		smsV.style.display = "none"
	}
})
var sendButton = document.getElementById('sendSMSButton')
// 发送验证码
function sendSMS(obj) {
	var phone = $("#stuPhone").val();
	var result = isPhoneNum(phone);
	if (result) {
		$.ajax({
			url: "https://skylab.org.cn/api/v1/sendSmsCode",
			data: { mobile: phone },
			dataType: "json",
			type: "post",
			async: true,
			cache: false,
			success: function (res) {
				if (res.code == 1016) {
					alert(res.message)
				}
				if (res.code == 1013) {
					alert(res.message)
				}
				if (res.code == 1014) {
					alert(res.message)
				}
				if (res.code == 2333) {
					setTime(sendButton);
				}
			},
			error: function () {
				alert('错误！')
			}
		})
		//开始倒计时
		setTime(sendButton);
	}
}

// 校验手机号
function isPhoneNum(obj) {
	//正则校验
	var reg = /^(?:(?:\+|00)86)?1\d{10}$/
	if (!reg.test(obj)) {
		phoneV.style.display = "block"
		return false;
	} else {
		phoneV.style.display = "true"
		return true;
	}
}
// 计时器
var countdown = 60;
function setTime(obj) {
	if (countdown == 0) {
		obj.disabled = false
		obj.value = "发送验证码";
		//60秒过后button上的文字初始化,计时器初始化;
		countdown = 60;
		return;
	} else {
		obj.disabled = true;
		obj.value = countdown + "s后重新发送";
		countdown--;
	}
	//每1000毫秒执行一次
	setTimeout(function () { setTime(sendButton) }, 1000);
}

function patternValidate(pattern, obj) {
	if (!pattern.test(obj.value)) {
		return false
	} else {
		return true
	}
}


window.onload = function () {
	// 绑定点击事件，发送数据到后台服务器
	document.getElementById("signUp").addEventListener('click', sendInfo);
	// 发送数据
	function sendInfo() {
		// 检查所有数据是否合法
		if (!checkFile(stuPhoto)) {
			photoV.style.display = "block"
			alert('请上传符合规定的图片！')
			return
		}
		if (!patternValidate(/^[\u4E00-\u9FA5]{2,10}$/, stuName)) {
			nameV.style.display = "block"
			alert('请输入你的真实姓名！')
			return
		}
		if (!patternValidate(/^女|男$/, stuGender)) {
			genderV.style.display = "block"
			alert('请选择你的真实性别！')
			return
		}
		if (!patternValidate(/^22(999|006|180|008)([0-9][0-9][0-9])$/, stuNum)) {
			numV.style.display = "block"
			alert('请输入你的真实学号！')
			return
		}
		if (!patternValidate(/^(1|2|3|4|5|6|7|11|12|13|14|15|16)$/, stuClass)) {
			classV.style.display = "block"
			alert('请选择你的真实班级！')
			return
		}
		if (!patternValidate(/^[0-9]{4}-(((0[13578]|(10|12))-(0[1-9]|[1-2][0-9]|3[0-1]))|(02-(0[1-9]|[1-2][0-9]))|((0[469]|11)-(0[1-9]|[1-2][0-9]|30)))$/, stuBirth)) {
			birthV.style.display = "block"
			alert('请选择你的真实出生日期！')
			return
		}
		if (!patternValidate(/^浙江|上海|北京|天津|重庆|黑龙江|吉林|辽宁|内蒙古|河北|新疆|甘肃|青海|陕西|宁夏|河南|山东|山西|安徽|湖北|湖南|江苏|四川|贵州|云南|广西|西藏|江西|广东|福建|台湾|海南|香港|澳门$/, stuProvince)) {
			provinceV.style.display = "block"
			alert('请选择你的家乡省份！')
			return
		}
		if (!patternValidate(/^[1-9][0-9]{4,10}$/, stuQq)) {
			qqV.style.display = "block"
			alert('请输入你的真实QQ号！')
			return
		}
		if (!patternValidate(/^计算机科学与技术|计科师范|软件工程|网络工程$/, stuMajor)) {
			majorV.style.display = "block"
			alert('请选择你的专业！')
			return
		}

		if (!patternValidate(/^(?:(?:\+|00)86)?1\d{10}$/, stuPhone)) {
			phoneV.style.display = "block"
			alert('请输入你的真实手机号码！')
			return
		}
		if (!patternValidate(/^[0-9]{4}$/, smscode)) {
			smsV.style.display = "block"
			alert('请输入真实的验证码！')
			return
		}

		// 发送ajax请求

		var signUpData = new FormData();
		signUpData.append('stuNum', stuNum.value)
		signUpData.append('stuName', stuName.value)
		signUpData.append('stuGender', stuGender.value)
		signUpData.append('major', stuMajor.value)
		signUpData.append('class', stuClass.value)
		signUpData.append('qq', stuQq.value)
		signUpData.append('mobile', stuPhone.value)
		signUpData.append('province', stuProvince.value)
		signUpData.append('birth', stuBirth.value)
		signUpData.append('smsCode', smscode.value)
		signUpData.append('photo', stuPhoto.files[0])
		signUpData.append('introduce', stuIntroduce.value)
		signUpData.append('imagine', stuImagine.value)
		signUpData.append('advantage', stuAdvantage.value)
		$.ajax({
			url: "https://skylab.org.cn/api/v1/signUp",
			contentType: false,
			data: signUpData,
			processData: false,
			type: "post",
			async: false,
			cache: false,
			success: function (data, statu) {
				if (data.code == 1015) {
					alert(data.message)
				}
				if (data.code == 1001) {
					alert(data.message)
				}
				if (data.code == 1005) {
					alert(data.message + "请稍后再试")
				}
				if (data.code == 1012) {
					alert(data.message)
				}
				if (data.code == 1013) {
					alert(data.message)
				}
				if (data.code == 1016) {
					alert(data.message)
				}
				if (data.code == 1002) {
					alert('你已经报名过！请勿重复报名！')
				}
				if (data.code == 2333) {
					processResponse()
				}
			},
			error: function (data, status) {
				alert(status, data)
			},
		})
	}

	function processResponse() {
		stuName.value = ""
		stuNum.value = ""
		stuGender.value = ""
		stuClass.value = ""
		stuBirth.value = ""
		stuProvince.value = ""
		stuQq.value = ""
		stuMajor.value = ""
		stuAdvantage.value = ""
		stuIntroduce.value = ""
		stuImagine.value = ""
		smscode.value = ""
		stuPhone.value = ""
		preview.src = ""
		alert('报名成功！')
	}
}
