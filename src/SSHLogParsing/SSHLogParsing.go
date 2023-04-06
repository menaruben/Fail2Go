package SSHLogParsing

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func get_failed_ip_addr(line string) string {
	fields := strings.Split(line, " ")
	var ip_addr string = fields[len(fields)-4]
	return ip_addr
}

func get_unique_strings(arr []string) []string {
	keys := make(map[string]bool)
	var unique_strings []string

	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			unique_strings = append(unique_strings, entry)
		}
	}
	return unique_strings
}

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
	var path_to_file string = "C:/ProgramData/ssh/logs/sshd.log"
	file, err := os.Open(path_to_file)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error reading from %s", path_to_file))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var failed_hosts []string

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Failed password for") {
			ip_addr := get_failed_ip_addr(scanner.Text())
			// fmt.Println(ip_addr)
			failed_hosts = append(failed_hosts, ip_addr)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	unique_failed_hosts := get_unique_strings(failed_hosts)

	var bannable_hosts []string
	for i := 0; i < len(unique_failed_hosts); i++ {
		if get_failed_login_count(unique_failed_hosts[i], failed_hosts) >= failed_login_limit {
			bannable_hosts = append(bannable_hosts, unique_failed_hosts[i])
		}
	}
	return bannable_hosts
}

// func main() {
// 	stdout := Get_bannable_hosts(2)
// 	fmt.Println(stdout)
// }
