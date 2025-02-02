# VhostFinder
This tool will identify virtual hosts by performing a similarity comparison. It will generate a baseline request to attempt to map a non-existent virtual host. From there it will iterate over the supplied domains and compare them for any differences. Any significant differences will result in a virtual host being detected.

# Install

```
go install -v github.com/wdahlenburg/VhostFinder@latest
```

# Usage

```
Usage: VhostFinder -ip 10.8.0.1 -wordlist domains.txt
  -domains string
      Optional domain or comma seperated list to append to a subdomain wordlist (Ex: example1.com,example2.com)
  -ip string
      IP Address to Fuzz
  -path string
      Custom path to send during fuzzing (default "/")
  -port int
      Port to use (default 443)
  -threads int
      Number of threads to use (default 10)
  -tls
      Use TLS (default true)
  -v  Verbose mode
  -verify
      Verify vhost is different than public url
  -wordlist string
      File of FQDNs or subdomain prefixes to fuzz for
```

### Examples:
```bash
  VhostFinder -ip 10.8.0.1 -wordlist domains.txt
  [!] Finding vhosts!
  [!] Obtaining baseline
  [+] host.example.com

  VhostFinder -ip 10.8.0.1 -wordlist subdomains.txt -domains host1.example.com -v
  [!] Finding vhosts!
  [!] Obtaining baseline
  [+] admin.host1.example.com
  [-] test.host1.example.com

  VhostFinder -ip 10.8.0.1 -wordlist subdomains.txt -domains host1.example.com,anotherdomain.net,host2.example.com -v
  [!] Finding vhosts!
  [!] Obtaining baseline
  [+] admin.host1.example.com
  [-] test.host1.example.com
  [-] admin.anotherdomain.net
  [-] test.anotherdomain.net
  [+] admin.host2.example.com
  [-] test.host2.example.com
```

# What is Virtual Host Fuzzing?

Essentially the following request is sent repeatedly to a particular IP:

```
GET / HTTP/1.1
Host: FUZZ
Connection: close


```

The host header is fuzzed based on user input, while all requests are sent to the same IP. 

