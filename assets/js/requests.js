$(document).ready(function(){
    $(".btn-light").click(function(){
        receiverID = $(this).parent().parent(".result").attr("uid");
        $.post('/newrequest', 
            {
                authenticity_token: $("input[name=authenticity_token]").val(),
                receiver: receiverID,
                topic: $("#SearchTopic").attr("value"),
            },
            function(data, status){
                if (status == "success"){
                    alert("Successfully sent a request");
                }else{
                    alert("Internal error. Please Try Again");
                }
            }
        );
    });
});
