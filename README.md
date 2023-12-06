[![progress-banner](https://backend.codecrafters.io/progress/grep/28176ce0-63c3-4817-aa12-e6df9c6ea2f8)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)
[![progress-banner](https://backend.codecrafters.io/progress/dns-server/79f0b411-29c2-4e7d-ad8a-01a8b699a0d0)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

# DNS Server Go
[![ci](https://github.com/ttiimmothy/dns-server-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ttiimmothy/dns-server-go/actions/workflows/ci.yml)

This is a starting point for Go solutions to the
["Build Your Own DNS server" Challenge](https://app.codecrafters.io/courses/dns-server/overview).

In this challenge, you'll build a DNS server that's capable of parsing and
creating DNS packets, responding to DNS queries, handling various record types
and doing recursive resolve. Along the way we'll learn about the DNS protocol,
DNS packet format, root servers, authoritative servers, forwarding servers,
various record types (A, AAAA, CNAME, etc) and more.

**Note**: If you're viewing this repo on GitHub, head over to
[codecrafters.io](https://codecrafters.io) to try the challenge.

# Passing the first stage

The entry point for your `your_server.sh` implementation is in `app/main.go`.
Study and uncomment the relevant code, and push your changes to pass the first
stage:

```sh
git add .
git commit -m "pass 1st stage" # any msg
git push origin master
```

Time to move on to the next stage!

# Stage 2 & beyond

Note: This section is for stages 2 and beyond.

1. Ensure you have `go (1.19)` installed locally
2. Run `./your_server.sh` to run your program, which is implemented in
   `app/main.go`.
3. Commit your changes and run `git push origin master` to submit your solution
   to CodeCrafters. Test output will be streamed to your terminal.

## License

DNS Server Go is licensed under [GNU General Public License v3.0](LICENSE).