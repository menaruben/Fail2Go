package winfw

import (
	"fmt"
	"log"
	"os/exec"
)

// Creates the firewall rule to block the IP address
func Ban_IP(ip string) {
	cmdlet_args := fmt.Sprintf("New-NetFirewallRule -Name %s -DisplayName %s -Direction Inbound -LocalPort Any -Protocol TCP -Action Block -RemoteAddress %s/32", ip, ip, ip)
	log.Printf("Executing: powershell.exe %s", cmdlet_args)

	ban_cmd := exec.Command("powershell.exe", cmdlet_args)

	if err := ban_cmd.Run(); err != nil {
		log.Fatalf("Error banning %s: %v", ip, err)
	}
	log.Printf("%s banned successfully", ip)
}

// Removes the firewall rule that blocks the IP address
func Unban_IP(ip string) {
	var cmdlet_args string = fmt.Sprintf("Remove-NetFirewallRule -Name %s", ip)
	log.Printf("Executing: powershell.exe %s", cmdlet_args)

	unban_cmd := exec.Command("powershell.exe", cmdlet_args)

	if err := unban_cmd.Run(); err != nil {
		log.Fatalf("Error unbanning %s: %v", ip, err)
	}
	log.Printf("%s unbanned successfully", ip)
}
