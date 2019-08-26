# Carnot

Small utility to copy/move files from a source directory to a target directory and organize them by date.

## Dependencies

This tool uses the [ExifTool](https://www.sno.phy.queensu.ca/~phil/exiftool/) to retrieve a file's meta data
in order to read out the content creation date. If it can't figure it out, it will fall back to the system's
file birth date.

## Installation

Download and install precombiled [releases](https://github.com/wejustdostuff/carnot/releases).

On Mac OS X you can also install the latest release via Homebrew:

```bash
brew install wejustdostuff/carnot/carnot
```

## Usage

```
Available Commands:
  Move        Move all files from the source directory to the target directory.
  copy        Copy all files from the source directory to the target directory.
  help        Help about any command
  list        List all files in the given source directory.
  version     Version will output the current build information

Flags:
      --colors              enable or disable color output (default true)
  -c, --config string       config file (default is $HOME/.carnot.yaml)
  -h, --help                help for carnot
      --log-format string   specify log format to use when logging to stderr [text or json] (default "text")
      --log-level string    specify log level to use when logging to stderr [error, info, debug] (default "info")

Use "carnot [command] --help" for more information about a command.
```

## Example

```bash
$ tree path/to/source
path/to/source
├── image.jpg
├── video.mov
├── text.txt
├── movie.mp4
└── other.null
```

Preview what the tool would do

```bash
$ carnot list -s path/to/source -t path/to/destination
SOURCE                      TARGET                                              DATE FIELD
 tmp/source/image.jpg   ->  path/to/destination/2009/09/24/20090924-132253.jpg  DateTimeOriginal
 tmp/source/video.mov   ->  path/to/destination/2017/09/09/20170909-213011.mov  DateTimeOriginal
 tmp/source/text.txt    ->  path/to/destination/2017/09/14/20170914-113347.txt  BirthTime
 tmp/source/movie.mp4   ->  path/to/destination/2018/01/18/20180118-092446.mp4  DateTimeOriginal
 tmp/source/other.null  ->  path/to/destination/2019/08/23/20190823-114320.null BirthTime
```

Move or copy the files to the destination folder

```bash
$ carnot move -s path/to/source -t path/to/destination
$ carnot copy -s path/to/source -t path/to/destination
```

## Contributing
In lieu of a formal style guide, take care to maintain the existing coding style.

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License
Copyright (c) We Just Do Stuff under the MIT license.
