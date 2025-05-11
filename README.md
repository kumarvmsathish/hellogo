# hellogo

This Project is about setting up local development with Golang and basic learning stuffs

## Authors

- [@sathish] (https://github.com/kumarvmsathish)

## Setup 

    Current Go version used - `go version go1.24.1 darwin/arm64`

  Run command to setup  

  ```bash
        go mod init github.com/kumarvmsathish/hellogo # {remote}/{username}/{projectname}
  ```  

### Run and build
  To run the application use command,
  ```bash
        go run main.go # run using file name - suitable for single files
        # or
        go run . # compiles all .go files
        # or
        go run my/cmd # single know package
        # or
        ./hellogo # use only after generating build using eg: `go build` command execution

  ```

  Go help run `go help run ` -- provides help commands for run

  Go build - it makes production ready build
  ```bash
        go build # builds package of the current directory and generates binary executable file
        # hellogo -- will be created as executable file

        go build && ./hellogo # compile the build file and run
  ```

  Go Install -- helpful for local executables to run directly
  ```bash
        go install # builds and installed the `hellogo` program globally in local machine and no need to run `go build` before -- not req
        cd ../ # navigate to check and run `hellogo` directly instead of using ./hellogo 
        hellogo
  ```
## Module and add Dependency
    Create modules using,
  ```bash
        cd ../ # go to parent directory -- we are creating new module separately
        # run below command for creating library/module
        mkdir mystrings
        cd mystrings
        go mod init github.com/kumarvmsathish/mystrings

  ```

      Accessing local module as remote module,
      Go to `go.mod` file and add below,
``` bash
      go.mod

      replace github.com/kumarvmsathish/mystrings v0.0.0 => ../mystrings
      require (
            github.com/kumarvmsathish/mystrings v0.0.0 
      )

```

      Import or download modules,
```bash
      go get github.com/wagslane/go-tinytime # downloads and installs in local machine
```





  
