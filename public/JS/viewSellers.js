

;(function() {

    var httpRequest;

    $(document).on("click", "li", selectSeller);
    
    function selectSeller() {
        
        var self = this;
        //sanity check
        console.log("seller selected");
        
        httpRequest = new XMLHttpRequest();
        // response handler
        httpRequest.onreadystatechange = readResponse;
        
        // requests
        httpRequest.open("POST", "/sendRequestToSeller");
        httpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        var s = $(self).attr("id");
        httpRequest.send("sellerName="+s);
        
        function readResponse() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    //TODO: do something after receiving response
                    $("div.alerts").html(
                    "<div class='alert alert-success'>"+
                    "Success! list sent to seller"+
                    "</div>"); 
                    console.log("request sent to seller!");
                }
            }
        }
    }
})();