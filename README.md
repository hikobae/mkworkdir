# mkworkdir

Make daily working directory.

## Customize

You can customize the behavior by setting the following environment variables.

- WORKDIR: root working directory path
- WORKDIR_OPEN: opening the working directory if it sets any value

## (Windows) How to not show the Command Prompt

### Build with `--ldflags "-H windowsgui"` options

like below:

```
> go build --ldflags "-H windowsgui -w -s"
```

### Use WSH

See `misc/wsh/README.md`

## License

MIT

## Author

TAKAHASHI Satoshi (hikobae@gmail.com)
