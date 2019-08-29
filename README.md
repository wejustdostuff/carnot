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
  copy        Copy all files from the source directory to the target directory.
  exif        List all files exif data in the given source directory.
  help        Help about any command
  list        List all files in the given source directory.
  move        Move all files from the source directory to the target directory.
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
├── audio
│   ├── audio.mp3
│   ├── audio.ogg
│   └── audio.wav
├── document
│   ├── document.doc
│   ├── document.docx
│   ├── document.odp
│   ├── document.ods
│   ├── document.odt
│   ├── document.pdf
│   ├── document.ppt
│   ├── document.xls
│   └── document.xlsx
├── image
│   ├── image.gif
│   ├── image.ico
│   ├── image.jpg
│   ├── image.png
│   ├── image.svg
│   └── image.tiff
├── other
│   ├── other.csv
│   └── other.html
└── video
    ├── video.3gp
    ├── video.avi
    ├── video.flv
    ├── video.mkv
    ├── video.mov
    ├── video.mp4
    ├── video.ogg
    └── video.wmv
```

Preview what the tool would do

```bash
$ carnot list -s path/to/source -t path/to/destination
SOURCE                      TARGET                                              DATE FIELD
 path/to/source/audio/audio.mp3         ->  path/to/destination/2019/08/27/20190827-092039.mp3   BirthTime         2019.08.27  09:20:39
 path/to/source/audio/audio.ogg         ->  path/to/destination/2019/08/27/20190827-092059.ogg   BirthTime         2019.08.27  09:20:59
 path/to/source/audio/audio.wav         ->  path/to/destination/2019/08/27/20190827-092047.wav   BirthTime         2019.08.27  09:20:47
 path/to/source/document/document.doc   ->  path/to/destination/2019/08/27/20190827-092126.doc   BirthTime         2019.08.27  09:21:26
 path/to/source/document/document.docx  ->  path/to/destination/2019/08/27/20190827-092132.docx  BirthTime         2019.08.27  09:21:32
 path/to/source/document/document.odp   ->  path/to/destination/2019/08/27/20190827-092234.odp   BirthTime         2019.08.27  09:22:34
 path/to/source/document/document.ods   ->  path/to/destination/2019/08/27/20190827-092223.ods   BirthTime         2019.08.27  09:22:23
 path/to/source/document/document.odt   ->  path/to/destination/2019/08/27/20190827-092215.odt   BirthTime         2019.08.27  09:22:15
 path/to/source/document/document.pdf   ->  path/to/destination/2019/08/27/20190827-092201.pdf   BirthTime         2019.08.27  09:22:01
 path/to/source/document/document.ppt   ->  path/to/destination/2019/08/27/20190827-092152.ppt   BirthTime         2019.08.27  09:21:52
 path/to/source/document/document.xls   ->  path/to/destination/2019/08/27/20190827-092140.xls   BirthTime         2019.08.27  09:21:40
 path/to/source/document/document.xlsx  ->  path/to/destination/2019/08/27/20190827-092145.xlsx  BirthTime         2019.08.27  09:21:45
 path/to/source/image/image.gif         ->  path/to/destination/2019/08/27/20190827-093035.gif   BirthTime         2019.08.27  09:30:35
 path/to/source/image/image.ico         ->  path/to/destination/2019/08/27/20190827-092326.ico   BirthTime         2019.08.27  09:23:26
 path/to/source/image/image.jpg         ->  path/to/destination/2014/09/21/20140921-160056.jpg   DateTimeOriginal  2014.09.21  16:00:56
 path/to/source/image/image.png         ->  path/to/destination/2019/08/27/20190827-093026.png   BirthTime         2019.08.27  09:30:26
 path/to/source/image/image.svg         ->  path/to/destination/2019/08/27/20190827-093042.svg   BirthTime         2019.08.27  09:30:42
 path/to/source/image/image.tiff        ->  path/to/destination/2019/08/27/20190827-093344.tiff  BirthTime         2019.08.27  09:33:44
 path/to/source/other/other.csv         ->  path/to/destination/2019/08/27/20190827-092341.csv   BirthTime         2019.08.27  09:23:41
 path/to/source/other/other.html        ->  path/to/destination/2019/08/27/20190827-092353.html  BirthTime         2019.08.27  09:23:53
 path/to/source/video/video.3gp         ->  path/to/destination/2019/08/27/20190827-093147.3gp   BirthTime         2019.08.27  09:31:47
 path/to/source/video/video.avi         ->  path/to/destination/2019/08/27/20190827-091940.avi   BirthTime         2019.08.27  09:19:40
 path/to/source/video/video.flv         ->  path/to/destination/2019/08/27/20190827-093129.flv   BirthTime         2019.08.27  09:31:29
 path/to/source/video/video.mkv         ->  path/to/destination/2019/08/27/20190827-093137.mkv   BirthTime         2019.08.27  09:31:37
 path/to/source/video/video.mov         ->  path/to/destination/2019/08/27/20190827-091923.mov   BirthTime         2019.08.27  09:19:23
 path/to/source/video/video.mp4         ->  path/to/destination/2019/08/27/20190827-093114.mp4   BirthTime         2019.08.27  09:31:14
 path/to/source/video/video.ogg         ->  path/to/destination/2019/08/27/20190827-092011.ogg   BirthTime         2019.08.27  09:20:11
 path/to/source/video/video.wmv         ->  path/to/destination/2019/08/27/20190827-092027.wmv   BirthTime         2019.08.27  09:20:27
```

Move or copy the files to the destination folder

```bash
$ carnot copy -s path/to/source -t path/to/destination
```

```bash
$ carnot move -s path/to/source -t path/to/destination
```

View results

```
$ tree path/to/target
path/to/target
├── 2014
│   └── 09
│       └── 21
│           └── 20140921-160056.jpg
└── 2019
    └── 08
        └── 27
            ├── 20190827-091923.mov
            ├── 20190827-091940.avi
            ├── 20190827-092011.ogg
            ├── 20190827-092027.wmv
            ├── 20190827-092039.mp3
            ├── 20190827-092047.wav
            ├── 20190827-092059.ogg
            ├── 20190827-092126.doc
            ├── 20190827-092132.docx
            ├── 20190827-092140.xls
            ├── 20190827-092145.xlsx
            ├── 20190827-092152.ppt
            ├── 20190827-092201.pdf
            ├── 20190827-092215.odt
            ├── 20190827-092223.ods
            ├── 20190827-092234.odp
            ├── 20190827-092326.ico
            ├── 20190827-092341.csv
            ├── 20190827-092353.html
            ├── 20190827-093026.png
            ├── 20190827-093035.gif
            ├── 20190827-093042.svg
            ├── 20190827-093114.mp4
            ├── 20190827-093129.flv
            ├── 20190827-093137.mkv
            ├── 20190827-093147.3gp
            └── 20190827-093344.tiff
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
