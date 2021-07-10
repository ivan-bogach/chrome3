package chrome3

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/knq/chromedp"
)

func InitHeadLess(pathForUserDataDir string) (context.Context, context.CancelFunc) {
	opts := []chromedp.ExecAllocatorOption{chromedp.Flag("no-sandbox", true), chromedp.Flag("headless", true), chromedp.Flag("disable-gpu", true)}
	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	return chromedp.NewContext(allocContext)
}

func Init(pathForUserDataDir string) (context.Context, context.CancelFunc) {
	opts := []chromedp.ExecAllocatorOption{chromedp.Flag("no-sandbox", true)}
	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	return chromedp.NewContext(allocContext)
}
func InitWithProxy(pathForUserDataDir, proxyName string) (context.Context, context.CancelFunc) {
	opts := []chromedp.ExecAllocatorOption{chromedp.Flag("no-sandbox", true), chromedp.ProxyServer(proxyName)}
	allocContext, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	return chromedp.NewContext(allocContext)
}

func checkConn(connEnabled *bool) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scriptCheckConn(), connEnabled),
	}
}

// Check connection
func CheckConn(ctxt context.Context) (bool, error) {
	var connEnabled bool
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, checkConn(&connEnabled)))
	if err != nil {
		return false, fmt.Errorf("this is an %s error: %v", "CheckConn", err)
	}
	return connEnabled, nil
}

func openURL(url string, message *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scriptOpenURL(url), message),
	}
}

func OpenURL(ctxt context.Context, url string, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Opening page url %s - ", url)
	}

	var message string
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, openURL(url, &message)))
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "OpenURL", err)
	}

	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("Ok!.")
	}
	return nil
}

func reload() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Reload(),
		chromedp.Sleep(5 * time.Second),
	}
}

func Reload(ctxt context.Context) error {
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, reload()))
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "Reload", err)
	}
	return nil
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}

func waitVisible(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(selector, chromedp.ByQuery),
	}
}

func WaitVisible(ctxt context.Context, selector string, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Wait visible css:' %s ' - ", selector)
	}
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, waitVisible(selector)))
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "WaitVisible", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("Ok!.")
	}
	return nil
}

func waitReady(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitReady(selector, chromedp.ByQuery),
	}
}

func WaitReady(ctxt context.Context, selector string, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Wait ready css:' %s ' - ", selector)
	}
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, waitReady(selector)))

	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "WaitReady", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("Ok!.")
	}
	return nil
}

func getString(jsString string, resultString *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scriptGetString(jsString), resultString),
	}
}

func GetString(ctxt context.Context, jsString string, resultString *string, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Getting a string ' %s  ' - ", jsString)
	}
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, getString(jsString, resultString)))
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "GetString", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("-ok")
	}
	return nil
}

func getStringsSlice(jsString string, resultSlice *[]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scriptGetStringsSlice(jsString), resultSlice),
	}
}

func GetStringsSlice(ctxt context.Context, jsString string, stringSlice *[]string, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Getting a strings slice ' %s  ' - ", jsString)
	}
	color.Green("")
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, getStringsSlice(jsString, stringSlice)))
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "GetStringsSlice", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("-ok")
	}
	return nil
}

func GetReader(ctxt context.Context, jsString string, needLog bool) (*strings.Reader, error) {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Getting a string ' %s  ' - ", jsString)
	}
	var resultString string
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, getString(jsString, &resultString)))
	if err != nil {
		return nil, fmt.Errorf("this is an %s error: %v", "GetReader", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("-ok")
	}
	return strings.NewReader(resultString), nil
}

func getBool(jsBool string, resultBool *bool) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scriptGetBool(jsBool), resultBool),
	}
}

func GetBool(ctxt context.Context, jsBool string, resultBool *bool, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Getting a string ' %s  ' - ", jsBool)
	}
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, getBool(jsBool, resultBool)))
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "GetBool", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("-ok")
	}
	return nil
}

func click(selector string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(selector, chromedp.ByQuery),
	}
}

func Click(ctxt context.Context, selector string, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Click selector: ' %s '  - ", selector)
	}
	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, waitVisible(selector)))
	if err != nil {
		return fmt.Errorf("this is an wait visible in %s error: %v", "Click", err)
	}
	err = chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, click(selector)))
	if err != nil {
		return fmt.Errorf("this is an click in %s error: %v", "Click", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("-ok")
	}
	return nil
}

func setInputValue(selector, value string, result *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EvaluateAsDevTools(scriptSetInputValue(selector, value), result),
	}
}

func SetInputValue(ctxt context.Context, selector, value string, needLog bool) error {
	if needLog {
		c := color.New(color.FgGreen)
		c.Printf("Setting an input >>>%s<<< value - >>>%s<<<", selector, value)
	}
	color.Green("")
	var resultOperation string

	err := chromedp.Run(ctxt, RunWithTimeOut(&ctxt, 60, setInputValue(selector, value, &resultOperation)))
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "SetInputValue", err)
	}
	if needLog {
		d := color.New(color.FgGreen, color.Bold)
		d.Println("-ok")
	}
	return nil
}

func WaitLoaded(ctxt context.Context) error {
	var loaded bool
	err := GetBool(ctxt, `document.readyState !== 'ready' && document.readyState !== 'complete'`, &loaded, false)
	if err != nil {
		return fmt.Errorf("this is an %s error: %v", "WaitLoaded", err)
	}
	fmt.Print("Wait")
	n := 0
	for loaded {
		if n > 60 {
			return fmt.Errorf("this is an minute passed %s error: %v", "WaitLoaded", err)
		}
		fmt.Print(".")
		time.Sleep(1 * time.Second)
		err = GetBool(ctxt, `document.readyState !== 'ready' && document.readyState !== 'complete'`, &loaded, false)
		if err != nil {
			return fmt.Errorf("this is an %s error: %v", "WaitLoaded", err)
		}
		n++
	}
	return nil
}

func parsePage(ctxt context.Context, js string) ([]string, error) {
	var strSl []string
	err := GetStringsSlice(ctxt, js, &strSl, false)
	if err != nil {
		return nil, fmt.Errorf("this is an %s error: %v", "parsePage", err)
	}
	return strSl, nil
}

func StringSliceFromPage(ctxt context.Context, url, js string, waitFor ...string) ([]string, error) {
	err := OpenURL(ctxt, url, false)
	if err != nil {
		return nil, fmt.Errorf("this is an %s error: %v", "StringSliceFromPage", err)
	}
	if len(waitFor) == 0 {
		err := WaitLoaded(ctxt)
		if err != nil {
			return nil, fmt.Errorf("this is an %s error: %v", "StringSliceFromPage", err)
		}
	} else {
		for _, w := range waitFor {
			err := WaitVisible(ctxt, w, false)
			if err != nil {
				return nil, fmt.Errorf("this is an %s error: %v", "StringSliceFromPage", err)
			}
		}
	}

	time.Sleep(3 * time.Second)
	newJSONSl, err := parsePage(ctxt, js)
	if err != nil {
		return nil, fmt.Errorf("this is an %s error: %v", "StringSliceFromPage", err)
	}
	return newJSONSl, nil
}

func parseStrPage(ctxt context.Context, js string) (string, error) {
	var str string
	err := GetString(ctxt, js, &str, false)
	if err != nil {
		return "", fmt.Errorf("this is an %s error: %v", "parseStrPage", err)
	}

	return str, nil
}

func StringFromPage(ctxt context.Context, url, js string, waitFor ...string) (string, error) {
	err := OpenURL(ctxt, url, false)
	if err != nil {
		return "", fmt.Errorf("this is an %s error: %v", "StringFromPage", err)
	}
	if len(waitFor) == 0 {
		err := WaitLoaded(ctxt)
		if err != nil {
			return "", fmt.Errorf("this is an %s error: %v", "StringFromPage", err)
		}
	} else {
		for _, w := range waitFor {
			err := WaitVisible(ctxt, w, false)
			if err != nil {
				return "", fmt.Errorf("this is an %s error: %v", "StringFromPage", err)
			}
		}
	}

	time.Sleep(3 * time.Second)
	newJSON, err := parseStrPage(ctxt, js)
	if err != nil {
		return "", fmt.Errorf("this is an %s error: %v", "StringFromPage", err)
	}
	return newJSON, nil
}
