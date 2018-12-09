$(document).ready(function(){
    $(".btn-sender").click(function(){
        requestID = $(this).parent().parent(".result").attr("uid");
        $.post('/updatesentrequest', 
            {
                authenticity_token: $("input[name=authenticity_token]").val(),
                request: requestID,
            },
            function(data, status){
                if (status == "success"){
                    alert("Successfully withdrawn a request");
                }else{
                    alert("Internal error. Please Try Again");
                }
                location.reload();
            }
        );
    });
    $(".btn-receiver").click(function(){
        requestID = $(this).parent().parent(".result").attr("uid");
        $.post('/updaterecrequest', 
            {
                authenticity_token: $("input[name=authenticity_token]").val(),
                request: requestID,
                decision: $(this).siblings("select").val(),
            },
            function(data, status){
                if (status == "success"){
                    alert("Successfully updated a request");
                }else{
                    alert("Internal error. Please Try Again");
                }
                location.reload();
            }
        );
    });
});
