$(document).ready(function(){
    $(".btn-sender").click(function(){
        receiverID = $(this).parent().parent(".result").attr("uid");
        $.post('/newrequest', 
            {
                authenticity_token: $("input[name=authenticity_token]").val(),
                receiver: receiverID,
            },
            function(data, status){
                if (status == "success"){
                    alert("Successfully sent a request")
                }else{
                    alert("Internal error. Please Try Again")
                }
            }
        );
    });
});
