package X17

func Monitor() {
	//http://192.168.101.90/cgi-bin/monitor.cgi?_=1671455860664
}

func KernelLog() {
	//http://192.168.101.90/cgi-bin/get_kernel_log.cgi?_=1671455898017
}

func ResetConf() {
	//function f_submit_reset() {
	//	if(confirm('Really reset all changes?')) {
	//		$("#cbi_body_bmminer_fieldset").hide();
	//		$("#cbi_reset_bmminer_fieldset").show();
	//
	//		setTimeout(function(){
	//			window.location.href="http://192.168.1.99/index.html";
	//		}, 90000);
	//
	//		$.ajax({
	//		url: '/cgi-bin/reset_conf.cgi',
	//			type: 'GET',
	//				dataType: 'json',
	//				timeout: 30000,
	//				cache: false,
	//				data: {},
	//		success: function(data) {
	//			},
	//		error: function() {
	//			}
	//		});
	//	}
	//}
}

func KillBmMiner() {
	//function f_submit_restore_exec() {
	//	$.ajax({
	//	url: '/cgi-bin/kill_bmminer.cgi',
	//		type: 'GET',
	//			dataType: 'text',
	//			timeout: 30000,
	//			cache: false,
	//			data: {},
	//	success: function(data) {
	//		},
	//	error: function() {
	//		}
	//	});
	//	ant_restore_config_form.submit();
	//}
}

func Reboot() {
	//function f_submit_reboot() {
	//	$("#cbi_apply_bmminer_fieldset").show();
	//
	//	setTimeout(function(){
	//		window.location.reload();
	//	}, 90000);
	//
	//	$.ajax({
	//	url: '/cgi-bin/reboot.cgi',
	//		type: 'GET',
	//			dataType: 'json',
	//			timeout: 30000,
	//			cache: false,
	//			data: {},
	//	success: function(data) {
	//		},
	//	error: function() {
	//		}
	//	});
	//}
}

func Blink() {
	//function startBlink() {
	//	$("#cbi_apply_miner_blink").show();
	//	$.ajax({
	//	url: '/cgi-bin/blink.cgi',
	//		type: 'POST',
	//			dataType: 'json',
	//			cache: false,
	//			timeout: 1000,
	//			data: {action: 'startBlink'},
	//	success: function(data) {
	//		},
	//	error: function(data) {
	//		}
	//	});
	//	$("#start-blink-btn").attr("disabled", true);
	//	$("#stop-blink-btn").attr("disabled", false);
	//	checkBlinking("start");
	//}
	//function stopBlink() {
	//	$("#cbi_apply_miner_blink").hide();
	//	$("#start-blink-btn").attr("disabled", false);
	//	$("#stop-blink-btn").attr("disabled", true);
	//	$.ajax({
	//	url: '/cgi-bin/blink.cgi',
	//		type: 'POST',
	//			dataType: 'json',
	//			cache: false,
	//			data: {action: 'stopBlink'},
	//	success: function(data) {
	//		},
	//	error: function(data) {
	//		}
	//	});
	//}
}

func CheckBlinkStatus() {
	//function checkBlinking(status)
	//{
	//	function check(status){
	//	$.ajax({
	//	url: '/cgi-bin/blink.cgi',
	//		type: 'POST',
	//			dataType: 'json',
	//			cache: false,
	//			data: {action: 'onPageLoaded'},
	//	success: function(data) {
	//			if (data.isBlinking) {
	//				$("#cbi_apply_miner_blink").show();
	//				$("#start-blink-btn").attr("disabled", true);
	//				$("#stop-blink-btn").attr("disabled", false);
	//			} else {
	//				if (status != "start") {
	//					$("#cbi_apply_miner_blink").hide();
	//					$("#start-blink-btn").attr("disabled", false);
	//					$("#stop-blink-btn").attr("disabled", true);
	//					clearInterval(timer);
	//				}
	//			}
	//		},
	//	error: function(data) {
	//		}
	//	});
	//}
	//	var timer = setInterval(check, 2000);
	//	check(status);
	//}
}

func SetConf() {
	//function f_get_miner_conf() {
	//	try
	//	{
	//		for(var i = 0; i < ant_data.pools.length; i++) {
	//		switch(i) {
	//		case 0:
	//			jQuery("#ant_pool1url").val(ant_data.pools[i].url);
	//			jQuery("#ant_pool1user").val(ant_data.pools[i].user);
	//			jQuery("#ant_pool1pw").val(ant_data.pools[i].pass);
	//			break;
	//		case 1:
	//			jQuery("#ant_pool2url").val(ant_data.pools[i].url);
	//			jQuery("#ant_pool2user").val(ant_data.pools[i].user);
	//			jQuery("#ant_pool2pw").val(ant_data.pools[i].pass);
	//			break;
	//		case 2:
	//			jQuery("#ant_pool3url").val(ant_data.pools[i].url);
	//			jQuery("#ant_pool3user").val(ant_data.pools[i].user);
	//			jQuery("#ant_pool3pw").val(ant_data.pools[i].pass);
	//			break;
	//		}
	//	}
	//		if(ant_data["bitmain-nobeeper"]) {
	//			document.getElementById("ant_beeper").checked = false;
	//		} else {
	//			document.getElementById("ant_beeper").checked = true;
	//		}
	//		if(ant_data["bitmain-notempoverctrl"]) {
	//			document.getElementById("ant_tempoverctrl").checked = false;
	//		} else {
	//			document.getElementById("ant_tempoverctrl").checked = true;
	//		}
	//		if(ant_data["bitmain-fan-ctrl"]) {
	//			document.getElementById("ant_fan_customize_switch").checked = true;
	//
	//		} else {
	//			document.getElementById("ant_fan_customize_switch").checked = false;
	//		}
	//
	//		if(miner_type["value"] == "S17 Pro") {
	//			if(ant_data["bitmain-work-mode"] == 0) {
	//				document.getElementById("ant_work_mode_0").checked = true;
	//			} else if(ant_data["bitmain-work-mode"] == 1) {
	//				document.getElementById("ant_work_mode_1").checked = true;
	//			} else if(ant_data["bitmain-work-mode"] == 2) {
	//				document.getElementById("ant_work_mode_2").checked = true;
	//			} else if(ant_data["bitmain-work-mode"] == 0xfe){
	//				document.getElementById("ant_work_mode_sleep").checked = true;
	//			}
	//		} else if(miner_type["value"] == "S17") {
	//			if(ant_data["bitmain-work-mode"] == 1) {
	//				document.getElementById("ant_work_mode_1").checked = true;
	//			} else if(ant_data["bitmain-work-mode"] == 2){
	//				document.getElementById("ant_work_mode_2").checked = true;
	//			} else if(ant_data["bitmain-work-mode"] == 0xfe){
	//				document.getElementById("ant_work_mode_sleep").checked = true;
	//			}
	//		}
	//
	//		document.getElementById("ant_voltage").value=ant_data["bitmain-voltage"];
	//		document.getElementById("ant_freq").value=ant_data["bitmain-freq"];
	//
	//		jQuery("#ant_fan_customize_value").val(ant_data["bitmain-fan-pwm"]);
	//
	//	}
	//	catch(err)
	//	{
	//		alert('Invalid Miner configuration file. Edit manually or reset to default.');
	//	}
	//}
	//function f_submit_miner_conf() {
	//	_ant_freq = "18:218.75:1106";
	//	_ant_voltage = "0725";
	//	try
	//	{
	//		_ant_freq = ant_data["bitmain-freq"];
	//		_ant_voltage = ant_data["bitmain-voltage"];
	//	}
	//	catch(err)
	//	{
	//		alert('Invalid Miner configuration file. Edit manually or reset to default.');
	//	}
	//
	//	_ant_pool1url = jQuery("#ant_pool1url").val();
	//	_ant_pool1user = jQuery("#ant_pool1user").val();
	//	_ant_pool1pw = jQuery("#ant_pool1pw").val();
	//	_ant_pool2url = jQuery("#ant_pool2url").val();
	//	_ant_pool2user = jQuery("#ant_pool2user").val();
	//	_ant_pool2pw = jQuery("#ant_pool2pw").val();
	//	_ant_pool3url = jQuery("#ant_pool3url").val();
	//	_ant_pool3user = jQuery("#ant_pool3user").val();
	//	_ant_pool3pw = jQuery("#ant_pool3pw").val();
	//	_ant_nobeeper = "false";
	//	_ant_notempoverctrl = "false";
	//	_ant_fan_customize_switch = "false";
	//	_ant_fan_customize_value = jQuery("#ant_fan_customize_value").val();
	//
	//	if(document.getElementById("ant_beeper").checked) {
	//		_ant_nobeeper = "false";
	//	} else {
	//		_ant_nobeeper = "true";
	//	}
	//	if(document.getElementById("ant_tempoverctrl").checked) {
	//		_ant_notempoverctrl = "false";
	//	} else {
	//		_ant_notempoverctrl = "true";
	//	}
	//
	//	if(document.getElementById("ant_fan_customize_switch").checked) {
	//		_ant_fan_customize_switch= "true";
	//
	//	} else {
	//		_ant_fan_customize_switch= "false";
	//	}
	//
	//	if(miner_type["value"] == "S17 Pro") {
	//		if(document.getElementById("ant_work_mode_1").checked) {
	//			_ant_work_mode = 1;
	//		} else if (document.getElementById("ant_work_mode_2").checked) {
	//			_ant_work_mode = 2;
	//		} else if (document.getElementById("ant_work_mode_0").checked){
	//			_ant_work_mode = 0;
	//		} else if (document.getElementById("ant_work_mode_sleep").checked){
	//			_ant_work_mode = 0xfe;
	//		}
	//	}
	//	else if(miner_type["value"] == "S17") {
	//		if(document.getElementById("ant_work_mode_1").checked) {
	//			_ant_work_mode = 1;
	//		} else if (document.getElementById("ant_work_mode_2").checked) {
	//			_ant_work_mode = 2;
	//		} else if (document.getElementById("ant_work_mode_sleep").checked){
	//			_ant_work_mode = 0xfe;
	//		}
	//	}
	//	_ant_work_mode = 0;
	//	_ant_voltage=jQuery("#ant_voltage").val();
	//	_ant_freq=jQuery("#ant_freq").val();
	//	jQuery("#cbi_apply_bmminer_fieldset").show();
	//
	//	jQuery.ajax({
	//	url: '/cgi-bin/set_miner_conf.cgi',
	//		type: 'POST',
	//			dataType: 'json',
	//			timeout: 30000,
	//			cache: false,
	//			data: {_ant_pool1url:_ant_pool1url, _ant_pool1user:_ant_pool1user, _ant_pool1pw:_ant_pool1pw,_ant_pool2url:_ant_pool2url, _ant_pool2user:_ant_pool2user, _ant_pool2pw:_ant_pool2pw,_ant_pool3url:_ant_pool3url, _ant_pool3user:_ant_pool3user, _ant_pool3pw:_ant_pool3pw, _ant_nobeeper:_ant_nobeeper, _ant_notempoverctrl:_ant_notempoverctrl,_ant_fan_customize_switch:_ant_fan_customize_switch,_ant_fan_customize_value:_ant_fan_customize_value, _ant_freq:_ant_freq, _ant_voltage:_ant_voltage, _ant_work_mode:_ant_work_mode},
	//	success: function(data) {
	//			window.location.reload();
	//		},
	//	error: function() {
	//			window.location.reload();
	//		}
	//	});
	//}

}
