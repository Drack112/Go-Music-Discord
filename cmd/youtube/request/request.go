package youtube

import (
    "encoding/json"
    "net"
    "net/url"
    "os/exec"
    "strings"
)

type Request struct {
    ReqURL string
    Title  string
    // StreamURL string
    nowPlaying chan bool
}

func isValidURL(reqUrl string) bool {
    uri, err := url.Parse(reqUrl)
    if err != nil {
        return false
    }
    // Check it's a valid domain name
    _, err = net.LookupHost(uri.Host)
    return err == nil
}

func New(query string, callback chan bool) (*Request, error) {
    var title string
    var reqUrl string
    if isValidURL(query) {
        output, err := exec.Command("youtube-dl", "-e", query).Output()
        if err != nil {
            return nil, err
        }
        title = string(output)
        reqUrl = query
    } else {
        output, err := exec.Command("youtube-dl", "-j", "ytsearch:"+query).Output()
        if err != nil {
            return nil, err
        }
        var info map[string]interface{}
        err = json.Unmarshal(output, &info)
        if err != nil {
            return nil, err
        }
        title = info["title"].(string)
        reqUrl = info["webpage_url"].(string)
    }
    return &Request{reqUrl, strings.TrimSuffix(string(title), "\n"), callback}, nil
}

func (r Request) GetStream() (string, error) {
    streamUrl, err := exec.Command("youtube-dl", "-f", "bestaudio", "-g", r.ReqURL).Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSuffix(string(streamUrl), "\n"), nil
}

func (r Request) NowPlaying() {
    r.nowPlaying <- true
}

func (r Request) Cancel() {
    r.nowPlaying <- false
}
