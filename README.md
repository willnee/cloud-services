# Go template
* Chứa sẵn code connect db, redis, mysql, mongo
* Tổ chhức source theo coding convention

### How to use:
- Create .env file from .env & config env values
- Fetch dependencies : `go mod tidy` `go mod download`
- Build app : `go build -o app`
- Run Server: `./app api`

1 Coding convention

1.1 Project structure

```
cmd // Chứa main run service
   └── api.go
   └── catalog.go
   └── root.go
di // Dependency Injection
   └── redisfx
      └── initialize.go
   └── dbfx
      └── initialize.go
pkg // Chứa common function
   config
      └── config.go
   constant
      └── constant.go
      └── errros.go
   redis
   └── redis.go
service // Chứa business logic tương ứng từng service
   payment
      delivery // handler nhận request, validate input
         grpc
            middleware
               └── middleware.go
            └── server.go
         http
            middleware
               └── middleware.go
            └── payment_handler.go
            └── route.go
      repository // Tương tác DB, 3rd api
         └── account.go
         └── transaction_test.go
         └── transaction.go
         └── user.go
      usecase // Phần xử lý logic, get/mix data từ db, api
         └── payment_ucase.go
         └── payment_ucase_test.go
```

1.2 HTTP response

1.2.1 HTTP code: `2xx`:success; `4xx`:client input error; `5xx`: server err

1.2.2 Response body
```json
{
   "status":  1, // 1: success, 0: failure
   "code": 1 // 1: success, <0 :failure  
   "message": "message",
   "data": {
      "products": []
   },
  "trace_id": "2109150921_fnlaw82nslda2"
}
```
### Step coding API
The step will be reverse direction of layer dependency: repositories -> usecases -> handlers.
### Rule of thumb
**Single responsibility** function should do one thing and do it well.

## <a name="commits"></a> Git Commit Guidelines

We have very precise rules over how our git commit messages can be formatted.
This leads to **more
readable messages** that are easy to follow when looking through the **project history**.
But also, we use the git commit messages to **generate the change log**.

The commit message formatting can be added using a typical git workflow or through the use of a CLI
wizard ([Commitizen](https://github.com/commitizen/cz-cli)).
To use the wizard, run `yarn run commit` in your terminal after staging your changes in git.

### Commit Message Format

Each commit message consists of a **header**, a **body** and a **footer**.  The header has a special
format that includes a **type**, a **scope** and a **subject**:

```
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

The **header** is mandatory and the **scope** of the header is optional.

Any line of the commit message cannot be longer 100 characters! This allows the message to be easier
to read on GitHub as well as in various git tools.

### Revert

If the commit reverts a previous commit, it should begin with `revert: `, followed by the header
of the reverted commit.
In the body it should say: `This reverts commit <hash>.`, where the hash is the SHA of the commit
being reverted.
A commit with this format is automatically created by the [`git revert`][git-revert] command.

### Type

Must be one of the following:

* **feat**: A new feature
* **fix**: A bug fix
* **docs**: Documentation only changes
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
* **refactor**: A code change that neither fixes a bug nor adds a feature
* **perf**: A code change that improves performance
* **test**: Adding missing or correcting existing tests
* **chore**: Changes to the build process or auxiliary tools and libraries such as documentation generation
* **chg**, **change**: A code change to implement a feature

### Scope

The scope could be anything specifying place of the commit change. For example `logger`,
`redis`, `context`, `errors`, etc...

You can use `*` when the change affects more than a single scope.

### Subject

The subject contains succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize first letter
* no dot (.) at the end

### Body

Just as in the **subject**, use the imperative, present tense: "change" not "changed" nor "changes".
The body should include the motivation for the change and contrast this with previous behavior.

### Footer

The footer should contain any information about **Breaking Changes** and is also the place to
[reference GitHub issues that this commit closes][closing-issues].

**Breaking Changes** should start with the word `BREAKING CHANGE:` with a space or two newlines.
The rest of the commit message is then used for this.


Example:
* feat(cart): implement cart repository
* test(product): add unit test for product handler
* fix(checkout): fix error for payment method
   
