#!/usr/bin/env sh
set -ex

case $1 in
  lint) #
    golangci-lint run
    ;;
  fix) #
    go fmt ./...
    ;;
  cover) #
    go test ./... -coverprofile=cover.out
    ;;
  mocks) #
    # Remove:
    find . -name '*_mock.go' \
    | while read -r v; do
      rm "$v";
    done

    # Create:
    grep -rl interface src \
    | grep -v '_mock.go' \
    | while read -r v; do
      dir="$(dirname "$v")"
      file="$(basename "$v")"
      file="$(echo "$file" | sed 's/...$//')_mock.go"
      mockgen -source "$v" -destination "$dir/mocks/$file";
    done
    ;;
  *)
    cat <<INFO

Available commands:
./project.sh
$(grep ') #' "$0" | grep -v 'grep $0')

INFO

    exit 1
    ;;
esac