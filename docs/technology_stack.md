# Technology Stack
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

## Infrastructure tools
| Category | Technology |
|:---|:---|
| Infrastructure as Code | [CloudFormation](https://aws.amazon.com/cloudformation/?nc1=h_ls) |
| AWS Mock| [localstack](https://localstack.cloud/)|
| AWS CLI| [aws-cli](https://github.com/aws/aws-cli), [awscli-local](https://github.com/localstack/awscli-local)|

## Client-side tools
WIP