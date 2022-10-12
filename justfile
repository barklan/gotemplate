myapp-up:
    set -a; . .env; set +a && go run ./cmd/myapp

up-watch:
    set -a; . .env; set +a && reflex -c reflex.conf

myapp-build:
    bash ./scripts/build.sh
