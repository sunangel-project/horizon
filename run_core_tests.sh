#!/bin/bash

go test
go test github.com/sunangel-project/horizon/test_core -v

exit $?
