package kitgen

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func Generate(inputPath string) error {

    fi, err := os.Stat(inputPath)
    if err != nil {
        return err
    }
    if fi.IsDir() {
        fmt.Println("dir?", inputPath)
    } else {
        runKitgen(inputPath)
    }

    return nil
}

func runKitgen(inputFile string) error {
    cmdArgs := []string{"-repo-layout", "flat", inputFile}
    fmt.Println("running: kitgen", strings.Join(cmdArgs, " "))
    cmd := exec.Command("kitgen", cmdArgs...)
    err := cmd.Run()
    if err != nil {
        return err
    }
    fmt.Println(cmd.Stdout)

    return nil
}
