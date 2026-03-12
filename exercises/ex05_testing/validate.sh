#!/bin/bash
echo "Running your tests..."
go test -v ./exercises/ex05_testing/... 2>&1
if [ $? -eq 0 ]; then
    echo "All tests passed!"
else
    echo "Some tests failed. Check the output above."
    exit 1
fi
