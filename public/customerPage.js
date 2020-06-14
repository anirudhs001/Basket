
;(function() {
    var httpRequest;
    $("#addItem").click(sendData)
    
    function sendData() {
        
        httpRequest = new XMLHttpRequest();
        //handle the response
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200 ) {

                    $("#list").append(httpRequest.responseText)
                    console.log("item added!")
                } else {
                    alert("there was a problem with the request")
                }
            }
        };
        //send form value to server
        httpRequest.open("POST", "/addItem", true);
        
        var s = $("#newItem").val();
        httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        httpRequest.send("item="+s)
        console.log("request sent:"+s)
    };

    $(document).on("click", ".listButton", delItem); //TODO why document?

    function delItem() {

        httpRequest = new XMLHttpRequest();
        //respone handler
        httpRequest.onreadystatechange = deleteItem;

        //send request
        var id = $(this).attr("id");
        httpRequest.open("POST", "/delItem", true);
        httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        httpRequest.send("itemID="+id);
        //sanity check
        console.log("request sent to delete item");

        function deleteItem() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200 ) {
                    $("#"+id).parent().remove()
                }
            }
        };
    }
}());
