function openSocket(a) {
    var socket = new WebSocket(a);
    socket.onclose = function () {
        setTimeout(function () {
            openSocket(a)
        }, 2E3)
    };
    socket.onmessage = function () {
        setTimeout(function(){
            location.reload();
        },2000);
    }
}

try {
    if (window.WebSocket) try {
        openSocket("ws://localhost:12450/reload")
    } catch (a) {
        console.error(a)
    } else console.log("Your browser does not support WebSockets.")
} catch (a) {
    console.error("Exception during connecting to Reload:", a)
}
;
