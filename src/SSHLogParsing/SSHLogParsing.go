package SSHLogParsing

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// returns the ip address inside the ssh log
func get_failed_ip_addr(line string) string {
	fields := strings.Split(line, " ")
	var ip_addr string = fields[len(fields)-4]
	return ip_addr
}

// returns an array of unique hosts that failed to login
func get_unique_strings(arr []string) []string {
	// Create an empty map with string keys and boolean values to track unique strings.
	keys := make(map[string]bool)
	var unique_strings []string

	// Loop through each entry in the original slice.
	for _, entry := range arr {
		// Check if the entry is already in the keys map.
		if _, value := keys[entry]; !value {
			// If entry isn't in the keys map, add it with a value of true
			// and append it to the unique_strings slice.
			keys[entry] = true
			unique_strings = append(unique_strings, entry)
		}
	}
	return unique_strings
}

// returns the value how many times a host has failed to login
func get_failed_login_count(value string, failed_hosts []string) int {
	count := 0
	for i := 0; i < len(failed_hosts); i++ {
		if value == failed_hosts[i] {
			count++
		}
	}
	return count
}

// returns all hosts which failed to login more or euqal to x amount of times
func Get_bannable_hosts(failed_login_limit int) []string {
	// var path_to_file string = "C:/ProgramData/ssh/logs/sshd.log"
	var path_to_file string = "./TestSSHLog.log"
	file, err := os.Open(path_to_file)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error reading from %s", path_to_file))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var failed_hosts []string

	// store all lines containing substring "Failed password for" to "failed_hosts"
	// and parse out ip address
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Failed password for") {
			ip_addr := get_failed_ip_addr(scanner.Text())
			failed_hosts = append(failed_hosts, ip_addr)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	unique_failed_hosts := get_unique_strings(failed_hosts)

	// count occurrunces of unique hosts in failed_hosts and if bigger than faled_login_limit
	// append to bannable_hosts array which gets returned at the end
	var bannable_hosts []string
	for i := 0; i < len(unique_failed_hosts); i++ {
		if get_failed_login_count(unique_failed_hosts[i], failed_hosts) >= failed_login_limit {
			bannable_hosts = append(bannable_hosts, unique_failed_hosts[i])
			log.Printf("Found bannable host: %s", unique_failed_hosts[i])
		}
	}
	return bannable_hosts
}
