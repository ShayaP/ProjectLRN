package actions

import(
    "strconv"

    "github.com/cileonard/lrn/models"

    "github.com/gobuffalo/uuid"
    "github.com/gobuffalo/pop"
)

var statusmap map[int]string

func init(){

    statusmap = map[int]string{
        1: "Accepted",
        2: "Pending",
        3: "Withdrawn",
        4: "Rejected",
    }
}


func SplitRequestsByStatus(requests models.Requests, isSender bool, tx *pop.Connection) ([][]string, error){
    listRequestStrings := make([][]string, 0, len(requests))
    for ind, req := range requests{
        user := &models.User{}
        var id uuid.UUID
        if (isSender){
            id = req.ReceiverID
        }else{
            id = req.SenderID
        }
        if err := tx.Find(user, id); err != nil{
            return nil, err
        }
        name := user.FirstName + " " + user.LastName
        date := req.CreatedAt.Format("1/2/06")
        rating := strconv.FormatFloat(float64(user.AvgRating), 'f', 2, 32)
        listRequestStrings[ind] = []string{name, rating, req.Topic, statusmap[req.Status], date}
    }
    return listRequestStrings, nil
}

