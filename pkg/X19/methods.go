package X19

import (
	"encoding/json"
	dac "github.com/xinsnake/go-http-digest-auth-client"
	"io"
	"net/http"
	"net/url"
)

func NewDevice() *Device {
	return &Device{
		Target: url.URL{
			Scheme: "http",
			Host:   "127.0.0.1",
			Path:   "/cgi-bin/"},
		User:    "root",
		Pass:    "root",
		Payload: "",
	}
}

// Blink включить или выключить моргание светодиодов
func (c *Device) Blink(state string) (Blink, error) {
	blink, _ := json.Marshal(Blink{Blink: state})
	c.Target.Path = CgiPath + BlinkUrl
	dr := dac.NewRequest(c.User, c.Pass, "POST", c.Target.String(), string(blink))
	response, err := dr.Execute()
	BlinkResponse := Blink{}
	if err != nil {
		return BlinkResponse, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &BlinkResponse.Response)
		switch BlinkResponse.Response.Code {
		case "B000":
			BlinkResponse.Response.Message = "Blink successfully enabled."
		case "B100":
			BlinkResponse.Response.Message = "Blink successfully disabled."
		default:
			BlinkResponse.Response.Message = "Got another code."
			return BlinkResponse, err
		}
	}
	return BlinkResponse, err
}

// GetBlinkStatus получить статус моргания светодиодов
func (c *Device) GetBlinkStatus() (Blink, error) {
	c.Target.Path = CgiPath + BlinkStatusUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	GetBlinkResp := Blink{}
	if err != nil {
		return GetBlinkResp, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &GetBlinkResp.BlinkStatus)
	}
	return GetBlinkResp, err
}

// Reboot выполнить перезагрузку
func (c *Device) Reboot() (string, error) {
	c.Target.Path = CgiPath + RebootUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	if err != nil {
		return "", err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		return string(bytes), err
	}
	return "", err
}

// GetStats получить статистику
func (c *Device) GetStats() (Stats, error) {
	c.Target.Path = CgiPath + StatsUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponseStats := Stats{}
	if err != nil {
		return ResponseStats, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &ResponseStats)
	}
	return ResponseStats, err
}

func (c *Device) GetSummary() (Summary, error) {
	c.Target.Path = CgiPath + SummaryUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponseSummary := Summary{}
	if err != nil {
		return ResponseSummary, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &ResponseSummary)
	}
	return ResponseSummary, err
}

// GetMinerConf получить информацию о конфигурации
func (c *Device) GetMinerConf() (MinerConf, error) {
	c.Target.Path = CgiPath + MinerConfUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), c.Payload)
	response, err := dr.Execute()
	ResponseMinerConf := MinerConf{}
	if err != nil {
		return ResponseMinerConf, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &ResponseMinerConf)
	}
	return ResponseMinerConf, err
}

// GetSystemInfo получить системную информацию
func (c *Device) GetSystemInfo() (SystemInfo, error) {
	c.Target.Path = CgiPath + SystemInfoUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponseSystemInfo := SystemInfo{}
	if err != nil {
		return ResponseSystemInfo, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &ResponseSystemInfo)
	}
	return ResponseSystemInfo, err
}

// GetNetworkInfo получить информацию о сетевых настройках
func (c *Device) GetNetworkInfo() (NetworkInfo, error) {
	c.Target.Path = CgiPath + NetworkInfoUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponseNetworkInfo := NetworkInfo{}
	if err != nil {
		return ResponseNetworkInfo, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &ResponseNetworkInfo)
	}
	return ResponseNetworkInfo, err
}

// GetChart получить информацию о чартах
func (c *Device) GetChart() (Chart, error) {
	c.Target.Path = CgiPath + ChartUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponseChart := Chart{}
	if err != nil {
		return ResponseChart, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &ResponseChart)
	}
	return ResponseChart, err
}

// GetPools получить информацию о прописанных пулах
func (c *Device) GetPools() (Pools, error) {
	c.Target.Path = CgiPath + PoolsUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponsePools := Pools{}
	if err != nil {
		return ResponsePools, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &ResponsePools)
	}
	return ResponsePools, err
}

// GetHistoryLog получить текущий лог
func (c *Device) GetHistoryLog() (string, error) {
	c.Target.Path = CgiPath + HistoryLogUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponseHistoryLog := Logs{}.HistoryLog
	if err != nil {
		return ResponseHistoryLog, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		ResponseHistoryLog = string(bytes)
	}
	return ResponseHistoryLog, err
}

// GetCurrentLog получить исторический лог
func (c *Device) GetCurrentLog() (string, error) {
	c.Target.Path = CgiPath + CurrentLogUrl
	dr := dac.NewRequest(c.User, c.Pass, "GET", c.Target.String(), "")
	response, err := dr.Execute()
	ResponseLog := Logs{}.CurrentLog
	if err != nil {
		return ResponseLog, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		ResponseLog = string(bytes)
	}
	return ResponseLog, err
}

// ChangeFanSetting Мизменить режим работы вентиляторов
func (c *Device) ChangeFanSetting(forceMode string, fanPwm string) (SetMinerConfResponse, error) {
	NewConfig, _ := json.Marshal(SetMinerConf{BitmainFanCtrl: forceMode, BitmainFanPwm: fanPwm})
	c.Target.Path = CgiPath + SetMinerConfUrl
	dr := dac.NewRequest(c.User, c.Pass, "POST", c.Target.String(), string(NewConfig))
	response, err := dr.Execute()
	Response := SetMinerConfResponse{}
	if err != nil {
		return Response, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &Response)
	}
	return Response, err
}

// SetWorkMode переключить режим работы
func (c *Device) SetWorkMode(mode string) (SetMinerConfResponse, error) {
	NewConfig, _ := json.Marshal(SetMinerConf{MinerMode: mode})
	c.Target.Path = CgiPath + SetMinerConfUrl
	dr := dac.NewRequest(c.User, c.Pass, "POST", c.Target.String(), string(NewConfig))
	response, err := dr.Execute()
	Response := SetMinerConfResponse{}
	if err != nil {
		return Response, err
	}
	if response.StatusCode == http.StatusOK {
		bytes, _ := io.ReadAll(response.Body)
		_ = json.Unmarshal(bytes, &Response)
	}
	return Response, err
}

func (c *Device) CalculateNominal() {
	// Пока что эта функция ничего не делает, но в будущем она будет расчитывать номинал.
}
