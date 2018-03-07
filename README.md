# apiDocs

## What is **apiDocs**?
**apiDocs** is an apiblueprint rendering microservice.<br>
It is used to render an apiblueprint file format to a html output. Besides that you can also specify a custom html template in your request. So that this rendering service serves as a microservice.

## How it works
**apiDocs** uses the [](https://github.com/bukalapak/snowboard) goLang rendering library/service. Thus, this microservice is only available on Linux-based systems because [snowboard](https://github.com/bukalapak/snowboard) has [drafter]() as a dependency.

## Installation
### Download
Download the latest recommended build:
```sh
$ wget https://github.com/playnet-public/apiDocs/releases/download/<version>/apiDocsd_<ostype>.tar.gz
```
Variable | Description 
--- | ---
`<version>` | Release tag name.
`<ostype>` | OS type either `x64` or `x86`.

Now untar the package you've downloaded:
```sh
$ tar -xzf apiDocsd_<ostype>.tar.gz
```

## Usage
Now that you have downloaded the package, you can start the microservice server by:
```sh
$ ./apiDocsd <parameter>
```

To customize the execution of the http server you can use several flags and environment variables:
### Flags
Flag | Data Type | Description
--- | --- | ---
`-help`, `-h` || Shows the help with all available commands.
`-address` | `string` | The address of the http server the service is bind to.
`-port` | `int` | The http server port the server listen on.
`-defaultTemplate` | `string` | The default template file for rendering a given apiblueprint.

### Environment Variables
Variable | Data Type | Description
--- | --- | ---
`APIDOCSD_ADDRESS` | `string` | Same as `-address` flag.
`APIDOCSD_PORT` | `int` | Same as `-port` flag.
`APIDOCSD_DEFAULTEMPLATE` | `string` | Same as `-defaultTemplate`.

### Examples
> #### Example 1
> A server hosted on localhost on port 99 with a default template located in local directory:
> ```sh
> $ ./apiDocsd -address localhost -port 99 -defaultTemplate alpha.html
> ```

> #### Example 2
> A server hosted on all linked ip addresses on default port 8088:
> ```sh
> $ ./apiDocsd -address 0.0.0.0
> ```
