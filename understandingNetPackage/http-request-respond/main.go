package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer li.Close()

	for true {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	//read request
	request(conn)
}

func request(conn net.Conn) {
	flag := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if flag == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// header are done
			break
		}
		flag++
	}
}

func mux(conn net.Conn, ln string) {
	uri := strings.Fields(ln)[1]
	method := strings.Fields(ln)[0]

	switch uri {
	case "/":
		if method == http.MethodGet {
			indexPage(conn)
		}
	case "/about":
		if method == http.MethodGet {
			aboutPage(conn)
		}
	case "/apply":
		if method == http.MethodPost {
			applyProcess(conn)
		} else if method == http.MethodGet {
			applyPage(conn)
		}
	default:
		notFoundPage(conn)
	}
}

func indexPage(conn net.Conn) {
	body := `
			<!DOCTYPE html>
				<html>
					<head>
						<title>Page Title</title>
					</head>
					<body>
						<h1>This is index page<h1>
						<a href="/">index<a>
						<a href="/about">about<a>
						<a href="/apply">apply<a>
					</body>
				</html>
			`
	setHeaders(conn, len(body))
	_, _ = fmt.Fprintln(conn, body)
}

func applyPage(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/contact">contact</a><br>
	<a href="/applyPage">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`

	setHeaders(conn, len(body))
	_, _ = fmt.Fprintln(conn, body)
}

func aboutPage(conn net.Conn) {
	body := `
			<!DOCTYPE html>
				<html>
					<head>
						<title>Page Title</title>
					</head>
					<body>
						<p>this is about page</p>
					</body>
				</html>
			`
	setHeaders(conn, len(body))
	_, _ = fmt.Fprintln(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	setHeaders(conn, len(body))
	_, _ = fmt.Fprintln(conn, body)
}

func notFoundPage(conn net.Conn) {
	body := `
			<!DOCTYPE html>
				<html>
					<head>
						<title>404 Not Found</title>
					</head>
					<body>
						<h1>404 page not found</h1>
					</body>
				</html>
			`
	setHeadersError(conn, len(body))
	_, _ = fmt.Fprintln(conn, body)
}

func setHeaders(conn net.Conn, contentLen int) {
	_, _ = fmt.Fprintln(conn, "HTTP/1.1 200 OK\r")
	_, _ = fmt.Fprintf(conn, "Content-Length: %d\r\n", contentLen)
	_, _ = fmt.Fprintln(conn, "Content-Type: text/html\r")
	_, _ = fmt.Fprintln(conn, "\r")
}

func setHeadersError(conn net.Conn, contentLen int) {
	_, _ = fmt.Fprintln(conn, "HTTP/1.1 404 Not Found\r")
	_, _ = fmt.Fprintf(conn, "Content-Length: %d\r\n", contentLen)
	_, _ = fmt.Fprintln(conn, "Content-Type: text/html\r")
	_, _ = fmt.Fprintln(conn, "\r")
}
