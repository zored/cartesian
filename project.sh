#!/usr/bin/env sh
set -ex

case $1 in
  test) #
    go test ./... -race
    ;;
  lint) #
    golangci-lint run
    ;;
  fix) #
    go fmt ./...
    ;;
  cover) #
    go test ./... -coverprofile=cover.out
    ;;
  mock*) #
    # Remove:
    find . -name '*_mock.go' \
    | while read -r v; do
      rm "$v";
    done

    # Create:
    grep -rl interface src \
    | grep -v '_mock.go' \
    | grep -v '_test.go' \
    | while read -r v; do
      dir="$(dirname "$v")"
      file="$(basename "$v")"
      file="$(echo "$file" | sed 's/...$//')_mock.go"
      out="$dir/mocks/$file"
      mockgen -source "$v" -destination "$out";
      if [ "$(wc -l "$out" | awk '{print $1}')" = 9 ]; then
        rm "$out"
      fi
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