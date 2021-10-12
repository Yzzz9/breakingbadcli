package main

import(
    "fmt"
    "os"
    "flag"
    "breakingbad/cli"
    "breakingbad/data"
    "sync"
)

func init() {
    const (
        episodes = "Search episode number[1-102], e.g. 60"
    )
    flag.IntVar(&data.Episode, "ep", -1, episodes)
}

func main() {
    flag.Parse()

    // if number of flags are zero, display message and exit
    if flag.NFlag() == 0 {
        fmt.Printf("Usage: %s [option]\n", os.Args[0])
        fmt.Println("Options :")
        flag.PrintDefaults()
        os.Exit(1)
    }

    fmt.Printf("Searching for episode %v ...\n", data.Episode)

    // Adding waitgroups
    var wg sync.WaitGroup
    
    wg.Add(1)

    // Calling method from cli package
    go cli.GetEpisodeFromApi(&wg)

    wg.Wait()
}
