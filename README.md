# GitHub Activity
Go Command Line Tool to Check A User's GitHub Activity

## Setup

Follow these steps to run the project locally.

**Prerequisites**

Make sure the following are installed.

- [Git](https://git-scm.com)
- [Go](https://go.dev/)

**Cloning the Repository**

```bash
git clone https://github.com/juliansommer/github-activity.git
cd github-activity
```

**Installing Dependencies**

```bash
go mod tidy
```

**Building the Application**
```bash
go build -o github-activity
```

**Running the Application**

```bash
# Using go run
go run main.go <username>

# Or using the built binary
./github-activity <username>

# Example:
go run main.go juliansommer
```

# TODO
- Add flags for page numbers
- Add output format options (JSON, table)
