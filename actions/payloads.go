package actions

import (
    "github.com/gobuffalo/buffalo"
    //import "encoding/json"
    "os/exec"
    "fmt"
)
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
    head := c.Request().Header
    eventType := head.Get("X-GitHub-Event")
    if eventType == "" {
        //Not a git event
        return c.Render(500, r.String("Not a GitHub Event"))
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

