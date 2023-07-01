![Coverage](https://raw.githubusercontent.com/nao1215/octocovs-central-repo/main/badges/nao1215/emigre/coverage.svg)
[![reviewdog](https://github.com/nao1215/emigre/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/emigre/actions/workflows/reviewdog.yml)
[![LinuxUnitTest](https://github.com/nao1215/emigre/actions/workflows/linux_test.yml/badge.svg)](https://github.com/nao1215/emigre/actions/workflows/linux_test.yml)
[![MacUnitTest](https://github.com/nao1215/emigre/actions/workflows/mac_test.yml/badge.svg)](https://github.com/nao1215/emigre/actions/workflows/mac_test.yml)
[![WindowsUnitTest](https://github.com/nao1215/emigre/actions/workflows/windows_test.yml/badge.svg)](https://github.com/nao1215/emigre/actions/workflows/windows_test.yml)
[![Gosec(Security)](https://github.com/nao1215/emigre/actions/workflows/security.yml/badge.svg)](https://github.com/nao1215/emigre/actions/workflows/security.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/emigre)](https://goreportcard.com/report/github.com/nao1215/emigre)

# emigre - Self-hosted image sharing social media
**work in progress**  
  
emigre is a self-hosted server and client that mimics an image posting service. The purpose of developing Emigre is to validate technology stacks that are not used in the company that I belong to. In other words, satisfying technical interests is more important than creating a popular service.

Emigre is expected to be implemented using the following technologies:
- Server-side: Golang
- Infrastructure: AWS (Amazon Web Services)
- Client-side: Android application or GUI written in Golang

## Support OS & golang version
- Linux 
- macOS
- Windows
- golang ver 1.18 or later

## Documentations
- [API Document](https://nao1215.github.io/emigre/index.html)
- [Go reference](https://pkg.go.dev/github.com/nao1215/emigre)

## Server-side tools
The server-side is implemented using Golang. The server-side is expected to be implemented using the following technologies:
| Category | Technology |
|:---|:---|
| Web framework | [Echo](https://echo.labstack.com/) |
| Database | [MySQL](https://www.mysql.com/) or [SQLite3](https://www.sqlite.org/index.html) |
| SQL(DDL, Query, etc) | Not use tool. Write it myself. |
| DB Accessor | [sqlc](https://github.com/kyleconroy/sqlc) |
| DB migration| TBD |
| Swagger generator | [echo-swagger](https://github.com/swaggo/echo-swagger) |
| ER diagram| [tbls](https://github.com/k1LoW/tbls)|
| Test coverage | [ocotcov](https://github.com/k1LoW/octocov)|
| Static Analysis| [reviewdog](https://github.com/reviewdog/reviewdog)|
| Test framework | [ginkgo](https://github.com/onsi/ginkgo)|
| Dependency Injection |[wire](https://github.com/google/wire)|

## What does Emigre mean?
The Emigre is derived from the Emigre document in SHADOW HEART (it is game game developed for the PlayStation 2). The Emigre document describes a secret technique that overcomes death and creates new life from nothingness. I write code out of technical interest, however I also desire for the software to endure.