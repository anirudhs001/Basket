
;(function() {

    var HttpRequest = new XMLHttpRequest()

    $(".listSellers").click(selectSeller)

    function selectSeller() {

        // response handler
        HttpRequest.onreadystatechange = readResponse;

        // requests
        HttpRequest.open("post", "/selectSeller");
        HttpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        var s = $(this).attr(id);
        HttpRequest.send("name="+s);

        function readResponse() {
            if (HttpRequest.readyState === XMLHttpRequest.DONE) {
                if (HttpRequest.status === 200) {
                    //TODO: do something after receiving response
                }
            }
        }

    }
}())