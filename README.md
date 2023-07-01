# [WIP] emigre - Self-hosted image sharing social media
emigre is a self-hosted server and client that mimics an image posting service. The purpose of developing Emigre is to validate technology stacks that are not used in the company that I belong to. In other words, satisfying technical interests is more important than creating a popular service.


Emigre is expected to be implemented using the following technologies:
- Server-side: Golang
- Infrastructure: AWS (Amazon Web Services)
- Client-side: Android application or GUI written in Golang

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