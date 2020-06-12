
;(function() {
    var httpRequest;
    $("#addItem").click(sendData)
    
    function sendData() {
        
        httpRequest = new XMLHttpRequest();
        //handle the response
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readystate == XMLHttpRequest.DONE) {
                console.log("response recieved")
                if (httpRequest.status == 200 ) {
                    alert("item added succesfully!")
                } else {
                    alert("there was a problem with the request")
                }
            }
        };
        
        //send form value to server
        httpRequest.open("POST", "/addItem", true);
        
        var s = $("#newItem").val();
        httpRequest.send("item="+s);
        console.log("request sent:"+s)
    };
}());