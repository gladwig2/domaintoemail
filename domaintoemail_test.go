package domaintoemail

import (
	"fmt"
	"testing"
	"flag"
	"os"
)

func TestGoodDomain(t *testing.T){
	domain := "google.com"
	email,err := get(domain)
	if err != nil {
		fmt.Printf("test error, for domain %s, email string = %s, err = %v\n",domain, email,err)
		t.Fail()
	}
	domain = "blacklistme.net"
	email,err = get(domain)
	if err != nil {
		fmt.Printf("test error, for domain %s, email string = %s, err = %v\n",domain, email,err)
		t.Fail()
	}
	domain = "ipspace.net"
	email,err = get(domain)
	if err != nil {
		fmt.Printf("test error, for domain %s, email string = %s, err = %v\n",domain, email,err)
		t.Fail()
	}
	// was going to test logfile existance, but naming is unclear and doesn't appear queryable
	//logfile := flag.Lookup("log_dir")
	//if logfile.Value != nil {
	//	fmt.Println("Logfile vaulue = ", logfile.Value)
	//	logname := logfile.Value		
	//} else {
	//logname 
	//fmt.Printf("logfile type = %T\n",logfile)
}

func TestBadDomain(t *testing.T){
	domain := "gooberdoober.com"
	_,err := get(domain)
	if err == nil {
		fmt.Printf("test error, for domain %s, should have failed but didnt, err = %v\n",domain,err)
		t.Fail()
	}
}

func TestMalformedDomain(t *testing.T){
	domain := "google@.com"
	_,err := get(domain)
	if err == nil {
		fmt.Printf("test error, for domain %s, should have failed but didnt, err = %v\n",domain,err)
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	flag.Usage = usage
	flag.Parse()
	os.Exit(m.Run())
}

// usage will return information on glog usage
func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}




