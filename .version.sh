# PREBUMP runs for all kinds of bump
PREBUMP=
  # it is usefull to make sure local is in sync with remote
  git fetch --tags origin master
  git pull origin master

# PREVERSION runs for any kinds of bump, it the last pre-hook
PREVERSION=
  # use it to declare tasks that should run for any kind of bump
  go vet ./...
  go fmt ./...
  go run main.go
  # finalize the changelog
  echo "about to changelog"
  changelog finalize --version !newversion!
  commit -q -m "changelog: !newversion!" -f change.log
  changelog md -o CHANGELOG.md --vars='{"name":"devopsdays-cli"}'
  commit -q -m "changelog: !newversion!" -f CHANGELOG.md
  go install --ldflags "-X main.VERSION=!newversion!"
  emd gen -in README.e.md > README.md
  commit -q -m "README: !newversion!" -f README.md

# POSTVERSION runs for any kind of bumps
POSTVERSION=
  # use it to sync your local to the remote
  git push
  git push --tags
