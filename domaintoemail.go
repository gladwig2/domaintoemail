// domaintoemail takes a domain (xyz.com) and returns the email address of the Tech contact
// it returns an error if whois is unable to find the domain, or the result does not have a "Tech Email:" entry
// usage:
//    out, err := domaintoemail.get("domain.com")
//
package domaintoemail
import (
	"fmt"
	"os/exec"
	"regexp"
	"github.com/golang/glog"
)

func get(domain string) (string,error) {
	out, err := exec.Command("/usr/bin/whois", domain).Output() 
	if err != nil { 
		msg := fmt.Sprintf("domaintoemail:  error executing whois command with domain=%s, returned error= %v\n", domain, err)
		glog.Error(msg)
		glog.Flush()
		return "", err
	}	
	re := regexp.MustCompile(`(?i)tech\w* email: ([A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,})`)
	m := re.FindAllStringSubmatch(string(out),-1)

	//fmt.Printf("entries = %d \n m = %s, %v \n", len(m),m, err)
	//fmt.Printf("m type = %T, err type = %T\n", m, err)

	if len(m) == 1 {
		glog.Infof("Success looked up domain = %s, found email = %s \n", domain, m[0][1])
		glog.Flush()
		return m[0][1], nil
	} else {
		msg := fmt.Sprintf("domaintoemail: Error when looking up domain =%s, expected 1 match, received %d\n",domain,len(m))
		glog.Error(msg)
		glog.Flush()
		return "", fmt.Errorf(msg)
	}
}

