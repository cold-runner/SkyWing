let count = 0
const firstChild = document.getElementById("first").children;
const secondChild = document.getElementById("second").children;
const thirdChild = document.getElementById("third").children;
const fourthChild = document.getElementById("fourth").children;

const changeButton = document.getElementById("change")
changeButton.addEventListener('click', function () {
	fn(function () {
		count++
		next(firstChild, count)
		next(secondChild, count)
		next(thirdChild, count)
		next(fourthChild, count)
	});
})
// 渐变透明度
function alphaPlay(obj, method) {
	var n = (method == "show") ? 0 : 100;
	var time = setInterval(function () {
		if (method == "show") {
			if (n < 100) {
				n += 10;
				if (window.ActiveXObject) {
					obj.style.cssText = "filter:alpha(opacity=" + n + ")";
				} else {
					(n == 100) ? obj.style.opacity = 1 : obj.style.opacity = "0." + n;
				}
			} else {
				clearTimeout(time);
			}
		} else {
			if (n > 0) {
				n -= 10;
				if (window.ActiveXObject) {
					obj.style.cssText = "filter:alpha(opacity=" + n + ")";
				} else {
					obj.style.opacity = "0." + n;
				}
			} else {
				clearTimeout(time);
			}
		}
	}, 65);
}
// 切换
function next(array, count) {
	lastIndex = count % array.length - 1
	if (lastIndex == -1) {
		lastIndex = array.length - 1
		targetIndex = 0
	} else {
		targetIndex = lastIndex + 1
	}
	for (let index = 0; index < array.length; index++) {
		if (index == lastIndex) {
			alphaPlay(array[index], 'hiden')
			alphaPlay(array[targetIndex], 'show')
		}
	}
}

// 一定时间内仅触发一次
function fn(callback) {
	fn.prototype.init(callback);
}
fn.prototype = {
	canclick: true,
	init: function (callback) {
		if (this.canclick) {
			this.canclick = false
			callback();
			setTimeout(function () {
				this.canclick = true
			}.bind(this), 1000)
		}
	}
}
