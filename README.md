```text
██╗   ██╗██████╗ ██╗ ██████╗████████╗
╚██╗ ██╔╝██╔══██╗██║██╔════╝╚══██╔══╝
 ╚████╔╝ ██║  ██║██║██║        ██║   
  ╚██╔╝  ██║  ██║██║██║        ██║   
   ██║   ██████╔╝██║╚██████╗   ██║   
   ╚═╝   ╚═════╝ ╚═╝ ╚═════╝   ╚═╝   
 ```

Ydict, another command-line youdao dictionary!

![](https://raw.githubusercontent.com/rockagen/ydict/master/snapshots/ydict.png)


## Features

* Chinese -> English
* English -> Chinese
* Show hints if word is not found
* Speech
* Show example sentences
* Vim support

## Installation


#### shell

Linux

`curl -L https://github.com/rockagen/ydict/raw/master/ydict -o ydict`

MacOS

`curl -L https://github.com/rockagen/ydict/raw/master/ydict_osx -o lyrics`

#### Using Go

```bash
go get github.com/rockagen/ydict
```

#### Integrate with Vim

To query words from Vim, you need another Vim plugin: [vim-ydict](https://github.com/rockagen/vim-ydict)

## Speech

Speech feature is available. You need to install mpg123 to enable this feature.

___NOTICE:___ Currently, speech feature is only available for MacOS/Linux.

#### Mac OS

```bash
brew install mpg123
```
#### Ubuntu

```bash
sudo apt-get install mpg123
```

#### CentOS

```bash
yum install -y mpg123
```

## Usage

1. Query

```text
ydict <word(s) to query>
```

2. Query with speetch (__Available for MacOS & Linux__)

```text
ydict <word(s) to query> -v
```

3. Query and show more example sentences

```text
ydict <word(s) to query> -m
```

4. Interaction mode

```text
ydict -i
```

## SOCKS5 proxy

You can use SOCKS5 proxy. At the same directory of ydict, just create a ```.env``` file:

```text
SOCKS5=127.0.0.1:7070
```

Now all the queries will go through the specified SOCKS5 proxy.

## Help

Just type "ydict" to get help.
  
## Licence

[MIT License](https://github.com/rockagen/ydict/blob/master/LICENSE)
