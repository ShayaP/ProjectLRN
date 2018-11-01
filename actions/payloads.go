package actions

import "github.com/gobuffalo/buffalo"
//import "encoding/json"
import "os/exec"
import "fmt"

type PushReq struct{
    Author  string
    Message string
    Repo    string
}

type GitPing struct{
    Zen     string
    Hook_ID string
    Hook    string
}

type Response struct{
    msg     string
}

func PushPayloadHandler(c buffalo.Context) error {
    head := c.Request.params.Header

    return c.Render(201, r.String("pong"))
    eventType := head["X-GitHub-Event"]
    if eventType == "" {
        //Not a git event
    }else{
        switch eventType{
            case "ping":
                output := Response{"pong"}
                return c.Render(201, r.JSON(output))
            case "push":
                //do push things
        }
    }
    req := &PushReq{}
    if err := c.Bind(req); err != nil {
        return c.Render(500, r.String(err.Error()))
    }

    pullCMD := exec.Command("git", "-C /home/www-go/go/src/github.com/cileonard/lrn", "pull", "git@github.com:CILeonard/lrn")

    if err := pullCMD.Run(); err != nil{
        fmt.Println("Couldnt pull the git")
    }

    return nil
	//return c.Render(200, r.HTML("index.html"))
}

