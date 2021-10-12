package cli

import(
    "fmt"
    "encoding/json"
    "os"
    "log"
    "strconv"
    "sync"
    "io/ioutil"
    "net/http"
    "breakingbad/data"
)

const (
    baseURL = "https://www.breakingbadapi.com/api/"
)

func GetEpisodeFromApi(wg *sync.WaitGroup) {

    var resp_body []data.Response

    episodeNum := "episodes/" + strconv.Itoa(data.Episode)
    
    // Get response from REST API
    resp, err := http.Get(baseURL + episodeNum)

    // Close response body at end of the function
    defer resp.Body.Close()

    // Handle the error
    if err != nil {
        log.Fatal(err)
    }

    // if response is proper print output
    if resp.StatusCode == 200 {
        // Read response body
        resp_data, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Failed to read data from Response with ERROR")
            log.Fatal(err)
        }
        // Converting JSON to slice of Response struct from package data
        json.Unmarshal(resp_data, &resp_body)

        for _, r_body := range resp_body {
            fmt.Printf("%20s : %v\n", "Episode ID", r_body.Eid)
            fmt.Printf("%20s : %s\n", "Title", r_body.Title)
            fmt.Printf("%20s : %s\n", "Season", r_body.Season)
            fmt.Printf("%20s : %s\n", "Episode Number", r_body.Epnum)
            fmt.Printf("%20s : %s\n", "Air Date", r_body.Airdate)
            fmt.Printf("%20s : %s\n", "Series", r_body.Series)
            fmt.Printf("%20s : ", "Characters")
            for _, ch := range r_body.Characters {
                fmt.Printf("%s, ", ch)
            }
        }
    } else {
        fmt.Println("Episode", data.Episode, "not Found!")
    }

    // Send waitgroup response
    wg.Done()
}
