#!/bin/sh

case $1 in
worker) ./main -isWorker;;
web) ./main;;
*)
  echo "Invalid option"
  exit 1
  ;;
esac
