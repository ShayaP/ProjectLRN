package actions

import(
    "github.com/cileonard/lrn/models"
)

func SplitRequestsByStatus(requests models.Requests, user *models.User) ([]models.Requests, error){
    for ind, req := range requests{
        if (req.Status % 2) == isTutor{ //
        }
    }
    return [](models.Requests){sentBy, receivedBy}, nil
}

