
;(function() {

    var httpRequest;
    
    $(".listSellers").click(selectSeller)
    
    function selectSeller() {
        //sanity check
        console.log("seller selected")
        
        httpRequest = new XMLHttpRequest();
        // response handler
        httpRequest.onreadystatechange = readResponse;

        // requests
        httpRequest.open("POST", "/sendRequestToSeller");
        httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        var s = $(this).children("li").attr("id");
        httpRequest.send("sellerName="+s);

        function readResponse() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    //TODO: do something after receiving response
                    console.log("request sent to seller!")
                }
            }
        }

    }
}());