package actions

import "github.com/gobuffalo/buffalo"
import "encoding/json"
import "os/exec"
import "fmt"

type PushReq struct{
    Author  string
    Message string
    Repo    string
}

func PushPayloadHandler(c buffalo.Context) error {
    req := &PushReq{}
    if err := c.Bind(req); err != nil {
        return err
    }

    pullCMD := exec.Command("git", "-C /home/www-go/go/src/github.com/cileonard/lrn", "pull", "git@github.com:CILeonard/lrn")

    if err := pullCMD.run(); err != nil{
        fmt.println("Couldnt pull the git")
    }

    return nil
	//return c.Render(200, r.HTML("index.html"))
}

