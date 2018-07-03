

$(function () {
    $('.ui.checkbox')
        .checkbox()
        ;
});


var ulborUriAddBntDisable = function () {
    setInterval(
        function () { 
            document.getElementById("ulborUriAddBnt").disabled = true;
            document.getElementById("progBar").style.display = 'block';
            $('#progBar')
                .progress('increment')
                ;
            setInterval(function () {
                $('#progBar')
                    .progress('increment')
                    ;
            }, 500)
        }, 500);
   
}
