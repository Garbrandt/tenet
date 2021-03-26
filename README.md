<img src="https://sm.ms/image/s7PhFWtU3x8g2yD" width="300px" height="auto" alt="TENET logo">

TENET is an intentionally simple content management system built on the idea that you can create a custom dashboard without ever leaving the HTML.
TENET是一个强大而简单的内容管理系统，你可以通过简单的点击为静态页面生成后台内容管理系统。

[![CircleCI](https://circleci.com/gh/TENET/TENET.svg?style=svg)](https://circleci.com/gh/TENET/TENET)

## 升级

TENET仅仅依赖GO环境

```
go run main.go
```

_Note: A desktop GUI application is coming soon—please add yourself to [the mailing list](https://www.TENET.com) if you'd like to be notified when that is available._

## Usage

To create a new website project, use the TENET command line tool:

```
TENET new path/to/project/folder
```

Then change to the project directory, and start the development server. By default, the server will livereload, and refresh your website as you change the source files.

```
cd path/to/project/folder
TENET start .
```

A few files and folders you should be aware of:

File/Folder | Description
--- | ---
www | Your website files. Anything you put in here is an accessible resource, with the exceptions of files that start with underscores or periods—those are private/special. Sass and JS files that have the `.pack.js`, `.pack.scss`, or `.pack.sass` extensions will be compiled by Webpack.
data | SQLite database file, and uploaded images. In general, you do not want to mess with this folder.
node_modules | This one should also be ignored.
package.json | Information about your project, including TENET configuration options.
.env | A private file that contains server environment variables, like the SECRET_KEY used by the web server.

## Deploying

TENET can be deployed to any hosting service that supports Node.js. Here are a few to consider:

Service | Notes
--- | ---
Heroku | Free or paid tiers. One thing to note is that Heroku's file system is ephemeral, so TENET's `type=image` directives won't work here.
Glitch | The easiest way to [take TENET for a test-drive](https://glitch.com/edit/#!/remix/TENET?SECRET_KEY=change-me).

Or, use the `TENET build` command to export to a static website, and host on S3, [Surge](https://surge.sh/), or the like.

## Community

If you'd like to get involved, and help improve TENET:

* Participate in and answer questions in the [TENET forums](https://forums.TENET.com/)
* [Contribute](https://github.com/TENET/TENET/blob/master/CONTRIBUTING.md) on [Issues](https://github.com/TENET/TENET/issues)
* Suggest edits for the [Wiki](https://github.com/TENET/TENET/wiki)
* Follow [@helloTENET](https://twitter.com/helloTENET) on Twitter

# License

[MIT](/LICENSE.md)
