
;(function() {
    var httpRequest;
    $("#add-item").click(sendData)
    
    function sendData() {
        
        httpRequest = new XMLHttpRequest();
        //handle the response
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200 ) {
                    var resp = httpRequest.responseText;
                    $("#items").prepend(resp);
                    console.log("item added!")
                } else {
                    alert("there was a problem with the request")
                }
            }
        };
        //send form value to server
        httpRequest.open("POST", "/addItem", true);
        
        var s = $("#new-item").val();
        httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        httpRequest.send("item="+s)
        console.log("request sent:"+s)
    };
    
    $(document).on("click", ".list-button", delItem); //TODO why document?
    
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

