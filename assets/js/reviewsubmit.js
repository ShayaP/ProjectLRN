$(document).ready(function(){
    $(".btn-light").click(function(){
        receiverID = $(this).parent().parent(".review").attr("uid");
        $.post('/reviewspage', 
            {
                authenticity_token: $("input[name=authenticity_token]").val(),
                rating: $(this).parent().siblings("td").children("input[type=number]").val(),
                review: $(this).parent().siblings("td").children("input[type=text]").val(),
                reviewedID: receiverID,
            },
            function(data, status){
                if (status == "success"){
                    alert("Successfully submitted a review");
                }else{
                    alert("Internal error. Please Try Again");
                }
                location.reload();
            }
        );
    });
});
