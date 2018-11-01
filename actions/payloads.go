package actions

import (
    "github.com/gobuffalo/buffalo"
    "os/exec"
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
                output := & Response{"pong"}
                return c.Render(201, r.JSON(output))
            case "push":
                fmt.Println("Push received")
                pullCMD := exec.Command("git", "-C /home/www-go/go/src/github.com/cileonard/lrn", "pull", "git@github.com:CILeonard/lrn.git")
                if err := pullCMD.Run(); err != nil {
                    fmt.Println("Error pulling")
                    return c.Render(500, r.String(err.Error()))
                }
                output := & Response{"Succesful Pull!"}
                return c.Render(201, r.JSON(output))
            default:
                return c.Render(500, r.String("Not a handled Event"))
        }
    }
}

