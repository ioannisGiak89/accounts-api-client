#!/usr/bin/env bash

cd "$SRC_DIR"

go mod tidy

echo
echo "=============================================================================="
echo "==> Running tests for Form3 Accounts API lib"
echo "=============================================================================="
echo

go test ./... -cover

echo
echo "=============================================================================="
echo "==> NOTE: Integration tests are making calls to fake API."
echo "==> To run the tests from your host change baseUrl to http://localhost:8080/ in form3Integration_test.go file"
echo "==> To run the tests within the docker container use"
echo "==> 'docker exec -it <container_name> bash' to log into the container and then run the tests"
echo "=============================================================================="

# Keep the container alive
tail -f /dev/null
