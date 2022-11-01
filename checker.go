package checklib

type CheckerInfo struct {
	Vulns      int  `json:"vulns"`
	Timeout    int  `json:"timeout"`
	AttackData bool `json:"attack_data"`
	Puts       int  `json:"puts"`
	Gets       int  `json:"gets"`
}

type Checker interface {
	Info() *CheckerInfo
	Check(c *C, host string)
	Put(c *C, host, flagID, flag string, vuln int)
	Get(c *C, host, flagID, flag string, vuln int)
}
