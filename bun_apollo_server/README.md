```
 ____                           
/ ___|  ___ _ ____   _____ _ __ 
\___ \ / _ \ '__\ \ / / _ \ '__|
 ___) |  __/ |   \ V /  __/ |   
|____/ \___|_|    \_/ \___|_|   
                                
```
# Apollo Server w/ Bun

## Setup Bun Environment w/ Docker
Optional: Pull the image with `docker pull oven/bun:latest`

Setup an alias to use Bun from the current directory:
```bash
alias bun="docker run --name bun-temp-runner --interactive --tty --rm --network host -v $(pwd):/app -w /app oven/bun bun"
```

To install dependencies:
```bash
bun install
```

To run the server:
```bash
bun main.ts
```

