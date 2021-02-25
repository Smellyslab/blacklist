package blacklist

import "net"



// comment
func BlacklistIP(name net.Conn, ip string) {
	if name.RemoteAddr().String() == ip {
		return
	}	
}


