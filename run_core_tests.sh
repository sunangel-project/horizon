#!/bin/bash

go test
go test codeberg.org/energiesandsuch/horizon/test_core -v

exit $?
