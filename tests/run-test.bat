@echo off
REM Set up Go environment
echo Initializing Go module...
go mod init github.com/eazycloudlife/terraform-azure-labels

echo Tidy up Go modules...
go mod tidy

echo Running Go tests with a timeout of 30 minutes...
go test -v -timeout 30m

REM Check if the Go test command was successful
if %errorlevel% neq 0 (
    echo Test failed.
    exit /b %errorlevel%
) else (
    echo Test passed successfully.
)
pause
